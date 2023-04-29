package colors

import "github.com/fatih/color"

var (
	Warning = color.New(color.FgYellow).Add(color.Bold)
	Errors  = color.New(color.FgRed).Add(color.Bold)
	Info    = color.New(color.FgBlue).Add(color.Bold)
)
