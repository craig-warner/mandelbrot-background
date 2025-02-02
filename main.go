package main

import (
	"fmt"
	"image/color"
	"os"
	"time"

	//	"github.com/hjson/hjson-go/v4"
	"encoding/json"

	"math"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	max_size = 10000
)

/*
JSON Structure Defining Color Settings
*/
const (
	all_colors_str = `[
      { "Ibits": 12,
	"blue_pos":    [  4,   5,   6,  7,  -1,  -1,  -1,  -1,  -1,  -1,  -1,  -1],
	"green_pos":   [ -1,  -1,  -1,   4,  5,   6,   7,  -1,  -1,  -1,  -1,  -1],
	"red_pos":     [ -1,  -1,  -1,  -1,  -1,  -1, -1,  -1,   4,   5,   6,   7],
	"default_color": [0,0,0] 
      },
      { "ibits": 9,
	"green_pos":    [  5,   6,   7,  -1,  -1,  -1,  -1,  -1,  -1],
	"blue_pos":   [ -1,  -1,  -1,   5,   6,   7,  -1,  -1,  -1],
	"red_pos":     [ -1,  -1,  -1,  -1,  -1,  -1,   5,   6,   7],
	"default_color": [0,0,0]
      },
      { "ibits": 9,
	"red_pos":    [  5,   6,   7,  -1,  -1,  -1,  -1,  -1,  -1],
	"blue_pos":   [ -1,  -1,  -1,   5,   6,   7,  -1,  -1,  -1],
	"green_pos":     [ -1,  -1,  -1,  -1,  -1,  -1,   5,   6,   7],
	"default_color":[0,0,0]
      },
      { "ibits": 9,
	"blue_pos":    [  3,   4,   5,  -1,  -1,  -1,  -1,  -1,  -1],
	"green_pos":   [ -1,  -1,  -1,   3,   4,   5,  -1,  -1,  -1],
	"red_pos":     [ -1,  -1,  -1,  -1,  -1,  -1,   3,   4,   5],
	"default_color": [0,0,0]
      },
      { "ibits": 9,
	"green_pos":    [  3,   4,   5,  -1,  -1,  -1,  -1,  -1,  -1],
	"blue_pos":   [ -1,  -1,  -1,   3,   4,   5,  -1,  -1,  -1],
	"red_pos":     [ -1,  -1,  -1,  -1,  -1,  -1,   3,   4,   5],
	"default_color": [0,0,0]
      },
      { "ibits": 9,
	"red_pos":    [  3,   4,   5,  -1,  -1,  -1,  -1,  -1,  -1],
	"blue_pos":   [ -1,  -1,  -1,   3,   4,   5,  -1,  -1,  -1],
	"green_pos":     [ -1,  -1,  -1,  -1,  -1,  -1,   3,   4,   5],
	"default_color": [0,0,0]
      },
      { "ibits": 6,
	"blue_pos":    [  2, 3,  4,   5,   6,   7],
	"green_pos":   [ 1,-1,-1, -1,-1,-1,-1],
	"red_pos":     [ -1,-1, -1,-1,-1,-1],
	"default_color": [0,0,0]
      },
      { "ibits": 6,
	"blue_pos":    [   2,   3,   4,  5,  6,  7],
	"red_pos":     [   2,   3,   4,  5,  6,  7],
	"green_pos":    [ -1,  -1,  -1, -1, -1, -1],
	"default_color": [0,0,0]
      },
      { "ibits": 6,
	"blue_pos":    [   2,   3,   4,  5,  6,  7],
	"red_pos":     [   2,   3,   4,  5,  6,  7],
	"green_pos":    [ -1,  -1,  -1, -1, -1, -1],
	"default_color": [113,1,147]
      },
      { "ibits": 6,
	"blue_pos":    [  0,  1,  2, 3, 4, 5],
	"red_pos":     [   2,   3,   4,  5,  6,  7],
	"green_pos":    [  0,  1,  2, 3, 4, 5],
	"default_color": [0,0,0]
	  },
      { "ibits": 6,
	"blue_pos":    [   -1,   -1,  -1,  -1,  -1, -1],
	"red_pos":     [   2,   3,   4,  5,  6,  7],
	"green_pos":    [  0,  1,  2, 3, 4, 5],
	"default_color": [0,0,0]
      },
      { "ibits": 6,
	"red_pos":    [  0,  1,  2, 3, 4, 5],
	"blue_pos":    [  0,  1,  2, 3, 4, 5],
	"green_pos":    [   2,   3,   4,  5,  6,  7],
	"default_color": [0,0,0]
      },
      { "ibits": 6,
	"blue_pos":     [   0,   1,   2,  3,  4,  5],
	"red_pos":     [   2,   3,   4,  5,  6,  7],
	"green_pos":     [   1,   2,   3,  4,  5,  6],
	"default_color": [0,0,0]
      },
      { "ibits": 6,
	"blue_pos":     [   2,   3,   4,  5,  6,  7],
	"red_pos":     [   2,   3,   4,  5,  6,  7],
	"green_pos":     [   2,   3,   4,  5,  6,  7],
	"default_color": [0,0,0]
      },
      { "ibits": 12,
	"blue_pos":     [   2,   3,   4,  5,  6,  7, -1 ,-1 ,-1,-1, -1, -1 ],
	"green_pos":     [  -1 ,-1 ,-1,-1, -1, -1,  2,   3,   4,  5,  6,  7],
	"red_pos":     [  -1,-1, -1, -1, -1, -1, -1 ,-1 ,-1,-1, -1, -1 ],
	"default_color": [0,0,0]
	  }
    ]`
	all_color_names_str = `[
	    "bold: blue,green,red",
	    "bold: green,blue,red",
	    "bold: red,blue,green",
	    "dim: blue,green,red",
	    "dim: green,blue,red",
	    "dim: red,blue,green",
	    "all blue",
	    "all purple",
	    "all purple - purple center",
	    "all maroon",
	    "all orange",
	    "all lime green",
	    "all gold",
	    "all Black",
	    "high resolution: blue, green"
	]`
)

