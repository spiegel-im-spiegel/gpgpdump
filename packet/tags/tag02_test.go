package tags

import (
	"testing"

	"github.com/spiegel-im-spiegel/gpgpdump/options"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/context"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/reader"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/values"

	openpgp "golang.org/x/crypto/openpgp/packet"
)

var (
	tag02Body1 = []byte{0x04, 0x01, 0x13, 0x08, 0x00, 0x06, 0x05, 0x02, 0x54, 0xc3, 0x08, 0xdf, 0x00, 0x0a, 0x09, 0x10, 0x31, 0xfb, 0xfd, 0xa9, 0x5f, 0xbb, 0xfa, 0x18, 0x36, 0x1f, 0x01, 0x00, 0xea, 0x1d, 0xa2, 0x14, 0x5b, 0x82, 0x06, 0xfd, 0xd5, 0xae, 0xc4, 0x9f, 0xd8, 0x14, 0x44, 0x41, 0xa4, 0xf5, 0x4f, 0x56, 0x69, 0xad, 0x9a, 0xb0, 0x44, 0xf3, 0xa3, 0x88, 0xb2, 0x60, 0xf4, 0x0c, 0x00, 0xfc, 0x0a, 0xd3, 0xc0, 0x23, 0xf3, 0xed, 0xcd, 0xaf, 0x9b, 0x19, 0x6f, 0xee, 0xc4, 0x65, 0x44, 0xb5, 0x08, 0xe8, 0x27, 0x6c, 0x3a, 0xa8, 0x6e, 0x3b, 0x52, 0x9f, 0x61, 0x7a, 0xea, 0xee, 0x27, 0x48}
	tag02Body2 = []byte{0x04, 0x00, 0x11, 0x08, 0x00, 0x1d, 0x16, 0x21, 0x04, 0x1b, 0x52, 0x02, 0xdb, 0x4a, 0x3e, 0xc7, 0x76, 0xf1, 0xe0, 0xad, 0x18, 0xb4, 0xda, 0x3b, 0xae, 0x7e, 0x20, 0xb8, 0x1c, 0x05, 0x02, 0x5a, 0x19, 0x0d, 0xe4, 0x00, 0x0a, 0x09, 0x10, 0xb4, 0xda, 0x3b, 0xae, 0x7e, 0x20, 0xb8, 0x1c, 0x73, 0x3c, 0x01, 0x00, 0x84, 0xef, 0xee, 0xae, 0x22, 0x69, 0x2e, 0xaf, 0x33, 0xb3, 0x85, 0xe1, 0xee, 0xaa, 0x5d, 0x2f, 0x7a, 0xd4, 0xae, 0xa3, 0x92, 0xd3, 0xe8, 0x73, 0xd4, 0xb0, 0x00, 0x3e, 0xc9, 0x2b, 0x80, 0xf7, 0x00, 0xff, 0x7c, 0xcc, 0xb9, 0xd2, 0x06, 0x48, 0xb3, 0x39, 0x58, 0x9b, 0xa8, 0x99, 0xc5, 0xc2, 0x53, 0x62, 0xbd, 0x8f, 0x16, 0x49, 0x73, 0xe0, 0x65, 0xfe, 0xa6, 0xf7, 0x18, 0x1a, 0x78, 0xff, 0x65, 0xe6}
	tag02Body3 = []byte{0x04, 0x18, 0x13, 0x08, 0x00, 0x0f, 0x05, 0x02, 0x54, 0xc3, 0x01, 0xbf, 0x02, 0x1b, 0x0c, 0x05, 0x09, 0x00, 0x09, 0x3a, 0x80, 0x00, 0x0a, 0x09, 0x10, 0x31, 0xfb, 0xfd, 0xa9, 0x5f, 0xbb, 0xfa, 0x18, 0xc6, 0x27, 0x01, 0x00, 0xc0, 0x2f, 0x76, 0x5a, 0x10, 0x6d, 0x1d, 0x22, 0x1e, 0x62, 0xc1, 0x9b, 0xbc, 0x62, 0xd1, 0x06, 0x4a, 0xf1, 0x3a, 0x47, 0x5a, 0xe9, 0x0b, 0xf1, 0x39, 0x6c, 0xe3, 0x67, 0xa0, 0x96, 0x3c, 0xd2, 0x01, 0x00, 0x8c, 0x59, 0x1c, 0x3a, 0x85, 0x0e, 0x1e, 0xd3, 0x98, 0x45, 0x13, 0x4d, 0x30, 0xe2, 0xb9, 0xa4, 0x15, 0x0e, 0x1b, 0x6d, 0x66, 0x1a, 0xa7, 0xe7, 0xd5, 0xe2, 0x51, 0x07, 0x95, 0x60, 0x87, 0x91}
	tag02Body4 = []byte{0x03, 0x05, 0x00, 0x36, 0x5e, 0xba, 0x44, 0x0f, 0x64, 0x8a, 0x1c, 0x9e, 0x4f, 0x74, 0x4d, 0x01, 0x01, 0x66, 0x36, 0x04, 0x00, 0x8f, 0x8c, 0x6b, 0x45, 0xa7, 0x65, 0xbd, 0x37, 0xf6, 0x76, 0x58, 0x85, 0x7c, 0x39, 0x66, 0x7a, 0xc5, 0xc1, 0x48, 0xf3, 0xb8, 0x85, 0x69, 0x7f, 0x22, 0x54, 0x71, 0x50, 0x0e, 0x97, 0xb2, 0x51, 0x77, 0x53, 0xa2, 0x22, 0xd4, 0x46, 0xec, 0x0c, 0x50, 0xbe, 0xee, 0xe6, 0xb0, 0xc2, 0x76, 0x08, 0xf0, 0x6b, 0x0e, 0x6c, 0xfc, 0xe6, 0xef, 0xcd, 0x10, 0x3d, 0x10, 0xfd, 0xb3, 0x87, 0x40, 0x20, 0x55, 0x6c, 0x06, 0xae, 0x41, 0xc5, 0x7c, 0x0d, 0x17, 0x75, 0x44, 0x32, 0x7d, 0x08, 0x41, 0x45, 0x95, 0xda, 0xd6, 0x57, 0x74, 0x58, 0x38, 0x72, 0x6e, 0xf7, 0x1f, 0x63, 0xce, 0xd8, 0x00, 0x1b, 0x25, 0x37, 0x23, 0xb2, 0x56, 0x1a, 0x02, 0x9a, 0xee, 0x5a, 0x57, 0xf7, 0xa3, 0xab, 0x2d, 0x89, 0x20, 0x85, 0x1c, 0xc5, 0xc0, 0xec, 0x64, 0xe9, 0x2f, 0x0b, 0xf5, 0x4b, 0x5f, 0x2b, 0x65, 0x39}
)

