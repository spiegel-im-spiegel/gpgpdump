package tags

import (
	"testing"

	"github.com/spiegel-im-spiegel/gpgpdump/options"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/context"

	openpgp "golang.org/x/crypto/openpgp/packet"
)

var (
	tag04Body1 = []byte{0x03, 0x00, 0x08, 0x11, 0xb4, 0xda, 0x3b, 0xae, 0x7e, 0x20, 0xb8, 0x1c, 0x01}
)

const (
	tag04Result1 = `One-Pass Signature Packet (tag 4) (13 bytes)
	03 00 08 11 b4 da 3b ae 7e 20 b8 1c 01
	Version: 3 (current)
		03
	Signiture Type: Signature of a binary document (0x00)
		00
	Hash Algorithm: SHA2-256 (hash 8)
		08
	Public-key Algorithm: DSA (Digital Signature Algorithm) (pub 17)
		11
	Key ID: 0xb4da3bae7e20b81c
	Encrypted session key: other than one pass signature (flag 0x01)
`
)

func TestTag04(t *testing.T) {
	op := &openpgp.OpaquePacket{Tag: 4, Contents: tag04Body1}
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
	if str != tag04Result1 {
		t.Errorf("Tag.String = \"%s\", want \"%s\".", str, tag04Result1)
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
