package ed448

import (
	"bytes"

	. "gopkg.in/check.v1"
)

func (s *Ed448Suite) Test_DecafDerivePrivate(c *C) {
	sym := [symKeyBytes]byte{
		0xd1, 0x50, 0x0e, 0xf4, 0x1c, 0xa7, 0xbe, 0xe3,
		0x7d, 0x2d, 0x95, 0x14, 0x9b, 0x75, 0xeb, 0xab,
		0xb0, 0x66, 0xc9, 0xe3, 0x66, 0x32, 0xd2, 0x12,
		0x23, 0x4f, 0xf2, 0x4a, 0x96, 0x94, 0x52, 0x3e,
	}

	expPub := []byte{
		0xb2, 0x95, 0x4a, 0x76, 0x1f, 0x3d, 0x98, 0x03,
		0xaa, 0x2b, 0xbc, 0x8c, 0x98, 0x08, 0x0e, 0x52,
		0x7d, 0x56, 0xb7, 0xb6, 0x14, 0x08, 0x93, 0x0b,
		0xc9, 0x0b, 0xf7, 0x89, 0x80, 0x4f, 0x4f, 0x2c,
		0x8c, 0x65, 0x37, 0xdd, 0xb9, 0xb0, 0xdf, 0x62,
		0xe1, 0xd6, 0x00, 0x9c, 0xee, 0x58, 0x85, 0x21,
		0x3a, 0x56, 0x7d, 0xb1, 0x16, 0x70, 0x0a, 0x57,
	}

	expPriv := []byte{
		0x78, 0x46, 0x08, 0x83, 0x64, 0xff, 0x17, 0x87,
		0x13, 0x3d, 0xa5, 0x9e, 0x71, 0x52, 0xdf, 0xe3,
		0xa5, 0x1b, 0xb6, 0xc7, 0x50, 0xc2, 0xbb, 0xfd,
		0x81, 0x49, 0x4e, 0x7e, 0x23, 0x44, 0x19, 0x3f,
		0x0e, 0xdd, 0x35, 0x10, 0x88, 0xf1, 0xc1, 0x9a,
		0xd1, 0x03, 0xbf, 0xf6, 0x99, 0x23, 0xf0, 0x91,
		0x05, 0xe3, 0x66, 0x30, 0xf8, 0xf0, 0x8f, 0x14,
	}

	pk, _ := decafCurve.decafDerivePrivateKey(sym)

	expSym := []byte{
		0xd1, 0x50, 0x0e, 0xf4, 0x1c, 0xa7, 0xbe, 0xe3,
		0x7d, 0x2d, 0x95, 0x14, 0x9b, 0x75, 0xeb, 0xab,
		0xb0, 0x66, 0xc9, 0xe3, 0x66, 0x32, 0xd2, 0x12,
		0x23, 0x4f, 0xf2, 0x4a, 0x96, 0x94, 0x52, 0x3e,
	}

	c.Assert(pk.symKey(), DeepEquals, expSym)
	c.Assert(pk.secretKey(), DeepEquals, expPriv)
	c.Assert(pk.publicKey(), DeepEquals, expPub)
}

