<p style="width:100%;text-align:center;font-size:40px;margin-bottom:30px;letter-spacing:5px;font-family:consolas;">
<span style="color:#ad7fa8;">\x1b</span><span style="color:#729fcf;">[</span><span style="color:#8ae234;">1;43</span><span style="color:#ef2929;">m</span><span style="font-weight:bold;background-color:#fce94f;color:#555753;">ANSI</span><span style="color:#ad7fa8;">\x1b</span><span style="color:#729fcf;">[</span><span style="color:#8ae234;">0</span><span style="color:#ef2929;">m</span>
</p>

<p style="width:100%;text-align:center;margin-bottom:30px;font-style:italic;">
A simple, lightweight Golang package for working with ANSI escape codes.
</p>

If you've ever worked on a CLI application before, you probably know the
struggle of working with ANSI escape codes. Often, you just want to quickly
apply a style to a string, but you either have to interrupt your workflow to
look up the escape sequence and codes, or you have to import a large package
that does more than you need.

The `ansi` package is extremely lightweight :feather: and has no external
dependencies :package: It allows you to quickly apply styles to strings using
**human-readable** style blocks. Because the styles are defined **inline**, there is no
need to call a function for each style, making it easy to apply/reset multiple
styles to a single string.

With `ansi`, the mess at the top of this page becomes: <span style="font-size:20px;letter-spacing:2px;font-family:consolas;">
<span style="color:#ad7fa8;">[</span><span style="color:#8ae234;">bold:bg-yellow</span><span style="color:#ad7fa8;">]</span><span style="font-weight:bold;background-color:#fce94f;color:#555753;">ANSI</span><span style="color:#ad7fa8;">[</span><span style="color:#ef2929;">/</span><span style="color:#ad7fa8;">]</span>
</span>

## Usage

Strings are formatted using "style blocks". A style block is delimited by square
brackets (`[]`) and contains a list of styles separated by colons. For
example, the style block `[red:bold]` applies the red and bold styles to any
text following the block.

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
ansi.Println("This is \\[red:bold]a normally formatted string because the style block is escaped\\[/]")

// We can avoid this in most situations by using a raw string (backticks) instead:
ansi.Println(`This is \[red:bold]a normally formatted string because the style block is escaped\[/]`)
```

For more info, check out the [style reference](#style-reference) below or take a look at our [examples](./examples/main.go).

## Style Reference

### Reset

| Code | ANSI Code | Description       |
| ---- | --------: | ----------------- |
| `/`  |       `0` | Resets all styles |

### Formatting

| Code           | ANSI Code | Example                                                                                                       |
| -------------- | --------: | ------------------------------------------------------------------------------------------------------------- |
| `bold`         |       `1` | <span style="font-weight:bold;">Sets the font weight to bold</span>                                           |
| `faint`, `dim` |       `2` | <span style="color:#d3d7cf;">Sets the text brightness to its faint/dim variant</span> \*                      |
| `italic`       |       `3` | <span style="font-style:italic;">Sets the font style to italic</span>                                         |
| `underline`    |       `4` | <span style="text-decoration:underline;">Sets the text decoration to underline</span>                         |
| `blink`        |       `5` | <span style="text-decoration:blink;">Sets the text to blink in and out</span>                                 |
| `invert`       |       `7` | <span style="backdrop-filter:invert(1);filter:invert(1);">Inverts the foreground and background colors</span> |
| `hidden`       |       `8` | <span style="opacity:0;">Sets the text to be hidden</span>                                                    |
| `strike`       |       `9` | <span style="text-decoration:line-through;">Sets the text decoration to line-through</span>                   |

> \* `faint` and `dim` do not work when using 8-bit (256) or 24-bit (RGB) color modes.

| Code             | ANSI Code | Example                                 |
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

| Code      |    ANSI Code | Reset Code                                                                                                                                                                     |
| --------- | -----------: | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `black`   |         `30` | <span style="color:#555753;">Sets the text foreground to black</span>                                                                                                          |
| `red`     |         `31` | <span style="color:#ef2929;">Sets the text foreground to red</span>                                                                                                            |
| `green`   |         `32` | <span style="color:#8ae234;">Sets the text foreground to green</span>                                                                                                          |
| `yellow`  |         `33` | <span style="color:#fce94f;">Sets the text foreground to yellow</span>                                                                                                         |
| `blue`    |         `34` | <span style="color:#729fcf;">Sets the text foreground to blue</span>                                                                                                           |
| `magenta` |         `35` | <span style="color:#ad7fa8;">Sets the text foreground to magenta</span>                                                                                                        |
| `cyan`    |         `36` | <span style="color:#34e2e2;">Sets the text foreground to cyan</span>                                                                                                           |
| `white`   |         `37` | <span style="color:#eeeeec;">Sets the text foreground to white</span>                                                                                                          |
| `R,G,B`   | `38;2;R;G;B` | Sets the text foreground to the given [24-bit (<span style="color:#ef2929;">R</span><span style="color:#8ae234;">G</span><span style="color:#729fcf;">B</span>) color][24-bit] |
| `N`       |     `38;5;N` | Sets the text foreground to the given [8-bit (256) color][8-bit]                                                                                                               |
| `/fg`     |         `39` | <span style="color:auto;">Resets the text foreground to the default color</span>                                                                                               |

### Background Colors

| Code         |    ANSI Code | Reset Code                                                                                                                                                                     |
| ------------ | -----------: | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `bg-black`   |         `40` | <span style="background-color:#555753;">Sets the text background to black</span>                                                                                               |
| `bg-red`     |         `41` | <span style="background-color:#ef2929;">Sets the text background to red</span>                                                                                                 |
| `bg-green`   |         `42` | <span style="background-color:#8ae234;color:#555753;">Sets the text background to green</span>                                                                                 |
| `bg-yellow`  |         `43` | <span style="background-color:#fce94f;color:#555753;">Sets the text background to yellow</span>                                                                                |
| `bg-blue`    |         `44` | <span style="background-color:#729fcf;">Sets the text background to blue</span>                                                                                                |
| `bg-magenta` |         `45` | <span style="background-color:#ad7fa8;">Sets the text background to magenta</span>                                                                                             |
| `bg-cyan`    |         `46` | <span style="background-color:#34e2e2;color:#555753;">Sets the text background to cyan</span>                                                                                  |
| `bg-white`   |         `47` | <span style="background-color:#eeeeec;color:#555753;">Sets the text background to white</span>                                                                                 |
| `bg-R,G,B`   | `48;2;R;G;B` | Sets the text background to the given [24-bit (<span style="color:#ef2929;">R</span><span style="color:#8ae234;">G</span><span style="color:#729fcf;">B</span>) color][24-bit] |
| `bg-N`       |     `48;5;N` | Sets the text background to the given [8-bit (256) color][8-bit]                                                                                                               |
| `/bg`        |         `49` | <span style="background-color:auto;">Resets the text background to the default color</span>                                                                                    |

[8-bit]: https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
[24-bit]: https://en.wikipedia.org/wiki/ANSI_escape_code#24-bit