const (
	tag02Redult1 = `Signature Packet (tag 2) (94 bytes)
	04 01 13 08 00 06 05 02 54 c3 08 df 00 0a 09 10 31 fb fd a9 5f bb fa 18 36 1f 01 00 ea 1d a2 14 5b 82 06 fd d5 ae c4 9f d8 14 44 41 a4 f5 4f 56 69 ad 9a b0 44 f3 a3 88 b2 60 f4 0c 00 fc 0a d3 c0 23 f3 ed cd af 9b 19 6f ee c4 65 44 b5 08 e8 27 6c 3a a8 6e 3b 52 9f 61 7a ea ee 27 48
	Version: 4 (current)
		04
	Signiture Type: Signature of a canonical text document (0x01)
		01
	Public-key Algorithm: ECDSA public key algorithm (pub 19)
		13
	Hash Algorithm: SHA2-256 (hash 8)
		08
	Hashed Subpacket (6 bytes)
		05 02 54 c3 08 df
		Signature Creation Time (sub 2): 2015-01-24T02:52:15Z
			54 c3 08 df
	Unhashed Subpacket (10 bytes)
		09 10 31 fb fd a9 5f bb fa 18
		Issuer (sub 16): 0x31fbfda95fbbfa18
	Hash left 2 bytes
		36 1f
	ECDSA value r (256 bits)
		ea 1d a2 14 5b 82 06 fd d5 ae c4 9f d8 14 44 41 a4 f5 4f 56 69 ad 9a b0 44 f3 a3 88 b2 60 f4 0c
	ECDSA value s (252 bits)
		0a d3 c0 23 f3 ed cd af 9b 19 6f ee c4 65 44 b5 08 e8 27 6c 3a a8 6e 3b 52 9f 61 7a ea ee 27 48
`
	tag02Redult2 = `Signature Packet (tag 2) (117 bytes)
	04 00 11 08 00 1d 16 21 04 1b 52 02 db 4a 3e c7 76 f1 e0 ad 18 b4 da 3b ae 7e 20 b8 1c 05 02 5a 19 0d e4 00 0a 09 10 b4 da 3b ae 7e 20 b8 1c 73 3c 01 00 84 ef ee ae 22 69 2e af 33 b3 85 e1 ee aa 5d 2f 7a d4 ae a3 92 d3 e8 73 d4 b0 00 3e c9 2b 80 f7 00 ff 7c cc b9 d2 06 48 b3 39 58 9b a8 99 c5 c2 53 62 bd 8f 16 49 73 e0 65 fe a6 f7 18 1a 78 ff 65 e6
	Version: 4 (current)
		04
	Signiture Type: Signature of a binary document (0x00)
		00
	Public-key Algorithm: DSA (Digital Signature Algorithm) (pub 17)
		11
	Hash Algorithm: SHA2-256 (hash 8)
		08
	Hashed Subpacket (29 bytes)
		16 21 04 1b 52 02 db 4a 3e c7 76 f1 e0 ad 18 b4 da 3b ae 7e 20 b8 1c 05 02 5a 19 0d e4
		Issuer Fingerprint (sub 33) (21 bytes)
			04 1b 52 02 db 4a 3e c7 76 f1 e0 ad 18 b4 da 3b ae 7e 20 b8 1c
			Version: 4 (need 20 octets length)
			Fingerprint (20 bytes)
				1b 52 02 db 4a 3e c7 76 f1 e0 ad 18 b4 da 3b ae 7e 20 b8 1c
		Signature Creation Time (sub 2): 2017-11-25T06:29:56Z
			5a 19 0d e4
	Unhashed Subpacket (10 bytes)
		09 10 b4 da 3b ae 7e 20 b8 1c
		Issuer (sub 16): 0xb4da3bae7e20b81c
	Hash left 2 bytes
		73 3c
	DSA value r (256 bits)
		84 ef ee ae 22 69 2e af 33 b3 85 e1 ee aa 5d 2f 7a d4 ae a3 92 d3 e8 73 d4 b0 00 3e c9 2b 80 f7
	DSA value s (255 bits)
		7c cc b9 d2 06 48 b3 39 58 9b a8 99 c5 c2 53 62 bd 8f 16 49 73 e0 65 fe a6 f7 18 1a 78 ff 65 e6
`
	tag02Redult3 = `Signature Packet (tag 2) (103 bytes)
	04 18 13 08 00 0f 05 02 54 c3 01 bf 02 1b 0c 05 09 00 09 3a 80 00 0a 09 10 31 fb fd a9 5f bb fa 18 c6 27 01 00 c0 2f 76 5a 10 6d 1d 22 1e 62 c1 9b bc 62 d1 06 4a f1 3a 47 5a e9 0b f1 39 6c e3 67 a0 96 3c d2 01 00 8c 59 1c 3a 85 0e 1e d3 98 45 13 4d 30 e2 b9 a4 15 0e 1b 6d 66 1a a7 e7 d5 e2 51 07 95 60 87 91
	Version: 4 (current)
		04
	Signiture Type: Subkey Binding Signature (0x18)
		18
	Public-key Algorithm: ECDSA public key algorithm (pub 19)
		13
	Hash Algorithm: SHA2-256 (hash 8)
		08
	Hashed Subpacket (15 bytes)
		05 02 54 c3 01 bf 02 1b 0c 05 09 00 09 3a 80
		Signature Creation Time (sub 2): 2015-01-24T02:21:51Z
			54 c3 01 bf
		Key Flags (sub 27) (1 bytes)
			0c
			Flag: This key may be used to encrypt communications.
			Flag: This key may be used to encrypt storage.
		Key Expiration Time (sub 9): 7 days after (2015-01-31T02:21:51Z)
			00 09 3a 80
	Unhashed Subpacket (10 bytes)
		09 10 31 fb fd a9 5f bb fa 18
		Issuer (sub 16): 0x31fbfda95fbbfa18
	Hash left 2 bytes
		c6 27
	ECDSA value r (256 bits)
		c0 2f 76 5a 10 6d 1d 22 1e 62 c1 9b bc 62 d1 06 4a f1 3a 47 5a e9 0b f1 39 6c e3 67 a0 96 3c d2
	ECDSA value s (256 bits)
		8c 59 1c 3a 85 0e 1e d3 98 45 13 4d 30 e2 b9 a4 15 0e 1b 6d 66 1a a7 e7 d5 e2 51 07 95 60 87 91
`
	tag02Redult4 = `Signature Packet (tag 2) (149 bytes)
	03 05 00 36 5e ba 44 0f 64 8a 1c 9e 4f 74 4d 01 01 66 36 04 00 8f 8c 6b 45 a7 65 bd 37 f6 76 58 85 7c 39 66 7a c5 c1 48 f3 b8 85 69 7f 22 54 71 50 0e 97 b2 51 77 53 a2 22 d4 46 ec 0c 50 be ee e6 b0 c2 76 08 f0 6b 0e 6c fc e6 ef cd 10 3d 10 fd b3 87 40 20 55 6c 06 ae 41 c5 7c 0d 17 75 44 32 7d 08 41 45 95 da d6 57 74 58 38 72 6e f7 1f 63 ce d8 00 1b 25 37 23 b2 56 1a 02 9a ee 5a 57 f7 a3 ab 2d 89 20 85 1c c5 c0 ec 64 e9 2f 0b f5 4b 5f 2b 65 39
	Version: 3 (old)
		03
	Hashed material (5 bytes)
		Signiture Type: Signature of a binary document (0x00)
			00
		Signature creation time: 1998-11-27T14:42:12Z
			36 5e ba 44
	Key ID: 0x0f648a1c9e4f744d
	Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
		01
	Hash Algorithm: MD5 (hash 1)
		01
	Hash left 2 bytes
		66 36
	RSA signature value m^d mod n (1024 bits)
		8f 8c 6b 45 a7 65 bd 37 f6 76 58 85 7c 39 66 7a c5 c1 48 f3 b8 85 69 7f 22 54 71 50 0e 97 b2 51 77 53 a2 22 d4 46 ec 0c 50 be ee e6 b0 c2 76 08 f0 6b 0e 6c fc e6 ef cd 10 3d 10 fd b3 87 40 20 55 6c 06 ae 41 c5 7c 0d 17 75 44 32 7d 08 41 45 95 da d6 57 74 58 38 72 6e f7 1f 63 ce d8 00 1b 25 37 23 b2 56 1a 02 9a ee 5a 57 f7 a3 ab 2d 89 20 85 1c c5 c0 ec 64 e9 2f 0b f5 4b 5f 2b 65 39
`
)

