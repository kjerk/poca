package main

import (
	"strconv"
	"strings"
	"unicode"
)

const ansiReset = "\x1b[0m"

var colors = map[string]uint8{
	"black":        30,
	"red":          31,
	"green":        32,
	"yellow":       33,
	"blue":         34,
	"magenta":      35,
	"cyan":         36,
	"white":        37,
	"darkgray":     90,
	"lightred":     91,
	"lightgreen":   92,
	"lightyellow":  93,
	"lightblue":    94,
	"lightmagenta": 95,
	"lightcyan":    96,
	"lightwhite":   97,
}

var styles = map[string]uint8{
	"bold":          1,
	"dim":           2,
	"italic":        3,
	"underline":     4,
	"blink":         5,
	"reverse":       7,
	"hidden":        8,
	"strikethrough": 9,
}

var symbols = map[string]string{
	"check":   "✓",
	"cross":   "✗",
	"arrow":   "→",
	"warning": "⚠",
	"info":    "ℹ",
}

// strip non-letters and lowercases, "dark-gray", "DarkGray", "dark_gray" all become "darkgray".
func normalize(input string) string {
	var builder strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) {
			builder.WriteRune(unicode.ToLower(r))
		}
	}
	return builder.String()
}

// ansiWrap wraps a message with ANSI escape codes built from the given code list.
func ansiWrap(message string, codes []uint8) string {
	if len(codes) == 0 {
		return message
	}

	parts := make([]string, len(codes))
	for i, code := range codes {
		parts[i] = strconv.FormatUint(uint64(code), 10)
	}

	return "\x1b[" + strings.Join(parts, ";") + "m" + message + ansiReset
}
