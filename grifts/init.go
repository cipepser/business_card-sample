package grifts

import (
	"github.com/cipepser/business_card/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
