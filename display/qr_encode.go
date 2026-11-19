package display

// A small QR Code Model 2 encoder, byte mode, error-correction level L,
// versions 1-5. It exists here so the QRCode component needs no dependency
// beyond my-ui's public ui API; longer values are rejected by qrEncode.
//
// The algorithm follows the ISO/IEC 18004 layout: function patterns, a
// zig-zag data placement, Reed-Solomon parity and the best of the eight
// data masks. Only the single-block versions are supported, which is
// enough for the URLs and short login tokens the component shows.

var qrDataCodewords = [5]int{19, 34, 55, 80, 108}
var qrECCCodewords = [5]int{7, 10, 15, 20, 26}

// qrEncode returns the dark modules of value, or nil when value is too
// long for the versions this encoder supports.
func qrEncode(value string) [][]bool {
	raw := []byte(value)
	for v := 1; v <= len(qrDataCodewords); v++ {
		dataCap := qrDataCodewords[v-1] * 8
		if 4+8+len(raw)*8 > dataCap {
			continue
		}
		bits := qrBits(raw, dataCap)
		divisor := rsDivisor(qrECCCodewords[v-1])
		code := append(bits, rsRemainder(bits, divisor)...)
		return qrMatrix(v, code)
	}
	return nil
}

// qrBits builds the data codewords: mode, count, payload, terminator and
// the alternating pad bytes.
func qrBits(raw []byte, cap int) []byte {
	buf := 0
	n := 0
	write := func(v, width int) {
		for i := width - 1; i >= 0; i-- {
			buf = buf<<1 | (v>>i)&1
			n++
		}
	}
	out := make([]byte, 0, cap/8)
	flush := func() {
		for n >= 8 {
			n -= 8
			out = append(out, byte(buf>>n))
			buf &= (1 << n) - 1
		}
	}
	write(0b0100, 4) // byte mode
	write(len(raw), 8)
	for _, b := range raw {
		write(int(b), 8)
	}
	end := cap - n
	if end > 4 {
		end = 4
	}
	write(0, end)
	flush()
	for i := 0; len(out)*8 < cap; i++ {
		if i%2 == 0 {
			out = append(out, 0xEC)
		} else {
			out = append(out, 0x11)
		}
	}
	return out
}

// gfMul multiplies two GF(256) elements in the field QR uses.
func gfMul(x, y int) int {
	z := 0
	for i := 7; i >= 0; i-- {
		z = (z << 1) ^ ((z >> 7) * 0x11D)
		z ^= ((y >> i) & 1) * x
	}
	return z & 0xFF
}

// rsDivisor is the Reed-Solomon generator polynomial of degree d.
func rsDivisor(d int) []int {
	result := make([]int, d)
	result[d-1] = 1
	root := 1
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			result[j] = gfMul(result[j], root)
			if j+1 < d {
				result[j] ^= result[j+1]
			}
		}
		root = gfMul(root, 2)
	}
	return result
}

// rsRemainder returns the d parity codewords of data.
func rsRemainder(data []byte, divisor []int) []byte {
	result := make([]int, len(divisor))
	for _, b := range data {
		factor := int(b) ^ result[0]
		copy(result, result[1:])
		result[len(result)-1] = 0
		for i, c := range divisor {
			result[i] ^= gfMul(c, factor)
		}
	}
	out := make([]byte, len(result))
	for i, v := range result {
		out[i] = byte(v)
	}
	return out
}

// qrMatrix lays out code and returns the dark modules of the best mask.
func qrMatrix(version int, code []byte) [][]bool {
	size := 17 + 4*version
	dark := make([][]bool, size)
	fn := make([][]bool, size)
	for i := range dark {
		dark[i] = make([]bool, size)
		fn[i] = make([]bool, size)
	}
	set := func(x, y int, d bool) {
		if x >= 0 && x < size && y >= 0 && y < size {
			dark[y][x], fn[y][x] = d, true
		}
	}
	isFn := func(x, y int) bool {
		return x >= 0 && x < size && y >= 0 && y < size && fn[y][x]
	}

	for _, f := range [][2]int{{3, 3}, {size - 4, 3}, {3, size - 4}} {
		for dy := -4; dy <= 4; dy++ {
			for dx := -4; dx <= 4; dx++ {
				d := max(abs(dx), abs(dy))
				set(f[0]+dx, f[1]+dy, d != 2 && d != 4)
			}
		}
	}

	align := qrAlignment(version, size)
	for _, ay := range align {
		for _, ax := range align {
			if isFn(ax, ay) {
				continue
			}
			for dy := -2; dy <= 2; dy++ {
				for dx := -2; dx <= 2; dx++ {
					set(ax+dx, ay+dy, max(abs(dx), abs(dy)) != 1)
				}
			}
		}
	}

	for i := 0; i < size; i++ {
		if !isFn(6, i) {
			set(6, i, i%2 == 0)
		}
		if !isFn(i, 6) {
			set(i, 6, i%2 == 0)
		}
	}

	bits := make([]bool, len(code)*8)
	for i, b := range code {
		for j := 0; j < 8; j++ {
			bits[i*8+j] = (b>>(7-j))&1 != 0
		}
	}
	i := 0
	for right := size - 1; right > 0; right -= 2 {
		if right == 6 {
			right = 5
		}
		for vert := 0; vert < size; vert++ {
			for j := 0; j < 2; j++ {
				x := right - j
				up := ((right + 1) & 2) == 0
				y := vert
				if up {
					y = size - 1 - vert
				}
				if !isFn(x, y) && i < len(bits) {
					dark[y][x] = bits[i]
					i++
				}
			}
		}
	}

	best, bestScore := 0, 0
	for mask := 0; mask < 8; mask++ {
		for y := range dark {
			for x := range dark {
				if !fn[y][x] {
					dark[y][x] = dark[y][x] != qrMasked(mask, x, y)
				}
			}
		}
		qrFormat(set, mask, size)
		score := qrPenalty(dark)
		if best == 0 || score < bestScore {
			best, bestScore = mask, score
		}
	}
	for y := range dark {
		for x := range dark {
			if !fn[y][x] {
				dark[y][x] = dark[y][x] != qrMasked(best, x, y)
			}
		}
	}
	qrFormat(set, best, size)
	return dark
}

