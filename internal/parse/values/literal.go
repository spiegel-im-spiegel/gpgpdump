package values

import (
	"fmt"

	"github.com/spiegel-im-spiegel/gpgpdump/items"
)

var literalFormatNames = Msgs{
	0x62: "binary",     //'b'
	0x74: "text",       //'t'
	0x75: "UTF-8 text", //'u'
	0x31: "local",      //'l' -- RFC 1991 incorrectly stated this local mode flag as '1' (ASCII numeral one). Both of these local modes are deprecated.
	0x6c: "local",      //'l'
}

// LiteralFormat is format of literal data
type LiteralFormat byte

// Get returns Item instance
func (l LiteralFormat) Get() *items.Item {
	return items.NewItem("Literal data format", l.String(), literalFormatNames.Get(int(l), "unknown"), "")
}

func (l LiteralFormat) String() string {
	return string(byte(l))
}

// LiteralFname is file name of literal data
type LiteralFname string

// Get returns Item instance
func (l LiteralFname) Get() *items.Item {
	return items.NewItem("File name", string(l), "", "")
}

//LiteralData returns new RawData instance for Literal data
func LiteralData(buf []byte, dump bool) *RawData {
	return NewRawData("Literal data", fmt.Sprintf("%d bytes", len(buf)), buf, dump)
}
