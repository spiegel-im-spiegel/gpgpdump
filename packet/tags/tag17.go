package tags

import (
	"github.com/spiegel-im-spiegel/gpgpdump/info"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/context"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/reader"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/values"
)

// tag17 class for User Attribute Packet
type tag17 tagInfo

//newTag17 return tag17 instance
func newTag17(cxt *context.Context, tag values.TagID, body []byte) Tags {
	return &tag17{cxt: cxt, tag: tag, reader: reader.New(body)}
}

// Parse parsing User Attribute Packet
func (t *tag17) Parse() (*info.Item, error) {
	rootInfo := t.tag.ToItem(t.reader, t.cxt.Debug())
	sp, err := t.reader.Read2EOF()
	if err != nil {
		return rootInfo, err
	}
	subpcket, err := newSubparser(t.cxt, t.tag, "Subpacket", sp)
	if err != nil {
		return rootInfo, err
	}
	itm, err := subpcket.Parse()
	if err != nil {
		return rootInfo, err
	}
	rootInfo.Add(itm)
	return rootInfo, nil
}

/* Copyright 2016 Spiegel
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