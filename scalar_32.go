package ed448

type Scalar [scalarWords]uint32

//See Goldilocks spec, "Public and private keys" section.
//This is equivalent to DESERMODq()
func (dst *Scalar) deserializeModQ(serial []byte) {
	barrettDeserializeAndReduce(dst[:], serial, &curvePrimeOrder)
	return
}

// Serializes an array of words into an array of bytes (little-endian)
func (src Scalar) serialize(dst []byte) {
	wordBytes := wordBits / 8

	for i := 0; i*wordBytes < len(dst); i++ {
		for j := 0; j < wordBytes; j++ {
			b := src[i] >> uint(8*j)
			dst[wordBytes*i+j] = byte(b)
		}
	}
}

func (dst *Scalar) scalarAdd(a, b *Scalar) {
	out := &Scalar{}
	var chain uint64

	for i := uintZero; i < scalarWords; i++ {
		chain += uint64(a[i]) + uint64(b[i])
		out[i] = uint32(chain)
		chain >>= wordBits
	}
	out.scalarSubExtra(out, scalarQ, uint32(chain))
	copy(dst[:], out[:])
}

func (dst *Scalar) scalarSubExtra(minuend *Scalar, subtrahend *Scalar, carry uint32) {
	out := &Scalar{}
	var chain int64

	for i := uintZero; i < scalarWords; i++ {
		chain += int64(minuend[i]) - int64(subtrahend[i])
		out[i] = uint32(chain)
		chain >>= wordBits
	}

	borrow := chain + int64(carry)
	chain = 0

	for i := uintZero; i < scalarWords; i++ {
		chain += int64(out[i]) + (int64(scalarQ[i]) & borrow)
		out[i] = uint32(chain)
		chain >>= wordBits
	}
	copy(dst[:], out[:])
}

func (dst *Scalar) scalarHalve(a, b *Scalar) {
	out := &Scalar{}
	mask := -(a[0] & 1)
	var chain uint64
	var i uint

	for i = 0; i < scalarWords; i++ {
		chain += uint64(a[i]) + uint64(b[i]&mask)
		out[i] = uint32(chain)
		chain >>= wordBits
	}
	for i = 0; i < scalarWords-1; i++ {
		out[i] = out[i]>>1 | out[i+1]<<(wordBits-1)
	}

	out[i] = out[i]>>1 | uint32(chain<<(wordBits-1))

	copy(dst[:], out[:])
}

func (dst *Scalar) montgomeryMultiply(x, y *Scalar) {
	out := &Scalar{}
	carry := uint32(0)

	for i := 0; i < scalarWords; i++ {
		chain := uint64(0)
		for j := 0; j < scalarWords; j++ {
			chain += uint64(x[i])*uint64(y[j]) + uint64(out[j])
			out[j] = uint32(chain)
			chain >>= wordBits
		}
		saved := uint32(chain)
		multiplicand := out[0] * montgomeryFactor
		chain = 0
		for j := 0; j < scalarWords; j++ {
			chain += uint64(multiplicand)*uint64(scalarQ[j]) + uint64(out[j])
			if j > 0 {
				out[j-1] = uint32(chain)
			}
			chain >>= wordBits
		}
		chain += uint64(saved) + uint64(carry)
		out[scalarWords-1] = uint32(chain)
		carry = uint32(chain >> wordBits)
	}
	out.scalarSubExtra(out, scalarQ, carry)
	copy(dst[:], out[:])
}

func (dst *Scalar) Mul(x, y ScalarI) {
	dst.montgomeryMultiply(x.(*Scalar), y.(*Scalar))
	dst.montgomeryMultiply(dst, scalarR2)
}

func (dst *Scalar) Sub(x, y ScalarI) {
	noExtra := uint32(0)
	dst.scalarSubExtra(x.(*Scalar), y.(*Scalar), noExtra)
}
