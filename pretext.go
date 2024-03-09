package pretext

import (
	"fmt"
	"regexp"
)

func Emit(text string, options ...option) string {
	for _, option := range options {
		text = option(text)
	}
	return fmt.Sprintf(`<pre style="font-family: monospace;">%s</pre>`, text)
}

var Options options

type option func(string) string

type options struct{}

func (options) Style(rawPattern, key, value string) option {
	pattern, err := regexp.Compile(rawPattern)
	if err != nil {
		return func(s string) string { return err.Error() }
	}
	return func(text string) string {
		return pattern.ReplaceAllStringFunc(text, func(match string) string {
			return fmt.Sprintf(`<span style="%s: %s;">%s</span>`, key, value, match)
		})
	}
}
func (options) Colorize(re, color string) option { return Options.Style(re, "color", color) }
func (options) Embolden(re string) option        { return Options.Style(re, "font-weight", "bold") }
func (options) Italicize(re string) option       { return Options.Style(re, "font-style", "italic") }
func (options) Underline(re string) option       { return Options.Style(re, "text-decoration", "underline") }
