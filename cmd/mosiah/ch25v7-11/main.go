package main

import (
	"fmt"

	"github.com/mdw-go/pretext"
)

const text = `
 5 And it came to pass that Mosiah did read...the records of Zeniff to his people...
     from the time they left the land of Zarahemla until they returned again.
 6 And he also read the account of Alma and his brethren, and all their afflictions, 
     from the time they left the land of Zarahemla until the time they returned again.
 7 And now, when Mosiah had made an end of reading the records...
     his people were struck with wonder and amazement.
     For they knew not what to think; 
 8 when they beheld those that had been delivered out of bondage 
     they were filled with exceedingly great joy.
 9 when they thought of their brethren who had been slain by the Lamanites 
     they were filled with sorrow, and even shed many tears of sorrow.
10 when they thought of the immediate goodness of God in delivering [their] brethren out of...bondage,
     they did raise their voices and give thanks to God.
11 when they thought upon the Lamanites, who were their brethren, of their sinful and polluted state, 
     they were filled with pain and anguish for the welfare of their souls.
`

func main() {
	o := pretext.Options
	c := pretext.Colors
	fmt.Println(
		pretext.Emit(text,
			o.Embolden("struck"),
			o.Embolden("filled"),
			o.Embolden("did raise"),

			o.Italicize("For they knew not what to think"),

			o.Embolden("wonder"), o.Colorize("wonder", c.Purple()),
			o.Embolden("amazement"), o.Colorize("amazement", c.Purple()),
			o.Embolden("great joy"), o.Colorize("great joy", c.Blue()),
			o.Embolden("sorrow"), o.Colorize("sorrow", c.Maroon()),
			o.Embolden("thanks"), o.Colorize("thanks", c.Orange()),
			o.Embolden("pain"), o.Colorize("pain", c.Red()),
			o.Embolden("anguish"), o.Colorize("anguish", c.Red()),
		),
	)
}