/*
JSON Structure Defining Display Sizes
*/
const (
	all_display_description_str = `[
            "128 x 108 (Test Size)",
            "640 x 360 (nHD)",
            "800 x 600 (SVGA)",
            "1024 x 768 (XGA)",
            "1280 x 720 (WXGA - 16:9)",
            "1280 x 800 (WXGA - 16:10)",
            "1280 x 1024 Super-eXtended Graphics Array (SXGA)",
            "1360 x 768 High Definition (HD - 1360 width)",
            "1366 x 768 High Definition (HD - 1366 width)",
            "1440 x 900 (WXGA+)",
            "1536 x 864 No Name",
            "1600 x 900 High Definition Plus (HD+)",
            "1600 x 1200 (UXGA)",
            "1680 x 1050 (WSXGA+)",
            "1920 x 1080 Full High Definition (FHD)",
            "1920 x 1200 Wide Ultra Extended Graphics Array (WUXGA)",
            "2048 x 1152 (QWXGA)",
            "2048 x 1536 (QXGA)",
            "2560 x 1080 (UWFHD)",
            "2560 x 1440 Quad High Definition (QHD)",
            "2560 x 1600 (WQXGA)",
            "3440 x 1440 Wide Quad High Definition (UWQHD)",
            "3840 x 2160 4K or Ultra High Definition (UHD)"
	]`
	all_display_x_dots_str = `[
            128,
            640,
            800,
            1024,
            1280,
            1280,
            1280,
            1360,
            1366,
            1440,
            1536,
            1600,
            1600,
            1680,
            1920,
            1920,
            2048,
            2048,
            2560,
            2560,
            2560,
            3440,
            3840
	]`
	all_display_y_dots_str = `[
            108,
            360,
            600,
            768,
            720,
            800,
            1024,
            768,
            768,
            900,
            864,
            900,
            1200,
            1050,
            1080,
            1200,
            1152,
            1536,
            1080,
            1440,
            1600,
            1440,
            2160
	]`
)

