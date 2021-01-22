package tags

import (
	"github.com/spiegel-im-spiegel/gpgpdump/parse/context"
	"github.com/spiegel-im-spiegel/gpgpdump/parse/reader"
	"github.com/spiegel-im-spiegel/gpgpdump/parse/result"
	"github.com/spiegel-im-spiegel/gpgpdump/parse/values"
)

//sub10 class for Placeholder for backward compatibility Sub-packet
type sub10 struct {
	subInfo
}

//newSub10 return sub10 instance
func newSub10(cxt *context.Context, subID values.SuboacketID, body []byte) Subs {
	return &sub10{subInfo{cxt: cxt, subID: subID, reader: reader.New(body)}}
}

// Parse parsing Placeholder for backward compatibility Sub-packet
func (s *sub10) Parse() (*result.Item, error) {
	rootInfo := s.ToItem()
	return rootInfo, nil
}

/* Copyright 2016-2019 Spiegel
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
