package grifts

import (
	"memorize/actions"
	
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
