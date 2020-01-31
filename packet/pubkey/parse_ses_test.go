package pubkey

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/spiegel-im-spiegel/gpgpdump/info"
	"github.com/spiegel-im-spiegel/gpgpdump/options"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/context"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/reader"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/values"
)

var (
	pubkeySes16  = []byte{16, 0x08, 0x00, 0xca, 0xc6, 0x90, 0x39, 0xfd, 0x65, 0x20, 0x48, 0xba, 0xe0, 0x13, 0xb3, 0xc4, 0x75, 0x3a, 0xf2, 0xab, 0x80, 0x4c, 0xd2, 0x72, 0xce, 0x57, 0x85, 0x28, 0x2c, 0xc9, 0x8f, 0x16, 0x7d, 0xb3, 0x57, 0xec, 0x5a, 0x18, 0x5f, 0x29, 0xa6, 0x1f, 0xca, 0xed, 0x4b, 0x27, 0x03, 0x16, 0x5b, 0x07, 0xbb, 0xef, 0xf0, 0x30, 0x13, 0x6c, 0x89, 0x4e, 0xdf, 0x85, 0x64, 0x68, 0x30, 0xc9, 0x01, 0xbf, 0x8a, 0xaf, 0xc4, 0x3b, 0xb8, 0xe1, 0x08, 0x79, 0xdf, 0x1e, 0xde, 0x00, 0x4c, 0x8e, 0x9d, 0x4f, 0x14, 0x31, 0xc9, 0x78, 0x88, 0x05, 0xdc, 0x35, 0x59, 0xbb, 0x84, 0xaa, 0x23, 0x9c, 0x1a, 0xc5, 0xd7, 0x7a, 0x48, 0xe3, 0xe6, 0x50, 0x62, 0x84, 0x90, 0xa2, 0x38, 0x57, 0x2b, 0x56, 0x07, 0x88, 0xfa, 0xc6, 0xa2, 0xf7, 0x9f, 0x9c, 0x26, 0x0f, 0x4a, 0x36, 0xa8, 0x28, 0x2f, 0xa9, 0xa2, 0x86, 0x04, 0x21, 0x5b, 0xc4, 0xbc, 0x93, 0xe6, 0x9f, 0xd0, 0x93, 0x7a, 0xd0, 0x39, 0x71, 0x57, 0xcf, 0x2b, 0xda, 0x1a, 0xc2, 0x4b, 0xe1, 0xe7, 0x4b, 0xd9, 0xa0, 0x44, 0xce, 0x70, 0xe6, 0xe7, 0x74, 0x4e, 0x9a, 0xde, 0x24, 0x99, 0xbd, 0x03, 0x3d, 0xdf, 0x6e, 0x1a, 0x13, 0xe5, 0x35, 0x39, 0x7f, 0x51, 0xbe, 0xd0, 0xbc, 0x9e, 0x42, 0x82, 0x7f, 0x0f, 0xef, 0xe9, 0xef, 0x87, 0x9a, 0x52, 0x39, 0x15, 0x59, 0xcd, 0x47, 0x79, 0xd5, 0x1d, 0xf8, 0x10, 0xcb, 0xa5, 0x47, 0x4a, 0x4c, 0x9c, 0x52, 0xee, 0x9f, 0x81, 0xd3, 0x21, 0x20, 0xd5, 0xae, 0x9e, 0x30, 0x47, 0xa3, 0xd0, 0x28, 0xb1, 0xb6, 0xf3, 0x2e, 0x3f, 0x9e, 0x6a, 0x7c, 0x21, 0x18, 0x99, 0x72, 0xb2, 0xc9, 0x83, 0x49, 0x8b, 0x33, 0xfa, 0xe9, 0x0a, 0x8e, 0x66, 0x48, 0x03, 0x27, 0xf9, 0x56, 0xb0, 0x2f, 0xa5, 0x7d, 0xd1, 0x47, 0xb0, 0x08, 0x00, 0xf4, 0xf1, 0x13, 0x6a, 0x0f, 0xa5, 0xc5, 0x0f, 0x07, 0xe2, 0x02, 0x81, 0xe6, 0xbb, 0xd6, 0xc8, 0x73, 0x07, 0x8b, 0x37, 0xc6, 0x1e, 0x27, 0x79, 0xab, 0xf8, 0xc0, 0x5a, 0x03, 0x62, 0xf2, 0xa7, 0x5a, 0x9f, 0xf0, 0x80, 0xcf, 0x84, 0x5b, 0xe7, 0xba, 0xa0, 0xa2, 0xc6, 0x1f, 0x5d, 0xe6, 0x87, 0x6a, 0x22, 0x7c, 0x91, 0x2b, 0x01, 0xd8, 0x78, 0xbb, 0xaf, 0xbe, 0xba, 0xfa, 0x4c, 0x00, 0xc5, 0xfa, 0xd4, 0xba, 0x21, 0xc3, 0xb6, 0xdf, 0xcc, 0xbe, 0xf9, 0x39, 0x12, 0xf4, 0x31, 0xc9, 0x4f, 0x14, 0x93, 0x04, 0x4f, 0x13, 0xe6, 0x58, 0x7d, 0xff, 0xc2, 0x2b, 0x5a, 0x90, 0x5e, 0x6d, 0xb2, 0x57, 0xb5, 0xe4, 0xa4, 0xd8, 0xf8, 0x98, 0xcf, 0x98, 0x7c, 0xbc, 0xf5, 0xa4, 0x7a, 0x30, 0x51, 0xc9, 0x19, 0x92, 0xf4, 0x4e, 0x94, 0x14, 0x0a, 0xd4, 0x99, 0xaf, 0xf6, 0x9f, 0x53, 0x46, 0xa6, 0xdc, 0x41, 0x0c, 0xf5, 0xb3, 0xe1, 0x33, 0xb1, 0x95, 0x56, 0x3b, 0x5a, 0xb5, 0xee, 0x7d, 0x8c, 0x09, 0x22, 0xb2, 0x24, 0x18, 0x3e, 0x4c, 0xac, 0xc5, 0xb7, 0x72, 0x1b, 0x8b, 0xb2, 0x81, 0x29, 0xa2, 0x58, 0x94, 0x94, 0x0b, 0xab, 0xba, 0x43, 0xdd, 0x79, 0xc1, 0xf3, 0x62, 0x4d, 0x55, 0x7c, 0x56, 0x3d, 0x22, 0xae, 0xca, 0xdc, 0x8c, 0x55, 0x31, 0xd4, 0xaf, 0xe3, 0xf9, 0x57, 0x3b, 0xc2, 0xd9, 0x05, 0xcf, 0xa2, 0x24, 0x54, 0xdc, 0x25, 0x63, 0x12, 0x93, 0xb4, 0x95, 0x64, 0x5d, 0x40, 0x79, 0x47, 0x2a, 0xbc, 0x5c, 0x97, 0x72, 0x52, 0x76, 0x4d, 0x74, 0xc4, 0x4a, 0x53, 0xa4, 0xe0, 0x60, 0x15, 0x9e, 0xbc, 0x27, 0x2f, 0x94, 0x1d, 0xbb, 0xc3, 0x65, 0xb3, 0xf8, 0x67, 0x2d, 0x37, 0x6c, 0x81, 0xf0, 0x33, 0xd7, 0xb2, 0xb9, 0x04, 0x92, 0x9c, 0x99, 0xdb, 0x61, 0x62, 0xc7, 0xe4}
	pubkeySes17  = []byte{17, 0x02, 0x03, 0x04}
	pubkeySes18a = []byte{0x12, 0x02, 0x03, 0x04, 0xc3, 0xe7, 0xd7, 0x2b, 0xaf, 0x25, 0x2a, 0x19, 0xf6, 0x27, 0x80, 0xea, 0x7c, 0x4f, 0x6d, 0xca, 0x61, 0x22, 0x5a, 0xe3, 0xad, 0x0c, 0xfb, 0xd9, 0xa2, 0xd5, 0xa4, 0x30, 0x9a, 0xf3, 0xee, 0x34, 0x54, 0xae, 0xa8, 0xf6, 0x46, 0xac, 0x8a, 0xae, 0x38, 0xa6, 0x4f, 0xf3, 0xf2, 0xee, 0x30, 0x40, 0x62, 0x5b, 0x07, 0xe7, 0x2b, 0xee, 0x9a, 0x90, 0xd4, 0x6f, 0x1e, 0xd7, 0xc3, 0x26, 0x21, 0xab, 0x30, 0x4a, 0xfe, 0x88, 0xa2, 0x9f, 0x0e, 0xab, 0xf3, 0xbe, 0x7a, 0x89, 0x27, 0x32, 0x38, 0xb8, 0x06, 0x75, 0xfc, 0xac, 0x3c, 0xd4, 0xba, 0x0f, 0x49, 0x64, 0x15, 0xaa, 0x48, 0x9a, 0xdb, 0xc1, 0x8a, 0x7b, 0x11, 0x76, 0xfb, 0x2f, 0xef, 0xef, 0xb0, 0x29, 0xa9, 0x24, 0x75, 0x6d, 0x69, 0x12, 0x4d}
	pubkeySes18b = []byte{0x12, 0x02, 0x03, 0x04}
	pubkeySes18c = []byte{0x12, 0x02, 0x03, 0x04, 0xc3, 0xe7, 0xd7, 0x2b, 0xaf, 0x25, 0x2a, 0x19, 0xf6, 0x27, 0x80, 0xea, 0x7c, 0x4f, 0x6d, 0xca, 0x61, 0x22, 0x5a, 0xe3, 0xad, 0x0c, 0xfb, 0xd9, 0xa2, 0xd5, 0xa4, 0x30, 0x9a, 0xf3, 0xee, 0x34, 0x54, 0xae, 0xa8, 0xf6, 0x46, 0xac, 0x8a, 0xae, 0x38, 0xa6, 0x4f, 0xf3, 0xf2, 0xee, 0x30, 0x40, 0x62, 0x5b, 0x07, 0xe7, 0x2b, 0xee, 0x9a, 0x90, 0xd4, 0x6f, 0x1e, 0xd7, 0xc3, 0x26, 0x21, 0xab, 0x30, 0x4a, 0xfe, 0x88, 0xa2, 0x9f, 0x0e, 0xab, 0xf3, 0xbe}
	pubkeySes19  = []byte{19, 0x02, 0x03, 0x04}
	pubkeySesUn  = []byte{100, 0x02, 0x03, 0x04}
)

