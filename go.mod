module github.com/spiegel-im-spiegel/gpgpdump

go 1.15

require (
	github.com/spf13/cobra v1.0.1-0.20201006035406-b97b5ead31f7
	github.com/spiegel-im-spiegel/errs v1.0.2
	github.com/spiegel-im-spiegel/gocli v0.10.3
	golang.org/x/crypto v0.0.0-20201012173705-84dcc777aaee
)

replace github.com/coreos/etcd v3.3.13+incompatible => github.com/coreos/etcd v3.3.25+incompatible
