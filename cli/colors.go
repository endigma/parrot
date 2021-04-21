package cli

import "github.com/fatih/color"

// colors
var color_blue *color.Color = color.New(color.FgBlue)
var color_green *color.Color = color.New(color.FgGreen)
var color_cyan *color.Color = color.New(color.FgCyan)
var color_red *color.Color = color.New(color.FgRed)

// bolds
var color_cyan_bold *color.Color = color.New(color.FgCyan, color.Bold)
var color_blue_bold *color.Color = color.New(color.FgBlue, color.Bold)
var color_green_bold *color.Color = color.New(color.FgGreen, color.Bold)
var color_red_bold *color.Color = color.New(color.FgRed, color.Bold)
