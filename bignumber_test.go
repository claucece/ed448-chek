package ed448

import (
	. "gopkg.in/check.v1"
)

func (s *Ed448Suite) Test_Equals(c *C) {
	n, _ := deserialize(serialized{
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	})

	c.Assert(n.equals(n), Equals, true)

	x := mustDeserialize(serialized{0x01, 0x01})
	y := mustDeserialize(serialized{0x01, 0x02})

	c.Assert(x.equals(y), Equals, false)
}

func (s *Ed448Suite) Test_DecafEquals(c *C) {
	x, _ := deserialize(serialized{
		0xf5, 0x81, 0x74, 0xd5, 0x7a, 0x33, 0x72,
		0x36, 0x3c, 0x0d, 0x9f, 0xcf, 0xaa, 0x3d,
		0xc1, 0x8b, 0x1e, 0xff, 0x7e, 0x89, 0xbf,
		0x76, 0x78, 0x63, 0x65, 0x80, 0xd1, 0x7d,
		0xd8, 0x4a, 0x87, 0x3b, 0x14, 0xb9, 0xc0,
		0xe1, 0x68, 0x0b, 0xbd, 0xc8, 0x76, 0x47,
		0xf3, 0xc3, 0x82, 0x90, 0x2d, 0x2f, 0x58,
		0xd2, 0x75, 0x4b, 0x39, 0xbc, 0xa8, 0x74,
	})

	y, _ := deserialize(serialized{
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	})

	c.Assert(x.decafEq(x), Equals, decafTrue)
	c.Assert(x.decafEq(y), Equals, decafFalse)
}

func (s *Ed448Suite) Test_SetBytes(c *C) {
	bs := []byte{0x0e}
	n := new(bigNumber).setBytes(bs)

	c.Assert(n, IsNil)

	bs = bytesFromHex(
		"e6f5b8ae49cef779e577dc29824eff453f1c4106030088115ea49b4ee8" +
			"4a7b7cdfe06e0d622fc55c7c559ab1f6c3ea3257c07979809026de")
	n = new(bigNumber).setBytes(bs)
	exp := &bigNumber{
		0x09026de, 0xc079798,
		0x3ea3257, 0x9ab1f6c,
		0x55c7c55, 0x0d622fc,
		0xcdfe06e, 0xe84a7b7,
		0xea49b4e, 0x0088115,
		0xc410603, 0xff453f1,
		0xc29824e, 0x79e577d,
		0xe49cef7, 0xe6f5b8a,
	}

	c.Assert(n, DeepEquals, exp)
}

func (s *Ed448Suite) Test_ZeroMask(c *C) {

	c.Assert(bigZero.zeroMask(), Equals, word(lmask))
	c.Assert(bigOne.zeroMask(), Equals, word(0x00))
}

func (s *Ed448Suite) Test_IsZero(c *C) {
	n := mustDeserialize(serialized{0x01})
	c.Assert(n.isZero(), Equals, false)

	n = mustDeserialize(serialized{0x00})
	c.Assert(n.isZero(), Equals, true)
}

func (s *Ed448Suite) Test_HighBit(c *C) {
	n := &bigNumber{0xdeadbeef}
	h := highBit(n)

	c.Assert(h, Equals, word(0x00))
}

func (s *Ed448Suite) Test_LowBit(c *C) {
	n := &bigNumber{0x01}
	l := lowBit(n)

	c.Assert(l, Equals, word(0xffffffff))

	n = &bigNumber{0x00}
	l = lowBit(n)

	c.Assert(l, Equals, word(0x00))
}

func (s *Ed448Suite) Test_Add(c *C) {
	x := mustDeserialize(serialized{0x57})
	y := mustDeserialize(serialized{0x83})
	exp := mustDeserialize(serialized{0xda})

	c.Assert(new(bigNumber).add(x, y), DeepEquals, exp)

	// radix
	x = mustDeserialize(serialized{
		0xff, 0xff, 0xff, 0xf0,
	})
	y = mustDeserialize(serialized{0x01})
	exp = mustDeserialize(serialized{
		0x00, 0x00, 0x00, 0xf1,
	})

	c.Assert(new(bigNumber).add(x, y), DeepEquals, exp)
}

func (s *Ed448Suite) Test_AddWord(c *C) {
	x := word(0x01)
	exp := mustDeserialize(serialized{0x01})

	c.Assert(new(bigNumber).addW(x), DeepEquals, exp)
}

