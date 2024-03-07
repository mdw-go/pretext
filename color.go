package pretext

type Color string

var Colors colors

// colors inspired by 'The 12-bit rainbow palette', by Kate Morley:
// https://iamkate.com/data/12-bit-rainbow/
type colors struct{}

func (colors) DeepMaroon() Color { return "#817" }
func (colors) Maroon() Color     { return "#a35" }
func (colors) Red() Color        { return "#c66" }
func (colors) Orange() Color     { return "#e94" }
func (colors) Yellow() Color     { return "#ed0" }
func (colors) LightGreen() Color { return "#9d5" }
func (colors) Green() Color      { return "#4d8" }
func (colors) Aqua() Color       { return "#2cb" }
func (colors) SkyBlue() Color    { return "#0bc" }
func (colors) Blue() Color       { return "#09c" }
func (colors) DarkBlue() Color   { return "#36b" }
func (colors) Purple() Color     { return "#639" }
