package main

import (
	"encoding/json"
	"fmt"

	"github.com/craig-warner/mandelbrot-background/pkg/ctlprint"
)

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
    	"num_images": 8,
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
    	"num_images": 11,
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
    		"version": "2.0",
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
		"Big Left (11 Images)",
		"First Center (5 Images)",
		"Big Left (14 Images)",
		"Big Center (17 Images)",
		"Diagonal (17 Images)"
	]`
)

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
	image_defined      int
	//	cur_min_x          float64
	//	cur_min_y          float64
	//	cur_span           float64
	threshold      float64
	fine_grain_pan bool
	pan_speed      float64
	all_min_x      []float64
	all_min_y      []float64
	all_span       []float64
	templates      []Template
	images         []Mandel
	desktop_x_dots []int
	desktop_y_dots []int
	cp             ctlprint.CtlPrint
}

type Point struct {
	x float64
	y float64
}

/*
 * Background functions
 */

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
	bg.cp.DbgPrint("template num %d", bg.template_num)
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

func (bg *Background) SetThreshold(threshold float64) {
	bg.threshold = threshold
}

func (bg *Background) GetThresholdString() string {
	str := ""
	str = fmt.Sprintf("%f", bg.threshold)
	return str
}

func NewBackground(cp ctlprint.CtlPrint) Background {
	bg := Background{
		template_num:       0,
		desktop_num:        0,
		color_theme_num:    0,
		zoom_magnification: 1,
		image_defined:      0,
		//		cur_min_x:          -1.0,
		//		cur_min_y:          -1.5,
		//		cur_span:           3.0,
		threshold:      10.0,
		fine_grain_pan: true,
		pan_speed:      0.1,
		cp:             cp,
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
	cp.DbgPrint("Templates: %+v", bg.templates)
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
	cp.DbgPrint("Templates: %+v", bg.desktop_y_dots)
	return bg
}
