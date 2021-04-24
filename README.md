# parrot

parrot is a simple calculator in go currently using very little unique code. I wish to implement my own query parser in the future. Currently, this project exists to replace many other cli calculators I've found that don't meet what I want them to.

my goal is for parrot to be a multi-tool of sorts for command-line maths and computations. I've added a few tools for programmers/designers already, but I'd like these to be modular and plentiful, so I'll be integrating a plug-in system at some point. 

# current features

- most basic maths
- live reloading config (courtesy of viper)
- variables
- constant system (poorly designed)
- HTML color converter

# planned features

- custom token based math parser to allow units and constants better
- reimplement plugins to be go plugin based, allowing addition of modules without recompilation of the base tool.
- implied multiplication (like "4x" -> "4 * x")
- implement the standard set of math functions
- reimplement constants and variables to be truly dynamic and modular (categories, multiple constant lists)
- more programmer tools
- dynamic config tool
- more config options

# libraries

- spf13viper (here to stay)
- sirupsen/logrus (here to stay)
- chzyer/readline (here to stay)
- fatih/color (here to stay)
- go-playground/colors.v1 (will probably be removed, but its a small library used in a tool so eh?)

# to deprecate
- Knetic/govaluate (will be removed)