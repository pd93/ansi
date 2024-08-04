package ansi

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const (
	ansiEscapeSequenceFormat = "\x1b[%sm"
	ansiCodeFormat           = "%d"
	ansiCode8bitFormat       = "%d;5;%d"
	ansiCode24bitFormat      = "%d;2;%d;%d;%d"
)

// ANSI escape codes
const (
	RESET            = 0
	BOLD             = 1
	FAINT            = 2
	ITALIC           = 3
	UNDERLINE        = 4
	BLINK            = 5
	INVERT           = 7
	HIDDEN           = 8
	STRIKE           = 9
	BOLD_FAINT_RESET = 22
	ITALIC_RESET     = 23
	UNDER_RESET      = 24
	BLINK_RESET      = 25
	INVERT_RESET     = 27
	HIDDEN_RESET     = 28
	STRIKE_RESET     = 29
	FG_BLACK         = 30
	FG_RED           = 31
	FG_GREEN         = 32
	FG_YELLOW        = 33
	FG_BLUE          = 34
	FG_MAGENTA       = 35
	FG_CYAN          = 36
	FG_WHITE         = 37
	FG_CUSTOM        = 38
	FG_DEFAULT       = 39
	BG_BLACK         = 40
	BG_RED           = 41
	BG_GREEN         = 42
	BG_YELLOW        = 43
	BG_BLUE          = 44
	BG_MAGENTA       = 45
	BG_CYAN          = 46
	BG_WHITE         = 47
	BG_CUSTOM        = 48
	BG_DEFAULT       = 49
)

var (
	styleRegex = regexp.MustCompile(`\\?\[[\/a-z0-9,\-:]+\]`)
	styleMap   = map[string]int{
		// Reset
		"/": RESET,
		// Formatting
		"bold":       BOLD,
		"faint":      FAINT,
		"dim":        FAINT,
		"italic":     ITALIC,
		"underline":  UNDERLINE,
		"blink":      BLINK,
		"invert":     INVERT,
		"hidden":     HIDDEN,
		"strike":     STRIKE,
		"/bold":      BOLD_FAINT_RESET,
		"/faint":     BOLD_FAINT_RESET,
		"/dim":       BOLD_FAINT_RESET,
		"/italic":    ITALIC_RESET,
		"/underline": UNDER_RESET,
		"/blink":     BLINK_RESET,
		"/invert":    INVERT_RESET,
		"/hidden":    HIDDEN_RESET,
		"/strike":    STRIKE_RESET,
		// Foreground colors
		"black":   FG_BLACK,
		"red":     FG_RED,
		"green":   FG_GREEN,
		"yellow":  FG_YELLOW,
		"blue":    FG_BLUE,
		"magenta": FG_MAGENTA,
		"cyan":    FG_CYAN,
		"white":   FG_WHITE,
		"/fg":     FG_DEFAULT,
		// Background colors
		"bg-black":   BG_BLACK,
		"bg-red":     BG_RED,
		"bg-green":   BG_GREEN,
		"bg-yellow":  BG_YELLOW,
		"bg-blue":    BG_BLUE,
		"bg-magenta": BG_MAGENTA,
		"bg-cyan":    BG_CYAN,
		"bg-white":   BG_WHITE,
		"/bg":        BG_DEFAULT,
	}
)

type StringOrBytes interface {
	string | []byte
}

func formatEscapeSequence(code int) string {
	return fmt.Sprintf(ansiCodeFormat, code)
}

func formatEscapeSequence8bit(n int, bg bool) string {
	if bg {
		return fmt.Sprintf(ansiCode8bitFormat, BG_CUSTOM, n)
	}
	return fmt.Sprintf(ansiCode8bitFormat, FG_CUSTOM, n)
}

func formatEscapeSequence24bit(r, g, b int, bg bool) string {
	if bg {
		return fmt.Sprintf(ansiCode24bitFormat, BG_CUSTOM, r, g, b)
	}
	return fmt.Sprintf(ansiCode24bitFormat, FG_CUSTOM, r, g, b)
}

func valid256(n int) bool {
	return n >= 0 && n < 256
}