/*
JSON Structure

	Templates{
		"version": "1.0",       // Template File Format Version Number
		"num_images": 17,        // Number of images in Background
		// Coloring Settings
		"rgb": "bgr",           // First color, second color, then last color used. Blue,Green,Red specified
		"bits_per_color": 4,    // Bits for each color intensity
		"brightness_shift": 3,  // Number of left shifts positions;  bits_per_color + brightness_shift <= 8
		// Size Settings
		"x_units": 8,           // tile units in x-dimension
		"y_units": 4,           // Tile units in y-dimension
		// Python math library Settings
		"high_precision": false,    // true implies high precesion math labrary (decimal) is used;
									//      - Currently this setting does not work for deep zooms
									// false implies the math library built into python is used.
		// Size and position of each image
		// Image positions are specified using a positive x, positive y, cartesion coordinate system
		// That a fancy way of saying the lower left corder of your monitor is x=0; y=0.
		// All positions are specified in tile units.
		"images": [
			{ "side_size": 1, "bg_x": 0, "bg_y": 0 },
			...
			{ "side_size": 4, "bg_x": 4, "bg_y": 0 }
		]
	}
*/
const (
	all_template_str = `[
		{ 
			"version": "1.0",
			"num_images": 17,
			"rgb": "bgr",
			"bits_per_color": 4,
			"brightness_shift": 4,
			"x_units": 8,
			"y_units": 4,
			"high_precision": false,
			"images": [
  				{ "side_size": 1, "bg_x": 0, "bg_y": 3 },
				{ "side_size": 1, "bg_x": 0, "bg_y": 2 },
				{ "side_size": 1, "bg_x": 0, "bg_y": 1 },
				{ "side_size": 1, "bg_x": 0, "bg_y": 0 },
				{ "side_size": 1, "bg_x": 1, "bg_y": 3 },
				{ "side_size": 1, "bg_x": 1, "bg_y": 2 },
				{ "side_size": 1, "bg_x": 1, "bg_y": 1 },
				{ "side_size": 1, "bg_x": 1, "bg_y": 0 },
				{ "side_size": 4, "bg_x": 2, "bg_y": 0 },
				{ "side_size": 1, "bg_x": 6, "bg_y": 3 },
				{ "side_size": 1, "bg_x": 6, "bg_y": 2 },
				{ "side_size": 1, "bg_x": 6, "bg_y": 1 },
				{ "side_size": 1, "bg_x": 6, "bg_y": 0 },
				{ "side_size": 1, "bg_x": 7, "bg_y": 3 },
				{ "side_size": 1, "bg_x": 7, "bg_y": 2 },
				{ "side_size": 1, "bg_x": 7, "bg_y": 1 },
				{ "side_size": 1, "bg_x": 7, "bg_y": 0 }
			]
		},
		{
    		"version": "1.0",
    		"num_images": 17,
    		"rgb": "bgr",
    		"bits_per_color": 4,
    		"brightness_shift": 3,
    		"x_units": 8,
    		"y_units": 4,
    		"high_precision": false,
    		"images": [
				{ "side_size": 1, "bg_x": 0, "bg_y": 0 },
				{ "side_size": 1, "bg_x": 0, "bg_y": 1 },
				{ "side_size": 1, "bg_x": 1, "bg_y": 0 },
				{ "side_size": 1, "bg_x": 0, "bg_y": 2 },
				{ "side_size": 1, "bg_x": 1, "bg_y": 1 },
				{ "side_size": 1, "bg_x": 2, "bg_y": 0 },
				{ "side_size": 1, "bg_x": 0, "bg_y": 3 },
				{ "side_size": 1, "bg_x": 1, "bg_y": 2 },
				{ "side_size": 1, "bg_x": 2, "bg_y": 1 },
				{ "side_size": 1, "bg_x": 3, "bg_y": 0 },
				{ "side_size": 1, "bg_x": 1, "bg_y": 3 },
				{ "side_size": 1, "bg_x": 2, "bg_y": 2 },
				{ "side_size": 1, "bg_x": 3, "bg_y": 1 },
				{ "side_size": 1, "bg_x": 2, "bg_y": 3 },
				{ "side_size": 1, "bg_x": 1, "bg_y": 2 },
				{ "side_size": 1, "bg_x": 3, "bg_y": 3 },
				{ "side_size": 4, "bg_x": 4, "bg_y": 0 }
    		]
		},
		{
			"version": "1.0",
			"num_images": 2,
			"rgb": "bgr",
			"bits_per_color": 4,
			"brightness_shift": 3,
			"x_units": 2,
			"y_units": 1,
			"high_precision": false,
			"images": [
				{
					"side_size": 1,
					"bg_x": 0,
					"bg_y": 0 
				},
				{
					"side_size": 1,
					"bg_x": 1,
					"bg_y": 0 
				}
			]
		} 
	]`
)
const (
	all_template_names_str = `[
		"Big Center (17 Images)",
		"Diagonal (17 Images)",
		"Two (2 Images)"
	]`
)

type tappableRaster struct {
	fyne.CanvasObject
	OnTapped func()
}

func NewTappableRaster(raster fyne.CanvasObject, onTapped func()) *tappableRaster {
	return &tappableRaster{CanvasObject: raster, OnTapped: onTapped}
}

func (t *tappableRaster) Tapped(ev *fyne.PointEvent) {
	fmt.Println("x,y:", ev.Position.X, ev.Position.Y)
	t.OnTapped()
}

// func DoRasterTap(ev *fyne.PointEvent) {
func DoRasterTap() {
	fmt.Println("Tapped")
}

//func (t *tappableRaster) pixelColor(x,y,w,h int) color.Color {
//	fmt.Println( "x,y",x,y,w,h)
//	return(color.Black)
//}

// Field Names MUST start with a capital letter
type MandelColor struct {
	Ibits         int
	Blue_pos      []int
	Red_pos       []int
	Green_pos     []int
	Default_color []uint8
}

type Mandel struct {
	up_to_date      bool
	size            int
	cur_x           int
	cur_y           int
	cur_granularity int
	tiles           [][]Color
	// Math
	//iterations int // Defined by Color
	threshold    float64
	span         float64
	span_one_dot float64
	min_x, min_y float64
	// Window
	cur_w, cur_h                          int
	black_out_top, black_out_left         int
	centering_top_adj, centering_left_adj int
	// Colors
	all_colors      []MandelColor
	all_color_names []string
	cur_color_num   int
	new_color_num   int
	// Zoom Tap
	cur_zoom float64
	new_zoom float64
	// Roam
	cur_roam_speed       int // 1 to 100 (fast)
	new_roam_speed       int
	cur_draw_speed       int
	new_draw_speed       int
	cur_pan_total_steps  int
	cur_zoom_total_steps int
	roam_tgt_x           float64
	roam_tgt_y           float64
	roam_tgt_span_adj    float64 // 0.1-0.99
	roam_step            int
}

