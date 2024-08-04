<div align="center">
  <img src="res/title.svg">
</div>

<p align="center">
  <i>A simple, lightweight Golang package for working with ANSI escape codes.</i>
</p>
<br>

If you've ever worked on a CLI application before, you probably know the
struggle of working with ANSI escape codes. Often, you just want to quickly
apply a style to a string, but you either have to interrupt your workflow to
look up the escape sequence and codes, or you have to import a large package
that does more than you need.

The `ansi` package is extremely lightweight :feather: and has no external
dependencies :package: It allows you to quickly apply styles to strings using
**human-readable** style blocks. Because the styles are defined **inline**,
there is no need to call a function for each style, making it easy to
apply/reset multiple styles to a single string.

With `ansi`, the mess at the top of this page becomes: <img src="res/ansi.svg">

## Usage

Strings are formatted using "style blocks". A style block is delimited by square
brackets (`[]`) and contains a list of styles separated by colons. For example,
the style block `[red:bold]` applies the red and bold styles to any text
following the block.

Styles can be reset by using the corresponding reset code or by using the global
reset (`[/]`) to reset all styles.

```go
import "github.com/pd93/ansi"

str := ansi.Parse("This is a [red:bold]red and bold[/] string")
fmt.Println(str)

// Or use the provided fmt package function wrappers
ansi.Println("This is a [red:bold]red and bold[/] string")
```

If you need to print a literal string that conflicts with a valid style block,
you can escape the block by prepending it with a backslash (`\`). For example,
`\[red]` will print `[red]` instead of applying the red style.

```go
// Note that we need to escape the escaping backslash otherwise it will be removed.
ansi.Println("This is \\[red:bold]an escaped style block\\[/]")

// We can avoid this in most situations by using a raw string (backticks) instead:
ansi.Println(`This is \[red:bold]an escaped style block\[/]`)
```

For more info, check out the [style reference](#style-reference) below or take a
look at our [examples](./examples/main.go).

## Style Reference

### Reset

| Code | ANSI Code | Description       |
| ---- | --------: | ----------------- |
| `/`  |       `0` | Resets all styles |

### Formatting

| Code           | ANSI Code | Description                                          |
| -------------- | --------: | ---------------------------------------------------- |
| `bold`         |       `1` | Sets the font weight to bold                         |
| `faint`, `dim` |       `2` | Sets the text brightness to its faint/dim variant \* |
| `italic`       |       `3` | Sets the font style to italic                        |
| `underline`    |       `4` | Sets the text decoration to underline                |
| `blink`        |       `5` | Sets the text to blink in and out                    |
| `invert`       |       `7` | Inverts the foreground and background colors</span>  |
| `hidden`       |       `8` | Sets the text to be hidden                           |
| `strike`       |       `9` | Sets the text decoration to line-through             |

> \* `faint` and `dim` do not work when using 8-bit (256) or 24-bit (RGB) color
> modes.

| Code             | ANSI Code | Description                             |
| ---------------- | --------: | --------------------------------------- |
| `/bold`          |      `22` | Resets the font weight \*\*             |
| `/faint`, `/dim` |      `22` | Resets the text brightness \*\*         |
| `/italic`        |      `23` | Resets the font style                   |
| `/underline`     |      `24` | Resets the underline text decoration    |
| `/blink`         |      `25` | Stops the text from blinking            |
| `/invert`        |      `27` | Resets the text inversion               |
| `/hidden`        |      `28` | Resets the text visibility              |
| `/strike`        |      `29` | Resets the line-through text decoration |

> \*\* The `/bold`, `/faint`, and `/dim` codes all equal because they share an
> ANSI code. This means it is not possible to reset them individually.

### Foreground Colors

| Code      |    ANSI Code | Description                                                        |
| --------- | -----------: | ------------------------------------------------------------------ |
| `black`   |         `30` | Sets the text foreground to black                                  |
| `red`     |         `31` | Sets the text foreground to red                                    |
| `green`   |         `32` | Sets the text foreground to green                                  |
| `yellow`  |         `33` | Sets the text foreground to yellow                                 |
| `blue`    |         `34` | Sets the text foreground to blue                                   |
| `magenta` |         `35` | Sets the text foreground to magenta                                |
| `cyan`    |         `36` | Sets the text foreground to cyan                                   |
| `white`   |         `37` | Sets the text foreground to white                                  |
| `R,G,B`   | `38;2;R;G;B` | Sets the text foreground to the given [24-bit (RGB) color][24-bit] |
| `N`       |     `38;5;N` | Sets the text foreground to the given [8-bit (256) color][8-bit]   |
| `/fg`     |         `39` | Resets the text foreground to the default color                    |

### Background Colors

| Code         |    ANSI Code | Description                                                        |
| ------------ | -----------: | ------------------------------------------------------------------ |
| `bg-black`   |         `40` | Sets the text background to black                                  |
| `bg-red`     |         `41` | Sets the text background to red                                    |
| `bg-green`   |         `42` | Sets the text background to green                                  |
| `bg-yellow`  |         `43` | Sets the text background to yellow                                 |
| `bg-blue`    |         `44` | Sets the text background to blue                                   |
| `bg-magenta` |         `45` | Sets the text background to magenta                                |
| `bg-cyan`    |         `46` | Sets the text background to cyan                                   |
| `bg-white`   |         `47` | Sets the text background to white                                  |
| `bg-R,G,B`   | `48;2;R;G;B` | Sets the text background to the given [24-bit (RGB) color][24-bit] |
| `bg-N`       |     `48;5;N` | Sets the text background to the given [8-bit (256) color][8-bit]   |
| `/bg`        |         `49` | Resets the text background to the default color                    |

[8-bit]: https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
[24-bit]: https://en.wikipedia.org/wiki/ANSI_escape_code#24-bit