func getANSICodeFromStyle(style string) (string, error) {
	// Named codes
	code, ok := styleMap[style]
	if ok {
		return formatEscapeSequence(code), nil
	}

	// Strip bg- prefix if present
	bg := false
	style, bg = strings.CutPrefix(style, "bg-")

	// 8-bit (256) color codes
	n, err := strconv.Atoi(style)
	if err == nil && valid256(n) {
		return formatEscapeSequence8bit(n, bg), nil
	}

	// 24-bit (RGB) color codes
	rgb := strings.Split(style, ",")
	if len(rgb) == 3 {
		r, err := strconv.Atoi(rgb[0])
		if err != nil || !valid256(r) {
			return "", fmt.Errorf("invalid 24-bit color code: %s", style)
		}
		g, err := strconv.Atoi(rgb[1])
		if err != nil || !valid256(g) {
			return "", fmt.Errorf("invalid 24-bit color code: %s", style)
		}
		b, err := strconv.Atoi(rgb[2])
		if err != nil || !valid256(b) {
			return "", fmt.Errorf("invalid 24-bit color code: %s", style)
		}
		return formatEscapeSequence24bit(r, g, b, bg), nil
	}

	// Unrecognized code
	return "", fmt.Errorf("unrecognized style: %s", style)
}

func parse(input string, delete bool) string {
	// Replace each matching style block with the equivalent ANSI escape sequence
	return styleRegex.ReplaceAllStringFunc(input, func(input string) string {
		var ansiCodes []string

		// Check if the style block is escaped
		if input, escaped := strings.CutPrefix(input, `\`); escaped {
			return input
		}

		// Trim the delimiters from the style block
		styles := strings.Trim(input, "[]")

		// Loop over each style in the set and convert it to an ANSI code
		// If any of the styles are invalid, return the input without modification
		for _, style := range strings.Split(styles, ":") {
			ansiCode, err := getANSICodeFromStyle(style)
			if err != nil {
				return input
			}
			ansiCodes = append(ansiCodes, ansiCode)
		}

		// If we're stripping the styles, return an empty string
		// We do this here to avoid returning an empty string if any of the styles were invalid
		if delete {
			return ""
		}

		return fmt.Sprintf(ansiEscapeSequenceFormat, strings.Join(ansiCodes, ";"))
	})
}

// Parse will search the given input for codes given in the [code] format and
// replace them with the equivalent ANSI escape sequence. For example, if the
// input string is "[red]hello[/]" the output will be "\x1b[31mhello\x1b[0m".
// When output to a terminal this would print the string "hello" in red. See the
// list of constants in this package for available codes. Multiple codes can be
// specified at once by using the [code1:code2] syntax. If any code in the set
// is not recognized, the entire set is returned as-is. If you want the literal
// string [red], you can escape the parsing by using \[red].
func Parse[T StringOrBytes](input T) T {
	return T(parse(string(input), false))
}

// Strip will search the given input for codes given in the [code] format and
// delete them. For example, if the input string is "[red]hello[/]" the output
// will be "hello". See the list of constants in this package for available
// codes. Multiple codes can be specified at once by using the [code1:code2]
// syntax. If any code in the set is not recognized, the entire set is returned
// as-is. If you want the literal string [red], you can escape the parsing by
// using \[red].
func Strip[T StringOrBytes](input T) T {
	return T(parse(string(input), true))
}

// Print is a wrapper around fmt.Print. It will call the standard library's
// Sprint function with the given arguments and then pass the result to
// ansi.Parse. The result of this is then printed to stdout.
func Print(a ...any) {
	fmt.Print(Parse(fmt.Sprint(a...)))
}

// Printf is a wrapper around fmt.Printf. It will call the standard library's
// Sprintf function with the given arguments and then pass the result to
// ansi.Parse. The result of this is then printed to stdout.
func Printf(format string, a ...any) {
	fmt.Print(Parse(fmt.Sprintf(format, a...)))
}

// Println is a wrapper around fmt.Println. It will call the standard library's
// Sprintln function with the given arguments and then pass the result to
// ansi.Parse. The result of this is then printed to stdout.
func Println(a ...any) {
	fmt.Print(Parse(fmt.Sprintln(a...)))
}

// Fprint is a wrapper around fmt.Fprint. It will call the standard library's
// Sprint function with the given arguments and then pass the result to
// ansi.Parse. The result of this is then sent to the given writer.
func Fprint(w io.Writer, a ...any) {
	fmt.Fprint(w, Parse(fmt.Sprint(a...)))
}

// Fprintf is a wrapper around fmt.Fprintf. It will call the standard library's
// Sprintf function with the given arguments and then pass the result to
// ansi.Parse. The result of this is then sent to the given writer.
func Fprintf(w io.Writer, format string, a ...any) {
	fmt.Fprint(w, Parse(fmt.Sprintf(format, a...)))
}

// Fprintln is a wrapper around fmt.Fprintln. It will call the standard
// library's Sprintln function with the given arguments and then pass the result
// to ansi.Parse. The result of this is then sent to the given writer.
func Fprintln(w io.Writer, a ...any) {
	fmt.Fprint(w, Parse(fmt.Sprintln(a...)))
}
