package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"

	//	"github.com/hjson/hjson-go/v4"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const DEBUG = false

const WINDOW_SIZE = 512

const (
	MAX_DISPLAY_SIZE = 10000
	MAX_IMAGES       = 256
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
/*
const (
	all_template_str = `[
		{
    		"version": "2.0",
    		"num_images": 1,
    		"x_units": 4,
    		"y_units": 2,
    		"high_precision": false,
    		"save_filename": "mbg_one",
    		"images": [
     		 { "side_size": 2, "bg_x": 1, "bg_y": 0 }
    		]
		},
		{ 	"version": "2.0",
    		"num_images": 2,
    		"x_units": 2,
    		"y_units": 1,
    		"high_precision": false,
    		"save_filename": "mbg_two",
    		"images": [
      			{ "side_size": 1, "bg_x": 0, "bg_y": 0 },
      			{ "side_size": 1, "bg_x": 1, "bg_y": 0 }
    		]
		},
		{
    		"version": "2.0",
    		"num_images": 3,
    		"x_units": 3,
    		"y_units": 2,
    		"high_precision": false,
    		"save_filename": "mbg_three",
    		"images": [
				{ "side_size": 1, "bg_x": 0,     "bg_y": 1      },
      			{ "side_size": 1, "bg_x": 0, "bg_y": 0 },
				 { "side_size": 2, "bg_x": 1, "bg_y": 0 }
    		]
  		},
		{
		"version":"2.0",
		"x_units": 16,
  		"y_units": 8,
  		"high_precision": false,
  		"save_filename": "mbg_8",
  		"images": [
    		{ "side_size": 2, "bg_x": 0, "bg_y": 6 },
    		{ "side_size": 2, "bg_x": 2, "bg_y": 6 },
    		{ "side_size": 2, "bg_x": 0, "bg_y": 4 },
    		{ "side_size": 2, "bg_x": 2, "bg_y": 4 },
    		{ "side_size": 4, "bg_x": 4, "bg_y": 4 },
    		{ "side_size": 4, "bg_x": 0, "bg_y": 0 },
    		{ "side_size": 4, "bg_x": 4, "bg_y": 0 },
    		{ "side_size": 8, "bg_x": 8, "bg_y": 0 }
  		]
		},
		{
		"version":"2.0",
		"x_units": 16,
  		"y_units": 8,
  		"high_precision": false,
  		"save_filename": "mbg_11",
  		"images": [
			{ "side_size": 1, "bg_x": 15, "bg_y": 7 },
    		{ "side_size": 1, "bg_x": 14, "bg_y": 7 },
    		{ "side_size": 1, "bg_x": 14, "bg_y": 6 },
    		{ "side_size": 1, "bg_x": 15, "bg_y": 6 },
    		{ "side_size": 2, "bg_x": 12, "bg_y": 6 },
    		{ "side_size": 2, "bg_x": 12, "bg_y": 4 },
    		{ "side_size": 2, "bg_x": 14, "bg_y": 4 },
    		{ "side_size": 4, "bg_x": 8, "bg_y": 4 },
    		{ "side_size": 4, "bg_x": 8, "bg_y": 0 },
    		{ "side_size": 4, "bg_x": 12, "bg_y": 0 },
    		{ "side_size": 8, "bg_x": 0, "bg_y": 0 }
  		]
		},
		{
    		"version": "2.0",
    		"num_images": 5,
    		"x_units": 4,
    		"y_units": 2,
    		"high_precision": false,
    		"save_filename": "mbg_5",
    		"images": [
      			{ "side_size": 2, "bg_x": 1, "bg_y": 0 },
      			{ "side_size": 1, "bg_x": 0, "bg_y": 1 },
      			{ "side_size": 1, "bg_x": 3, "bg_y": 1 },
      			{ "side_size": 1, "bg_x": 0, "bg_y": 0 },
      			{ "side_size": 1, "bg_x": 3, "bg_y": 0 }
    		]
		},
		{
			"num_images": 14,
  			"x_units": 32,
  			"y_units": 16,
  			"high_precision": false,
  			"save_filename": "mbg_14",
  			"images": [
				{ "side_size": 1, "bg_x": 31, "bg_y": 15 },
      			{ "side_size": 1, "bg_x": 30, "bg_y": 15 },
      			{ "side_size": 1, "bg_x": 30, "bg_y": 14 },
      			{ "side_size": 1, "bg_x": 31, "bg_y": 14 },
      			{ "side_size": 2, "bg_x": 28, "bg_y": 14 },
      			{ "side_size": 2, "bg_x": 28, "bg_y": 12 },
      			{ "side_size": 2, "bg_x": 30, "bg_y": 12 },
      			{ "side_size": 4, "bg_x": 24, "bg_y": 12 },
      			{ "side_size": 4, "bg_x": 24, "bg_y": 8 },
      			{ "side_size": 4, "bg_x": 28, "bg_y": 8 },
      			{ "side_size": 8, "bg_x": 16, "bg_y": 8 },
      			{ "side_size": 8, "bg_x": 16, "bg_y": 0 },
      			{ "side_size": 8, "bg_x": 24, "bg_y": 0 },
      			{ "side_size": 16, "bg_x": 0, "bg_y": 0 }
  			]
		},
		{ 	"version":"2.0",
			"num_images":17,
			"x_units":8,
			"y_units":4,
			"high_precision":false,
  			"save_filename": "mbg_17",
			"images": [ { "side_size": 1, "bg_x": 0, "bg_y": 3 },
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
		{ 	"version": "1.0",
    		"num_images": 17,
    		"x_units": 8,
    		"y_units": 4,
    		"high_precision": false,
  			"save_filename": "mbg_diag17",
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
		}
	]`
)
const (
	all_template_names_str = `[
		"One (1 Images)",
		"Two (2 Images)",
		"Three (3 Images)",
		"Big Right (8 Images)",
		"Big Right (11 Images)",
		"First Center (5 Images)",
		"Big Right (14 Images)",
		"Big Center (17 Images)",
		"Diagonal (17 Images)"
	]`
)
*/

func DbgPrint(str ...interface{}) {
	if DEBUG {
		fmt.Println(str...)
		return
	}
}

type tappableRaster struct {
	fyne.CanvasObject
	OnTapped func()
}

func NewTappableRaster(raster fyne.CanvasObject, onTapped func()) *tappableRaster {
	return &tappableRaster{CanvasObject: raster, OnTapped: onTapped}
}

//func (t *tappableRaster) Tapped(ev *fyne.PointEvent) {
//	fmt.Println("x,y:", ev.Position.X, ev.Position.Y)
//	t.OnTapped()
//}

// func DoRasterTap(ev *fyne.PointEvent) {
//func DoRasterTap() {
//	fmt.Println("Tapped")
//}

//func (t *tappableRaster) pixelColor(x,y,w,h int) color.Color {
//	fmt.Println( "x,y",x,y,w,h)
//	return(color.Black)
//}

// Field Names MUST start with a capital letter
/*
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
	threshold    float64
	span         float64
	span_one_dot float64
	min_x, min_y float64
	// Window
	cur_w, cur_h                          int
	black_out_top, black_out_left         int
	centering_top_adj, centering_left_adj int
	// Colors
	cur_color_num int
	all_colors    []MandelColor
}

type Color struct {
	red   uint8
	green uint8
	blue  uint8
}
*/

/*
type ImageTemplate struct {
	Side_size int `json:"side_size"`
	Bg_x      int `json:"bg_x"`
	Bg_y      int `json:"bg_y"`
}

type Template struct {
	Version       string          `json:"version"`
	Num_images    int             `json:"num_images"`
	X_units       int             `json:"x_units"`
	Y_units       int             `json:"y_units"`
	Save_filename string          `json:"save_filename"`
	Images        []ImageTemplate `json:"images"`
}

type Background struct {
	template_num       int
	desktop_num        int
	color_theme_num    int
	zoom_magnification int
	zoom_in            bool
	image_defined      int
	//	cur_min_x          float64
	//	cur_min_y          float64
	//	cur_span           float64
	pan_speed      float64
	all_min_x      []float64
	all_min_y      []float64
	all_span       []float64
	templates      []Template
	images         []Mandel
	desktop_x_dots []int
	desktop_y_dots []int
}

type Point struct {
	x float64
	y float64
}

*/
/*
 * General Functions
 */

/*
func NewColor(r, g, b uint8) Color {
	c := Color{
		red:   r,
		green: g,
		blue:  b,
	}
	return c
}
*/

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

/*
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
	// DEBUG
	//fmt.Printf("AdvanceToNextDot %+v", m)
	//fmt.Printf("AdvanceToNextDot cur_x:%d, cur_y:%d,size:%d,cur_gran:%d", m.cur_x, m.cur_y, m.size, m.cur_granularity)
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
	if (w > MAX_DISPLAY_SIZE) || (h > MAX_DISPLAY_SIZE) {
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

	//fmt.Printf("DrawOne px:%d,py:%d,w:%d,h:%d", px, py, w, h)
	// Black out or color
	black_color := color.RGBA{0, 0, 0, 0xff}
	if use_py < m.black_out_top {
		// Top
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
	bsize := 16
	return bsize
}
*/

/*
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
*/

/*
func (m *Mandel) PercentCalced() float64 {
	return (float64(m.cur_y) / float64(m.size))
}

func (m *Mandel) UpdateSome() {
	// Update one Dot and advance
	bsize := m.CalcBundleSize()
	for b := 0; b < bsize; b++ {
		//	fmt.Printf("UpdateSome:b=%d,bsize=%d", b, bsize)
		m.CalcOneDot()
		m.AdvanceToNextDot()
	}
	// Stall longer for courser granularities
	//for d := 0; d < (101 - m.cur_draw_speed); d++ {
	//	time.Sleep(time.Nanosecond * 100000)
	//}
}

func (m *Mandel) SetColorTheme(color_theme_num int) {
	m.up_to_date = false
	m.cur_color_num = color_theme_num
}

func (m *Mandel) FetchOnePoint(px, py int) (r, g, b uint8) {
	return m.tiles[px][py].red, m.tiles[px][py].green, m.tiles[px][py].blue
}

func (m *Mandel) AdjustPreview(delta_x, delta_y, delta_span float64) {
	m.min_x += delta_x
	m.min_y += delta_y
	m.AdjustZoom(delta_span)
	m.up_to_date = false
}

func (m *Mandel) AdjustZoom(adj float64) {
	center_x := m.min_x + (m.span / 2.0)
	center_y := m.min_y + (m.span / 2.0)
	DbgPrint("center_x:%f,center_y:%f\n", center_x, center_y)
	// Reduce span
	m.span = m.span * adj
	m.span_one_dot = m.span / float64(m.size)
	// Set upper left point
	m.min_x = center_x - (m.span / 2.0)
	m.min_y = center_y - (m.span / 2.0)
}

/*
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

		fmt.Println(upper_left_pnt)
		fmt.Println(upper_left_iters)
		fmt.Println(upper_right_pnt)
		fmt.Println(upper_right_iters)
		fmt.Println(lower_left_pnt)
		fmt.Println(lower_left_iters)
		fmt.Println(lower_right_pnt)
		fmt.Println(lower_right_iters)

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
*/

/*
func RoamCalcDistance(x, y float64) float64 {
	distance := math.Sqrt(x*x + y*y)
	return (distance)
}
*/

/*
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
*/

/*
func NewMandel(min_x, min_y, span float64, size, color_theme_num int) Mandel {
	//	var lcl_all_colors []MandelColor
	m := Mandel{
		size:            size,
		cur_x:           0,
		cur_y:           0,
		cur_granularity: 1, // Go straight to highest resolution
		up_to_date:      false,
		// Math
		//span:      3.0,
		span:      span,
		threshold: 10000000.0,
		//threshold: 10.0,
		//		min_x:     -1.0,
		min_x: min_x,
		//		max_x: 2.0,
		//min_y: -1.5,
		min_y: min_y,
		//	max_y: 1.5,
		//Window
		cur_w: size,
		cur_h: size,
		// Color
		cur_color_num: color_theme_num,
	}
	m.span_one_dot = m.span / float64(m.size)
	m.tiles = make([][]Color, MAX_DISPLAY_SIZE)
	for i := 0; i < MAX_DISPLAY_SIZE; i++ {
		m.tiles[i] = make([]Color, MAX_DISPLAY_SIZE)
	}
	err := json.Unmarshal([]byte(all_colors_str), &m.all_colors)
	if err != nil {
		fmt.Printf("Unable to marshal JSON due to %s", err)
		panic(1)
	}
	//fmt.Println(m.all_colors)
	//fmt.Println("extra",lcl_all_colors[0].Ibits)
	//fmt.Println("extra",lcl_all_colors[0].Blue_pos)
	//err = json.Unmarshal([]byte(all_color_names_str), &m.all_color_names)
	//if err != nil {
	//	fmt.Printf("Unable to marshal JSON due to %s", err)
	//	panic(1)
	//}
	//fmt.Printf("colors: %+v", bg.templates)
	//fmt.Println(m.all_color_names)
	return m
}
*/

/*
 * Background functions
 */

/*
func (bg *Background) PixelsPerUnit() int {
	pixels_per_unit := 1
	width := bg.desktop_x_dots[bg.desktop_num]
	height := bg.desktop_y_dots[bg.desktop_num]
	pixels_per_unit_x := int(width / bg.templates[bg.template_num].X_units)
	pixels_per_unit_y := int(height / bg.templates[bg.template_num].Y_units)
	if pixels_per_unit_x < pixels_per_unit_y {
		// use x
		pixels_per_unit = pixels_per_unit_x
	} else {
		pixels_per_unit = pixels_per_unit_y
	}
	return pixels_per_unit

}

func (bg *Background) CalcPadX() int {
	width := bg.desktop_x_dots[bg.desktop_num]
	pixels_per_unit := bg.PixelsPerUnit()
	xunits := bg.templates[bg.template_num].X_units
	xpad := int((width - xunits*pixels_per_unit) / 2) // Might be shifted left one pixel
	return (xpad)
}

func (bg *Background) CalcPady() int {
	height := bg.desktop_y_dots[bg.desktop_num]
	pixels_per_unit := bg.PixelsPerUnit()
	yunits := bg.templates[bg.template_num].Y_units
	ypad := int((height - yunits*pixels_per_unit) / 2) // Might be shifted up one pixel
	return (ypad)
}

func (bg *Background) TotalImages() int {
	total_images := bg.templates[bg.template_num].Num_images
	return total_images
}

func (bg *Background) PathImageString() string {
	str := fmt.Sprintf("Zoom Path Points Defined: %d (out of %d)", bg.image_defined, bg.TotalImages())
	DbgPrint("template num %d", bg.template_num)
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

func (bg *Background) AdjustPreview(m *Mandel, delta_x, delta_y, delta_span float64) {
	m.AdjustPreview(delta_x, delta_y, delta_span)
	// bg.cur_min_x += delta_x
	// bg.cur_min_y += delta_y
	// bg.cur_span *= delta_span
}

func (bg *Background) PanUp(m *Mandel) {
	bg.AdjustPreview(m, 0.0, (m.span * bg.pan_speed), 1.0)
}
func (bg *Background) PanDown(m *Mandel) {
	bg.AdjustPreview(m, 0.0, -(m.span * bg.pan_speed), 1.0)
}
func (bg *Background) PanLeft(m *Mandel) {
	bg.AdjustPreview(m, -(m.span * bg.pan_speed), 0.0, 1.0)
}
func (bg *Background) PanRight(m *Mandel) {
	bg.AdjustPreview(m, m.span*bg.pan_speed, 0.0, 1.0)
}
func (bg *Background) PanZoomIn(m *Mandel) {
	bg.AdjustPreview(m, 0.0, 0.0, 1.0-0.5*(float64(bg.zoom_magnification)/float64(10.0)))
}
func (bg *Background) PanZoomOut(m *Mandel) {
	bg.AdjustPreview(m, 0.0, 0.0, 1.0+(float64(bg.zoom_magnification)/float64(10.0)))

}

func NewBackground() Background {
	bg := Background{
		template_num:       0,
		desktop_num:        0,
		color_theme_num:    0,
		zoom_magnification: 1,
		zoom_in:            true,
		image_defined:      0,
		//		cur_min_x:          -1.0,
		//		cur_min_y:          -1.5,
		//		cur_span:           3.0,
		pan_speed: 0.1,
	}
	for i := 0; i < MAX_IMAGES; i++ {
		bg.all_min_x = append(bg.all_min_x, float64(-1.0))
		bg.all_min_y = append(bg.all_min_y, float64(-1.5))
		bg.all_span = append(bg.all_span, float64(3.0))
	}
	// Read in all the templates
	err := json.Unmarshal([]byte(all_template_str), &bg.templates)
	if err != nil {
		fmt.Printf("Unable to marshal JSON due to %s", err)
		panic(1)
	}
	DbgPrint("Templates: %+v", bg.templates)
	// Read in Desktop size
	err = json.Unmarshal([]byte(all_display_x_dots_str), &bg.desktop_x_dots)
	if err != nil {
		fmt.Printf("Unable to marshal JSON due to %s", err)
		panic(1)
	}
	err = json.Unmarshal([]byte(all_display_y_dots_str), &bg.desktop_y_dots)
	if err != nil {
		fmt.Printf("Unable to marshal JSON due to %s", err)
		panic(1)
	}
	DbgPrint("Templates: %+v", bg.desktop_y_dots)
	return bg
}
*/
/*
 * Basic Image Functions
 */
func TranferMandelToImage(new_mandelbrot Mandel, mbg_image *image.RGBA, pos_x, pos_y int) {
	for px := 0; px < new_mandelbrot.size; px++ {
		for py := 0; py < new_mandelbrot.size; py++ {
			red, green, blue := new_mandelbrot.FetchOnePoint(px, py)
			mbg_image.Set(pos_x+px, pos_y+py, color.RGBA{red, green, blue, 0xff})
		}
	}
}

/*
func encodeImage(filename string, img image.Image, format string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	//switch format {
	//case "bmp":
		//		return bmp.Encode(file, img, nil)
		//	case "jpeg", "jpg":
		//		return jpeg.Encode(file, img, nil)
		//	case "png":
		//		return png.Encode(file, img)
		//	case "gif":
		//		return gif.Encode(file, img, nil)
	//default:
	//	return fmt.Errorf("unsupported format: %s", format)
	//}
}
*/

/*
 * Main
 */

func main() {

	zoomPathString := "Empty"
	zoomPathLabel := widget.NewLabel(zoomPathString)
	//zoomMagString := "Empty"
	//zoomMagLabel := widget.NewLabel(zoomMagString)
	colOneContent := container.New(layout.NewVBoxLayout())
	//widthCtlText := canvas.NewText("===========================================", color.Black)

	myApp := app.New()
	myWindow := myApp.NewWindow("Mandelbrot Background")
	myWindow.SetPadded(false)

	// Resize ignored by Mobile Platforms
	// - Mobile platforms are always full screen
	// - 27 is a hack determined by Ubuntu/Gnome
	//myWindow.Resize(fyne.NewSize(256, (256 + 27)))
	myWindow.Resize(fyne.NewSize(WINDOW_SIZE, (WINDOW_SIZE + 27)))

	// Control Menu Set up
	//	menuItemGenerate := fyne.NewMenuItem("Generate Background", func() {
	//		fmt.Println("In Generate Background")
	//	})
	menuItemQuit := fyne.NewMenuItem("Quit", func() {
		//fmt.Println("In DoQuit:")
		os.Exit(0)
	})
	//	menuControl:= fyne.NewMenu("Control", menuItemColor, menuItemZoom, menuItemQuit);
	//menuControl := fyne.NewMenu("Control", menuItemGenerate, menuItemQuit)
	menuControl := fyne.NewMenu("Control", menuItemQuit)
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
	myMandel := NewMandel(-1.0, -1.5, 3.0, 256, 0)
	myMandel.ResetSpan()
	myMandel.ResetWindow(256, 256)
	// Raster
	myRaster := canvas.NewRasterWithPixels(myMandel.DrawOneDot)

	// Content

	selectBackgroundTemplateText := canvas.NewText("Select a background temple", color.Black)
	selectBackgroundTemplateChoicesStrings := bg.GetTemplateChoicesStrings()
	selectBackgroundTemplateChoices := widget.NewSelect(selectBackgroundTemplateChoicesStrings, func(s string) {
		if DEBUG {
			fmt.Println("Select Background Template Callback:", s)
		}
		for i := 0; i < len(selectBackgroundTemplateChoicesStrings); i++ {
			if selectBackgroundTemplateChoicesStrings[i] == s {
				// New Template
				bg.template_num = i
				bg.image_defined = 0
				zoomPathString = bg.PathImageString()
				zoomPathLabel.SetText(zoomPathString)
				break
			}
		}
	})
	selectBackgroundTemplateChoices.SetSelectedIndex(0)

	selectDesktopSizeText := canvas.NewText("Select your desktop size ", color.Black)
	selectDesktopSizeChoicesStrings := bg.GetDesktopChiocesStrings()
	selectDesktopSizeChoices := widget.NewSelect(selectDesktopSizeChoicesStrings, func(s string) {
		if DEBUG {
			fmt.Println("Select Desktop Size Callback:", s)
		}
		for i := 0; i < len(selectDesktopSizeChoicesStrings); i++ {
			if selectDesktopSizeChoicesStrings[i] == s {
				bg.desktop_num = i
				break
			}
		}
	})
	selectDesktopSizeChoices.SetSelectedIndex(0)

	selectColorPreferenceText := canvas.NewText("Select your color preference", color.Black)
	selectColorPreferenceChoicesStrings := bg.GetColorChiocesStrings()
	selectColorPreferenceChoices := widget.NewSelect(selectColorPreferenceChoicesStrings, func(s string) {
		if DEBUG {
			fmt.Println("Select Color Preference Callback:", s)
		}
		for i := 0; i < len(selectColorPreferenceChoicesStrings); i++ {
			if selectColorPreferenceChoicesStrings[i] == s {
				bg.color_theme_num = i
				myMandel.SetColorTheme(bg.color_theme_num)
				myRaster.Refresh()
				break
			}
		}
	})
	selectColorPreferenceChoices.SetSelectedIndex(0)

	zoomMagnificationText := canvas.NewText("Zoom in Magnification (1x to 2x)", color.Black)
	zoomMagnificationSlider := widget.NewSlider(1.0, 10.0)
	zoomMagnificationSlider.SetValue(2.0)
	zoomMagnificationSlider.OnChanged = func(f float64) {
		if DEBUG {
			fmt.Println("Zoom Magnification Callback:", f)
		}
		bg.zoom_magnification = int(f)
	}
	//zoomInText := canvas.NewText("Zoom in", color.Black)
	//zoomInCheckBox := widget.NewCheck("Zoom In", func(b bool) {
	//	fmt.Println("Zoom In Callback:", b)
	//	bg.zoom_in = b
	//})
	//zoomInCheckBox.SetChecked(true)
	//zoomContent := container.New(layout.NewHBoxLayout(), zoomMagnificationText, zoomMagnificationSlider, zoomInText, zoomInCheckBox)

	//zoomMagString = bg.PathImageString()
	//zoomMagLabel.SetText(zoomPathString)

	//zoomContent := container.New(layout.NewHBoxLayout(), zoomMagnificationText, zoomMagnificationSlider, zoomMagLabel)
	zoomContent := container.New(layout.NewHBoxLayout(), zoomMagnificationText, zoomMagnificationSlider)

	panCheckBox := widget.NewCheck("Fine Grained Pan", func(b bool) {
		if DEBUG {
			fmt.Println("Zoom In Callback:", b)
		}
		if b {
			bg.pan_speed = 0.01
		} else {
			bg.pan_speed = 0.1
		}
	})
	panCheckBox.SetChecked(true)

	addResetContent := container.New(layout.NewHBoxLayout())
	addImageBtn := widget.NewButton("Add Image", func() {
		// Check
		if bg.image_defined >= bg.TotalImages() {
			if DEBUG {
				fmt.Println("Too many images defined")
			}
			var popup *widget.PopUp
			all_defined_label := widget.NewLabel("All Images Already Defined")
			popUpContent := container.NewVBox(
				all_defined_label,
				widget.NewButton("Ok", func() {
					popup.Hide()
				}),
			)
			popup = widget.NewModalPopUp(popUpContent, myWindow.Canvas())
			popup.Show()
			return
		}
		if DEBUG {
			fmt.Println("Add Image")
		}
		bg.image_defined++
		zoomPathString = bg.PathImageString()
		zoomPathLabel.SetText(zoomPathString)
		zoomPathLabel.Refresh()
		bg.all_min_x[bg.image_defined] = myMandel.min_x
		bg.all_min_y[bg.image_defined] = myMandel.min_y
		bg.all_span[bg.image_defined] = myMandel.span
		if DEBUG {
			fmt.Printf("zoomPathString: %d", bg.image_defined)
		}
		// FIXME: Add Image
	})
	resetPathBtn := widget.NewButton("Reset", func() {
		if DEBUG {
			fmt.Println("Reset")
		}
		bg.image_defined = 0
		myMandel.ResetSpan()
		myMandel.up_to_date = false
		zoomPathString = bg.PathImageString()
		zoomPathLabel.SetText(zoomPathString)
		zoomPathLabel.Refresh()
	})
	addResetContent.Add(addImageBtn)
	addResetContent.Add(resetPathBtn)

	zoomPathString = bg.PathImageString()
	zoomPathLabel.SetText(zoomPathString)

	// Column One
	//colOneContent := container.New(layout.NewVBoxLayout())
	//colOneContent.Add(widthCtlText)
	colOneContent.Add(selectBackgroundTemplateText)
	colOneContent.Add(selectBackgroundTemplateChoices)
	colOneContent.Add(selectDesktopSizeText)
	colOneContent.Add(selectDesktopSizeChoices)
	colOneContent.Add(selectColorPreferenceText)
	colOneContent.Add(selectColorPreferenceChoices)
	colOneContent.Add(zoomContent)
	colOneContent.Add(panCheckBox)
	colOneContent.Add(addResetContent)
	colOneContent.Add(zoomPathLabel)

	previewText := canvas.NewText("Preview", color.Black)
	myRaster.SetMinSize(fyne.NewSize(256, 256))
	previewContent := container.New(layout.NewHBoxLayout())
	previewContent.Add(layout.NewSpacer())
	previewContent.Add(myRaster)
	previewContent.Add(layout.NewSpacer())

	panControlContent := container.New(layout.NewHBoxLayout())
	panUpBtn := widget.NewButton("Up", func() {
		if DEBUG {
			fmt.Println("Up")
		}
		bg.PanUp(&myMandel)
	})
	panDownBtn := widget.NewButton("Down", func() {
		if DEBUG {
			fmt.Println("Down")
		}
		bg.PanDown(&myMandel)
	})
	panLeftBtn := widget.NewButton("Left", func() {
		if DEBUG {
			fmt.Println("Left")
		}
		bg.PanLeft(&myMandel)
	})
	panRightBtn := widget.NewButton("Right", func() {
		if DEBUG {
			fmt.Println("Right")
		}
		bg.PanRight(&myMandel)
	})
	panZoomInBtn := widget.NewButton("Zoom In", func() {
		if DEBUG {
			fmt.Println("Zoom In")
		}
		bg.PanZoomIn(&myMandel)
	})
	panZoomOutBtn := widget.NewButton("Zoom Out", func() {
		if DEBUG {
			fmt.Println("Zoom In")
		}
		bg.PanZoomOut(&myMandel)
	})
	panControlContent.Add(panUpBtn)
	panControlContent.Add(panDownBtn)
	panControlContent.Add(panLeftBtn)
	panControlContent.Add(panRightBtn)
	panControlContent.Add(panZoomInBtn)
	panControlContent.Add(panZoomOutBtn)

	colTwoContent := container.New(layout.NewVBoxLayout())
	colTwoContent.Add(layout.NewSpacer())
	colTwoContent.Add(previewText)
	colTwoContent.Add(previewContent)
	colTwoContent.Add(panControlContent)
	colTwoContent.Add(layout.NewSpacer())

	// Botton Content Creation
	imageGenerationProgressBar := widget.NewProgressBar()
	backgroundGenerationProgressBar := widget.NewProgressBar()

	//bottomContent := container.New(layout.NewVBoxLayout())
	generateBtn := widget.NewButton("Generate Background", func() {
		// Check
		if bg.image_defined != bg.TotalImages() {
			if DEBUG {
				fmt.Println("Location for all images is not defined")
			}
			var popup *widget.PopUp
			all_defined_label := widget.NewLabel("Location for all images is not defined")
			popUpContent := container.NewVBox(
				all_defined_label,
				widget.NewButton("Ok", func() {
					popup.Hide()
				}),
			)
			popup = widget.NewModalPopUp(popUpContent, myWindow.Canvas())
			popup.Show()
			return
		}
		if DEBUG {
			fmt.Println("Generate Background")
		}
		mbg_image := image.NewRGBA(image.Rect(0, 0, bg.desktop_x_dots[bg.desktop_num], bg.desktop_y_dots[bg.desktop_num]))
		if DEBUG {
			fmt.Println("Making Black")
		}
		for px := 0; px < bg.desktop_x_dots[bg.desktop_num]; px++ {
			for py := 0; py < bg.desktop_y_dots[bg.desktop_num]; py++ {
				mbg_image.Set(px, py, color.RGBA{0x0, 0x0, 0x0, 0xff})
			}
		}
		// Draw mandelbrots
		for i_num := 0; i_num < bg.TotalImages(); i_num++ {
			// Generate a Mandelbrot
			imageGenerationProgressBar.SetValue(float64(0))
			new_mandelbrot := NewMandel(
				bg.all_min_x[i_num],
				bg.all_min_y[i_num],
				bg.all_span[i_num],
				bg.templates[bg.template_num].Images[i_num].Side_size*bg.PixelsPerUnit(),
				bg.color_theme_num,
			)
			for {
				imageGenerationProgressBar.SetValue(new_mandelbrot.PercentCalced())
				if new_mandelbrot.up_to_date {
					break
				} else {
					//fmt.Print(".")
					new_mandelbrot.UpdateSome()
				}
			}
			bg.images = append(bg.images, new_mandelbrot)

			pixels_per_unit := bg.PixelsPerUnit()
			// Calculate the padx and pady
			padx := bg.CalcPadX()
			pady := bg.CalcPady()
			// Calculate the x and y units
			pos_x := bg.templates[bg.template_num].Images[i_num].Bg_x*pixels_per_unit + padx
			pos_y := bg.templates[bg.template_num].Images[i_num].Bg_y*pixels_per_unit + pady
			TranferMandelToImage(new_mandelbrot, mbg_image, pos_x, pos_y)
			// update progress bar
			backgroundGenerationProgressBar.SetValue(float64(i_num) / float64(bg.TotalImages()-1))
		}
		// Save the image
		//filename_save := bg.templates[bg.template_num].Name + ".bmp"
		//default_file_name_save := "mbg"+ ".png"
		file_name_save := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
			if err != nil {
				fmt.Println("Error in Save")
				return
			}

			if uc == nil {
				// user canceled
				return
			}

			buf := new(bytes.Buffer)
			err_png := png.Encode(buf, mbg_image)
			png_bytes := buf.Bytes()
			if err_png != nil {
				fmt.Println("Error in Save: Converting")
				return
			}

			// save file
			//os.WriteFile(uc.URI().Path(), bmp.Encode(mbg_image), 0644)
			//os.WriteFile(uc.URI().Path(), []byte{0xff}, 0644)
			os.WriteFile(uc.URI().Path(), png_bytes, 0644)

			//win.SetTitle(win.Title() + " - " + write.URI().Name())

			//defer uc.Close()
			// Save the image
			//err = bmp.Encode(uc, mbg_image)
			//if err != nil {
			//	fmt.Println("Error in Save")
			//	return
			//}
		}, myWindow)
		save_filename := bg.templates[bg.template_num].Save_filename + ".png"
		file_name_save.SetFileName(save_filename)
		file_name_save.SetOnClosed(func() { fmt.Println("Save Closed") })
		file_name_save.Show()
	})
	backgroundGenerationProgressText := canvas.NewText("Background Generation Progress", color.Black)
	backgroundProgrogressContent := container.New(layout.NewHBoxLayout())
	backgroundProgrogressContent.Add(backgroundGenerationProgressText)
	backgroundProgrogressContent.Add(backgroundGenerationProgressBar)
	imageGenerationProgressText := canvas.NewText("Image Generation Progress", color.Black)
	imageGenerationProgrogressContent := container.New(layout.NewHBoxLayout())
	imageGenerationProgrogressContent.Add(imageGenerationProgressText)
	imageGenerationProgrogressContent.Add(imageGenerationProgressBar)

	colOneContent.Add(generateBtn)
	colOneContent.Add(backgroundProgrogressContent)
	colOneContent.Add(imageGenerationProgrogressContent)

	topContent := container.New(layout.NewHBoxLayout())
	topContent.Add(layout.NewSpacer())
	topContent.Add(colOneContent)
	topContent.Add(colTwoContent)
	topContent.Add(layout.NewSpacer())

	wholeContent := container.New(layout.NewVBoxLayout())
	wholeContent.Add(layout.NewSpacer())
	wholeContent.Add(topContent)
	wholeContent.Add(layout.NewSpacer())
	//wholeContent.Add(bottomContent)

	myWindow.SetContent(wholeContent)

	go func() {
		for {
			//fmt.Println(myMandel)
			if !myMandel.up_to_date {
				myMandel.UpdateSome()
				myRaster.Refresh()
			} else {
				time.Sleep(time.Nanosecond * 100000000)
			}
		}
	}()

	myWindow.ShowAndRun()
}