func (s *Ed448Suite) Test_DecafDerivePrivateWithZeroSymKey(c *C) {
	sym := [symKeyBytes]byte{
		0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	expPub := []byte{
		0x50, 0xee, 0x17, 0xd7, 0x4a, 0xbd, 0x73, 0x1e,
		0x76, 0x47, 0x3d, 0x52, 0x30, 0xd3, 0x7a, 0x35,
		0xf6, 0x35, 0x9c, 0x0e, 0x19, 0xd2, 0x80, 0x88,
		0x74, 0x4e, 0x82, 0x5e, 0x9b, 0x4e, 0xbd, 0x83,
		0x48, 0x6d, 0xec, 0x5b, 0x8f, 0x17, 0xa8, 0x87,
		0xeb, 0x39, 0x93, 0x2f, 0x43, 0x57, 0x90, 0xef,
		0xa9, 0xcd, 0x95, 0xf5, 0xea, 0xdc, 0xb9, 0x5a,
	}

	expPriv := []byte{
		0xfd, 0x7e, 0xbe, 0x65, 0xe4, 0xd1, 0xa2, 0x5c,
		0x61, 0x89, 0x36, 0x66, 0xc1, 0x5a, 0x52, 0x24,
		0x2d, 0x9c, 0xd0, 0x78, 0x4a, 0x2d, 0x56, 0x0e,
		0xb2, 0xf4, 0x4c, 0xfa, 0x73, 0x5c, 0x49, 0xed,
		0x2b, 0xd0, 0x72, 0x0d, 0xf4, 0x2c, 0x92, 0x13,
		0xf3, 0xb4, 0x41, 0x7d, 0x7f, 0x13, 0x13, 0x1d,
		0xef, 0x94, 0x67, 0x7b, 0xd3, 0x1e, 0xb5, 0x0e,
	}

	pk, err := decafCurve.decafDerivePrivateKey(sym)

	c.Assert(pk.secretKey(), DeepEquals, expPriv)
	c.Assert(pk.publicKey(), DeepEquals, expPub)
	c.Assert(err, IsNil)
}

func (s *Ed448Suite) Test_DecafGenerateKeys(c *C) {
	buffer := make([]byte, symKeyBytes)
	buffer[0] = 0x10
	r := bytes.NewReader(buffer[:])

	privKey, err := decafCurve.decafGenerateKeys(r)

	expSym := make([]byte, symKeyBytes)
	expSym[0] = 0x10

	expPub := []byte{
		0x50, 0xee, 0x17, 0xd7, 0x4a, 0xbd, 0x73, 0x1e,
		0x76, 0x47, 0x3d, 0x52, 0x30, 0xd3, 0x7a, 0x35,
		0xf6, 0x35, 0x9c, 0x0e, 0x19, 0xd2, 0x80, 0x88,
		0x74, 0x4e, 0x82, 0x5e, 0x9b, 0x4e, 0xbd, 0x83,
		0x48, 0x6d, 0xec, 0x5b, 0x8f, 0x17, 0xa8, 0x87,
		0xeb, 0x39, 0x93, 0x2f, 0x43, 0x57, 0x90, 0xef,
		0xa9, 0xcd, 0x95, 0xf5, 0xea, 0xdc, 0xb9, 0x5a,
	}

	expPriv := []byte{
		0xfd, 0x7e, 0xbe, 0x65, 0xe4, 0xd1, 0xa2, 0x5c,
		0x61, 0x89, 0x36, 0x66, 0xc1, 0x5a, 0x52, 0x24,
		0x2d, 0x9c, 0xd0, 0x78, 0x4a, 0x2d, 0x56, 0x0e,
		0xb2, 0xf4, 0x4c, 0xfa, 0x73, 0x5c, 0x49, 0xed,
		0x2b, 0xd0, 0x72, 0x0d, 0xf4, 0x2c, 0x92, 0x13,
		0xf3, 0xb4, 0x41, 0x7d, 0x7f, 0x13, 0x13, 0x1d,
		0xef, 0x94, 0x67, 0x7b, 0xd3, 0x1e, 0xb5, 0x0e,
	}

	c.Assert(err, IsNil)
	c.Assert(privKey.symKey(), DeepEquals, expSym)
	c.Assert(privKey.publicKey(), DeepEquals, expPub)
	c.Assert(privKey.secretKey(), DeepEquals, expPriv)
}

func (s *Ed448Suite) Test_DecafComputeSharedSecret(c *C) {
	pub := publicKey([pubKeyBytes]byte{
		0x2a, 0x51, 0xd9, 0x1f, 0xc1, 0x7c, 0xd5, 0xc6,
		0xff, 0x0c, 0x0c, 0x78, 0xee, 0xc7, 0x8e, 0x21,
		0x00, 0xc1, 0x91, 0x1d, 0xc6, 0x25, 0x12, 0x02,
		0x4c, 0xbf, 0xcc, 0x3f, 0xc5, 0xa3, 0xaf, 0xd9,
		0x4e, 0xbd, 0x4d, 0xb9, 0x73, 0xef, 0x46, 0xa7,
		0x70, 0x6e, 0x85, 0x38, 0x87, 0x0c, 0xaf, 0xf6,
		0xf0, 0x05, 0xa9, 0x27, 0x00, 0xcf, 0xf7, 0x7e,
	})

	priv := &privateKey{
		// priv
		0x78, 0x46, 0x08, 0x83, 0x64, 0xff, 0x17, 0x87,
		0x13, 0x3d, 0xa5, 0x9e, 0x71, 0x52, 0xdf, 0xe3,
		0xa5, 0x1b, 0xb6, 0xc7, 0x50, 0xc2, 0xbb, 0xfd,
		0x81, 0x49, 0x4e, 0x7e, 0x23, 0x44, 0x19, 0x3f,
		0x0e, 0xdd, 0x35, 0x10, 0x88, 0xf1, 0xc1, 0x9a,
		0xd1, 0x03, 0xbf, 0xf6, 0x99, 0x23, 0xf0, 0x91,
		0x05, 0xe3, 0x66, 0x30, 0xf8, 0xf0, 0x8f, 0x14,
		// pub
		0xb2, 0x95, 0x4a, 0x76, 0x1f, 0x3d, 0x98, 0x03,
		0xaa, 0x2b, 0xbc, 0x8c, 0x98, 0x08, 0x0e, 0x52,
		0x7d, 0x56, 0xb7, 0xb6, 0x14, 0x08, 0x93, 0x0b,
		0xc9, 0x0b, 0xf7, 0x89, 0x80, 0x4f, 0x4f, 0x2c,
		0x8c, 0x65, 0x37, 0xdd, 0xb9, 0xb0, 0xdf, 0x62,
		0xe1, 0xd6, 0x00, 0x9c, 0xee, 0x58, 0x85, 0x21,
		0x3a, 0x56, 0x7d, 0xb1, 0x16, 0x70, 0x0a, 0x57,
		// sym
		0xd1, 0x50, 0x0e, 0xf4, 0x1c, 0xa7, 0xbe, 0xe3,
		0x7d, 0x2d, 0x95, 0x14, 0x9b, 0x75, 0xeb, 0xab,
		0xb0, 0x66, 0xc9, 0xe3, 0x66, 0x32, 0xd2, 0x12,
		0x23, 0x4f, 0xf2, 0x4a, 0x96, 0x94, 0x52, 0x3e,
	}

	exp := []byte{
		0xcd, 0xd4, 0x17, 0xbb, 0xcd, 0x0e, 0xd3, 0xf6,
		0x65, 0x60, 0xd0, 0x16, 0xca, 0x88, 0x0a, 0x23,
		0x8e, 0x36, 0x5e, 0x0c, 0xd6, 0x47, 0xce, 0xbf,
		0x36, 0x55, 0x15, 0xf6, 0xc2, 0x1d, 0x5c, 0xce,
		0x06, 0x85, 0x74, 0x84, 0xe0, 0xa1, 0xdb, 0x3d,
		0x07, 0x6e, 0x21, 0x28, 0x43, 0x21, 0x16, 0x98,
		0x7f, 0x95, 0x44, 0x8c, 0x1a, 0x6a, 0xa4, 0xd1,
	}

	shared, ok := decafCurve.decafComputeSecret(priv, pub)

	c.Assert(shared, DeepEquals, exp)
	c.Assert(ok, DeepEquals, decafTrue)
}
func (s *Ed448Suite) Test_DecafDeriveNonce(c *C) {
	msg := []byte("Hello, world!")

	symKey := [symKeyBytes]byte{
		0xd1, 0x50, 0x0e, 0xf4, 0x1c, 0xa7, 0xbe, 0xe3,
		0x7d, 0x2d, 0x95, 0x14, 0x9b, 0x75, 0xeb, 0xab,
		0xb0, 0x66, 0xc9, 0xe3, 0x66, 0x32, 0xd2, 0x12,
		0x23, 0x4f, 0xf2, 0x4a, 0x96, 0x94, 0x52, 0x3e,
	}

	expectedNonce := &scalar{
		0x358a567a, 0x1b623fd7, 0x1c37439c, 0xb8713abb,
		0x6e7de9d9, 0xb3c4e14e, 0x04ac6658, 0x49706bbb,
		0x13b50aa4, 0x091d9be6, 0x366bc5c8, 0x301b1d05,
		0x73bed629, 0x34663085,
	}

	nonce := decafDeriveNonce(msg, symKey[:])

	c.Assert(nonce, DeepEquals, expectedNonce)
}

func (s *Ed448Suite) Test_DecafDeriveChallenge(c *C) {
	msg := []byte("Hello, world!")

	pubKey := [pubKeyBytes]byte{
		0xb2, 0x95, 0x4a, 0x76, 0x1f, 0x3d, 0x98, 0x03,
		0xaa, 0x2b, 0xbc, 0x8c, 0x98, 0x08, 0x0e, 0x52,
		0x7d, 0x56, 0xb7, 0xb6, 0x14, 0x08, 0x93, 0x0b,
		0xc9, 0x0b, 0xf7, 0x89, 0x80, 0x4f, 0x4f, 0x2c,
		0x8c, 0x65, 0x37, 0xdd, 0xb9, 0xb0, 0xdf, 0x62,
		0xe1, 0xd6, 0x00, 0x9c, 0xee, 0x58, 0x85, 0x21,
		0x3a, 0x56, 0x7d, 0xb1, 0x16, 0x70, 0x0a, 0x57,
	}

	tmpSignature := [fieldBytes]uint8{
		0xee, 0xec, 0x0c, 0xa7, 0x39, 0x65, 0x3c, 0x35,
		0xe2, 0x28, 0xd3, 0xc8, 0xc1, 0x07, 0x96, 0xeb,
		0x06, 0xe8, 0x14, 0x05, 0x62, 0x52, 0xab, 0x6c,
		0x63, 0xf1, 0x4f, 0x55, 0xb3, 0xea, 0x9b, 0x1d,
		0xbf, 0xe7, 0xb7, 0xec, 0x8b, 0x52, 0x43, 0x46,
		0x35, 0xd5, 0xd5, 0xbb, 0xbb, 0xea, 0xfe, 0x7e,
		0xcd, 0xc8, 0xd6, 0xf2, 0x7c, 0x71, 0x87, 0x61,
	}

	expectedChallenge := &scalar{
		0x4b91949b, 0x8366b93a, 0xea749b37, 0x94751b8c,
		0xe11471b6, 0xae84f274, 0xa0f9e9df, 0x6f6684b1,
		0x1b7da377, 0xd27c53a6, 0x140a1a92, 0x5c6e4d8a,
		0xa8aeeb36, 0x0733d7f9,
	}

	challenge := decafDeriveChallenge(pubKey[:], tmpSignature, msg)

	c.Assert(challenge, DeepEquals, expectedChallenge)
}

func (s *Ed448Suite) Test_DecafDeriveTemporarySignature(c *C) {

	nonce := &scalar{
		0x358a567a, 0x1b623fd7, 0x1c37439c, 0xb8713abb,
		0x6e7de9d9, 0xb3c4e14e, 0x04ac6658, 0x49706bbb,
		0x13b50aa4, 0x091d9be6, 0x366bc5c8, 0x301b1d05,
		0x73bed629, 0x34663085,
	}

	exp := [fieldBytes]byte{
		0xee, 0xec, 0x0c, 0xa7, 0x39, 0x65, 0x3c, 0x35,
		0xe2, 0x28, 0xd3, 0xc8, 0xc1, 0x07, 0x96, 0xeb,
		0x06, 0xe8, 0x14, 0x05, 0x62, 0x52, 0xab, 0x6c,
		0x63, 0xf1, 0x4f, 0x55, 0xb3, 0xea, 0x9b, 0x1d,
		0xbf, 0xe7, 0xb7, 0xec, 0x8b, 0x52, 0x43, 0x46,
		0x35, 0xd5, 0xd5, 0xbb, 0xbb, 0xea, 0xfe, 0x7e,
		0xcd, 0xc8, 0xd6, 0xf2, 0x7c, 0x71, 0x87, 0x61,
	}

	sig := decafCurve.decafDeriveTemporarySignature(nonce)

	c.Assert(sig, DeepEquals, exp)
}

func (s *Ed448Suite) Test_DecafSign(c *C) {
	msg := []byte("Hello, world!")
	k := privateKey([privKeyBytes]byte{
		//secret
		0x78, 0x46, 0x08, 0x83, 0x64, 0xff, 0x17, 0x87,
		0x13, 0x3d, 0xa5, 0x9e, 0x71, 0x52, 0xdf, 0xe3,
		0xa5, 0x1b, 0xb6, 0xc7, 0x50, 0xc2, 0xbb, 0xfd,
		0x81, 0x49, 0x4e, 0x7e, 0x23, 0x44, 0x19, 0x3f,
		0x0e, 0xdd, 0x35, 0x10, 0x88, 0xf1, 0xc1, 0x9a,
		0xd1, 0x03, 0xbf, 0xf6, 0x99, 0x23, 0xf0, 0x91,
		0x05, 0xe3, 0x66, 0x30, 0xf8, 0xf0, 0x8f, 0x14,
		//public
		0xb2, 0x95, 0x4a, 0x76, 0x1f, 0x3d, 0x98, 0x03,
		0xaa, 0x2b, 0xbc, 0x8c, 0x98, 0x08, 0x0e, 0x52,
		0x7d, 0x56, 0xb7, 0xb6, 0x14, 0x08, 0x93, 0x0b,
		0xc9, 0x0b, 0xf7, 0x89, 0x80, 0x4f, 0x4f, 0x2c,
		0x8c, 0x65, 0x37, 0xdd, 0xb9, 0xb0, 0xdf, 0x62,
		0xe1, 0xd6, 0x00, 0x9c, 0xee, 0x58, 0x85, 0x21,
		0x3a, 0x56, 0x7d, 0xb1, 0x16, 0x70, 0x0a, 0x57,
		//symmetric
		0xd1, 0x50, 0x0e, 0xf4, 0x1c, 0xa7, 0xbe, 0xe3,
		0x7d, 0x2d, 0x95, 0x14, 0x9b, 0x75, 0xeb, 0xab,
		0xb0, 0x66, 0xc9, 0xe3, 0x66, 0x32, 0xd2, 0x12,
		0x23, 0x4f, 0xf2, 0x4a, 0x96, 0x94, 0x52, 0x3e,
	})

	expectedSignature := [signatureBytes]byte{
		0xee, 0xec, 0x0c, 0xa7, 0x39, 0x65, 0x3c, 0x35,
		0xe2, 0x28, 0xd3, 0xc8, 0xc1, 0x07, 0x96, 0xeb,
		0x06, 0xe8, 0x14, 0x05, 0x62, 0x52, 0xab, 0x6c,
		0x63, 0xf1, 0x4f, 0x55, 0xb3, 0xea, 0x9b, 0x1d,
		0xbf, 0xe7, 0xb7, 0xec, 0x8b, 0x52, 0x43, 0x46,
		0x35, 0xd5, 0xd5, 0xbb, 0xbb, 0xea, 0xfe, 0x7e,
		0xcd, 0xc8, 0xd6, 0xf2, 0x7c, 0x71, 0x87, 0x61,
		0xfa, 0x77, 0xed, 0x08, 0x51, 0x91, 0xc4, 0x85,
		0x74, 0x28, 0xdd, 0xa0, 0xed, 0xbc, 0x88, 0x71,
		0xbd, 0xc3, 0x34, 0x9a, 0xce, 0xee, 0x1a, 0xab,
		0x4c, 0xa2, 0x37, 0xea, 0xb4, 0xea, 0xd2, 0x8d,
		0x25, 0xf1, 0x10, 0x86, 0xc0, 0x60, 0xeb, 0xb3,
		0xb0, 0x9a, 0xaa, 0x8a, 0x4b, 0x00, 0x9e, 0xf1,
		0x93, 0x25, 0xfe, 0x78, 0x0f, 0xdd, 0xa1, 0x3a,
	}

	signature, err := decafCurve.decafSign(msg, &k)

	c.Assert(err, IsNil)
	c.Assert(signature, DeepEquals, expectedSignature)
}

func (s *Ed448Suite) Test_DecafVerify(c *C) {
	msg := []byte("Hello, world!")

	k := publicKey([pubKeyBytes]byte{
		0xb2, 0x95, 0x4a, 0x76, 0x1f, 0x3d, 0x98, 0x03,
		0xaa, 0x2b, 0xbc, 0x8c, 0x98, 0x08, 0x0e, 0x52,
		0x7d, 0x56, 0xb7, 0xb6, 0x14, 0x08, 0x93, 0x0b,
		0xc9, 0x0b, 0xf7, 0x89, 0x80, 0x4f, 0x4f, 0x2c,
		0x8c, 0x65, 0x37, 0xdd, 0xb9, 0xb0, 0xdf, 0x62,
		0xe1, 0xd6, 0x00, 0x9c, 0xee, 0x58, 0x85, 0x21,
		0x3a, 0x56, 0x7d, 0xb1, 0x16, 0x70, 0x0a, 0x57,
	})

	signature := [signatureBytes]byte{
		0xee, 0xec, 0x0c, 0xa7, 0x39, 0x65, 0x3c, 0x35,
		0xe2, 0x28, 0xd3, 0xc8, 0xc1, 0x07, 0x96, 0xeb,
		0x06, 0xe8, 0x14, 0x05, 0x62, 0x52, 0xab, 0x6c,
		0x63, 0xf1, 0x4f, 0x55, 0xb3, 0xea, 0x9b, 0x1d,
		0xbf, 0xe7, 0xb7, 0xec, 0x8b, 0x52, 0x43, 0x46,
		0x35, 0xd5, 0xd5, 0xbb, 0xbb, 0xea, 0xfe, 0x7e,
		0xcd, 0xc8, 0xd6, 0xf2, 0x7c, 0x71, 0x87, 0x61,
		0xfa, 0x77, 0xed, 0x08, 0x51, 0x91, 0xc4, 0x85,
		0x74, 0x28, 0xdd, 0xa0, 0xed, 0xbc, 0x88, 0x71,
		0xbd, 0xc3, 0x34, 0x9a, 0xce, 0xee, 0x1a, 0xab,
		0x4c, 0xa2, 0x37, 0xea, 0xb4, 0xea, 0xd2, 0x8d,
		0x25, 0xf1, 0x10, 0x86, 0xc0, 0x60, 0xeb, 0xb3,
		0xb0, 0x9a, 0xaa, 0x8a, 0x4b, 0x00, 0x9e, 0xf1,
		0x93, 0x25, 0xfe, 0x78, 0x0f, 0xdd, 0xa1, 0x3a,
	}

	valid, err := decafCurve.decafVerify(signature, msg, &k)

	c.Assert(valid, Equals, true)
	c.Assert(err, IsNil)

	// unverifiable
	k1 := publicKey([pubKeyBytes]byte{
		0xb4, 0x95, 0x4a, 0x76, 0x1f, 0x3d, 0x98, 0x03,
		0xaa, 0x2b, 0xbc, 0x8c, 0x98, 0xe6, 0x0e, 0x52,
		0x7d, 0xf6, 0xb7, 0xb6, 0xfe, 0xe8, 0x93, 0x0b,
		0xc9, 0x0b, 0xf7, 0x89, 0x80, 0x4f, 0x4f, 0x2c,
		0x8c, 0x65, 0x37, 0xdd, 0xb9, 0xb0, 0xdf, 0x62,
		0xe1, 0xd6, 0x00, 0x9c, 0xee, 0x58, 0x85, 0x21,
		0x3a, 0x56, 0x7d, 0xb1, 0x16, 0x70, 0x0a, 0x57,
	})

	valid, err = decafCurve.decafVerify(signature, msg, &k1)

	c.Assert(valid, Equals, false)
	c.Assert(err, ErrorMatches, "unable to verify given signature")

	// invalid input
	signature = [112]byte{
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	}

	invalid, err := decafCurve.decafVerify(signature, msg, &k)
	c.Assert(invalid, Equals, false)
	c.Assert(err, ErrorMatches, "unable to decode given point")
}