func (s *Ed448Suite) Test_Subtraction(c *C) {
	x := mustDeserialize(serialized{0xda})
	y := mustDeserialize(serialized{0x83})
	exp := mustDeserialize(serialized{0x57})

	c.Assert(new(bigNumber).sub(x, y).strongReduce(), DeepEquals, exp)

	x = mustDeserialize(serialized{
		0x00, 0x00, 0x00, 0xf1,
	})
	y = mustDeserialize(serialized{0x01})
	exp = mustDeserialize(serialized{
		0xff, 0xff, 0xff, 0xf0,
	})

	c.Assert(new(bigNumber).sub(x, y).strongReduce(), DeepEquals, exp)
}

func (s *Ed448Suite) Test_SubWord(c *C) {
	x := mustDeserialize(serialized{0x01})
	y := word(0x01)
	exp := mustDeserialize(serialized{0x00})

	c.Assert(x.subW(y), DeepEquals, exp)
}

func (s *Ed448Suite) Test_SubWithDifferentBias(c *C) {
	x := mustDeserialize(serialized{0xff})
	y := mustDeserialize(serialized{0xff})
	exp := &bigNumber{
		0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff,
		0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff,
		0xffffffe, 0xfffffff, 0xfffffff, 0xfffffff,
		0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff,
	}

	c.Assert(new(bigNumber).subXBias(x, y, word(2)), DeepEquals, exp)
}

func (s *Ed448Suite) Test_Multiplication(c *C) {
	x := mustDeserialize(serialized{0x02})
	y := mustDeserialize(serialized{0x03})
	exp := mustDeserialize(serialized{0x06})

	c.Assert(new(bigNumber).mulCopy(x, y), DeepEquals, exp)

	x = mustDeserialize(serialized{0x10})
	y = mustDeserialize(serialized{0x0e})
	exp = mustDeserialize(serialized{0xe0})

	c.Assert(new(bigNumber).mul(x, y), DeepEquals, exp)
}

func (s *Ed448Suite) Test_MulWithDConstant(c *C) {
	x := mustDeserialize(serialized{0x02})
	exp := &bigNumber{
		0xffecead, 0xfffffff, 0xfffffff, 0xfffffff,
		0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff,
		0xffffffe, 0xfffffff, 0xfffffff, 0xfffffff,
		0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff,
	}

	c.Assert(new(bigNumber).mulWSignedCurveConstant(x, edwardsD), DeepEquals, exp)
}

func (s *Ed448Suite) Test_SquareN(c *C) {
	gx := mustDeserialize(serialized{
		0x9f, 0x93, 0xed, 0x0a, 0x84, 0xde, 0xf0,
		0xc7, 0xa0, 0x4b, 0x3f, 0x03, 0x70, 0xc1,
		0x96, 0x3d, 0xc6, 0x94, 0x2d, 0x93, 0xf3,
		0xaa, 0x7e, 0x14, 0x96, 0xfa, 0xec, 0x9c,
		0x70, 0xd0, 0x59, 0x3c, 0x5c, 0x06, 0x5f,
		0x24, 0x33, 0xf7, 0xad, 0x26, 0x6a, 0x3a,
		0x45, 0x98, 0x60, 0xf4, 0xaf, 0x4f, 0x1b,
		0xff, 0x92, 0x26, 0xea, 0xa0, 0x7e, 0x29,
	})

	exp := gx.copy()
	for i := 0; i < 5; i++ {
		exp = new(bigNumber).square(exp)
	}

	n := new(bigNumber).squareN(gx, 5)

	c.Assert(n.equals(exp), Equals, true)

	exp = gx.copy()
	for i := 0; i < 6; i++ {
		exp = new(bigNumber).square(exp)
	}

	n = n.squareN(gx, 6)

	c.Assert(n.equals(exp), Equals, true)
}