type Color struct {
	red   uint8
	green uint8
	blue  uint8
}

type ImageTemplate struct {
	side_size int
	bg_x      int
	bg_y      int
}

type Template struct {
	version          string
	num_images       int
	rgb              string
	bits_per_color   int
	brightness_shift int
	x_units          int
	y_units          int
	high_precision   bool
	images           []ImageTemplate
}

type Background struct {
	template_num       int
	desktop_num        int
	color_theme_num    int
	zoom_magnification int
	zoom_in            bool
	image_defined      int
	all_min_x          []float64
	all_min_y          []float64
	all_span           []float64
	templates          []Template
}

type Point struct {
	x float64
	y float64
}

/*
 * General Functions
 */

func NewColor(r, g, b uint8) Color {
	c := Color{
		red:   r,
		green: g,
		blue:  b,
	}
	return c
}

func NewPoint(set_x, set_y float64) Point {
	p := Point{
		x: set_x,
		y: set_y,
	}
	return p
}

/*
 * Mandel functions
 */

func (m *Mandel) CalcIterationsOneXY(c, di float64) int {
	newA := 0.0
	newBi := 0.0
	a := 0.0
	bi := 0.0
	iterations := 1 << (m.all_colors[m.cur_color_num].Ibits)
	for i := 0; i < iterations; i++ {
		if i == 0 {
			a = c
			bi = di
		} else {
			newA = a*a - bi*bi - c
			newBi = 2.0*a*bi - di
			a = newA
			bi = newBi
			if a > m.threshold {
				return i
			}
		}
	}
	return 0
}

func (m *Mandel) CalcIterationsOnePoint(p Point) int {
	iters := m.CalcIterationsOneXY(p.x, p.y)
	return iters
}

func (m *Mandel) CalcOnePointRGB(p Point) (red_color uint8, green_color uint8, blue_color uint8) {

	iters := m.CalcIterationsOneXY(p.x, p.y)

	red_color = 0
	green_color = 0
	blue_color = 0
	red_adj := 0
	green_adj := 0
	blue_adj := 0
	if iters == 0 {
		red_color = uint8(m.all_colors[m.cur_color_num].Default_color[0])
		green_color = uint8(m.all_colors[m.cur_color_num].Default_color[1])
		blue_color = uint8(m.all_colors[m.cur_color_num].Default_color[2])
	} else {
		for i := 0; i < m.all_colors[m.cur_color_num].Ibits; i++ {
			if (iters & (1 << i)) != 0 {
				red_adj = m.all_colors[m.cur_color_num].Red_pos[i]
				green_adj = m.all_colors[m.cur_color_num].Green_pos[i]
				blue_adj = m.all_colors[m.cur_color_num].Blue_pos[i]
				if red_adj > 0 {
					red_color |= 1 << (red_adj)
				}
				if green_adj > 0 {
					green_color |= 1 << (green_adj)
				}
				if blue_adj > 0 {
					blue_color |= 1 << (blue_adj)
				}
			}
		}
	}
	return
}

func (m *Mandel) CalcOnePointColor(p Point) (c Color) {
	red, green, blue := m.CalcOnePointRGB(p)
	c = NewColor(red, green, blue)
	return
}

func (m *Mandel) CalcOneDot() {
	var p Point

	realx := m.min_x + float64(m.cur_x)*m.span_one_dot
	realy := m.min_y + m.span - float64(m.cur_y)*m.span_one_dot

	p = NewPoint(realx, realy)

	color := m.CalcOnePointColor(p)

	m.tiles[m.cur_x][m.cur_y].red = color.red
	m.tiles[m.cur_x][m.cur_y].green = color.green
	m.tiles[m.cur_x][m.cur_y].blue = color.blue
}

func (m *Mandel) AdvanceToNextDot() {
	// FIXME
	if !m.up_to_date {
		m.cur_x = (m.cur_x + m.cur_granularity) % m.size
		if m.cur_x == 0 {
			m.cur_y = (m.cur_y + m.cur_granularity) % m.size
			if m.cur_y == 0 {
				if m.cur_granularity == 1 {
					m.up_to_date = true
				} else {
					m.cur_granularity = m.cur_granularity >> 1
				}
			}
		}
	}
}

func (m *Mandel) ResetSpan() {
	m.span = 3.0
	m.min_x = -1.0
	//	m.max_x= 2.0
	m.min_y = -1.5
	//	m.max_y= 1.5
	m.span_one_dot = m.span / float64(m.size)
}

