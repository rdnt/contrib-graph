package graph

const (
	Color0 = "#ebedf0"
	Color1 = "#9be9a8"
	Color2 = "#40c463"
	Color3 = "#30a14e"
	Color4 = "#216e39"

	ColorDark0 = "#2d333b"
	ColorDark1 = "#0e4429"
	ColorDark2 = "#006d32"
	ColorDark3 = "#26a641"
	ColorDark4 = "#39d353"

	HalloweenColor0 = "#ebedf0"
	HalloweenColor1 = "#ffee4a"
	HalloweenColor2 = "#ffc501"
	HalloweenColor3 = "#fe9600"
	HalloweenColor4 = "#03001c"

	HalloweenDarkColor0 = "#161b22"
	HalloweenDarkColor1 = "#631c03"
	HalloweenDarkColor2 = "#bd561d"
	HalloweenDarkColor3 = "#fa7a18"
	HalloweenDarkColor4 = "#fddf68"
)

type Theme func(string) string

var DarkTheme = func(color string) string {
	switch color {
	case Color0:
		return ColorDark0
	case Color1:
		return ColorDark1
	case Color2:
		return ColorDark2
	case Color3:
		return ColorDark3
	case Color4:
		return ColorDark4
	default:
		return "#2d333b"
	}
}

var LightTheme = func(color string) string {
	switch color {
	case Color0:
		return Color0
	case Color1:
		return Color1
	case Color2:
		return Color2
	case Color3:
		return Color3
	case Color4:
		return Color4
	}
	return "#ebedf0"
}

var HalloweenDarkTheme = func(color string) string {
	switch color {
	case Color0:
		return HalloweenDarkColor0
	case Color1:
		return HalloweenDarkColor1
	case Color2:
		return HalloweenDarkColor2
	case Color3:
		return HalloweenDarkColor3
	case Color4:
		return HalloweenDarkColor4
	}
	return "#2d333b"
}

var HalloweenLightTheme = func(color string) string {
	switch color {
	case Color0:
		return HalloweenColor0
	case Color1:
		return HalloweenColor1
	case Color2:
		return HalloweenColor2
	case Color3:
		return HalloweenColor3
	case Color4:
		return HalloweenColor4
	}
	return "#ebedf0"
}
