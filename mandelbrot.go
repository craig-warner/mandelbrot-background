package main

import (
	"github.com/craig-warner/mandelbrot-background/pkg/ctlprint"

	"encoding/json"
	"fmt"
	"image/color"
)

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
	cp            ctlprint.CtlPrint
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
	m.cp.DbgPrint("center_x:%f,center_y:%f\n", center_x, center_y)
	// Reduce span
	m.span = m.span * adj
	m.span_one_dot = m.span / float64(m.size)
	// Set upper left point
	m.min_x = center_x - (m.span / 2.0)
	m.min_y = center_y - (m.span / 2.0)
}

func NewMandel(min_x, min_y, span float64, size, color_theme_num int, cp ctlprint.CtlPrint) Mandel {
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
		cp:            cp,
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
	return m
}
