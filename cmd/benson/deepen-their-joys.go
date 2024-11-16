package main

import (
	"fmt"

	"github.com/mdw-go/pretext"
)

const text = `
"Men and women who turn their lives 
over to God will find out that He can 
 make a lot more out of their lives 
      than they can. He will

       deepen their joys,
       expand their vision,
      quicken their minds,
   strengthen their muscles,
         lift their spirits,
     multiply their blessings,
     increase their opportunities,
      comfort their souls,
        raise  up   friends, and
         pour  out  peace.
   
  Whoever will lose his life in God 
   will find he has eternal life."

        <a href="https://www.churchofjesuschrist.org/study/ensign/1988/12/jesus-christ-gifts-and-expectations?lang=eng">―Ezra Taft Benson</a>
`

func main() {
	fmt.Println(
		pretext.Emit(text,
			pretext.Options.Colorize("Men", pretext.Colors.Blue()),
			pretext.Options.Colorize("women", pretext.Colors.Red()),

			pretext.Options.Colorize("God", pretext.Colors.Green()),
			pretext.Options.Embolden("God"),
			pretext.Options.Embolden("He"),
			pretext.Options.Colorize("He", pretext.Colors.Green()),
			pretext.Options.Italicize("a lot more"),

			pretext.Options.Embolden("deepen"),
			pretext.Options.Italicize("expand"),
			pretext.Options.Underline("quicken"),
			pretext.Options.Embolden("strengthen"),
			pretext.Options.Italicize("lift"),
			pretext.Options.Underline("multiply"),
			pretext.Options.Embolden("increase"),
			pretext.Options.Italicize("comfort"),
			pretext.Options.Underline("raise"),
			pretext.Options.Embolden("pour"),

			pretext.Options.Colorize("joys", pretext.Colors.Blue()),
			pretext.Options.Colorize("vision", pretext.Colors.Yellow()),
			pretext.Options.Colorize("minds", pretext.Colors.Maroon()),
			pretext.Options.Colorize("muscles", pretext.Colors.Orange()),
			pretext.Options.Colorize("spirits", pretext.Colors.Green()),
			pretext.Options.Colorize("blessings", pretext.Colors.Purple()),
			pretext.Options.Colorize("opportunities", pretext.Colors.Aqua()),
			pretext.Options.Colorize("souls", pretext.Colors.LightGreen()),
			pretext.Options.Colorize("friends", pretext.Colors.Red()),
			pretext.Options.Colorize("peace", pretext.Colors.DarkBlue()),
			pretext.Options.Colorize(`their`, "grey"),

			pretext.Options.Colorize("lose", "grey"),
			pretext.Options.Embolden("eternal life"),
			pretext.Options.Colorize("eternal life", pretext.Colors.Green()),
		),
	)
}
