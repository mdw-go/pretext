package pretext

import (
	"fmt"
	"regexp"
)

func Emit(text string, options ...option) string {
	for _, option := range options {
		text = option(text)
	}
	return `<pre style="font-family: monospace;">` + text + "</pre>"
}

var Options options

type option func(string) string

type options struct{}

func (options) Embolden(rawPattern string) option {
	pattern, err := regexp.Compile(rawPattern)
	if err != nil {
		return func(s string) string { return err.Error() }
	}
	return func(text string) string {
		return pattern.ReplaceAllStringFunc(text, func(match string) string {
			return `<span style="font-weight: bold;">` + match + "</span>"
		})
	}
}
func (options) Italicize(rawPattern string) option {
	pattern, err := regexp.Compile(rawPattern)
	if err != nil {
		return func(s string) string { return err.Error() }
	}
	return func(text string) string {
		return pattern.ReplaceAllStringFunc(text, func(match string) string {
			return `<span style="font-style: italic;">` + match + "</span>"
		})
	}
}
func (options) Underline(rawPattern string) option {
	pattern, err := regexp.Compile(rawPattern)
	if err != nil {
		return func(s string) string { return err.Error() }
	}
	return func(text string) string {
		return pattern.ReplaceAllStringFunc(text, func(match string) string {
			return `<span style="text-decoration: underline;">` + match + "</span>"
		})
	}
}
func (options) Colorize(rawPattern string, color Color) option {
	pattern, err := regexp.Compile(rawPattern)
	if err != nil {
		return func(s string) string { return err.Error() }
	}
	return func(text string) string {
		return pattern.ReplaceAllStringFunc(text, func(match string) string {
			return fmt.Sprintf(`<span style="color: %s;">`, color) + match + "</span>"
		})
	}
}
