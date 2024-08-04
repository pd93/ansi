package main

import (
	"fmt"

	"github.com/pd93/ansi"
)

func main() {
	// Apply a style to a string
	str := ansi.Parse("[magenta]Hello, world![/]")
	fmt.Println(str)

	// Use the Print wrappers to combine fmt.Print and ansi.Parse
	ansi.Println("[blue]Hello, world![/]")

	// Use the Printf wrappers to combine print formatting and ANSI parsing
	ansi.Printf("[cyan]%s[/]\n", "Hello, world!")

	// Apply multiple styles to a string
	ansi.Println("[green:strike]Hello, world![/]")

	// Remove a single style from a string
	ansi.Println("[yellow:strike]Hello[/strike], world![/]")

	// Apply a different styles to different parts of a string
	ansi.Println("[yellow]Hello,[/] [red]world![/]")

	// Apply an 8-bit (256) color
	ansi.Println("[208]Hello, world![/]")

	// Apply a 24-bit (RGB) color
	ansi.Println("[255,0,0]Hello, world![/]")

	// Apply a 24-bit (RGB) background color
	ansi.Println("[bg-255,0,0:bold]Hello, world![/]")

	// Apply a foreground and background color
	ansi.Println("[255,0,0:bg-255,0,0]Hello, world![/]")

	// Escape the style block so it prints as-is
	ansi.Println(`\[red]Hello, world!\[/]`)

	// Strip the styles from a string
	str = ansi.Strip("[red]Hello, world![/]")
	fmt.Println(str)
}
