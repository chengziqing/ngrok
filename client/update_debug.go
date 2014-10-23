// +build release autoupdate

package client

import (
	"github.com/chengziqing/ngrok/client/mvc"
)

// no auto-updating in debug mode
func autoUpdate(state mvc.State, token string) {
}
