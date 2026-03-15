package main

// LineTemplate - {Prefix}( Message )({FillChar}...){Suffix}{EndCap}
// [─][ Section Title ][───────────────────────────────────────][┐]
type LineTemplate struct {
	Prefix   string
	Suffix   string
	FillChar string
	EndCap   string
}

// SectionTemplate defines the visual structure of a header or footer.
type SectionTemplate struct {
	LeadLine     *LineTemplate // Optional decorative line before the text line
	MainLine     LineTemplate  // The line carrying the text content
	DefaultColor string        // Color applied if none specified by the user
	TextAlign    string        // "left", "center", "right"
}

// Preset combines header and footer settings into a cohesive style for a section.
type Preset struct {
	Header SectionTemplate
	Footer SectionTemplate
}

var presets = map[string]Preset{
	// ─ Section Title ─────────────────────────────────────────┐
	// ... Section Content ...
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
	// ... Section Content ...
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