func (m *Mandel) ResetWindow(w, h int) {
	// Check
	if (w > max_size) || (h > max_size) {
		fmt.Println("Monitor is too big")
		panic(1)
	}
	// New Window Size
	m.cur_w = w
	m.cur_h = h
	// New Mandelbrot Size
	max_val := 0
	min_val := 0
	// Choose the smallest so it looks okay on Mobile platform
	if w > h {
		max_val = w
		min_val = h
	} else {
		max_val = h
		min_val = w
	}
	max_mult64 := (max_val / 64) * 64
	min_mult64 := (min_val / 64) * 64
	// scale
	m.size = max_mult64
	m.span_one_dot = m.span / float64(m.size)
	// Blackout and Center
	if max_val == w {
		// wider than tall
		m.black_out_left = (w - m.size) >> 1
		m.centering_left_adj = 0
		m.centering_top_adj = (m.size - min_mult64) >> 1
		m.black_out_top = (h - min_mult64) >> 1
	} else {
		// taller than wide
		m.black_out_left = (w - min_mult64) >> 1
		m.centering_left_adj = (m.size - min_mult64) >> 1
		m.centering_top_adj = 0
		m.black_out_top = (h - m.size) >> 1
	}
}

func (m *Mandel) DrawOneDot(px, py, w, h int) color.Color {
	use_px := 0
	use_py := 0
	use_px = px
	use_py = py
	if (w != m.cur_w) || (h != m.cur_h) {
		m.ResetWindow(w, h)
	}

	// color_px
	color_px := use_px - m.black_out_left + m.centering_left_adj
	color_py := use_py - m.black_out_top + m.centering_top_adj

	// Black out or color
	black_color := color.RGBA{0, 0, 0, 0xff}
	if use_py < m.black_out_top {
		// Top
		return (black_color)
	} else if use_py >= (m.black_out_top + m.size) {
		// Bottom
		return (black_color)
	} else if use_px < m.black_out_left {
		// Left
		return (black_color)
	} else if use_py >= (m.black_out_left + m.size) {
		// Right
		return (black_color)
	} else {
		return (m.DrawOneDotNotBlack(color_px, color_py))
	}
}
func (m *Mandel) DrawOneDotNotBlack(use_px, use_py int) color.Color {
	//fmt.Println("px:",px,"py:",py,"w:",w,"h:",h)
	idx_x := 0
	idx_y := 0
	gran := 64
	if m.up_to_date {
		idx_x = use_px
		idx_y = use_py
	} else {
		if m.cur_granularity == 64 {
			gran = 64
		} else if use_py < m.cur_y {
			gran = m.cur_granularity
		} else {
			gran = m.cur_granularity * 2
		}
		if gran == 0 {
			panic(1)
		}
		idx_x = (use_px / gran) * gran
		idx_y = (use_py / gran) * gran
	}
	ret_red := uint8(m.tiles[idx_x][idx_y].red)
	ret_green := uint8(m.tiles[idx_x][idx_y].green)
	ret_blue := uint8(m.tiles[idx_x][idx_y].blue)
	ret_color := color.RGBA{ret_red, ret_green, ret_blue, 0xff}
	return (ret_color)
}

func (m *Mandel) Status() {
	fmt.Println(m.up_to_date, m.cur_granularity, m.cur_x, m.cur_y)
}

func (m *Mandel) CalcBundleSize() int {
	bsize := 0
	if m.cur_granularity == 64 {
		bsize = 4
	} else if m.cur_granularity == 32 {
		bsize = 16
	} else if m.cur_granularity == 16 {
		bsize = 64
	} else if m.cur_granularity == 8 {
		bsize = 256
	} else if m.cur_granularity == 4 {
		bsize = 1024
	} else if m.cur_granularity == 2 {
		bsize = 4096
	} else if m.cur_granularity == 1 {
		bsize = 4096 * 4
	}
	return bsize
}

func (m *Mandel) UpdateSome() {

	// Update one Dot and advance
	bsize := m.CalcBundleSize()
	for b := 0; b < bsize; b++ {
		m.CalcOneDot()
		m.AdvanceToNextDot()
	}
	// Stall longer for courser granularities
	for d := 0; d < (101 - m.cur_draw_speed); d++ {
		time.Sleep(time.Nanosecond * 100000)
	}
}