func (s *Ed448Suite) Test_InverseSquareRoot(c *C) {
	x := mustDeserialize(serialized{
		0x9f, 0x93, 0xed, 0x0a, 0x84, 0xde, 0xf0,
		0xc7, 0xa0, 0x4b, 0x3f, 0x03, 0x70, 0xc1,
		0x96, 0x3d, 0xc6, 0x94, 0x2d, 0x93, 0xf3,
		0xaa, 0x7e, 0x14, 0x96, 0xfa, 0xec, 0x9c,
		0x70, 0xd0, 0x59, 0x3c, 0x5c, 0x06, 0x5f,
		0x24, 0x33, 0xf7, 0xad, 0x26, 0x6a, 0x3a,
		0x45, 0x98, 0x60, 0xf4, 0xaf, 0x4f, 0x1b,
		0xff, 0x92, 0x26, 0xea, 0xa0, 0x7e, 0x29,
	})

	x.isr(x)

	n := bytesFromHex("04027d13a34bbe052fdf4247b02a4a3406268203a09076e56" +
		"dee9dc2b699c4abc66f2832a677dfd0bf7e70ee72f01db170839717d1c64f02")
	exp := new(bigNumber).setBytes(n)

	c.Assert(x.equals(exp), Equals, true)
}

func (s *Ed448Suite) Test_Invert(c *C) {
	n := &bigNumber{}
	x := &bigNumber{
		0x4516644, 0x1430f14, 0x72318d2, 0xb1c2096,
		0x32e3855, 0x1c1105f, 0xbf1556f, 0xbb9f535,
		0xe3d45c0, 0xe954acd, 0xcba31b2, 0x5b931f9,
		0x0920cdd, 0x64f93a9, 0x2d91281, 0x674f3d0,
	}

	exp := &bigNumber{
		0x3509cef, 0x92c009c, 0x4116af4, 0x4bd5cae,
		0x5c60b66, 0x1da9fbd, 0xe925340, 0x2fffa3f,
		0xdd725b2, 0xc2ae8ae, 0xf4808a9, 0x40ed04c,
		0x864dc36, 0x6821f90, 0x8099dc5, 0xcf9ca3d,
	}
	n = invert(x)

	c.Assert(n, DeepEquals, exp)
}

func (s *Ed448Suite) Test_Negate(c *C) {
	bs := bytesFromHex("e6f5b8ae49cef779e577dc29824eff453f1c4106030088115" +
		"ea49b4ee84a7b7cdfe06e0d622fc55c7c559ab1f6c3ea3257c07979809026de")

	n := new(bigNumber).setBytes(bs)
	out := new(bigNumber).neg(n)

	bs = bytesFromHex("190a4751b63108861a8823d67db100bac0e3bef9fcff77eea1" +
		"5b64b017b58483201f91f29dd03aa383aa654e093c15cda83f86867f6fd921")
	exp := new(bigNumber).setBytes(bs)

	c.Assert(out, DeepEquals, exp)
}

func (s *Ed448Suite) Test_ConditionalNegateNumber(c *C) {
	bs := bytesFromHex("e6f5b8ae49cef779e577dc29824eff453f1c4106030088115" +
		"ea49b4ee84a7b7cdfe06e0d622fc55c7c559ab1f6c3ea3257c07979809026de")
	n := new(bigNumber).setBytes(bs)

	bs = bytesFromHex("190a4751b63108861a8823d67db100bac0e3bef9fcff77eea1" +
		"5b64b017b58483201f91f29dd03aa383aa654e093c15cda83f86867f6fd921")
	exp := new(bigNumber).setBytes(bs)

	c.Assert(n.copy().conditionalNegate(lmask), DeepEquals, exp)
	c.Assert(n.copy().conditionalNegate(0), DeepEquals, n)
}

func (s *Ed448Suite) Test_DecafConditionalNegateNumber(c *C) {
	bs := bytesFromHex("e6f5b8ae49cef779e577dc29824eff453f1c4106030088115" +
		"ea49b4ee84a7b7cdfe06e0d622fc55c7c559ab1f6c3ea3257c07979809026de")
	x := new(bigNumber).setBytes(bs)

	bs = bytesFromHex("190a4751b63108861a8823d67db100bac0e3bef9fcff77eea1" +
		"5b64b017b58483201f91f29dd03aa383aa654e093c15cda83f86867f6fd921")
	exp := new(bigNumber).setBytes(bs)

	x.decafCondNegate(lmask)

	c.Assert(x, DeepEquals, exp)

	x = &bigNumber{}

	n, _ := deserialize(serialized{
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	})

	x.decafCondNegate(lmask)

	// 0 mod p = n1
	c.Assert(x, DeepEquals, n)
}

