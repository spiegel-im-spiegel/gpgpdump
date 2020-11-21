package facade

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/gocli/signal"
	"github.com/spiegel-im-spiegel/gpgpdump/ecode"
	"github.com/spiegel-im-spiegel/gpgpdump/fetch"
	"github.com/spiegel-im-spiegel/gpgpdump/parse"
)

//newHkpCmd returns cobra.Command instance for show sub-command
func newFetchCmd(ui *rwi.RWI) *cobra.Command {
	fetchCmd := &cobra.Command{
		Use:     "fetch [flags] <URL>",
		Aliases: []string{"fch", "f"},
		Short:   "Dumps OpenPGP packets form Web",
		Long:    "Dumps OpenPGP packets form Web.",
		RunE: func(cmd *cobra.Command, args []string) error {
			cxt := parseContext(cmd)
			//user id
			if len(args) == 0 {
				return debugPrint(ui, ecode.ErrNoURL)
			}
			urlStr := args[0]

			//options
			rawFlag, err := cmd.Flags().GetBool("raw")
			if err != nil {
				return debugPrint(ui, errs.New("error in --raw option", errs.WithCause(err)))
			}

			//Fetch OpenPGP packets
			resp, err := fetch.New(
				fetch.WithContext(signal.Context(context.Background(), os.Interrupt)),
			).Get(urlStr)
			if err != nil {
				return debugPrint(ui, err)
			}
			defer resp.Close()
			if rawFlag {
				return debugPrint(ui, ui.WriteFrom(resp))
			}

			//parse OpenPGP packets
			p, err := parse.New(cxt, resp)
			if err != nil {
				return debugPrint(ui, err)
			}
			res, err := p.Parse()
			if err != nil {
				return debugPrint(ui, err)
			}
			r, err := marshalPacketInfo(res)
			if err != nil {
				return debugPrint(ui, err)
			}
			return debugPrint(ui, ui.WriteFrom(r))
		},
	}
	fetchCmd.Flags().BoolP("raw", "", false, "output raw text from GitHub")

	return fetchCmd
}

/* Copyright 2019,2020 Spiegel
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
