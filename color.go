package pretext

var Colors colors

// colors inspired by 'The 12-bit rainbow palette', by Kate Morley:
// https://iamkate.com/data/12-bit-rainbow/
type colors struct{}

func (colors) DeepMaroon() string { return "#817" }
func (colors) Maroon() string     { return "#a35" }
func (colors) Red() string        { return "#c66" }
func (colors) Orange() string     { return "#e94" }
func (colors) Yellow() string     { return "#ed0" }
func (colors) LightGreen() string { return "#9d5" }
func (colors) Green() string      { return "#4d8" }
func (colors) Aqua() string       { return "#2cb" }
func (colors) SkyBlue() string    { return "#0bc" }
func (colors) Blue() string       { return "#09c" }
func (colors) DarkBlue() string   { return "#36b" }
func (colors) Purple() string     { return "#639" }