// qrAlignment returns the alignment pattern centers of a version.
func qrAlignment(version, size int) []int {
	if version == 1 {
		return nil
	}
	num := version/7 + 2
	step := (version*4 + num*2 + 1) / (num*2 - 2) * 2
	positions := []int{6}
	for i := num - 2; i >= 0; i-- {
		positions = append(positions, size-7-step*i)
	}
	return positions
}

// qrMasked reports whether mask m flips module (x, y).
func qrMasked(m, x, y int) bool {
	switch m {
	case 0:
		return (x+y)%2 == 0
	case 1:
		return y%2 == 0
	case 2:
		return x%3 == 0
	case 3:
		return (x+y)%3 == 0
	case 4:
		return (x/3+y/2)%2 == 0
	case 5:
		return x*y%2+x*y%3 == 0
	case 6:
		return (x*y%2+x*y%3)%2 == 0
	default:
		return ((x+y)%2+x*y%3)%2 == 0
	}
}

// qrFormat writes the 15 format bits for error level L and mask m.
func qrFormat(set func(x, y int, d bool), m, size int) {
	data := 1<<3 | m // level L is 01
	rem := data
	for i := 0; i < 10; i++ {
		rem = (rem << 1) ^ ((rem >> 9) * 0x537)
	}
	bits := ((data << 10) | rem) ^ 0x5412
	at := func(i int) bool { return bits&(1<<uint(i)) != 0 }
	for i := 0; i <= 5; i++ {
		set(8, i, at(i))
	}
	set(8, 7, at(6))
	set(8, 8, at(7))
	set(7, 8, at(8))
	for i := 9; i < 15; i++ {
		set(14-i, 8, at(i))
	}
	for i := 0; i < 8; i++ {
		set(size-1-i, 8, at(i))
	}
	for i := 8; i < 15; i++ {
		set(8, size-15+i, at(i))
	}
	set(8, size-8, true) // the fixed dark module
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// qrPenalty scores a masked matrix; the lowest score is the best mask.
func qrPenalty(m [][]bool) int {
	n := len(m)
	score := 0
	pattern := [11]bool{true, false, true, true, true, false, true, false, false, false, false}
	runScore := func(row []bool) int {
		s := 0
		for i, v := range row {
			run := 1
			for i+run < len(row) && row[i+run] == v {
				run++
			}
			if run >= 5 {
				s += 3 + run - 5
			}
			for j := 0; j+11 <= len(row); j++ {
				match, rev := true, true
				for k := 0; k < 11; k++ {
					if row[j+k] != pattern[k] {
						match = false
					}
					if row[j+k] != pattern[10-k] {
						rev = false
					}
				}
				if match || rev {
					s += 40
				}
			}
		}
		return s
	}
	for y := 0; y < n; y++ {
		score += runScore(m[y])
	}
	for x := 0; x < n; x++ {
		col := make([]bool, n)
		for y := 0; y < n; y++ {
			col[y] = m[y][x]
		}
		score += runScore(col)
	}
	for y := 0; y < n-1; y++ {
		for x := 0; x < n-1; x++ {
			if m[y][x] == m[y][x+1] && m[y][x] == m[y+1][x] && m[y][x] == m[y+1][x+1] {
				score += 3
			}
		}
	}
	dark := 0
	for _, row := range m {
		for _, v := range row {
			if v {
				dark++
			}
		}
	}
	pct := dark * 100 / (n * n)
	k := pct - 50
	if k < 0 {
		k = -k
	}
	score += k / 5 * 10
	return score
}