const (
	pubkeySesResult16 = `
	ElGamal g^k mod p (2048 bits)
		ca c6 90 39 fd 65 20 48 ba e0 13 b3 c4 75 3a f2 ab 80 4c d2 72 ce 57 85 28 2c c9 8f 16 7d b3 57 ec 5a 18 5f 29 a6 1f ca ed 4b 27 03 16 5b 07 bb ef f0 30 13 6c 89 4e df 85 64 68 30 c9 01 bf 8a af c4 3b b8 e1 08 79 df 1e de 00 4c 8e 9d 4f 14 31 c9 78 88 05 dc 35 59 bb 84 aa 23 9c 1a c5 d7 7a 48 e3 e6 50 62 84 90 a2 38 57 2b 56 07 88 fa c6 a2 f7 9f 9c 26 0f 4a 36 a8 28 2f a9 a2 86 04 21 5b c4 bc 93 e6 9f d0 93 7a d0 39 71 57 cf 2b da 1a c2 4b e1 e7 4b d9 a0 44 ce 70 e6 e7 74 4e 9a de 24 99 bd 03 3d df 6e 1a 13 e5 35 39 7f 51 be d0 bc 9e 42 82 7f 0f ef e9 ef 87 9a 52 39 15 59 cd 47 79 d5 1d f8 10 cb a5 47 4a 4c 9c 52 ee 9f 81 d3 21 20 d5 ae 9e 30 47 a3 d0 28 b1 b6 f3 2e 3f 9e 6a 7c 21 18 99 72 b2 c9 83 49 8b 33 fa e9 0a 8e 66 48 03 27 f9 56 b0 2f a5 7d d1 47 b0
	ElGamal m * y^k mod p (2048 bits)
		f4 f1 13 6a 0f a5 c5 0f 07 e2 02 81 e6 bb d6 c8 73 07 8b 37 c6 1e 27 79 ab f8 c0 5a 03 62 f2 a7 5a 9f f0 80 cf 84 5b e7 ba a0 a2 c6 1f 5d e6 87 6a 22 7c 91 2b 01 d8 78 bb af be ba fa 4c 00 c5 fa d4 ba 21 c3 b6 df cc be f9 39 12 f4 31 c9 4f 14 93 04 4f 13 e6 58 7d ff c2 2b 5a 90 5e 6d b2 57 b5 e4 a4 d8 f8 98 cf 98 7c bc f5 a4 7a 30 51 c9 19 92 f4 4e 94 14 0a d4 99 af f6 9f 53 46 a6 dc 41 0c f5 b3 e1 33 b1 95 56 3b 5a b5 ee 7d 8c 09 22 b2 24 18 3e 4c ac c5 b7 72 1b 8b b2 81 29 a2 58 94 94 0b ab ba 43 dd 79 c1 f3 62 4d 55 7c 56 3d 22 ae ca dc 8c 55 31 d4 af e3 f9 57 3b c2 d9 05 cf a2 24 54 dc 25 63 12 93 b4 95 64 5d 40 79 47 2a bc 5c 97 72 52 76 4d 74 c4 4a 53 a4 e0 60 15 9e bc 27 2f 94 1d bb c3 65 b3 f8 67 2d 37 6c 81 f0 33 d7 b2 b9 04 92 9c 99 db 61 62 c7 e4
`
	pubkeySesResult17 = `
	Multi-precision integers of DSA (3 bytes)
		02 03 04
`
	pubkeySesResult18a = `
	ECDH EC point (uncompressed format) (515 bits)
		04 c3 e7 d7 2b af 25 2a 19 f6 27 80 ea 7c 4f 6d ca 61 22 5a e3 ad 0c fb d9 a2 d5 a4 30 9a f3 ee 34 54 ae a8 f6 46 ac 8a ae 38 a6 4f f3 f2 ee 30 40 62 5b 07 e7 2b ee 9a 90 d4 6f 1e d7 c3 26 21 ab
	symmetric key (encoded) (48 bytes)
		4a fe 88 a2 9f 0e ab f3 be 7a 89 27 32 38 b8 06 75 fc ac 3c d4 ba 0f 49 64 15 aa 48 9a db c1 8a 7b 11 76 fb 2f ef ef b0 29 a9 24 75 6d 69 12 4d
`
	// 	pubkeySesResult18b = "\n"
	// 	pubkeySesResult18c = `
	// 	ECDH EC point (uncompressed format) (515 bits)
	// 		04 c3 e7 d7 2b af 25 2a 19 f6 27 80 ea 7c 4f 6d ca 61 22 5a e3 ad 0c fb d9 a2 d5 a4 30 9a f3 ee 34 54 ae a8 f6 46 ac 8a ae 38 a6 4f f3 f2 ee 30 40 62 5b 07 e7 2b ee 9a 90 d4 6f 1e d7 c3 26 21 ab
	// `
	pubkeySesResult19 = `
	Multi-precision integers of ECDSA (3 bytes)
		02 03 04
`
	pubkeySesResultUnknown = `
	Multi-precision integers of Unknown (pub 100) (3 bytes)
		02 03 04
`
)

