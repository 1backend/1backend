/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package boot

import "os"

type Options struct {
	Test      bool
	ServerUrl string
	SelfUrl   string
}

func (o *Options) LoadEnvars() {
	if o.ServerUrl == "" {
		o.ServerUrl = os.Getenv("OB_SERVER_URL")
	}

	if o.ServerUrl == "" {
		o.ServerUrl = "http://127.0.0.1:11337"
	}

	if o.SelfUrl == "" {
		o.SelfUrl = os.Getenv("OB_SELF_URL")
	}

	if !o.Test && os.Getenv("OB_TEST") == "true" {
		o.Test = true
	}
}
