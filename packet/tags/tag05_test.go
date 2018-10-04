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
	tag05Body1 = []byte{0x04, 0x5a, 0xfa, 0x85, 0x65, 0x01, 0x08, 0x00, 0x8a, 0xf8, 0x48, 0xe4, 0xc2, 0x21, 0x23, 0xc4, 0x47, 0xe9, 0x84, 0x13, 0x4d, 0x80, 0x88, 0xad, 0x28, 0x9e, 0x34, 0x32, 0xeb, 0xb2, 0x53, 0x9e, 0xd2, 0xb3, 0x64, 0xc4, 0xf4, 0xe9, 0xf7, 0xc0, 0x5d, 0x8a, 0xe6, 0x01, 0xe1, 0x4f, 0x04, 0x2b, 0xf2, 0x08, 0x38, 0xc1, 0x4d, 0x65, 0x49, 0x3b, 0x25, 0x22, 0xf1, 0xbf, 0x18, 0x0d, 0x31, 0x63, 0x9e, 0xd5, 0x04, 0xcf, 0x4c, 0x99, 0x57, 0xe1, 0xa8, 0x2e, 0xa0, 0xcb, 0x3f, 0xac, 0x25, 0x00, 0x85, 0x00, 0x3d, 0xaa, 0x61, 0xe2, 0xba, 0xe5, 0xe3, 0xf0, 0x1f, 0xf2, 0x36, 0x9a, 0x97, 0x12, 0x7b, 0x0b, 0x91, 0xb0, 0xc2, 0x99, 0xb4, 0x5a, 0x6a, 0x3a, 0x40, 0xec, 0x76, 0x76, 0xa8, 0x19, 0x4d, 0x8b, 0xed, 0xe0, 0xb5, 0x1f, 0xb3, 0xef, 0x65, 0xfa, 0x55, 0x2e, 0xf2, 0xa6, 0xd4, 0x2d, 0x7a, 0x76, 0xd9, 0x7a, 0x3a, 0x3e, 0xe9, 0x39, 0xd3, 0x20, 0xb0, 0x3b, 0x7d, 0xc8, 0xbb, 0xff, 0xa0, 0xc8, 0x88, 0x82, 0xeb, 0xcf, 0x75, 0xf4, 0xac, 0x07, 0x41, 0x76, 0xf6, 0x66, 0xfb, 0x24, 0x87, 0x00, 0x8e, 0x94, 0x73, 0xd7, 0x2c, 0x3e, 0x98, 0xfe, 0x46, 0xa0, 0x2e, 0x75, 0x3f, 0xd1, 0xeb, 0x35, 0x8d, 0xf1, 0x7b, 0x9f, 0x07, 0xb4, 0x81, 0xe4, 0xdd, 0x56, 0x31, 0x1f, 0xb7, 0xbe, 0xbd, 0xf2, 0xd8, 0x64, 0xa8, 0x4e, 0xb8, 0xa1, 0x4d, 0x56, 0x58, 0x74, 0xbc, 0x38, 0xba, 0x76, 0xb9, 0x23, 0x58, 0x50, 0x49, 0x4a, 0xce, 0xb5, 0x76, 0x6b, 0x15, 0x40, 0xbb, 0xbf, 0x27, 0x03, 0xf5, 0x50, 0x09, 0xbb, 0x07, 0x23, 0x23, 0x2e, 0x55, 0x07, 0x51, 0x91, 0xa0, 0xee, 0xfe, 0xf5, 0xb1, 0x74, 0x51, 0xbc, 0x7f, 0x78, 0x51, 0x12, 0x8b, 0xde, 0x71, 0x22, 0x6a, 0x0b, 0x1c, 0xb6, 0x33, 0x5f, 0x72, 0xaf, 0xeb, 0x26, 0xb0, 0x8b, 0x00, 0x11, 0x01, 0x00, 0x01, 0x00, 0x07, 0xfb, 0x07, 0x33, 0xd7, 0xc9, 0x9a, 0xd6, 0x6a, 0xb8, 0xc8, 0xa1, 0xa4, 0x79, 0xae, 0x35, 0x20, 0xb2, 0x99, 0x5d, 0xd3, 0x6b, 0x2f, 0x58, 0xc8, 0xa4, 0xd8, 0x92, 0x0b, 0xca, 0x9e, 0x50, 0x5c, 0x92, 0xe8, 0xd3, 0xf0, 0xbd, 0x1d, 0x34, 0xf4, 0x68, 0x09, 0xcf, 0x7d, 0xb4, 0x66, 0x7c, 0x48, 0x08, 0xdb, 0xa3, 0x84, 0x5d, 0x4b, 0x2a, 0xe5, 0x97, 0x9e, 0x8a, 0xdb, 0xad, 0x50, 0xcc, 0x14, 0xec, 0xde, 0xc3, 0xce, 0xc2, 0xa0, 0xd4, 0x1d, 0x6c, 0x42, 0xf8, 0x69, 0xb9, 0xc6, 0x25, 0x21, 0x11, 0xa8, 0xb8, 0x27, 0x37, 0xb6, 0x6e, 0x10, 0x8c, 0xcf, 0x70, 0x7b, 0xfd, 0x57, 0xa5, 0x36, 0xba, 0x18, 0x0d, 0xc6, 0xb8, 0x5a, 0xb3, 0x26, 0x5d, 0xec, 0x36, 0x79, 0xe3, 0x13, 0x59, 0x3e, 0x37, 0xcd, 0xcb, 0x77, 0x70, 0x9d, 0x8a, 0xc6, 0xf7, 0x23, 0x97, 0x3b, 0xbf, 0xdb, 0xc4, 0x86, 0x84, 0x03, 0x81, 0x7f, 0xf0, 0x54, 0xff, 0x92, 0x7d, 0x9b, 0x88, 0x9b, 0x9f, 0x28, 0x8b, 0xcd, 0x85, 0x76, 0xa6, 0xdb, 0x52, 0x9c, 0x8f, 0x33, 0x9e, 0x34, 0x79, 0x6f, 0xdd, 0x99, 0x2e, 0x3d, 0xca, 0x76, 0x74, 0x09, 0x71, 0x67, 0x7e, 0xbc, 0x51, 0x9f, 0x6b, 0x25, 0x48, 0x68, 0x7d, 0x3d, 0xa8, 0x8b, 0xfe, 0x92, 0x42, 0x9f, 0xc4, 0x93, 0xc6, 0x76, 0xc2, 0xb7, 0xd9, 0xc2, 0x75, 0xd1, 0xad, 0xe4, 0x28, 0x36, 0x26, 0x15, 0xb8, 0xa1, 0x4d, 0x4d, 0xab, 0xe5, 0xf1, 0x80, 0x41, 0x82, 0x56, 0xe4, 0x9c, 0x74, 0x2e, 0xee, 0xbf, 0xba, 0x8b, 0xee, 0xb2, 0x05, 0xa2, 0xb9, 0x08, 0x2f, 0x8f, 0xfe, 0xae, 0x64, 0x8e, 0x03, 0x6d, 0xc4, 0x79, 0x7b, 0xe3, 0xe1, 0x33, 0x5b, 0x1c, 0x71, 0x0e, 0xce, 0x6d, 0xa2, 0x92, 0x41, 0xb6, 0xde, 0xfd, 0x1a, 0x5d, 0x48, 0x3c, 0x19, 0x14, 0x58, 0x49, 0x04, 0x00, 0xb6, 0x95, 0xff, 0x19, 0x0d, 0x79, 0xd7, 0x72, 0xd0, 0x51, 0xe4, 0x87, 0x08, 0xff, 0xce, 0x03, 0x36, 0x9b, 0x05, 0xb5, 0x2b, 0xf4, 0x63, 0xaa, 0xf0, 0x78, 0xe0, 0x7d, 0xc3, 0xd0, 0xc2, 0xe5, 0x61, 0x88, 0x36, 0x22, 0x30, 0x6d, 0xe4, 0x9b, 0xf9, 0x80, 0x70, 0xd8, 0xd2, 0xa4, 0xf4, 0x8f, 0xd7, 0x3b, 0xe5, 0x69, 0xef, 0xe9, 0x61, 0x50, 0x8d, 0x36, 0xef, 0x77, 0x84, 0xff, 0xa9, 0x92, 0x83, 0xb1, 0xd8, 0x65, 0x4f, 0x4b, 0x62, 0xb4, 0x34, 0x03, 0xc4, 0x4b, 0x81, 0xba, 0xa3, 0x37, 0xe2, 0xb8, 0x06, 0xcf, 0x40, 0x8b, 0x7a, 0x4b, 0x03, 0xd7, 0xfa, 0xac, 0xbc, 0x73, 0x60, 0x8d, 0x1c, 0x32, 0xe5, 0x58, 0x41, 0x86, 0x7e, 0x5e, 0x1f, 0x0e, 0x3d, 0x53, 0x42, 0xf0, 0x2f, 0x7e, 0x28, 0x9a, 0x76, 0x40, 0xd8, 0x6a, 0x64, 0x76, 0x57, 0x69, 0xfe, 0x64, 0x68, 0x31, 0x70, 0x79, 0x04, 0x00, 0xc2, 0xd8, 0xc8, 0x85, 0xb0, 0x0a, 0x1b, 0xea, 0x8e, 0x06, 0xa7, 0x1a, 0x38, 0x4d, 0xb4, 0x6f, 0x2e, 0x90, 0x20, 0x7d, 0xfb, 0xf2, 0x4f, 0xd5, 0x5b, 0xbf, 0x7c, 0x81, 0x15, 0x3c, 0x4b, 0xfa, 0x21, 0xb0, 0xc3, 0x46, 0xb1, 0x4f, 0x25, 0xe8, 0xaf, 0x2e, 0x0d, 0xe0, 0xeb, 0xb1, 0x96, 0x06, 0xa3, 0x0c, 0xb7, 0x35, 0xaa, 0xbd, 0x6d, 0x55, 0x7f, 0xc4, 0x07, 0xd0, 0x1d, 0x1f, 0x67, 0x95, 0x73, 0x86, 0xba, 0x67, 0xcc, 0xad, 0x6a, 0xf3, 0x97, 0xa1, 0xf6, 0x65, 0xfa, 0xaa, 0xeb, 0x24, 0xd9, 0xb2, 0x30, 0x63, 0xa3, 0xdc, 0x9e, 0x2f, 0x89, 0xf6, 0xe9, 0x52, 0x20, 0x7f, 0x72, 0x82, 0x9a, 0x9f, 0xa0, 0x1d, 0xf6, 0x18, 0xe1, 0xfb, 0x48, 0xab, 0xf3, 0x46, 0x34, 0x4b, 0x4e, 0x8a, 0x31, 0x48, 0xd3, 0x3d, 0x74, 0x31, 0x30, 0xf8, 0x63, 0x7c, 0x47, 0xf7, 0x3c, 0x92, 0xd0, 0x23, 0x03, 0xfe, 0x2a, 0xe3, 0xab, 0x31, 0x74, 0x13, 0x51, 0xc4, 0xc0, 0x5e, 0xb5, 0xec, 0xac, 0x3c, 0xcc, 0xc6, 0xc7, 0x6a, 0x8c, 0xe3, 0xb1, 0x81, 0x06, 0xc5, 0x9b, 0xd2, 0x26, 0xdc, 0x0c, 0xde, 0x67, 0x6e, 0xcb, 0x10, 0x0d, 0x01, 0x23, 0x91, 0x2c, 0x68, 0x90, 0x71, 0x9a, 0x3d, 0xb7, 0xc4, 0xd2, 0x64, 0x18, 0xb5, 0x61, 0xd1, 0x77, 0x0a, 0xd5, 0x4e, 0xda, 0xcb, 0x57, 0x65, 0x8f, 0xb7, 0xac, 0x3d, 0x5a, 0x41, 0x64, 0x87, 0xc5, 0xb8, 0x4d, 0x86, 0x11, 0x9d, 0xaf, 0xc0, 0x97, 0x67, 0x9e, 0xd6, 0xab, 0x7e, 0xb7, 0xc2, 0x2e, 0x1e, 0xa7, 0x15, 0x63, 0xe7, 0x2f, 0x83, 0x13, 0xcf, 0x96, 0xd9, 0x14, 0xed, 0x1d, 0x45, 0xa4, 0x46, 0x83, 0x0d, 0x47, 0xb3, 0x1a, 0xb4, 0xef, 0x45, 0xbb, 0xa7, 0xf3, 0xae, 0x12, 0x6f, 0x40, 0xaa, 0xfd, 0xd2, 0x68, 0x80, 0xc8, 0xdb, 0x60, 0x89, 0xb6, 0x86, 0x7f, 0x2a, 0x90, 0x1a, 0x23, 0x91, 0xd2, 0x12, 0xb2, 0x65, 0x48, 0xf6, 0xa8, 0xee, 0x1e, 0x82, 0xc1, 0x81, 0x5d, 0x52}
	tag05Body2 = []byte{0x04, 0x5a, 0xfc, 0x9e, 0x70, 0x01, 0x08, 0x00, 0xc0, 0x9e, 0x6d, 0x1a, 0xba, 0xef, 0xab, 0xa8, 0x12, 0xfb, 0x1d, 0xd0, 0x0f, 0xa3, 0xe5, 0x8a, 0x7c, 0xbe, 0x3b, 0x84, 0x73, 0x73, 0x07, 0xa7, 0x42, 0xdf, 0x84, 0x8b, 0x59, 0x2d, 0x54, 0x1a, 0x33, 0xd1, 0x3d, 0x01, 0xf8, 0x22, 0xa7, 0xfa, 0xff, 0x7b, 0x2a, 0x23, 0x65, 0xe9, 0xc0, 0x93, 0xb4, 0xca, 0x3b, 0x35, 0xcb, 0x75, 0x8a, 0xa1, 0xb6, 0xf5, 0x25, 0x0a, 0x5c, 0x75, 0xf3, 0x35, 0xd9, 0x6e, 0x0e, 0xeb, 0xd9, 0x38, 0xcc, 0xae, 0x45, 0x50, 0xc6, 0x01, 0x40, 0x34, 0xce, 0x8f, 0x82, 0xc9, 0x41, 0xe2, 0xdf, 0xc5, 0x3b, 0x49, 0x07, 0xdb, 0x13, 0x1e, 0x34, 0x83, 0x30, 0xdc, 0xd1, 0x97, 0xeb, 0x07, 0xef, 0xbb, 0x02, 0x9b, 0x46, 0xe3, 0xa7, 0x23, 0x03, 0xad, 0x1c, 0xb8, 0x01, 0xa7, 0x9f, 0x3f, 0x4b, 0x48, 0x98, 0x31, 0xa1, 0xb9, 0xd2, 0x94, 0xce, 0x89, 0x35, 0xb3, 0x01, 0x85, 0x3f, 0xd7, 0x0d, 0x71, 0x65, 0x6b, 0x08, 0x67, 0x83, 0x2d, 0x8e, 0x3b, 0xce, 0x2c, 0x3f, 0x8c, 0x08, 0x8a, 0xb3, 0xa1, 0x97, 0xe8, 0x33, 0xde, 0x25, 0x9b, 0x1e, 0x75, 0x90, 0xf8, 0xa4, 0xee, 0xef, 0x8d, 0x3e, 0xdb, 0x8d, 0xa7, 0x6f, 0x10, 0xf6, 0x83, 0x7d, 0xd8, 0x0e, 0xf8, 0xb7, 0x37, 0x80, 0x50, 0xc5, 0xb1, 0x82, 0xee, 0xfa, 0xd5, 0x60, 0x20, 0xc6, 0xb8, 0x3f, 0x01, 0xf0, 0x6d, 0x1d, 0xcc, 0x3b, 0xdb, 0xe3, 0xbb, 0x32, 0x30, 0x40, 0xd3, 0xe8, 0xe6, 0xe5, 0x2f, 0x04, 0x67, 0x13, 0x03, 0xaa, 0x21, 0x81, 0x6f, 0x87, 0x72, 0x99, 0x5f, 0x8f, 0x57, 0x94, 0x22, 0xd4, 0xa2, 0x97, 0x33, 0x47, 0x22, 0xaf, 0x37, 0xdd, 0x67, 0x07, 0x96, 0xfb, 0x3e, 0x37, 0xca, 0x95, 0x22, 0x9c, 0x46, 0x2e, 0xd1, 0x67, 0x91, 0xf1, 0xd6, 0xf5, 0x14, 0x89, 0x57, 0xa1, 0xfb, 0x00, 0x11, 0x01, 0x00, 0x01, 0x00, 0x07, 0xfe, 0x25, 0xde, 0xc0, 0x0a, 0xb2, 0x58, 0x2e, 0xc2, 0xa3, 0xc0, 0xb5, 0x72, 0xd3, 0xb0, 0x60, 0x8f, 0xe2, 0xc8, 0xb0, 0x00, 0xf1, 0x85, 0xdb, 0x2a, 0x5a, 0x6e, 0x81, 0xab, 0xb8, 0x03, 0xbe, 0x76, 0x4c, 0x5b, 0xc6, 0x07, 0xde, 0x16, 0x4a, 0x3a, 0x82, 0x02, 0x60, 0x1d, 0x87, 0x8a, 0xf6, 0xae, 0xd3, 0xab, 0xb3, 0x0a, 0x77, 0x8f, 0x0b, 0x8b, 0x91, 0xe2, 0x0e, 0xbf, 0x43, 0xc0, 0x78, 0xe9, 0xcc, 0x6e, 0xe4, 0x06, 0x20, 0xb6, 0x17, 0x1f, 0xe8, 0x46, 0xe2, 0x37, 0x1a, 0xbd, 0x87, 0x23, 0x16, 0x0e, 0xa5, 0xa2, 0x8a, 0x66, 0x47, 0xaa, 0xab, 0x1d, 0xba, 0x5b, 0x84, 0xed, 0x8a, 0x2c, 0xd0, 0x14, 0x73, 0x44, 0x23, 0x30, 0xfc, 0x69, 0x34, 0xfd, 0xcb, 0x3d, 0x8a, 0x1a, 0x7d, 0xfb, 0xfb, 0x6f, 0x4e, 0x52, 0xee, 0x65, 0x3e, 0x6e, 0xfb, 0xa2, 0x02, 0x31, 0xf9, 0x8d, 0x66, 0x7e, 0x0c, 0x76, 0x9f, 0x7f, 0x62, 0xbf, 0x53, 0x69, 0x6d, 0xf6, 0x92, 0xbe, 0xdc, 0x51, 0x96, 0xcb, 0x8c, 0x84, 0x89, 0x7e, 0xb8, 0x44, 0x85, 0x60, 0xc5, 0xee, 0xa3, 0x0e, 0x12, 0xa4, 0x96, 0x77, 0xdb, 0x99, 0x48, 0xa4, 0xd4, 0x40, 0xfa, 0xd5, 0x34, 0x76, 0xdf, 0x65, 0x28, 0x4a, 0x24, 0xf9, 0x3e, 0x52, 0xed, 0xe3, 0x0c, 0x36, 0x1d, 0x3e, 0x0e, 0xc8, 0x8c, 0xe7, 0x1d, 0x17, 0xa2, 0xba, 0x09, 0x48, 0xf2, 0x13, 0x34, 0x4c, 0x11, 0x15, 0x4c, 0x49, 0x31, 0x83, 0x35, 0xa4, 0x9a, 0x11, 0x02, 0x1d, 0xbf, 0x2c, 0xc1, 0xd3, 0x10, 0xab, 0xc4, 0xcd, 0xbe, 0xea, 0xa7, 0x71, 0x9d, 0x70, 0x3e, 0xeb, 0x0e, 0x11, 0x52, 0x1d, 0x18, 0x65, 0x47, 0xda, 0x77, 0x32, 0x75, 0x3a, 0x17, 0x7c, 0x84, 0xd3, 0x9c, 0xc9, 0x7b, 0xe8, 0x91, 0xb1, 0xbf, 0x67, 0x13, 0x0c, 0x17, 0xdb, 0x55, 0xb1, 0x04, 0x00, 0xc8, 0x6e, 0xe8, 0xa3, 0xab, 0x46, 0x6e, 0x4b, 0x8b, 0xa2, 0xb7, 0xdb, 0x34, 0xf8, 0x0e, 0xd8, 0x3e, 0xba, 0xb5, 0xf8, 0xd2, 0x08, 0xd6, 0xc0, 0x53, 0xe0, 0xf9, 0xdb, 0xca, 0x84, 0xa6, 0xf2, 0xae, 0xe5, 0x4d, 0x48, 0x73, 0xf7, 0xe3, 0x0b, 0xb4, 0x23, 0xc3, 0xb1, 0xad, 0xf0, 0x09, 0x5a, 0xfc, 0xe5, 0x1d, 0xfc, 0x63, 0xeb, 0xaf, 0x71, 0xc6, 0x49, 0xad, 0x61, 0x3d, 0x32, 0x96, 0x65, 0x89, 0x4e, 0xce, 0xfa, 0x12, 0x7f, 0xd7, 0x0e, 0x82, 0xed, 0xf0, 0xfb, 0x8e, 0x9f, 0x54, 0x5c, 0xd8, 0x44, 0xd5, 0x26, 0x74, 0x9e, 0x72, 0x34, 0xb2, 0x3e, 0xf4, 0xbd, 0x7e, 0x38, 0x0f, 0xfe, 0x02, 0xed, 0x2f, 0x27, 0x6f, 0xeb, 0xbc, 0x19, 0xba, 0xc9, 0xf1, 0x9e, 0x81, 0xb0, 0x24, 0x29, 0x3f, 0x0c, 0x73, 0x9d, 0xba, 0x26, 0x49, 0x34, 0x06, 0xbb, 0xa9, 0x73, 0x6d, 0x1d, 0x98, 0x53, 0x04, 0x00, 0xf6, 0x04, 0xea, 0xc7, 0xd0, 0x1a, 0x82, 0xeb, 0x67, 0xde, 0x12, 0x7a, 0xff, 0x5a, 0x4b, 0x9d, 0x21, 0x9e, 0x3d, 0xd2, 0x8e, 0x76, 0xc4, 0x6d, 0x9b, 0x0e, 0xf9, 0x11, 0xd8, 0x91, 0x30, 0x07, 0x8c, 0x98, 0x52, 0x4c, 0x2a, 0x66, 0x1e, 0x50, 0xf5, 0xf3, 0xf9, 0xef, 0xcc, 0x9f, 0x42, 0x86, 0x1e, 0x5f, 0xbb, 0x33, 0x2e, 0x31, 0x2a, 0x0d, 0x02, 0x81, 0x7e, 0x06, 0x68, 0x1b, 0xe7, 0x8f, 0x3a, 0xe4, 0x76, 0xfd, 0xbe, 0xa4, 0x8f, 0x75, 0x71, 0xad, 0x9a, 0x7b, 0x3f, 0x71, 0x09, 0x3f, 0xc2, 0xd1, 0x69, 0x9f, 0x99, 0x22, 0x4a, 0xcc, 0x2d, 0x9a, 0xb9, 0x63, 0x23, 0x37, 0xb2, 0xad, 0x76, 0xf2, 0x01, 0x1a, 0x33, 0xa1, 0xf2, 0xee, 0x72, 0x68, 0xa7, 0x2a, 0xfd, 0xaa, 0xac, 0x2a, 0x42, 0xc0, 0x98, 0x41, 0x49, 0x74, 0x4f, 0x7e, 0x8a, 0x31, 0x99, 0x2b, 0xca, 0x8f, 0x7a, 0xb9, 0x03, 0xff, 0x52, 0x22, 0x89, 0x52, 0x98, 0x1c, 0xf5, 0xd5, 0x74, 0x96, 0x77, 0x1c, 0xc5, 0x43, 0x53, 0x7e, 0x2b, 0x4f, 0x37, 0xdd, 0x30, 0x55, 0xdd, 0x26, 0xce, 0xfe, 0x2b, 0x65, 0x77, 0x26, 0x1e, 0x4d, 0xd2, 0xda, 0xcb, 0x41, 0x1b, 0x22, 0xb9, 0xb2, 0x7c, 0x3f, 0x2a, 0xd9, 0x1c, 0xc5, 0xea, 0xad, 0xb6, 0xae, 0xd6, 0x05, 0x62, 0x47, 0x0d, 0xe2, 0x1b, 0x23, 0x4b, 0xbd, 0xeb, 0x84, 0xca, 0x6d, 0xdb, 0x09, 0x51, 0xb0, 0x31, 0x76, 0x6b, 0x46, 0xd8, 0xfd, 0xcc, 0xcb, 0x8e, 0x40, 0x60, 0x1f, 0xcb, 0x52, 0x61, 0xcd, 0xd3, 0x7d, 0xda, 0x6d, 0x62, 0x2d, 0x69, 0x68, 0x37, 0x5d, 0xc2, 0x8e, 0xb9, 0x70, 0xe5, 0x58, 0xb6, 0x65, 0x1d, 0x6e, 0x93, 0x20, 0x50, 0x75, 0x2b, 0xc6, 0x88, 0x4b, 0x23, 0x42, 0xbe, 0x2f, 0x44, 0xc6, 0xd1, 0xd5, 0x9a, 0xbd, 0x27, 0x5b, 0x01, 0x2e, 0x3f, 0xf2, 0x3d, 0x0c}
	tag05Body3 = []byte{0x04, 0x5b, 0x1a, 0x4e, 0x1d, 0x16, 0x09, 0x2b, 0x06, 0x01, 0x04, 0x01, 0xda, 0x47, 0x0f, 0x01, 0x01, 0x07, 0x40, 0xc6, 0xae, 0xd8, 0x56, 0x62, 0x34, 0x73, 0xe7, 0xf1, 0x86, 0xff, 0x5f, 0x09, 0xdd, 0xd2, 0xc2, 0xb5, 0x48, 0xbd, 0x78, 0x94, 0x90, 0xa8, 0xd2, 0xfd, 0x9c, 0xfc, 0xc6, 0x69, 0x15, 0xfb, 0x86, 0x00, 0x00, 0xff, 0x50, 0x5e, 0xcc, 0x13, 0x31, 0x23, 0x59, 0x49, 0xc2, 0xcc, 0x48, 0x1d, 0x7c, 0xe8, 0x39, 0x85, 0xac, 0x36, 0x2f, 0x76, 0xff, 0x5a, 0xe5, 0xd6, 0x09, 0x68, 0xc6, 0xe7, 0xde, 0xcb, 0x00, 0x5c, 0x10, 0x55}
)