func (m *Mandel) RoamTgtScreenTwo(x, y float64) bool {

	new_span := 3.0
	for i := 0; i < m.cur_zoom_total_steps; i++ {
		new_span = new_span * m.roam_tgt_span_adj
	}
	half_new_span := new_span / 2.0

	upper_left_pnt := NewPoint((x - half_new_span), (y + half_new_span))
	upper_right_pnt := NewPoint((x + half_new_span), (y + half_new_span))
	lower_left_pnt := NewPoint((x - half_new_span), (y - half_new_span))
	lower_right_pnt := NewPoint((x + half_new_span), (y - half_new_span))

	upper_left_iters := m.CalcIterationsOnePoint(upper_left_pnt)
	upper_right_iters := m.CalcIterationsOnePoint(upper_right_pnt)
	lower_left_iters := m.CalcIterationsOnePoint(lower_left_pnt)
	lower_right_iters := m.CalcIterationsOnePoint(lower_right_pnt)

	f64_upper_left_iters := float64(upper_left_iters)
	f64_upper_right_iters := float64(upper_right_iters)
	f64_lower_left_iters := float64(lower_left_iters)
	f64_lower_right_iters := float64(lower_right_iters)

	/*
		fmt.Println(upper_left_pnt)
		fmt.Println(upper_left_iters)
		fmt.Println(upper_right_pnt)
		fmt.Println(upper_right_iters)
		fmt.Println(lower_left_pnt)
		fmt.Println(lower_left_iters)
		fmt.Println(lower_right_pnt)
		fmt.Println(lower_right_iters)
	*/

	same_cnt := 0
	if upper_left_iters == upper_right_iters {
		same_cnt++
	}
	if upper_left_iters == lower_left_iters {
		same_cnt++
	}
	if upper_left_iters == lower_right_iters {
		same_cnt++
	}
	if upper_right_iters == lower_right_iters {
		same_cnt++
	}
	if upper_right_iters == lower_left_iters {
		same_cnt++
	}
	if lower_left_iters == lower_right_iters {
		same_cnt++
	}
	//fmt.Println("Screen 2: ", same_cnt)

	iterbits := m.all_colors[m.cur_color_num].Ibits
	max_iters := (1 << iterbits)
	f64_max_iters := float64(max_iters)
	good_pnt := 0.7
	if (f64_upper_left_iters / f64_max_iters) > good_pnt {
		return true
	} else if (f64_upper_right_iters / f64_max_iters) > good_pnt {
		return true
	} else if (f64_lower_left_iters / f64_max_iters) > good_pnt {
		return true
	} else if (f64_lower_right_iters / f64_max_iters) > good_pnt {
		return true
	} else if same_cnt > 2 {
		return false
	} else {
		return true
	}
}

func RoamCalcDistance(x, y float64) float64 {
	distance := math.Sqrt(x*x + y*y)
	return (distance)
}

// Must not be in center
func (m *Mandel) RoamTgtScreenOne(x, y float64) bool {
	distance := RoamCalcDistance(x, y)
	if distance < 1.5 {
		return false
	} else {
		if distance > 2.5 {
			return false
		} else {
			return true
		}
	}
}

func (m *Mandel) RoamGenNewTgt() {
	new_x := 0.0
	new_y := 0.0
	found_good_tgt := false
	for found_good_tgt == false {
		new_x = float64(rand.Intn(100))/100.0*3 - 1
		new_y = float64(rand.Intn(100))/100.0*3 - 1.5
		if m.RoamTgtScreenOne(new_x, new_y) {
			if m.RoamTgtScreenTwo(new_x, new_y) {
				found_good_tgt = true
			}
		}
	}
	m.roam_tgt_x = new_x
	m.roam_tgt_y = new_y
}

func (m *Mandel) RoamDelay() {
	for delays := 0; delays < (101 - m.cur_roam_speed); delays++ {
		time.Sleep(time.Nanosecond * 100000000)
	}
}

func (m *Mandel) RoamAdjustSetMinXMinY(imageCenter Point) {
	m.min_x = imageCenter.x - (m.span / 2.0)
	m.min_y = imageCenter.y - (m.span / 2.0)
}

func (m *Mandel) RoamAdjustPanTo() {
	percentPanned := float64(m.roam_step) / float64(m.cur_pan_total_steps)
	new_x := 0.5 + (m.roam_tgt_x-0.5)*percentPanned
	new_y := m.roam_tgt_y * percentPanned
	imageCenter := NewPoint(new_x, new_y)
	m.RoamAdjustSetMinXMinY(imageCenter)
}

func (m *Mandel) RoamAdjustPanFrom() {
	percentPanned := float64(m.roam_step) / float64(m.cur_pan_total_steps)
	new_x := m.roam_tgt_x - (m.roam_tgt_x-0.5)*percentPanned
	new_y := m.roam_tgt_y - m.roam_tgt_y*percentPanned
	imageCenter := NewPoint(new_x, new_y)
	m.RoamAdjustSetMinXMinY(imageCenter)
}

func (m *Mandel) RoamAdjustZoomIn() {
	// Reduce span
	m.span = m.span * m.roam_tgt_span_adj
	m.span_one_dot = m.span / float64(m.size)
	// Set upper left point
	imageCenter := NewPoint(m.roam_tgt_x, m.roam_tgt_y)
	m.RoamAdjustSetMinXMinY(imageCenter)
}

func (m *Mandel) RoamAdjustZoomOut() {
	// Increase span
	new_span := 3.0
	for i := 0; i < (m.cur_zoom_total_steps - m.roam_step); i++ {
		new_span = new_span * m.roam_tgt_span_adj
	}
	m.span = new_span
	m.span_one_dot = m.span / float64(m.size)
	// Set upper left point
	imageCenter := NewPoint(m.roam_tgt_x, m.roam_tgt_y)
	m.RoamAdjustSetMinXMinY(imageCenter)
}

func (m *Mandel) FrcRedraw() {
	m.up_to_date = false
	m.cur_granularity = 64
}