func TestTag02a(t *testing.T) {
	op := &openpgp.OpaquePacket{Tag: 2, Contents: tag02Body1}
	cxt := context.NewContext(options.New(
		options.Set(options.DebugOpt, true),
		options.Set(options.IntegerOpt, true),
		options.Set(options.MarkerOpt, true),
		options.Set(options.LiteralOpt, true),
		options.Set(options.PrivateOpt, true),
		options.Set(options.UTCOpt, true),
	))
	i, err := NewTag(op, cxt).Parse()
	if err != nil {
		t.Errorf("NewTag() = %v, want nil error.", err)
		return
	}
	if cxt.AlgMode() != context.ModeNotSpecified {
		t.Errorf("Options.Mode = %v, want \"%v\".", cxt.AlgMode(), context.ModeNotSpecified)
	}
	str := i.String()
	if str != tag02Redult1 {
		t.Errorf("Tag.String = \"%s\", want \"%s\".", str, tag02Redult1)
	}
}

func TestTag02b(t *testing.T) {
	op := &openpgp.OpaquePacket{Tag: 2, Contents: tag02Body2}
	cxt := context.NewContext(options.New(
		options.Set(options.DebugOpt, true),
		options.Set(options.IntegerOpt, true),
		options.Set(options.MarkerOpt, true),
		options.Set(options.LiteralOpt, true),
		options.Set(options.PrivateOpt, true),
		options.Set(options.UTCOpt, true),
	))
	i, err := NewTag(op, cxt).Parse()
	if err != nil {
		t.Errorf("NewTag() = %v, want nil error.", err)
		return
	}
	if cxt.AlgMode() != context.ModeNotSpecified {
		t.Errorf("Options.Mode = %v, want \"%v\".", cxt.AlgMode(), context.ModeNotSpecified)
	}
	str := i.String()
	if str != tag02Redult2 {
		t.Errorf("Tag.String = \"%s\", want \"%s\".", str, tag02Redult2)
	}
}

