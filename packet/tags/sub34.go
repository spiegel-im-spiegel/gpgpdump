package tags

import (
	"github.com/spiegel-im-spiegel/gpgpdump/info"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/context"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/reader"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/values"
)

//subReserved class for Preferred AEAD Algorithms Sub-packet
type sub34 struct {
	subInfo
}

//newSubReserved return sub34 instance
func newSub34(cxt *context.Context, subID values.SuboacketID, body []byte) Subs {
	return &sub34{subInfo{cxt: cxt, subID: subID, reader: reader.New(body)}}
}

// Parse parsing Preferred AEAD Algorithms Sub-packet
func (s *sub34) Parse() (*info.Item, error) {
	rootInfo := s.ToItem()
	return rootInfo, nil
}

/* Copyright 2018,2019 Spiegel
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