func (s *Ed448Suite) Test_ConditionalSwap(c *C) {
	bs := bytesFromHex("e6f5b8ae49cef779e577dc29824eff453f1c4106030088115" +
		"ea49b4ee84a7b7cdfe06e0d622fc55c7c559ab1f6c3ea3257c07979809026de")
	x := new(bigNumber).setBytes(bs)

	bs = bytesFromHex("190a4751b63108861a8823d67db100bac0e3bef9fcff77eea1" +
		"5b64b017b58483201f91f29dd03aa383aa654e093c15cda83f86867f6fd921")
	y := new(bigNumber).setBytes(bs)

	a := x.copy()
	b := y.copy()
	a.conditionalSwap(b, lmask)

	c.Assert(a, DeepEquals, y)
	c.Assert(b, DeepEquals, x)

	a.conditionalSwap(b, 0)
	c.Assert(a, DeepEquals, y)
	c.Assert(b, DeepEquals, x)
}

func (s *Ed448Suite) Test_ConditionalSelect(c *C) {
	bs := bytesFromHex("e6f5b8ae49cef779e577dc29824eff453f1c4106030088115" +
		"ea49b4ee84a7b7cdfe06e0d622fc55c7c559ab1f6c3ea3257c07979809026de")
	x := new(bigNumber).setBytes(bs)

	bs = bytesFromHex("190a4751b63108861a8823d67db100bac0e3bef9fcff77eea1" +
		"5b64b017b58483201f91f29dd03aa383aa654e093c15cda83f86867f6fd921")
	y := new(bigNumber).setBytes(bs)

	c.Assert(constantTimeSelect(x, y, lmask), DeepEquals, x)
	c.Assert(constantTimeSelect(x, y, 0), DeepEquals, y)

}

func (s *Ed448Suite) Test_DecafConstTimeSel(c *C) {
	x := &bigNumber{
		0x08db85c2, 0x0fd2361e, 0x0ce2105d, 0x06a17729,
		0x0e3ca84d, 0x0a137aa5, 0x0985ee61, 0x05a26d64,
		0x0734c5f3, 0x0da853af, 0x01d955b7, 0x03160ecd,
		0x0a59046d, 0x0c32cf71, 0x98dce72d, 0x00007fff,
	}

	y := &bigNumber{
		0x07247a3d, 0x002dc9e1, 0x031defa2, 0x095e88d6,
		0x01c357b2, 0x05ec855a, 0x067a119e, 0x0a5d929b,
		0x08cb3a0b, 0x0257ac50, 0x0e26aa48, 0x0ce9f132,
		0x05a6fb92, 0x03cd308e, 0x072318d2, 0x0fff8007,
	}

	exp := &bigNumber{
		0x07247a3d, 0x002dc9e1, 0x031defa2, 0x095e88d6,
		0x01c357b2, 0x05ec855a, 0x067a119e, 0x0a5d929b,
		0x08cb3a0b, 0x0257ac50, 0x0e26aa48, 0x0ce9f132,
		0x05a6fb92, 0x03cd308e, 0x072318d2, 0x0fff8007,
	}

	n := &bigNumber{}
	n.decafConstTimeSel(x, y, word(lmask))

	c.Assert(n, DeepEquals, exp)

	n.decafConstTimeSel(x, y, word(0))

	c.Assert(n, DeepEquals, n)
}

func (s *Ed448Suite) Test_StrongReduce(c *C) {
	n, _ := deserialize(serialized{
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	})

	//p = p mod p = 0
	n.strongReduce()

	c.Assert(n, DeepEquals, bigZero)

	n = mustDeserialize(serialized{
		0xf5, 0x81, 0x74, 0xd5, 0x7a, 0x33, 0x72,
		0x36, 0x3c, 0x0d, 0x9f, 0xcf, 0xaa, 0x3d,
		0xc1, 0x8b, 0x1e, 0xff, 0x7e, 0x89, 0xbf,
		0x76, 0x78, 0x63, 0x65, 0x80, 0xd1, 0x7d,
		0xd8, 0x4a, 0x87, 0x3b, 0x14, 0xb9, 0xc0,
		0xe1, 0x68, 0x0b, 0xbd, 0xc8, 0x76, 0x47,
		0xf3, 0xc3, 0x82, 0x90, 0x2d, 0x2f, 0x58,
		0xd2, 0x75, 0x4b, 0x39, 0xbc, 0xa8, 0x74,
	})

	n.strongReduce()

	c.Assert(n, DeepEquals, mustDeserialize(serialized{
		0xf5, 0x81, 0x74, 0xd5, 0x7a, 0x33, 0x72,
		0x36, 0x3c, 0x0d, 0x9f, 0xcf, 0xaa, 0x3d,
		0xc1, 0x8b, 0x1e, 0xff, 0x7e, 0x89, 0xbf,
		0x76, 0x78, 0x63, 0x65, 0x80, 0xd1, 0x7d,
		0xd8, 0x4a, 0x87, 0x3b, 0x14, 0xb9, 0xc0,
		0xe1, 0x68, 0x0b, 0xbd, 0xc8, 0x76, 0x47,
		0xf3, 0xc3, 0x82, 0x90, 0x2d, 0x2f, 0x58,
		0xd2, 0x75, 0x4b, 0x39, 0xbc, 0xa8, 0x74,
	}))
}

