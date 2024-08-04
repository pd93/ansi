package ansi

import (
	"testing"
)

var tests = []struct {
	name                   string
	input                  string
	expectedParsedOutput   string
	expectedStrippedOutput string
}{
	// Formatting
	{
		name:                   "bold",
		input:                  "[bold]foo[/]bar",
		expectedParsedOutput:   "\x1b[1mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "faint",
		input:                  "[faint]foo[/]bar",
		expectedParsedOutput:   "\x1b[2mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "dim",
		input:                  "[dim]foo[/]bar",
		expectedParsedOutput:   "\x1b[2mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "italic",
		input:                  "[italic]foo[/]bar",
		expectedParsedOutput:   "\x1b[3mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "underline",
		input:                  "[underline]foo[/]bar",
		expectedParsedOutput:   "\x1b[4mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "blink",
		input:                  "[blink]foo[/]bar",
		expectedParsedOutput:   "\x1b[5mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "invert",
		input:                  "[invert]foo[/]bar",
		expectedParsedOutput:   "\x1b[7mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "hidden",
		input:                  "[hidden]foo[/]bar",
		expectedParsedOutput:   "\x1b[8mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "strike",
		input:                  "[strike]foo[/]bar",
		expectedParsedOutput:   "\x1b[9mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	// Formatting reset
	{
		name:                   "/bold",
		input:                  "[/bold]foo[/]bar",
		expectedParsedOutput:   "\x1b[22mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "/faint",
		input:                  "[/faint]foo[/]bar",
		expectedParsedOutput:   "\x1b[22mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "/dim",
		input:                  "[/dim]foo[/]bar",
		expectedParsedOutput:   "\x1b[22mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "/italic",
		input:                  "[/italic]foo[/]bar",
		expectedParsedOutput:   "\x1b[23mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "/underline",
		input:                  "[/underline]foo[/]bar",
		expectedParsedOutput:   "\x1b[24mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "/blink",
		input:                  "[/blink]foo[/]bar",
		expectedParsedOutput:   "\x1b[25mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "/invert",
		input:                  "[/invert]foo[/]bar",
		expectedParsedOutput:   "\x1b[27mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "/hidden",
		input:                  "[/hidden]foo[/]bar",
		expectedParsedOutput:   "\x1b[28mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "/strike",
		input:                  "[/strike]foo[/]bar",
		expectedParsedOutput:   "\x1b[29mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	// Foreground colors
	{
		name:                   "black",
		input:                  "[black]foo[/]bar",
		expectedParsedOutput:   "\x1b[30mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "red",
		input:                  "[red]foo[/]bar",
		expectedParsedOutput:   "\x1b[31mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "green",
		input:                  "[green]foo[/]bar",
		expectedParsedOutput:   "\x1b[32mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "yellow",
		input:                  "[yellow]foo[/]bar",
		expectedParsedOutput:   "\x1b[33mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "blue",
		input:                  "[blue]foo[/]bar",
		expectedParsedOutput:   "\x1b[34mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "magenta",
		input:                  "[magenta]foo[/]bar",
		expectedParsedOutput:   "\x1b[35mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "cyan",
		input:                  "[cyan]foo[/]bar",
		expectedParsedOutput:   "\x1b[36mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "white",
		input:                  "[white]foo[/]bar",
		expectedParsedOutput:   "\x1b[37mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "24-bit style",
		input:                  "[255,0,0]foo[/]bar",
		expectedParsedOutput:   "\x1b[38;2;255;0;0mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "8-bit style",
		input:                  "[208]foo[/]bar",
		expectedParsedOutput:   "\x1b[38;5;208mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "reset foreground",
		input:                  "[/fg]foo[/]bar",
		expectedParsedOutput:   "\x1b[39mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	// Background colors
	{
		name:                   "bg black",
		input:                  "[bg-black]foo[/]bar",
		expectedParsedOutput:   "\x1b[40mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "bg red",
		input:                  "[bg-red]foo[/]bar",
		expectedParsedOutput:   "\x1b[41mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "bg green",
		input:                  "[bg-green]foo[/]bar",
		expectedParsedOutput:   "\x1b[42mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "bg yellow",
		input:                  "[bg-yellow]foo[/]bar",
		expectedParsedOutput:   "\x1b[43mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "bg blue",
		input:                  "[bg-blue]foo[/]bar",
		expectedParsedOutput:   "\x1b[44mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "bg magenta",
		input:                  "[bg-magenta]foo[/]bar",
		expectedParsedOutput:   "\x1b[45mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "bg cyan",
		input:                  "[bg-cyan]foo[/]bar",
		expectedParsedOutput:   "\x1b[46mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "bg white",
		input:                  "[bg-white]foo[/]bar",
		expectedParsedOutput:   "\x1b[47mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "bg 24-bit style",
		input:                  "[bg-255,0,0]foo[/]bar",
		expectedParsedOutput:   "\x1b[48;2;255;0;0mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "bg 8-bit style",
		input:                  "[bg-208]foo[/]bar",
		expectedParsedOutput:   "\x1b[48;5;208mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "bg reset foreground",
		input:                  "[/bg]foo[/]bar",
		expectedParsedOutput:   "\x1b[49mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	// Misc tests
	{
		name:                   "multiple styles",
		input:                  "[red:faint]foo[/]bar",
		expectedParsedOutput:   "\x1b[31;2mfoo\x1b[0mbar",
		expectedStrippedOutput: "foobar",
	},
	{
		name:                   "invalid style",
		input:                  "[invalid]foo[/]bar",
		expectedParsedOutput:   "[invalid]foo\x1b[0mbar",
		expectedStrippedOutput: "[invalid]foobar",
	},
	{
		name:                   "invalid 8-bit style",
		input:                  "[300]foo[/]bar",
		expectedParsedOutput:   "[300]foo\x1b[0mbar",
		expectedStrippedOutput: "[300]foobar",
	},
	{
		name:                   "invalid 24-bit style",
		input:                  "[300,0,0]foo[/]bar",
		expectedParsedOutput:   "[300,0,0]foo\x1b[0mbar",
		expectedStrippedOutput: "[300,0,0]foobar",
	},
	{
		name:                   "multiple styles, one invalid",
		input:                  "[red:invalid]foo[/]bar",
		expectedParsedOutput:   "[red:invalid]foo\x1b[0mbar",
		expectedStrippedOutput: "[red:invalid]foobar",
	},
	{
		name:                   "escaped style",
		input:                  `\[red]foo[/]bar`,
		expectedParsedOutput:   "[red]foo\x1b[0mbar",
		expectedStrippedOutput: "[red]foobar",
	},
	{
		name:                   "escaped invalid style",
		input:                  `\[invalid]foo[/]bar`,
		expectedParsedOutput:   "[invalid]foo\x1b[0mbar",
		expectedStrippedOutput: "[invalid]foobar",
	},
	{
		name:                   "multiple styles, escaped and one invalid",
		input:                  `\[red:invalid]foo[/]bar`,
		expectedParsedOutput:   "[red:invalid]foo\x1b[0mbar",
		expectedStrippedOutput: "[red:invalid]foobar",
	},
}

func assertEqual[T comparable](t *testing.T, expected, got T) {
	if got != expected {
		t.Errorf("values did not match\nexpected: %v\ngot:      %v", expected, got)
	}
}

func TestParse(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assertEqual(t, test.expectedParsedOutput, Parse(test.input))
		})
	}
}

func TestStrip(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assertEqual(t, test.expectedStrippedOutput, Strip(test.input))
		})
	}
}

func BenchmarkParse(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			Parse(test.input)
		})
	}
}