func NewMandel() Mandel {
	//	var lcl_all_colors []MandelColor
	m := Mandel{
		size:            256,
		cur_x:           0,
		cur_y:           0,
		cur_granularity: 64,
		up_to_date:      false,
		// Math
		//span:      3.0,
		span:      3.0,
		threshold: 1000.0,
		//		min_x:     -1.0,
		min_x: -1.0,
		//		max_x: 2.0,
		//min_y: -1.5,
		min_y: -1.5,
		//	max_y: 1.5,
		//Window
		cur_w: 256,
		cur_h: 256,
		// Color
	}
	m.span_one_dot = m.span / float64(m.size)
	m.tiles = make([][]Color, max_size)
	for i := 0; i < max_size; i++ {
		m.tiles[i] = make([]Color, max_size)
	}
	err := json.Unmarshal([]byte(all_colors_str), &m.all_colors)
	if err != nil {
		fmt.Printf("Unable to marshal JSON due to %s", err)
		panic(1)
	}
	//fmt.Println(m.all_colors)
	//fmt.Println("extra",lcl_all_colors[0].Ibits)
	//fmt.Println("extra",lcl_all_colors[0].Blue_pos)
	err = json.Unmarshal([]byte(all_color_names_str), &m.all_color_names)
	if err != nil {
		fmt.Printf("Unable to marshal JSON due to %s", err)
		panic(1)
	}
	//fmt.Println(m.all_color_names)
	return m
}

/*
 * Background functions
 */

func (bg *Background) TotalImages() int {
	total_images := bg.templates[bg.template_num].num_images
	return total_images
}

func (bg *Background) PathImageString() string {
	str := fmt.Sprintf("Zoom Path Points Defined: %d (out of %d)", bg.image_defined, bg.TotalImages())
	return str
}

func (bg *Background) GetTemplateChoicesStrings() []string {
	var choices []string
	err := json.Unmarshal([]byte(all_template_names_str), &choices)
	if err != nil {
		fmt.Printf("Unable to marshal JSON due to %s", err)
		panic(1)
	}
	return choices
}

func (bg *Background) GetDesktopChiocesStrings() []string {
	var choices []string
	err := json.Unmarshal([]byte(all_display_description_str), &choices)
	if err != nil {
		fmt.Printf("Unable to marshal JSON due to %s", err)
		panic(1)
	}
	return choices
}

func (bg *Background) GetColorChiocesStrings() []string {
	var choices []string
	err := json.Unmarshal([]byte(all_color_names_str), &choices)
	if err != nil {
		fmt.Printf("Unable to marshal JSON due to %s", err)
		panic(1)
	}
	return choices
}

func NewBackground() Background {
	bg := Background{
		template_num:       0,
		desktop_num:        0,
		color_theme_num:    0,
		zoom_magnification: 1,
		zoom_in:            true,
		image_defined:      0,
		//all_min_x: []float64
		//all_min_y []float64
		//all_span  []float64
	}
	// Read in all the templates
	err := json.Unmarshal([]byte(all_template_str), &bg.templates)
	if err != nil {
		fmt.Printf("Unable to marshal JSON due to %s", err)
		panic(1)
	}
	return bg
}

