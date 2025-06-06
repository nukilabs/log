package log

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Styles defines the styles for the text logger.
type Styles struct {
	// Timestamp is the style for timestamps.
	Timestamp lipgloss.Style

	// Index is the style for indices.
	Index lipgloss.Style

	// Prefix is the style for prefix.
	Prefix lipgloss.Style

	// Message is the style for messages.
	Message lipgloss.Style

	// Key is the style for keys.
	Key lipgloss.Style

	// Value is the style for values.
	Value lipgloss.Style

	// Separator is the style for separators.
	Separator lipgloss.Style

	// Levels are the styles for each level.
	Levels map[Level]lipgloss.Style

	// Keys overrides styles for specific keys.
	Keys map[string]lipgloss.Style

	// Values overrides value styles for specific keys.
	Values map[string]lipgloss.Style
}

// DefaultStyles returns the default styles.
func DefaultStyles() *Styles {
	return &Styles{
		Timestamp: lipgloss.NewStyle(),
		Index:     lipgloss.NewStyle().Faint(true),
		Prefix:    lipgloss.NewStyle().Bold(true).Faint(true),
		Message:   lipgloss.NewStyle(),
		Key:       lipgloss.NewStyle().Faint(true),
		Value:     lipgloss.NewStyle(),
		Separator: lipgloss.NewStyle().Faint(true),
		Levels: map[Level]lipgloss.Style{
			DebugLevel: lipgloss.NewStyle().
				SetString(strings.ToUpper(DebugLevel.String())).
				Bold(true).
				MaxWidth(4).
				Foreground(lipgloss.Color("146")),
			HintLevel: lipgloss.NewStyle().
				SetString(strings.ToUpper(HintLevel.String())).
				Bold(true).
				MaxWidth(4).
				Foreground(lipgloss.Color("75")),
			InfoLevel: lipgloss.NewStyle().
				SetString(strings.ToUpper(InfoLevel.String())).
				Bold(true).
				MaxWidth(4).
				Foreground(lipgloss.Color("86")),
			WarnLevel: lipgloss.NewStyle().
				SetString(strings.ToUpper(WarnLevel.String())).
				Bold(true).
				MaxWidth(4).
				Foreground(lipgloss.Color("157")),
			CartLevel: lipgloss.NewStyle().
				SetString(strings.ToUpper(CartLevel.String())).
				Bold(true).
				MaxWidth(4).
				Foreground(lipgloss.Color("153")),
			MissLevel: lipgloss.NewStyle().
				SetString(strings.ToUpper(MissLevel.String())).
				Bold(true).
				MaxWidth(4).
				Foreground(lipgloss.Color("219")),
			ErrorLevel: lipgloss.NewStyle().
				SetString(strings.ToUpper(ErrorLevel.String())).
				Bold(true).
				MaxWidth(4).
				Foreground(lipgloss.Color("205")),
			DoneLevel: lipgloss.NewStyle().
				SetString(strings.ToUpper(DoneLevel.String())).
				Bold(true).
				MaxWidth(4).
				Foreground(lipgloss.Color("47")),
			FatalLevel: lipgloss.NewStyle().
				SetString(strings.ToUpper(FatalLevel.String())).
				Bold(true).
				MaxWidth(4).
				Foreground(lipgloss.Color("134")),
		},
		Keys:   map[string]lipgloss.Style{},
		Values: map[string]lipgloss.Style{},
	}
}
