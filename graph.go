// Package graph allows rendering an isometric projection of a GitHub user's contribution history.
package graph

import (
	"github.com/teacat/noire"
	"io"
	"math"

	svg "github.com/ajstarks/svgo/float"
	"github.com/akrennmair/slice"
)

// Graph is a graph instance that is used to render a contributions graph.
type Graph struct {
	cols []col
}

type ContributionDay struct {
	Count int
	Color string
}

// NewGraph creates a graph instance with the contributions-per-day that are
// passed. Graph.Render can then be used to
// render the graph.
func NewGraph(contribs []ContributionDay) *Graph {
	var cols []col
	var total int

	for k, entry := range contribs {
		total += entry.Count

		z := float64(entry.Count*2 + 2)

		i := float64(k / 7)
		j := float64(k % 7)

		size := float64(10)
		space := float64(12)

		var pts []vector3
		// top
		pts = append(pts, vector3{X: space * i, Y: space * j, Z: -z})
		pts = append(pts, vector3{X: space*i + size, Y: space * j, Z: -z})
		pts = append(pts, vector3{X: space*i + size, Y: space*j + size, Z: -z})
		pts = append(pts, vector3{X: space * i, Y: space*j + size, Z: -z})
		// bottom
		pts = append(pts, vector3{X: space * i, Y: space * j, Z: 0})
		pts = append(pts, vector3{X: space*i + size, Y: space * j, Z: 0})
		pts = append(pts, vector3{X: space*i + size, Y: space*j + size, Z: 0})
		pts = append(pts, vector3{X: space * i, Y: space*j + size, Z: 0})

		cols = append(cols, col{pts: pts, count: entry.Count, color: entry.Color})
	}

	return &Graph{
		cols: cols,
	}
}

// Render writes the generated SVG to the given io.Writer in the appropriate Theme.
func (g *Graph) Render(f io.WriteCloser, theme Theme) error {
	canvas := svg.New(f)

	canvas.Start(840, 400)

	for _, c := range g.cols {
		// each column has 3 visible faces.
		// p1 is the outline of the column (all 3 visible faces), p2 is the bottom 2 faces and 3 is the bottom right face
		// they are rendered that way to avoid weird spacing artifacts between the faces
		p1 := []vector3{c.pts[0], c.pts[1], c.pts[5], c.pts[6], c.pts[7], c.pts[3]}
		p2 := []vector3{c.pts[3], c.pts[2], c.pts[1], c.pts[5], c.pts[6], c.pts[7]}
		p3 := []vector3{c.pts[2], c.pts[1], c.pts[5], c.pts[6]}

		// project the 3d coordinates to 2d space using an isometric projection
		p1iso := isometricProjection(p1)
		p2iso := isometricProjection(p2)
		p3iso := isometricProjection(p3)

		// horizontal & vertical offsets to center the whole chart
		h, v := float64(180), float64(25)

		// colors used for the 3 visible faces
		c1, c2, c3 := faceColors(c.color, theme)

		xs := slice.Map(p1iso, func(vec vector2) float64 { return vec.X + h })
		ys := slice.Map(p1iso, func(vec vector2) float64 { return vec.Y + v })
		canvas.Polygon(xs, ys, "fill:"+c1)

		xs = slice.Map(p2iso, func(vec vector2) float64 { return vec.X + h })
		ys = slice.Map(p2iso, func(vec vector2) float64 { return vec.Y + v })
		canvas.Polygon(xs, ys, "fill:"+c2)

		xs = slice.Map(p3iso, func(vec vector2) float64 { return vec.X + h })
		ys = slice.Map(p3iso, func(vec vector2) float64 { return vec.Y + v })
		canvas.Polygon(xs, ys, "fill:"+c3)
	}

	canvas.End()

	return f.Close()
}

func isometricProjection(v []vector3) []vector2 {
	return slice.Map(v, func(v vector3) vector2 {
		x, y := spaceToIso(v.X, v.Y, v.Z)

		return vector2{
			X: x,
			Y: y,
		}
	})
}

func spaceToIso(x, y, z float64) (h, v float64) {
	x, y = x+z, y+z

	h = (x - y) * math.Sqrt(3) / 2
	v = (x + y) / 2

	return h, v
}

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

	ColorCustom0 = "#161b22"
	ColorCustom1 = "#631c03"
	ColorCustom2 = "#bd561d"
	ColorCustom3 = "#fa7a18"
	ColorCustom4 = "#fddf68"
)

func faceColors(color string, theme Theme) (string, string, string) {
	switch theme {
	case Dark:
		switch color {
		case Color0:
			return ColorDark0, "#030a12", "#171e26"
		case Color1:
			return ColorDark1, "#001b00", "#002f12"
		case Color2:
			return ColorDark2, "#004307", "#00571b"
		case Color3:
			return ColorDark3, "#007d1a", "#11912e"
		case Color4:
			return ColorDark4, "#10a92c", "#24bd40"
		}
		return "#2d333b", "#030a12", "#171e26"
	case DarkHalloween:
		switch color {
		case HalloweenColor0:
			return HalloweenDarkColor0, "#000000", "#00050c"
		case HalloweenColor1:
			return HalloweenDarkColor1, "#390000", "#4d0800"
		case HalloweenColor2:
			return HalloweenDarkColor2, "#942d00", "#a84109"
		case HalloweenColor3:
			return HalloweenDarkColor3, "#d05000", "#e46403"
		case HalloweenColor4:
			return HalloweenDarkColor4, "#d3b53e", "#e7c952"
		}
		return "#2d333b", "#030a12", "#171e26"
	case Custom:
		c1 := noire.NewHex(color)
		switch color {
		case Color0:
			c1 = noire.NewHex(ColorCustom0)
		case Color1:
			c1 = noire.NewHex(ColorCustom1)
		case Color2:
			c1 = noire.NewHex(ColorCustom2)
		case Color3:
			c1 = noire.NewHex(ColorCustom3)
		case Color4:
			c1 = noire.NewHex(ColorCustom4)
		}
		return c1.HTML(), c1.Shade(0.35).HTML(), c1.Shade(0.2).HTML()
	case LightHalloween:
		switch color {
		case HalloweenColor0:
			return HalloweenColor0, "#c2c5c8", "#d6d9dc"
		case HalloweenColor1:
			return HalloweenColor1, "#d5c522", "#e9d936"
		case HalloweenColor2:
			return HalloweenColor2, "#d59d00", "#e9b100"
		case HalloweenColor3:
			return HalloweenColor3, "#d56e00", "#e98200"
		case HalloweenColor4:
			return HalloweenColor4, "#000001", "#000007"
		}
		return "#ebedf0", "#c2c5c8", "#d6d9dc"
	default:
		fallthrough
	case Light:
		switch color {
		case Color0:
			return Color0, "#c2c5c8", "#d6d9dc"
		case Color1:
			return Color1, "#73c080", "#87d494"
		case Color2:
			return Color2, "#199b3c", "#2daf50"
		case Color3:
			return Color3, "#077725", "#1b8b39"
		case Color4:
			return Color4, "#004410", "#0c5824"
		}
		return "#ebedf0", "#c2c5c8", "#d6d9dc"
	}
}