func TestPubkeySes(t *testing.T) {
	testCases := []struct {
		content []byte
		res     string
		err     error
	}{
		{content: pubkeySes16, res: pubkeySesResult16, err: nil},
		{content: pubkeySes17, res: pubkeySesResult17, err: nil},
		{content: pubkeySes18a, res: pubkeySesResult18a, err: nil},
		{content: pubkeySes18b, res: "", err: io.ErrUnexpectedEOF},
		{content: pubkeySes18c, res: "", err: io.ErrUnexpectedEOF},
		{content: pubkeySes19, res: pubkeySesResult19, err: nil},
		{content: pubkeySesUn, res: pubkeySesResultUnknown, err: nil},
	}
	for _, tc := range testCases {
		parent := info.NewItem()
		cxt := context.New(options.New(
			options.Set(options.DEBUG, true),
			options.Set(options.GDUMP, true),
			options.Set(options.INTEGER, true),
			options.Set(options.LITERAL, true),
			options.Set(options.MARKER, true),
			options.Set(options.PRIVATE, true),
			options.Set(options.UTC, true),
		))
		if err := New(cxt, values.PubID(tc.content[0]), reader.New(tc.content[1:])).ParseSes(parent); !errors.Is(err, tc.err) {
			t.Errorf("Parse() = \"%v\", want \"%v\".", err, tc.err)
		} else if err != nil {
			fmt.Printf("Info: %+v\n", err)
		} else {
			str := parent.String()
			if str != tc.res {
				t.Errorf("Parse() = \"%v\", want \"%v\".", str, tc.res)
			}
		}
	}
}

/* Copyright 2016-2020 Spiegel
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