func (s *Ed448Suite) Test_Serialize(c *C) {
	invalid := [55]byte{}

	c.Assert(func() { serialize(invalid[:], bigOne) }, Panics, "Failed to serialize")

	dst := [fieldBytes]byte{}

	serialize(dst[:], bigOne)
	c.Assert(dst, DeepEquals, [fieldBytes]byte{1})

	n := &bigNumber{
		0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff,
		0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff,
		0xffffffe, 0xfffffff, 0xfffffff, 0xfffffff,
		0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff,
	}

	serialize(dst[:], n)

	//0 because serialize reduces mod p
	c.Assert(dst, DeepEquals, [fieldBytes]byte{})

	exp := [fieldBytes]byte{
		0xf5, 0x81, 0x74, 0xd5, 0x7a, 0x33, 0x72,
		0x36, 0x3c, 0x0d, 0x9f, 0xcf, 0xaa, 0x3d,
		0xc1, 0x8b, 0x1e, 0xff, 0x7e, 0x89, 0xbf,
		0x76, 0x78, 0x63, 0x65, 0x80, 0xd1, 0x7d,
		0xd8, 0x4a, 0x87, 0x3b, 0x14, 0xb9, 0xc0,
		0xe1, 0x68, 0x0b, 0xbd, 0xc8, 0x76, 0x47,
		0xf3, 0xc3, 0x82, 0x90, 0x2d, 0x2f, 0x58,
		0xd2, 0x75, 0x4b, 0x39, 0xbc, 0xa8, 0x74,
	}

	x := &bigNumber{
		0x57481f5, 0x72337ad, 0xf0d3c36, 0x3daacf9,
		0xf1e8bc1, 0xbf897ef, 0x5637876, 0x7dd1806,
		0xb874ad8, 0xc0b9143, 0xd0b68e1, 0x4776c8b,
		0x082c3f3, 0x582f2d9, 0x94b75d2, 0x74a8bc3,
	}

	serialize(dst[:], x)

	c.Assert(dst, DeepEquals, exp)

}

func (s *Ed448Suite) Test_Deserialize(c *C) {
	ser := serialized{0x01}
	n, ok := deserialize(ser)

	c.Assert(n, DeepEquals, bigOne)
	c.Assert(ok, Equals, true)

	ser = serialized{
		0xf5, 0x81, 0x74, 0xd5, 0x7a, 0x33, 0x72,
		0x36, 0x3c, 0x0d, 0x9f, 0xcf, 0xaa, 0x3d,
		0xc1, 0x8b, 0x1e, 0xff, 0x7e, 0x89, 0xbf,
		0x76, 0x78, 0x63, 0x65, 0x80, 0xd1, 0x7d,
		0xd8, 0x4a, 0x87, 0x3b, 0x14, 0xb9, 0xc0,
		0xe1, 0x68, 0x0b, 0xbd, 0xc8, 0x76, 0x47,
		0xf3, 0xc3, 0x82, 0x90, 0x2d, 0x2f, 0x58,
		0xd2, 0x75, 0x4b, 0x39, 0xbc, 0xa8, 0x74,
	}

	n, ok = deserialize(ser)

	c.Assert(n, DeepEquals, &bigNumber{
		0x57481f5, 0x72337ad, 0xf0d3c36, 0x3daacf9,
		0xf1e8bc1, 0xbf897ef, 0x5637876, 0x7dd1806,
		0xb874ad8, 0xc0b9143, 0xd0b68e1, 0x4776c8b,
		0x082c3f3, 0x582f2d9, 0x94b75d2, 0x74a8bc3,
	})
	c.Assert(ok, Equals, true)

	ser = serialized{
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	}

	n, ok = deserialize(ser)
	c.Assert(n, DeepEquals, &bigNumber{
		0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff,
		0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff,
		0xffffffe, 0xfffffff, 0xfffffff, 0xfffffff,
		0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff,
	})
	c.Assert(ok, Equals, false)
}

