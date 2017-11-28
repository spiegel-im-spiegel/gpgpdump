package gpgpdump

import (
	"bytes"
	"io"

	"github.com/spiegel-im-spiegel/gpgpdump/info"
	"github.com/spiegel-im-spiegel/gpgpdump/options"
	"github.com/spiegel-im-spiegel/gpgpdump/packet"
)

//Parse returns packet info (from io.Reader stream).
func Parse(r io.Reader, o *options.Options) (*info.Info, error) {
	parser, err := packet.NewParser(r, o)
	if err != nil {
		return nil, err
	}
	return parser.Parse()
}

//ParseByte returns packet info (from []byte data).
func ParseByte(data []byte, o *options.Options) (*info.Info, error) {
	return Parse(bytes.NewReader(data), o)
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