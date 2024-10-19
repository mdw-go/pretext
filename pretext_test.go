package pretext_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/mdw-go/pretext"
	"github.com/mdw-go/testing/should"
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

const expectedOutput = `<pre style="font-family: monospace; padding: 0.5em;">
   Lorem <span style="text-decoration: underline;">ipsum</span> dolor sit amet,
  consectetur adipiscing elit,
sed do <span style="color: #09c;">eiusmod</span> tempor incididunt
<span style="font-weight: bold;">ut</span> labore et dolore magna aliqua.
    Ut enim ad minim veniam,
quis nostrud exercitation ullamco
  laboris nisi <span style="font-weight: bold;">ut</span> aliquip ex ea
        commodo <span style="font-style: italic;">consequat</span>.
</pre>`

func Test(t *testing.T) {
	result := pretext.Emit(sampleText,
		pretext.Options.Underline(`ipsum`),
		pretext.Options.Embolden(`ut`),
		pretext.Options.Italicize(`consequat`),
		pretext.Options.Colorize(`eiusmod`, pretext.Colors.Blue()),
	)
	should.So(t, strings.TrimSpace(result), should.Equal, strings.TrimSpace(expectedOutput))
	html("example-lorem-ipsum.html", "Lorem Ipsum", result)
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
const expectedOutput2 = `<pre style="font-family: monospace; padding: 0.5em;">
And the <span style="color: #09c;">Spirit</span> of the Lord shall rest upon him,
    the <span style="color: #09c;">spirit</span> of wisdom and <span style="color: #9d5;">understanding</span>,
    the <span style="color: #09c;">spirit</span> of counsel and might,
    the <span style="color: #09c;">spirit</span> of knowledge and the <span style="color: #c66;">fear of the Lord</span>,
And shall make him
      of quick <span style="color: #9d5;">understanding</span> in the <span style="color: #c66;">fear of the Lord</span>,
and he shall not <span style="color: #a35;">judge</span> after the sight   of his eyes,
       neither <span style="color: #a35;">reprove</span> after the hearing of his ears.
But with <span style="color: #4d8;">righteousness</span> shall he <span style="color: #a35;">judge</span> the <span style="color: #e94;">poor</span>,
and <span style="color: #a35;">reprove</span> with equity for the <span style="color: #e94;">meek</span> of the <span style="color: #2cb;">earth</span>...
     and <span style="color: #4d8;">righteousness</span> shall be the girdle of his loins,
     and <span style="color: #4d8;">faithfulness</span>           the girdle of his reins.
</pre>`

func Test2(t *testing.T) {
	result := pretext.Emit(sampleText2,
		pretext.Options.Colorize(`Spirit|spirit`, pretext.Colors.Blue()),
		pretext.Options.Colorize(`fear of the Lord`, pretext.Colors.Red()),
		pretext.Options.Colorize(`understanding`, pretext.Colors.LightGreen()),
		pretext.Options.Colorize(`judge|reprove`, pretext.Colors.Maroon()),
		pretext.Options.Colorize(`righteousness`, pretext.Colors.Green()),
		pretext.Options.Colorize(`faithfulness`, pretext.Colors.Green()),
		pretext.Options.Colorize(`meek|poor`, pretext.Colors.Orange()),
		pretext.Options.Colorize(`earth`, pretext.Colors.Aqua()),
	)
	should.So(t, strings.TrimSpace(result), should.Equal, strings.TrimSpace(expectedOutput2))
	html("example-2-nephi-21.html", "2 Nephi 21:2-5", result)
	t.Log(result)
}

func html(path, title, text string) {
	content := fmt.Sprintf(`<html><head><title>%s</title></head><body><h1>%s</h1>%s</body></html>`, title, title, text)
	err := os.WriteFile(path, []byte(content), 0600)
	if err != nil {
		panic(err)
	}
}
