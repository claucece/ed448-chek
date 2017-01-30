package ed448

// ModQ produces a byte array mod Q (prime order)
func ModQ(serial []byte) []byte {
	words := scalar32{}
	words.deserializeModQ(serial)
	out := make([]byte, fieldBytes)
	words.serialize(out)
	return out
}

// PointMul multiplies a point x with a scalar y
// PointMul automatically reduces the output by P
func PointMul(x [fieldBytes]byte, y [fieldBytes]byte) (out []byte) {
	desX, okX := deserialize(x)
	desY, okY := deserialize(y)
	if !(okX && okY) {
		return nil
	}
	desZ := &bigNumber{}
	desZ.mulCopy(desX, desY)
	out = make([]byte, fieldBytes)
	serialize(out, desZ)
	return out
}

// PointAddition adds two Ed448 points
// Inputs should never be >= prime P. If they are, PointAddition returns nil.
// PointAddition automatically reduces the output by P
func PointAddition(x [fieldBytes]byte, y [fieldBytes]byte) (out []byte) {
	desX, okX := deserialize(x)
	desY, okY := deserialize(y)
	if !(okX && okY) {
		return nil
	}
	desZ := &bigNumber{}
	desZ.add(desX, desY)
	out = make([]byte, fieldBytes)
	serialize(out, desZ)
	return out
}