const (
	tag05Redult1 = `Secret-Key Packet (tag 5) (938 bytes)
	04 5a fa 85 65 01 08 00 8a f8 48 e4 c2 21 23 c4 47 e9 84 13 4d 80 88 ad 28 9e 34 32 eb b2 53 9e d2 b3 64 c4 f4 e9 f7 c0 5d 8a e6 01 e1 4f 04 2b f2 08 38 c1 4d 65 49 3b 25 22 f1 bf 18 0d 31 63 9e d5 04 cf 4c 99 57 e1 a8 2e a0 cb 3f ac 25 00 85 00 3d aa 61 e2 ba e5 e3 f0 1f f2 36 9a 97 12 7b 0b 91 b0 c2 99 b4 5a 6a 3a 40 ec 76 76 a8 19 4d 8b ed e0 b5 1f b3 ef 65 fa 55 2e f2 a6 d4 2d 7a 76 d9 7a 3a 3e e9 39 d3 20 b0 3b 7d c8 bb ff a0 c8 88 82 eb cf 75 f4 ac 07 41 76 f6 66 fb 24 87 00 8e 94 73 d7 2c 3e 98 fe 46 a0 2e 75 3f d1 eb 35 8d f1 7b 9f 07 b4 81 e4 dd 56 31 1f b7 be bd f2 d8 64 a8 4e b8 a1 4d 56 58 74 bc 38 ba 76 b9 23 58 50 49 4a ce b5 76 6b 15 40 bb bf 27 03 f5 50 09 bb 07 23 23 2e 55 07 51 91 a0 ee fe f5 b1 74 51 bc 7f 78 51 12 8b de 71 22 6a 0b 1c b6 33 5f 72 af eb 26 b0 8b 00 11 01 00 01 00 07 fb 07 33 d7 c9 9a d6 6a b8 c8 a1 a4 79 ae 35 20 b2 99 5d d3 6b 2f 58 c8 a4 d8 92 0b ca 9e 50 5c 92 e8 d3 f0 bd 1d 34 f4 68 09 cf 7d b4 66 7c 48 08 db a3 84 5d 4b 2a e5 97 9e 8a db ad 50 cc 14 ec de c3 ce c2 a0 d4 1d 6c 42 f8 69 b9 c6 25 21 11 a8 b8 27 37 b6 6e 10 8c cf 70 7b fd 57 a5 36 ba 18 0d c6 b8 5a b3 26 5d ec 36 79 e3 13 59 3e 37 cd cb 77 70 9d 8a c6 f7 23 97 3b bf db c4 86 84 03 81 7f f0 54 ff 92 7d 9b 88 9b 9f 28 8b cd 85 76 a6 db 52 9c 8f 33 9e 34 79 6f dd 99 2e 3d ca 76 74 09 71 67 7e bc 51 9f 6b 25 48 68 7d 3d a8 8b fe 92 42 9f c4 93 c6 76 c2 b7 d9 c2 75 d1 ad e4 28 36 26 15 b8 a1 4d 4d ab e5 f1 80 41 82 56 e4 9c 74 2e ee bf ba 8b ee b2 05 a2 b9 08 2f 8f fe ae 64 8e 03 6d c4 79 7b e3 e1 33 5b 1c 71 0e ce 6d a2 92 41 b6 de fd 1a 5d 48 3c 19 14 58 49 04 00 b6 95 ff 19 0d 79 d7 72 d0 51 e4 87 08 ff ce 03 36 9b 05 b5 2b f4 63 aa f0 78 e0 7d c3 d0 c2 e5 61 88 36 22 30 6d e4 9b f9 80 70 d8 d2 a4 f4 8f d7 3b e5 69 ef e9 61 50 8d 36 ef 77 84 ff a9 92 83 b1 d8 65 4f 4b 62 b4 34 03 c4 4b 81 ba a3 37 e2 b8 06 cf 40 8b 7a 4b 03 d7 fa ac bc 73 60 8d 1c 32 e5 58 41 86 7e 5e 1f 0e 3d 53 42 f0 2f 7e 28 9a 76 40 d8 6a 64 76 57 69 fe 64 68 31 70 79 04 00 c2 d8 c8 85 b0 0a 1b ea 8e 06 a7 1a 38 4d b4 6f 2e 90 20 7d fb f2 4f d5 5b bf 7c 81 15 3c 4b fa 21 b0 c3 46 b1 4f 25 e8 af 2e 0d e0 eb b1 96 06 a3 0c b7 35 aa bd 6d 55 7f c4 07 d0 1d 1f 67 95 73 86 ba 67 cc ad 6a f3 97 a1 f6 65 fa aa eb 24 d9 b2 30 63 a3 dc 9e 2f 89 f6 e9 52 20 7f 72 82 9a 9f a0 1d f6 18 e1 fb 48 ab f3 46 34 4b 4e 8a 31 48 d3 3d 74 31 30 f8 63 7c 47 f7 3c 92 d0 23 03 fe 2a e3 ab 31 74 13 51 c4 c0 5e b5 ec ac 3c cc c6 c7 6a 8c e3 b1 81 06 c5 9b d2 26 dc 0c de 67 6e cb 10 0d 01 23 91 2c 68 90 71 9a 3d b7 c4 d2 64 18 b5 61 d1 77 0a d5 4e da cb 57 65 8f b7 ac 3d 5a 41 64 87 c5 b8 4d 86 11 9d af c0 97 67 9e d6 ab 7e b7 c2 2e 1e a7 15 63 e7 2f 83 13 cf 96 d9 14 ed 1d 45 a4 46 83 0d 47 b3 1a b4 ef 45 bb a7 f3 ae 12 6f 40 aa fd d2 68 80 c8 db 60 89 b6 86 7f 2a 90 1a 23 91 d2 12 b2 65 48 f6 a8 ee 1e 82 c1 81 5d 52
	Version: 4 (current)
		04
	Public-Key
		Public key creation time: 2018-05-15T06:59:49Z
			5a fa 85 65
		Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
			01
		RSA public modulus n (2048 bits)
			8a f8 48 e4 c2 21 23 c4 47 e9 84 13 4d 80 88 ad 28 9e 34 32 eb b2 53 9e d2 b3 64 c4 f4 e9 f7 c0 5d 8a e6 01 e1 4f 04 2b f2 08 38 c1 4d 65 49 3b 25 22 f1 bf 18 0d 31 63 9e d5 04 cf 4c 99 57 e1 a8 2e a0 cb 3f ac 25 00 85 00 3d aa 61 e2 ba e5 e3 f0 1f f2 36 9a 97 12 7b 0b 91 b0 c2 99 b4 5a 6a 3a 40 ec 76 76 a8 19 4d 8b ed e0 b5 1f b3 ef 65 fa 55 2e f2 a6 d4 2d 7a 76 d9 7a 3a 3e e9 39 d3 20 b0 3b 7d c8 bb ff a0 c8 88 82 eb cf 75 f4 ac 07 41 76 f6 66 fb 24 87 00 8e 94 73 d7 2c 3e 98 fe 46 a0 2e 75 3f d1 eb 35 8d f1 7b 9f 07 b4 81 e4 dd 56 31 1f b7 be bd f2 d8 64 a8 4e b8 a1 4d 56 58 74 bc 38 ba 76 b9 23 58 50 49 4a ce b5 76 6b 15 40 bb bf 27 03 f5 50 09 bb 07 23 23 2e 55 07 51 91 a0 ee fe f5 b1 74 51 bc 7f 78 51 12 8b de 71 22 6a 0b 1c b6 33 5f 72 af eb 26 b0 8b
		RSA public encryption exponent e (17 bits)
			01 00 01
	Secret-Key (the secret-key data is not encrypted.)
		RSA secret exponent d (2043 bits)
			07 33 d7 c9 9a d6 6a b8 c8 a1 a4 79 ae 35 20 b2 99 5d d3 6b 2f 58 c8 a4 d8 92 0b ca 9e 50 5c 92 e8 d3 f0 bd 1d 34 f4 68 09 cf 7d b4 66 7c 48 08 db a3 84 5d 4b 2a e5 97 9e 8a db ad 50 cc 14 ec de c3 ce c2 a0 d4 1d 6c 42 f8 69 b9 c6 25 21 11 a8 b8 27 37 b6 6e 10 8c cf 70 7b fd 57 a5 36 ba 18 0d c6 b8 5a b3 26 5d ec 36 79 e3 13 59 3e 37 cd cb 77 70 9d 8a c6 f7 23 97 3b bf db c4 86 84 03 81 7f f0 54 ff 92 7d 9b 88 9b 9f 28 8b cd 85 76 a6 db 52 9c 8f 33 9e 34 79 6f dd 99 2e 3d ca 76 74 09 71 67 7e bc 51 9f 6b 25 48 68 7d 3d a8 8b fe 92 42 9f c4 93 c6 76 c2 b7 d9 c2 75 d1 ad e4 28 36 26 15 b8 a1 4d 4d ab e5 f1 80 41 82 56 e4 9c 74 2e ee bf ba 8b ee b2 05 a2 b9 08 2f 8f fe ae 64 8e 03 6d c4 79 7b e3 e1 33 5b 1c 71 0e ce 6d a2 92 41 b6 de fd 1a 5d 48 3c 19 14 58 49
		RSA secret prime value p (1024 bits)
			b6 95 ff 19 0d 79 d7 72 d0 51 e4 87 08 ff ce 03 36 9b 05 b5 2b f4 63 aa f0 78 e0 7d c3 d0 c2 e5 61 88 36 22 30 6d e4 9b f9 80 70 d8 d2 a4 f4 8f d7 3b e5 69 ef e9 61 50 8d 36 ef 77 84 ff a9 92 83 b1 d8 65 4f 4b 62 b4 34 03 c4 4b 81 ba a3 37 e2 b8 06 cf 40 8b 7a 4b 03 d7 fa ac bc 73 60 8d 1c 32 e5 58 41 86 7e 5e 1f 0e 3d 53 42 f0 2f 7e 28 9a 76 40 d8 6a 64 76 57 69 fe 64 68 31 70 79
		RSA secret prime value q (p < q) (1024 bits)
			c2 d8 c8 85 b0 0a 1b ea 8e 06 a7 1a 38 4d b4 6f 2e 90 20 7d fb f2 4f d5 5b bf 7c 81 15 3c 4b fa 21 b0 c3 46 b1 4f 25 e8 af 2e 0d e0 eb b1 96 06 a3 0c b7 35 aa bd 6d 55 7f c4 07 d0 1d 1f 67 95 73 86 ba 67 cc ad 6a f3 97 a1 f6 65 fa aa eb 24 d9 b2 30 63 a3 dc 9e 2f 89 f6 e9 52 20 7f 72 82 9a 9f a0 1d f6 18 e1 fb 48 ab f3 46 34 4b 4e 8a 31 48 d3 3d 74 31 30 f8 63 7c 47 f7 3c 92 d0 23
		RSA u, the multiplicative inverse of p, mod q (1022 bits)
			2a e3 ab 31 74 13 51 c4 c0 5e b5 ec ac 3c cc c6 c7 6a 8c e3 b1 81 06 c5 9b d2 26 dc 0c de 67 6e cb 10 0d 01 23 91 2c 68 90 71 9a 3d b7 c4 d2 64 18 b5 61 d1 77 0a d5 4e da cb 57 65 8f b7 ac 3d 5a 41 64 87 c5 b8 4d 86 11 9d af c0 97 67 9e d6 ab 7e b7 c2 2e 1e a7 15 63 e7 2f 83 13 cf 96 d9 14 ed 1d 45 a4 46 83 0d 47 b3 1a b4 ef 45 bb a7 f3 ae 12 6f 40 aa fd d2 68 80 c8 db 60 89 b6 86
		Checksum
			7f 2a
	Unknown data (18 bytes)
		90 1a 23 91 d2 12 b2 65 48 f6 a8 ee 1e 82 c1 81 5d 52
`
	tag05Redult2 = `Secret-Key Packet (tag 5) (920 bytes)
	04 5a fc 9e 70 01 08 00 c0 9e 6d 1a ba ef ab a8 12 fb 1d d0 0f a3 e5 8a 7c be 3b 84 73 73 07 a7 42 df 84 8b 59 2d 54 1a 33 d1 3d 01 f8 22 a7 fa ff 7b 2a 23 65 e9 c0 93 b4 ca 3b 35 cb 75 8a a1 b6 f5 25 0a 5c 75 f3 35 d9 6e 0e eb d9 38 cc ae 45 50 c6 01 40 34 ce 8f 82 c9 41 e2 df c5 3b 49 07 db 13 1e 34 83 30 dc d1 97 eb 07 ef bb 02 9b 46 e3 a7 23 03 ad 1c b8 01 a7 9f 3f 4b 48 98 31 a1 b9 d2 94 ce 89 35 b3 01 85 3f d7 0d 71 65 6b 08 67 83 2d 8e 3b ce 2c 3f 8c 08 8a b3 a1 97 e8 33 de 25 9b 1e 75 90 f8 a4 ee ef 8d 3e db 8d a7 6f 10 f6 83 7d d8 0e f8 b7 37 80 50 c5 b1 82 ee fa d5 60 20 c6 b8 3f 01 f0 6d 1d cc 3b db e3 bb 32 30 40 d3 e8 e6 e5 2f 04 67 13 03 aa 21 81 6f 87 72 99 5f 8f 57 94 22 d4 a2 97 33 47 22 af 37 dd 67 07 96 fb 3e 37 ca 95 22 9c 46 2e d1 67 91 f1 d6 f5 14 89 57 a1 fb 00 11 01 00 01 00 07 fe 25 de c0 0a b2 58 2e c2 a3 c0 b5 72 d3 b0 60 8f e2 c8 b0 00 f1 85 db 2a 5a 6e 81 ab b8 03 be 76 4c 5b c6 07 de 16 4a 3a 82 02 60 1d 87 8a f6 ae d3 ab b3 0a 77 8f 0b 8b 91 e2 0e bf 43 c0 78 e9 cc 6e e4 06 20 b6 17 1f e8 46 e2 37 1a bd 87 23 16 0e a5 a2 8a 66 47 aa ab 1d ba 5b 84 ed 8a 2c d0 14 73 44 23 30 fc 69 34 fd cb 3d 8a 1a 7d fb fb 6f 4e 52 ee 65 3e 6e fb a2 02 31 f9 8d 66 7e 0c 76 9f 7f 62 bf 53 69 6d f6 92 be dc 51 96 cb 8c 84 89 7e b8 44 85 60 c5 ee a3 0e 12 a4 96 77 db 99 48 a4 d4 40 fa d5 34 76 df 65 28 4a 24 f9 3e 52 ed e3 0c 36 1d 3e 0e c8 8c e7 1d 17 a2 ba 09 48 f2 13 34 4c 11 15 4c 49 31 83 35 a4 9a 11 02 1d bf 2c c1 d3 10 ab c4 cd be ea a7 71 9d 70 3e eb 0e 11 52 1d 18 65 47 da 77 32 75 3a 17 7c 84 d3 9c c9 7b e8 91 b1 bf 67 13 0c 17 db 55 b1 04 00 c8 6e e8 a3 ab 46 6e 4b 8b a2 b7 db 34 f8 0e d8 3e ba b5 f8 d2 08 d6 c0 53 e0 f9 db ca 84 a6 f2 ae e5 4d 48 73 f7 e3 0b b4 23 c3 b1 ad f0 09 5a fc e5 1d fc 63 eb af 71 c6 49 ad 61 3d 32 96 65 89 4e ce fa 12 7f d7 0e 82 ed f0 fb 8e 9f 54 5c d8 44 d5 26 74 9e 72 34 b2 3e f4 bd 7e 38 0f fe 02 ed 2f 27 6f eb bc 19 ba c9 f1 9e 81 b0 24 29 3f 0c 73 9d ba 26 49 34 06 bb a9 73 6d 1d 98 53 04 00 f6 04 ea c7 d0 1a 82 eb 67 de 12 7a ff 5a 4b 9d 21 9e 3d d2 8e 76 c4 6d 9b 0e f9 11 d8 91 30 07 8c 98 52 4c 2a 66 1e 50 f5 f3 f9 ef cc 9f 42 86 1e 5f bb 33 2e 31 2a 0d 02 81 7e 06 68 1b e7 8f 3a e4 76 fd be a4 8f 75 71 ad 9a 7b 3f 71 09 3f c2 d1 69 9f 99 22 4a cc 2d 9a b9 63 23 37 b2 ad 76 f2 01 1a 33 a1 f2 ee 72 68 a7 2a fd aa ac 2a 42 c0 98 41 49 74 4f 7e 8a 31 99 2b ca 8f 7a b9 03 ff 52 22 89 52 98 1c f5 d5 74 96 77 1c c5 43 53 7e 2b 4f 37 dd 30 55 dd 26 ce fe 2b 65 77 26 1e 4d d2 da cb 41 1b 22 b9 b2 7c 3f 2a d9 1c c5 ea ad b6 ae d6 05 62 47 0d e2 1b 23 4b bd eb 84 ca 6d db 09 51 b0 31 76 6b 46 d8 fd cc cb 8e 40 60 1f cb 52 61 cd d3 7d da 6d 62 2d 69 68 37 5d c2 8e b9 70 e5 58 b6 65 1d 6e 93 20 50 75 2b c6 88 4b 23 42 be 2f 44 c6 d1 d5 9a bd 27 5b 01 2e 3f f2 3d 0c
	Version: 4 (current)
		04
	Public-Key
		Public key creation time: 2018-05-16T21:11:12Z
			5a fc 9e 70
		Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
			01
		RSA public modulus n (2048 bits)
			c0 9e 6d 1a ba ef ab a8 12 fb 1d d0 0f a3 e5 8a 7c be 3b 84 73 73 07 a7 42 df 84 8b 59 2d 54 1a 33 d1 3d 01 f8 22 a7 fa ff 7b 2a 23 65 e9 c0 93 b4 ca 3b 35 cb 75 8a a1 b6 f5 25 0a 5c 75 f3 35 d9 6e 0e eb d9 38 cc ae 45 50 c6 01 40 34 ce 8f 82 c9 41 e2 df c5 3b 49 07 db 13 1e 34 83 30 dc d1 97 eb 07 ef bb 02 9b 46 e3 a7 23 03 ad 1c b8 01 a7 9f 3f 4b 48 98 31 a1 b9 d2 94 ce 89 35 b3 01 85 3f d7 0d 71 65 6b 08 67 83 2d 8e 3b ce 2c 3f 8c 08 8a b3 a1 97 e8 33 de 25 9b 1e 75 90 f8 a4 ee ef 8d 3e db 8d a7 6f 10 f6 83 7d d8 0e f8 b7 37 80 50 c5 b1 82 ee fa d5 60 20 c6 b8 3f 01 f0 6d 1d cc 3b db e3 bb 32 30 40 d3 e8 e6 e5 2f 04 67 13 03 aa 21 81 6f 87 72 99 5f 8f 57 94 22 d4 a2 97 33 47 22 af 37 dd 67 07 96 fb 3e 37 ca 95 22 9c 46 2e d1 67 91 f1 d6 f5 14 89 57 a1 fb
		RSA public encryption exponent e (17 bits)
			01 00 01
	Secret-Key (the secret-key data is not encrypted.)
		RSA secret exponent d (2046 bits)
			25 de c0 0a b2 58 2e c2 a3 c0 b5 72 d3 b0 60 8f e2 c8 b0 00 f1 85 db 2a 5a 6e 81 ab b8 03 be 76 4c 5b c6 07 de 16 4a 3a 82 02 60 1d 87 8a f6 ae d3 ab b3 0a 77 8f 0b 8b 91 e2 0e bf 43 c0 78 e9 cc 6e e4 06 20 b6 17 1f e8 46 e2 37 1a bd 87 23 16 0e a5 a2 8a 66 47 aa ab 1d ba 5b 84 ed 8a 2c d0 14 73 44 23 30 fc 69 34 fd cb 3d 8a 1a 7d fb fb 6f 4e 52 ee 65 3e 6e fb a2 02 31 f9 8d 66 7e 0c 76 9f 7f 62 bf 53 69 6d f6 92 be dc 51 96 cb 8c 84 89 7e b8 44 85 60 c5 ee a3 0e 12 a4 96 77 db 99 48 a4 d4 40 fa d5 34 76 df 65 28 4a 24 f9 3e 52 ed e3 0c 36 1d 3e 0e c8 8c e7 1d 17 a2 ba 09 48 f2 13 34 4c 11 15 4c 49 31 83 35 a4 9a 11 02 1d bf 2c c1 d3 10 ab c4 cd be ea a7 71 9d 70 3e eb 0e 11 52 1d 18 65 47 da 77 32 75 3a 17 7c 84 d3 9c c9 7b e8 91 b1 bf 67 13 0c 17 db 55 b1
		RSA secret prime value p (1024 bits)
			c8 6e e8 a3 ab 46 6e 4b 8b a2 b7 db 34 f8 0e d8 3e ba b5 f8 d2 08 d6 c0 53 e0 f9 db ca 84 a6 f2 ae e5 4d 48 73 f7 e3 0b b4 23 c3 b1 ad f0 09 5a fc e5 1d fc 63 eb af 71 c6 49 ad 61 3d 32 96 65 89 4e ce fa 12 7f d7 0e 82 ed f0 fb 8e 9f 54 5c d8 44 d5 26 74 9e 72 34 b2 3e f4 bd 7e 38 0f fe 02 ed 2f 27 6f eb bc 19 ba c9 f1 9e 81 b0 24 29 3f 0c 73 9d ba 26 49 34 06 bb a9 73 6d 1d 98 53
		RSA secret prime value q (p < q) (1024 bits)
			f6 04 ea c7 d0 1a 82 eb 67 de 12 7a ff 5a 4b 9d 21 9e 3d d2 8e 76 c4 6d 9b 0e f9 11 d8 91 30 07 8c 98 52 4c 2a 66 1e 50 f5 f3 f9 ef cc 9f 42 86 1e 5f bb 33 2e 31 2a 0d 02 81 7e 06 68 1b e7 8f 3a e4 76 fd be a4 8f 75 71 ad 9a 7b 3f 71 09 3f c2 d1 69 9f 99 22 4a cc 2d 9a b9 63 23 37 b2 ad 76 f2 01 1a 33 a1 f2 ee 72 68 a7 2a fd aa ac 2a 42 c0 98 41 49 74 4f 7e 8a 31 99 2b ca 8f 7a b9
		RSA u, the multiplicative inverse of p, mod q (1023 bits)
			52 22 89 52 98 1c f5 d5 74 96 77 1c c5 43 53 7e 2b 4f 37 dd 30 55 dd 26 ce fe 2b 65 77 26 1e 4d d2 da cb 41 1b 22 b9 b2 7c 3f 2a d9 1c c5 ea ad b6 ae d6 05 62 47 0d e2 1b 23 4b bd eb 84 ca 6d db 09 51 b0 31 76 6b 46 d8 fd cc cb 8e 40 60 1f cb 52 61 cd d3 7d da 6d 62 2d 69 68 37 5d c2 8e b9 70 e5 58 b6 65 1d 6e 93 20 50 75 2b c6 88 4b 23 42 be 2f 44 c6 d1 d5 9a bd 27 5b 01 2e 3f f2
		Checksum
			3d 0c
`
	tag05Redult3 = `Secret-Key Packet (tag 5) (88 bytes)
	04 5b 1a 4e 1d 16 09 2b 06 01 04 01 da 47 0f 01 01 07 40 c6 ae d8 56 62 34 73 e7 f1 86 ff 5f 09 dd d2 c2 b5 48 bd 78 94 90 a8 d2 fd 9c fc c6 69 15 fb 86 00 00 ff 50 5e cc 13 31 23 59 49 c2 cc 48 1d 7c e8 39 85 ac 36 2f 76 ff 5a e5 d6 09 68 c6 e7 de cb 00 5c 10 55
	Version: 4 (current)
		04
	Public-Key
		Public key creation time: 2018-06-08T09:36:29Z
			5b 1a 4e 1d
		Public-key Algorithm: EdDSA (pub 22)
			16
		ECC Curve OID: ed25519 (256bits key size)
			2b 06 01 04 01 da 47 0f 01
		EdDSA EC point (40 || compressd format) (263 bits)
			40 c6 ae d8 56 62 34 73 e7 f1 86 ff 5f 09 dd d2 c2 b5 48 bd 78 94 90 a8 d2 fd 9c fc c6 69 15 fb 86
	Secret-Key (the secret-key data is not encrypted.)
		EdDSA secret key (255 bits)
			50 5e cc 13 31 23 59 49 c2 cc 48 1d 7c e8 39 85 ac 36 2f 76 ff 5a e5 d6 09 68 c6 e7 de cb 00 5c
		Checksum
			10 55
`
)

