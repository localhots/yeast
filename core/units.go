package core

import (
	"code.google.com/p/go.net/context"
	"github.com/localhots/yeast/units/aggregator"
	"github.com/localhots/yeast/units/input"
	"github.com/localhots/yeast/units/logger"
	"github.com/localhots/yeast/units/power"
	"github.com/localhots/yeast/units/sleep"
	"github.com/localhots/yeast/units/uuid"
)

type (
	UnitsDict map[string]func(context.Context) context.Context
)

var (
	Units = UnitsDict{
		"aggregator":      aggregator.Call,
		"logger":          logger.Call,
		"power2":          power.Power2,
		"power3":          power.Power3,
		"power5":          power.Power5,
		"input_from_flag": input.FromFlag,
		"sleep":           sleep.Call,
		"uuid":            uuid.Call,
	}
)
