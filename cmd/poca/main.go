package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/spf13/pflag"
	"golang.org/x/term"
)

var version = "0.0.0" // Build time

const (
	fillChar         = "─"
	sectionOpenChar  = "┐"
	sectionCloseChar = "┘"
)

var (
	colorFlag        = pflag.String("color", "", "Text color")
	styleFlags       = pflag.StringSlice("style", nil, "Text style (can be repeated)")
	symbolFlag       = pflag.String("symbol", "", "Prepend a symbol")
	sectionStartFlag = pflag.Bool("section-start", false, "Print section start divider")
	sectionEndFlag   = pflag.Bool("section-end", false, "Print section end divider")
	headerFlag       = pflag.BoolP("header", "h", false, "Print header (alias for section-start)")
	footerFlag       = pflag.Bool("footer", false, "Print footer (alias for section-end)")
	widthFlag        = pflag.Int("width", 0, "Override terminal width")
	versionFlag      = pflag.Bool("version", false, "Print version and exit")
)

func buildSectionDivider(message string, width int, terminator string) string {
	if message == "" {
		return strings.Repeat(fillChar, width-1) + terminator
	}

	prefix := fillChar + " " + message + " " + fillChar
	remaining := width - utf8.RuneCountInString(prefix) - 1
	if remaining > 0 {
		return prefix + strings.Repeat(fillChar, remaining) + terminator
	}
	return prefix[:width-1] + terminator
}

func determineWidth() int {
	if *widthFlag > 0 {
		return *widthFlag
	}

	if term.IsTerminal(int(os.Stdout.Fd())) {
		if w, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil {
			return w
		}
	}

	return 80
}

func main() {
	pflag.Parse()

	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	// Handle aliases
	if *headerFlag {
		*sectionStartFlag = true
	}
	if *footerFlag {
		*sectionEndFlag = true
	}

	message := strings.Join(pflag.Args(), " ")

	width := determineWidth()

	if *sectionStartFlag {
		message = buildSectionDivider(message, width, sectionOpenChar)
	} else if *sectionEndFlag {
		if *colorFlag == "" {
			*colorFlag = "blue"
		}
		message = buildSectionDivider("", width, sectionCloseChar)
	} else if *symbolFlag != "" {
		if sym, ok := symbols[normalize(*symbolFlag)]; ok {
			message = sym + " " + message
		} else {
			fmt.Fprintf(os.Stderr, "Unknown symbol: %s\n", *symbolFlag)
			os.Exit(1)
		}
	}

	isInteractive := term.IsTerminal(int(os.Stdout.Fd()))

	if isInteractive && (*colorFlag != "" || len(*styleFlags) > 0) {
		var codes []uint8

		if *colorFlag != "" {
			if code, ok := colors[normalize(*colorFlag)]; ok {
				codes = append(codes, code)
			} else {
				fmt.Fprintf(os.Stderr, "Unknown color: %s\n", *colorFlag)
				os.Exit(1)
			}
		}

		for _, s := range *styleFlags {
			if code, ok := styles[normalize(s)]; ok {
				codes = append(codes, code)
			} else {
				fmt.Fprintf(os.Stderr, "Unknown style: %s\n", s)
				os.Exit(1)
			}
		}

		fmt.Println(ansiWrap(message, codes))
	} else {
		fmt.Println(message)
	}
}