func TestTag05(t *testing.T) {
	testCases := []struct {
		tag     uint8
		content []byte
		ktm     []byte
		cxt     context.SymAlgMode
		res     string
	}{
		{tag: 5, content: tag05Body1, ktm: nil, cxt: context.ModeNotSpecified, res: tag05Redult1},
		{tag: 5, content: tag05Body2, ktm: nil, cxt: context.ModeNotSpecified, res: tag05Redult2},
		{tag: 5, content: tag05Body3, ktm: nil, cxt: context.ModeNotSpecified, res: tag05Redult3},
	}
	for _, tc := range testCases {
		op := &openpgp.OpaquePacket{Tag: tc.tag, Contents: tc.content}
		cxt := context.NewContext(options.New(
			options.Set(options.DebugOpt, true),
			options.Set(options.IntegerOpt, true),
			options.Set(options.MarkerOpt, true),
			options.Set(options.LiteralOpt, true),
			options.Set(options.PrivateOpt, true),
			options.Set(options.UTCOpt, true),
		))
		if tc.ktm != nil {
			tm, _ := values.NewDateTime(reader.New(tc.ktm), cxt.UTC())
			cxt.KeyCreationTime = tm
		}
		i, err := NewTag(op, cxt).Parse()
		if err != nil {
			t.Errorf("NewTag() = %v, want nil error.", err)
			return
		}
		if cxt.AlgMode() != tc.cxt {
			t.Errorf("Options.Mode = %v, want \"%v\".", cxt.AlgMode(), tc.cxt)
		}
		res := i.String()
		if res != tc.res {
			t.Errorf("Tag.String = \"%s\", want \"%s\".", res, tc.res)
		}
	}
}

/* Copyright 2017,2018 Spiegel
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