func (s *Ed448Suite) Test_MustDeserialize(c *C) {
	n := serialized{
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	}

	c.Assert(func() { mustDeserialize(n) }, Panics, "Failed to deserialize")
}

func (s *Ed448Suite) Test_DsaLikeSerialize(c *C) {
	n := &bigNumber{
		0x04ced9a5, 0x0a48906a, 0x09f09413, 0x0e0fe326,
		0x007b11db, 0x0687875e, 0x0044482c, 0x0c6e93b7,
		0x006cde64, 0x0a3a5d6e, 0x0938e74e, 0x0930cb3d,
		0x088d7575, 0x006de50e, 0x0075b92c, 0x085247d5,
	}

	out := make([]byte, 57)
	exp := []byte{
		0xa5, 0xd9, 0xce, 0xa4, 0x06, 0x89, 0xa4, 0x13,
		0x94, 0xf0, 0x69, 0x32, 0xfe, 0xe0, 0xdb, 0x11,
		0x7b, 0xe0, 0x75, 0x78, 0x68, 0x2c, 0x48, 0x44,
		0x70, 0x3b, 0xe9, 0xc6, 0x64, 0xde, 0x6c, 0xe0,
		0xd6, 0xa5, 0xa3, 0x4e, 0xe7, 0x38, 0xd9, 0xb3,
		0x0c, 0x93, 0x75, 0x75, 0x8d, 0xe8, 0x50, 0xde,
		0x06, 0x2c, 0xb9, 0x75, 0x50, 0x7d, 0x24, 0x85,
		0x0,
	}

	dsaLikeSerialize(out, n)

	c.Assert(out, DeepEquals, exp)
}

func (s *Ed448Suite) Test_DsaLikeDeserialize(c *C) {
	ser := []byte{
		0xa5, 0xd9, 0xce, 0xa4, 0x06, 0x89, 0xa4, 0x13,
		0x94, 0xf0, 0x69, 0x32, 0xfe, 0xe0, 0xdb, 0x11,
		0x7b, 0xe0, 0x75, 0x78, 0x68, 0x2c, 0x48, 0x44,
		0x70, 0x3b, 0xe9, 0xc6, 0x64, 0xde, 0x6c, 0xe0,
		0xd6, 0xa5, 0xa3, 0x4e, 0xe7, 0x38, 0xd9, 0xb3,
		0x0c, 0x93, 0x75, 0x75, 0x8d, 0xe8, 0x50, 0xde,
		0x06, 0x2c, 0xb9, 0x75, 0x50, 0x7d, 0x24, 0x85, 0x0,
	}

	exp := &bigNumber{
		0x04ced9a5, 0x0a48906a, 0x09f09413, 0x0e0fe326,
		0x007b11db, 0x0687875e, 0x0044482c, 0x0c6e93b7,
		0x006cde64, 0x0a3a5d6e, 0x0938e74e, 0x0930cb3d,
		0x088d7575, 0x006de50e, 0x0075b92c, 0x085247d5,
	}

	dst := &bigNumber{}
	ok := dsaLikeDeserialize(dst, ser)

	c.Assert(dst, DeepEquals, exp)
	c.Assert(ok, Equals, decafTrue)
}

func (s *Ed448Suite) Test_ReturnsTheStringRepresentation(c *C) {
	n := &bigNumber{}
	str := n.String()

	exp := "[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, " +
		"0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, " +
		"0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, " +
		"0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, " +
		"0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, " +
		"0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, " +
		"0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0" +
		"}"

	c.Assert(str, DeepEquals, exp)
}

func (s *Ed448Suite) Test_ReturnsLimbs(c *C) {
	n := bigOne
	str := n.limbs()

	exp := []word{
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	c.Assert(str, DeepEquals, exp)
}
