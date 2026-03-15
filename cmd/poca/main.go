package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/spf13/pflag"
	"golang.org/x/term"
)

var version = "0.0.0" // Set at build time

var (
	colorFlag        = pflag.StringP("color", "c", "", "Text color")
	styleFlags       = pflag.StringSliceP("style", "s", nil, "Text style (can be repeated)")
	symbolFlag       = pflag.StringP("symbol", "S", "", "Prepend a symbol")
	sectionStartFlag = pflag.Bool("section-start", false, "Print section start divider")
	sectionEndFlag   = pflag.Bool("section-end", false, "Print section end divider")
	headerFlag       = pflag.BoolP("header", "h", false, "Print header (alias for section-start)")
	footerFlag       = pflag.BoolP("footer", "f", false, "Print footer (alias for section-end)")
	presetFlag       = pflag.StringP("preset", "p", "default", "Preset style name")
	widthFlag        = pflag.IntP("width", "w", 0, "Override terminal width")
	versionFlag      = pflag.Bool("version", false, "Print version and exit")
)

// renderLine builds a single line from a LineTemplate, filling to the target width.
// If text is non-empty, it is placed after the prefix with fill on either side per alignment.
func renderLine(template LineTemplate, text string, width int, align string) string {
	prefixLen := utf8.RuneCountInString(template.Prefix)
	suffixLen := utf8.RuneCountInString(template.Suffix)
	endCapLen := utf8.RuneCountInString(template.EndCap)
	fillSpace := width - prefixLen - suffixLen - endCapLen

	if fillSpace < 0 {
		fillSpace = 0
	}

	if text == "" {
		fill := strings.Repeat(template.FillChar, fillSpace)
		return template.Prefix + fill + template.Suffix + template.EndCap
	}

	textWithPad := " " + text + " "
	textLen := utf8.RuneCountInString(textWithPad)
	remaining := fillSpace - textLen

	if remaining < 0 {
		remaining = 0
	}

	var fill string
	switch align {
	case "center":
		leftFill := remaining / 2
		rightFill := remaining - leftFill
		fill = strings.Repeat(template.FillChar, leftFill) + textWithPad + strings.Repeat(template.FillChar, rightFill)
	case "right":
		fill = strings.Repeat(template.FillChar, remaining) + textWithPad
	default: // left
		fill = template.FillChar + textWithPad + strings.Repeat(template.FillChar, remaining-1)
	}

	return template.Prefix + fill + template.Suffix + template.EndCap
}

// renderSection renders a full section (header or footer) using its template.
func renderSection(section SectionTemplate, text string, width int) string {
	var lines []string

	if section.LeadLine != nil {
		lines = append(lines, renderLine(*section.LeadLine, "", width, section.TextAlign))
	}

	lines = append(lines, renderLine(section.MainLine, text, width, section.TextAlign))

	return strings.Join(lines, "\n")
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

	// Env vars as defaults when flags aren't explicitly set
	if !pflag.CommandLine.Changed("preset") {
		if env := os.Getenv("POCA_PRESET"); env != "" {
			*presetFlag = env
		}
	}

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

	preset, ok := presets[normalize(*presetFlag)]
	if !ok {
		fmt.Fprintf(os.Stderr, "Unknown preset: %s\n", *presetFlag)
		os.Exit(1)
	}

	message := strings.Join(pflag.Args(), " ")
	width := determineWidth()

	if *sectionStartFlag {
		if *colorFlag == "" && preset.Header.DefaultColor != "" {
			*colorFlag = preset.Header.DefaultColor
		}
		message = renderSection(preset.Header, message, width)
	} else if *sectionEndFlag {
		if *colorFlag == "" && preset.Footer.DefaultColor != "" {
			*colorFlag = preset.Footer.DefaultColor
		}
		message = renderSection(preset.Footer, "", width)
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
