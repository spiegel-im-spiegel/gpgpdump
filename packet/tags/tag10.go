package tags

import (
	"github.com/spiegel-im-spiegel/gpgpdump/info"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/context"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/reader"
	"github.com/spiegel-im-spiegel/gpgpdump/packet/values"
)

// tag10 class for Marker Packet (Obsolete Literal Packet; cannot debug)
type tag10 tagInfo

//newTag10 return tag10 instance
func newTag10(cxt *context.Context, tag values.TagID, body []byte) Tags {
	return &tag10{cxt: cxt, tag: tag, reader: reader.New(body)}
}

// Parse parsing Marker Packet
func (t *tag10) Parse() (*info.Item, error) {
	rootInfo := t.tag.ToItem(t.reader, t.cxt.Debug())
	rootInfo.Add(values.RawData(t.reader, "Literal data", t.cxt.Marker()))
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
