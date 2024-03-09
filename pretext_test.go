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

const sampleText2 = `
And the Spirit of the Lord shall rest upon him,
    the spirit of wisdom and understanding,
    the spirit of counsel and might,
    the spirit of knowledge and the fear of the Lord,
And shall make him
      of quick understanding in the fear of the Lord,
and he shall not judge after the sight   of his eyes,
       neither reprove after the hearing of his ears.
But with righteousness shall he judge the poor,
and reprove with equity for the meek of the earth...
     and righteousness shall be the girdle of his loins,
     and faithfulness           the girdle of his reins.
`

func Test2(t *testing.T) {
	t.Log(pretext.Emit(sampleText2,
		pretext.Options.Colorize(`Spirit|spirit`, pretext.Colors.Blue()),
		pretext.Options.Colorize(`fear of the Lord`, pretext.Colors.Red()),
		pretext.Options.Colorize(`understanding`, pretext.Colors.LightGreen()),
		pretext.Options.Colorize(`judge|reprove`, pretext.Colors.Maroon()),
		pretext.Options.Colorize(`righteousness`, pretext.Colors.Green()),
		pretext.Options.Colorize(`faithfulness`, pretext.Colors.Green()),
		pretext.Options.Colorize(`meek|poor`, pretext.Colors.Orange()),
		pretext.Options.Colorize(`earth`, pretext.Colors.Aqua()),
	))
}