/*
 * Main
 */

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Mandelbrot Background")
	myWindow.SetPadded(false)

	// Resize ignored by Mobile Platforms
	// - Mobile platforms are always full screen
	// - 27 is a hack determined by Ubuntu/Gnome
	myWindow.Resize(fyne.NewSize(256, (256 + 27)))

	// Control Menu Set up
	menuItemGenerate := fyne.NewMenuItem("Generate Background", func() {
		fmt.Println("In Generate Background")
	})
	menuItemQuit := fyne.NewMenuItem("Quit", func() {
		//fmt.Println("In DoQuit:")
		os.Exit(0)
	})
	//	menuControl:= fyne.NewMenu("Control", menuItemColor, menuItemZoom, menuItemQuit);
	menuControl := fyne.NewMenu("Control", menuItemGenerate, menuItemQuit)
	// About Menu Set up
	menuItemAbout := fyne.NewMenuItem("About...", func() {
		dialog.ShowInformation("About Mandelbrot Background v1.0.0", "Author: Craig Warner \n\ngithub.com/craig-warner/mandelbrot-background", myWindow)
	})
	menuHelp := fyne.NewMenu("Help ", menuItemAbout)
	mainMenu := fyne.NewMainMenu(menuControl, menuHelp)
	myWindow.SetMainMenu(mainMenu)

	// Background
	bg := NewBackground()

	// Mandelbrot
	myMandel := NewMandel()

	// Content

	selectBackgroundTemplateText := canvas.NewText("Select a background temple", color.Black)
	selectBackgroundTemplateChoicesStrings := bg.GetTemplateChoicesStrings()
	selectBackgroundTemplateChoices := widget.NewSelect(selectBackgroundTemplateChoicesStrings, func(s string) {
		fmt.Println("Select Background Template Callback:", s)
		for i := 0; i < len(selectBackgroundTemplateChoicesStrings); i++ {
			if selectBackgroundTemplateChoicesStrings[i] == s {
				bg.template_num = i
				break
			}
		}
	})
	selectDesktopSizeText := canvas.NewText("Select your desktop size ", color.Black)
	selectDesktopSizeChoicesStrings := bg.GetDesktopChiocesStrings()
	selectDesktopSizeChoices := widget.NewSelect(selectDesktopSizeChoicesStrings, func(s string) {
		fmt.Println("Select Desktop Size Callback:", s)
		for i := 0; i < len(selectDesktopSizeChoicesStrings); i++ {
			if selectDesktopSizeChoicesStrings[i] == s {
				bg.desktop_num = i
				break
			}
		}
	})
	selectColorPreferenceText := canvas.NewText("Select your color preference", color.Black)
	selectColorPreferenceChoicesStrings := bg.GetColorChiocesStrings()
	selectColorPreferenceChoices := widget.NewSelect(selectColorPreferenceChoicesStrings, func(s string) {
		fmt.Println("Select Color Preference Callback:", s)
		for i := 0; i < len(selectColorPreferenceChoicesStrings); i++ {
			if selectColorPreferenceChoicesStrings[i] == s {
				bg.color_theme_num = i
				break
			}
		}
	})
	zoomMagnificationText := canvas.NewText("Zoom in Magnification (1x to 10x)", color.Black)

	zoomMagnificationSlider := widget.NewSlider(1.0, 10.0)
	zoomMagnificationSlider.OnChanged = func(f float64) {
		fmt.Println("Zoom Magnification Callback:", f)
		bg.zoom_magnification = int(f)
	}
	zoomInText := canvas.NewText("Zoom in", color.Black)
	zoomInCheckBox := widget.NewCheck("Zoom In", func(b bool) {
		fmt.Println("Zoom In Callback:", b)
		bg.zoom_in = b
	})
	zoomContent := container.New(layout.NewHBoxLayout(), zoomMagnificationText, zoomMagnificationSlider, zoomInText, zoomInCheckBox)

	addResetContent := container.New(layout.NewHBoxLayout())
	addImageBtn := widget.NewButton("Add Image", func() {
		fmt.Println("Add Image")
		bg.image_defined++
		//zoomPathText.Text = bg.PathImageString()
		// FIXME: Add Image
	})
	resetPathBtn := widget.NewButton("Reset", func() {
		fmt.Println("Reset")
		bg.image_defined = 0
		//	FIXME: Reset
	})
	addResetContent.Add(addImageBtn)
	addResetContent.Add(resetPathBtn)

	zoomPathText := canvas.NewText("Zoom Path Points Defined: 0", color.Black)

	// Column One
	colOneContent := container.New(layout.NewVBoxLayout())
	colOneContent.Add(selectBackgroundTemplateText)
	colOneContent.Add(selectBackgroundTemplateChoices)
	colOneContent.Add(selectDesktopSizeText)
	colOneContent.Add(selectDesktopSizeChoices)
	colOneContent.Add(selectColorPreferenceText)
	colOneContent.Add(selectColorPreferenceChoices)
	colOneContent.Add(zoomContent)
	colOneContent.Add(addResetContent)
	colOneContent.Add(zoomPathText)

	colTwoContent := container.New(layout.NewVBoxLayout())

	myRaster := canvas.NewRasterWithPixels(myMandel.DrawOneDot)
	colTwoContent.Add(myRaster)

	topContent := container.New(layout.NewHBoxLayout())
	topContent.Add(colOneContent)
	topContent.Add(colTwoContent)

	// Botton Content Creation
	bottomContent := container.New(layout.NewVBoxLayout())
	generateBtn := widget.NewButton("Generate Background", func() {
		fmt.Println("Generate Background")
		// FIXME: Generate Background
	})
	backgroundGenerationProgressText := canvas.NewText("Background Generation Progress", color.Black)
	backgroundGenerationProgressBar := widget.NewProgressBar()
	backgroundProgrogressContent := container.New(layout.NewHBoxLayout())
	backgroundProgrogressContent.Add(backgroundGenerationProgressText)
	backgroundProgrogressContent.Add(backgroundGenerationProgressBar)
	imageGenerationProgressText := canvas.NewText("Image Generation Progress", color.Black)
	imageGenerationProgressBar := widget.NewProgressBar()
	imageGenerationProgrogressContent := container.New(layout.NewHBoxLayout())
	imageGenerationProgrogressContent.Add(imageGenerationProgressText)
	imageGenerationProgrogressContent.Add(imageGenerationProgressBar)

	bottomContent.Add(generateBtn)
	bottomContent.Add(backgroundProgrogressContent)
	bottomContent.Add(imageGenerationProgrogressContent)

	wholeContent := container.New(layout.NewVBoxLayout())
	wholeContent.Add(topContent)
	wholeContent.Add(bottomContent)

	myWindow.SetContent(wholeContent)

	//go func() {
	//}()

	myWindow.ShowAndRun()
}
