package pretext_test

import (
	"testing"

	"github.com/mdwhatcott/pretext"
)

const sampleText = `

   Lorem ipsum dolor sit amet,
  consectetur adipiscing elit,
sed do eiusmod tempor incididunt
ut labore et dolore magna aliqua.
    Ut enim ad minim veniam,
quis nostrud exercitation ullamco
  laboris nisi ut aliquip ex ea
        commodo consequat.

`

func Test(t *testing.T) {
	result := pretext.Emit(sampleText,
		pretext.Options.Underline(`ipsum`),
		pretext.Options.Embolden(`ut`),
		pretext.Options.Italicize(`consequat`),
		pretext.Options.Colorize(`eiusmod`, pretext.Colors.Blue()),
	)
	t.Log(result)
}
