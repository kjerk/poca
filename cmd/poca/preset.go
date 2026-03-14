package main

// LineTemplate defines how a single line of a section is rendered.
type LineTemplate struct {
	Prefix   string // Fixed string at the start of the line
	Suffix   string // Fixed string at the end of the line
	FillChar string // Character used to fill remaining width
	EndCap   string // Single character at the far right (after fill)
}

// SectionTemplate defines the visual structure of a header or footer.
type SectionTemplate struct {
	LeadLine     *LineTemplate // Optional decorative line before the text line
	MainLine     LineTemplate  // The line carrying the text content
	DefaultColor string        // Color applied if none specified by the user
	TextAlign    string        // "left", "center", "right"
}

// Preset is a named pair of header and footer styles.
type Preset struct {
	Header SectionTemplate
	Footer SectionTemplate
}

var presets = map[string]Preset{
	// ─ Section Title ─────────────────────────────────────────┐
	// ─────────────────────────────────────────────────────────┘
	"default": {
		Header: SectionTemplate{
			MainLine: LineTemplate{
				FillChar: "─",
				EndCap:   "┐",
			},
			TextAlign: "left",
		},
		Footer: SectionTemplate{
			MainLine: LineTemplate{
				FillChar: "─",
				EndCap:   "┘",
			},
			DefaultColor: "blue",
			TextAlign:    "left",
		},
	},

	// ━ Section Title ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
	"bold": {
		Header: SectionTemplate{
			MainLine: LineTemplate{
				FillChar: "━",
				EndCap:   "┓",
			},
			TextAlign: "left",
		},
		Footer: SectionTemplate{
			MainLine: LineTemplate{
				FillChar: "━",
				EndCap:   "┛",
			},
			DefaultColor: "blue",
			TextAlign:    "left",
		},
	},

	// ═ Section Title ═════════════════════════════════════════╗
	// ═════════════════════════════════════════════════════════╝
	"double": {
		Header: SectionTemplate{
			MainLine: LineTemplate{
				FillChar: "═",
				EndCap:   "╗",
			},
			TextAlign: "left",
		},
		Footer: SectionTemplate{
			MainLine: LineTemplate{
				FillChar: "═",
				EndCap:   "╝",
			},
			DefaultColor: "blue",
			TextAlign:    "left",
		},
	},

	// ═══ Section Title ════════════════════════════════════════
	// ══════════════════════════════════════════════════════════
	"simple": {
		Header: SectionTemplate{
			MainLine: LineTemplate{
				FillChar: "═",
			},
			TextAlign: "left",
		},
		Footer: SectionTemplate{
			MainLine: LineTemplate{
				FillChar: "═",
			},
			DefaultColor: "blue",
			TextAlign:    "left",
		},
	},

	// ─────────────────────────────────────────────────────────╮
	// ─ Section Title ─────────────────────────────────────────┤
	// ─────────────────────────────────────────────────────────╯
	"boxed": {
		Header: SectionTemplate{
			LeadLine: &LineTemplate{
				FillChar: "─",
				EndCap:   "╮",
			},
			MainLine: LineTemplate{
				FillChar: "─",
				EndCap:   "┤",
			},
			TextAlign: "left",
		},
		Footer: SectionTemplate{
			MainLine: LineTemplate{
				FillChar: "─",
				EndCap:   "╯",
			},
			DefaultColor: "blue",
			TextAlign:    "left",
		},
	},

	// ░░░▒▒▒▓▓▓███━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━███▓▓▓▒▒▒░░░
	// ░░░▒▒▒▓▓▓███  Section Title                ███▓▓▓▒▒▒░░░
	// ░░░▒▒▒▓▓▓███━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━███▓▓▓▒▒▒░░░
	"gradient": {
		Header: SectionTemplate{
			LeadLine: &LineTemplate{
				Prefix:   "░░░▒▒▒▓▓▓███",
				Suffix:   "███▓▓▓▒▒▒░░░",
				FillChar: "━",
			},
			MainLine: LineTemplate{
				Prefix:   "░░░▒▒▒▓▓▓███",
				Suffix:   "███▓▓▓▒▒▒░░░",
				FillChar: " ",
			},
			TextAlign: "left",
		},
		Footer: SectionTemplate{
			MainLine: LineTemplate{
				Prefix:   "░░░▒▒▒▓▓▓███",
				Suffix:   "███▓▓▓▒▒▒░░░",
				FillChar: "━",
			},
			DefaultColor: "blue",
			TextAlign:    "left",
		},
	},
}
