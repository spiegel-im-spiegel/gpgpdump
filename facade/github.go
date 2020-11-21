package facade

import (
	"bytes"
	"context"
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/gocli/signal"
	"github.com/spiegel-im-spiegel/gpgpdump/ecode"
	"github.com/spiegel-im-spiegel/gpgpdump/fetch"
	"github.com/spiegel-im-spiegel/gpgpdump/github"
	"github.com/spiegel-im-spiegel/gpgpdump/parse"
)

//newHkpCmd returns cobra.Command instance for show sub-command
func newGitHubCmd(ui *rwi.RWI) *cobra.Command {
	githubCmd := &cobra.Command{
		Use:     "github [flags] <GitHub user ID>",
		Aliases: []string{"gh", "g"},
		Short:   "Dumps OpenPGP keys registered on GitHub",
		Long:    "Dumps OpenPGP keys registered on GitHub.",
		RunE: func(cmd *cobra.Command, args []string) error {
			cxt := parseContext(cmd)
			//user id
			if len(args) == 0 {
				return debugPrint(ui, ecode.ErrGitHubUserID)
			}
			userID := args[0]

			//options
			keyid, err := cmd.Flags().GetString("keyid")
			if err != nil {
				return debugPrint(ui, errs.New("error in --keyid option", errs.WithCause(err)))
			}
			rawFlag, err := cmd.Flags().GetBool("raw")
			if err != nil {
				return debugPrint(ui, errs.New("error in --raw option", errs.WithCause(err)))
			}

			//Fetch OpenPGP packets
			resp, err := github.GetKey(
				fetch.New(
					fetch.WithContext(signal.Context(context.Background(), os.Interrupt)),
				),
				userID,
				keyid,
			)
			if err != nil {
				if errors.Is(err, ecode.ErrArmorText) {
					return debugPrint(ui, ui.WriteFrom(bytes.NewReader(resp)))
				}
				return debugPrint(ui, err)
			}
			if rawFlag {
				return debugPrint(ui, ui.WriteFrom(bytes.NewReader(resp)))
			}

			//parse OpenPGP packets
			p, err := parse.NewBytes(cxt, resp)
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
	githubCmd.Flags().StringP("keyid", "", "", "OpenPGP key ID")
	githubCmd.Flags().BoolP("raw", "", false, "output raw text from GitHub")

	return githubCmd
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