func TestTag02c(t *testing.T) {
	op := &openpgp.OpaquePacket{Tag: 2, Contents: tag02Body3}
	cxt := context.NewContext(options.New(
		options.Set(options.DebugOpt, true),
		options.Set(options.IntegerOpt, true),
		options.Set(options.MarkerOpt, true),
		options.Set(options.LiteralOpt, true),
		options.Set(options.PrivateOpt, true),
		options.Set(options.UTCOpt, true),
	))
	tm, _ := values.NewDateTime(reader.New([]byte{0x54, 0xc3, 0x01, 0xbf}), cxt.UTC())
	cxt.KeyCreationTime = tm
	i, err := NewTag(op, cxt).Parse()
	if err != nil {
		t.Errorf("NewTag() = %v, want nil error.", err)
		return
	}
	if cxt.AlgMode() != context.ModeNotSpecified {
		t.Errorf("Options.Mode = %v, want \"%v\".", cxt.AlgMode(), context.ModeNotSpecified)
	}
	str := i.String()
	if str != tag02Redult3 {
		t.Errorf("Tag.String = \"%s\", want \"%s\".", str, tag02Redult3)
	}
}

func TestTag02d(t *testing.T) {
	op := &openpgp.OpaquePacket{Tag: 2, Contents: tag02Body4}
	cxt := context.NewContext(options.New(
		options.Set(options.DebugOpt, true),
		options.Set(options.IntegerOpt, true),
		options.Set(options.MarkerOpt, true),
		options.Set(options.LiteralOpt, true),
		options.Set(options.PrivateOpt, true),
		options.Set(options.UTCOpt, true),
	))
	tm, _ := values.NewDateTime(reader.New([]byte{0x54, 0xc3, 0x01, 0xbf}), cxt.UTC())
	cxt.KeyCreationTime = tm
	i, err := NewTag(op, cxt).Parse()
	if err != nil {
		t.Errorf("NewTag() = %v, want nil error.", err)
		return
	}
	if cxt.AlgMode() != context.ModeNotSpecified {
		t.Errorf("Options.Mode = %v, want \"%v\".", cxt.AlgMode(), context.ModeNotSpecified)
	}
	str := i.String()
	if str != tag02Redult4 {
		t.Errorf("Tag.String = \"%s\", want \"%s\".", str, tag02Redult4)
	}
}

/* Copyright 2017 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */