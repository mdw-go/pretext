package main

import (
	"fmt"

	"github.com/mdw-go/pretext"
)

const text = `
INSERT FORMATTED TEXT HERE
`

func main() {
	fmt.Println(
		pretext.Emit(text,
			pretext.Options.Colorize("stuff", pretext.Colors.Blue()),
		),
	)
}
