package main

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
	  },
      { "ibits": 6,
	"blue_pos":     [   -1,   -1,   -1,  -1,  -1,  -1],
	"red_pos":     [   2,   3,   4,  5,  6,  7],
	"green_pos":     [   2,   3,   4,  5,  6,  7],
	"default_color": [0,0,0]
	  },
      { "ibits": 6,
	"blue_pos":     [   -1,   -1,   -1,  -1,  -1,  -1],
	"red_pos":     [   7,   6,   5,  4,  3,  2],
	"green_pos":     [   7,   6,   5,  4,  3, 2],
	"default_color": [0,0,0]
	  },
      { "ibits": 6,
	"blue_pos":     [   2,   3,   4,  5,  6,  7],
	"red_pos":     [   2,   3,   4,  5,  6,  7],
	"green_pos":     [   -1,   -1,   -1,  -1,  -1,  -1],
	"default_color": [0,0,0]
	  },
      { "ibits": 6,
	"blue_pos":     [   7,   6,   5,  4,  3, 2],
	"red_pos":     [   7,   6,   5,  4,  3,  2],
	"green_pos":     [   -1,   -1,   -1,  -1,  -1,  -1],
	"default_color": [0,0,0]
	  },
      { "ibits": 6,
	"blue_pos":     [   2,   3,   4,  5,  6,  7],
	"red_pos":     [   -1,   -1,   -1,  -1,  -1,  -1],
	"green_pos":     [   1,   2,   3,  4,  5,  6],
	"default_color": [0,0,0]
	  },
      { "ibits": 6,
	"blue_pos":     [   7,   6,   5,  4,  3, 2],
	"red_pos":     [   -1,   -1,   -1,  -1,  -1,  -1],
	"green_pos":     [   6,   5,   4,  3,  2,  1],
	"default_color": [0,0,0]
	  },
      { "ibits": 6,
	"blue_pos":     [   1,   2,   3,  4,  5,  6],
	"red_pos":     [   2,   3,   4,  5,  6,  7],
	"green_pos":     [   0,   1,   2,  3,  4,  5],
	"default_color": [0,0,0]
	  },
      { "ibits": 6,
	"blue_pos":     [   6,   5,   4,  3,  2,  1],
	"red_pos":     [   7,   6,   5,  4,  3,  2],
	"green_pos":     [   5,   4,   3,  2,  1,  0],
	"default_color": [0,0,0]
	  },
      { "Ibits": 12,
	"blue_pos":    [  6,  7,  -1,  -1,  -1,  -1,  6,   7, -1,  -1, -1, -1],
	"green_pos":   [ -1, -1,   6,   7,  -1,  -1,  -1, -1,  6,  7,  -1, -1],
	"red_pos":     [ -1, -1,  -1,  -1,   6,   7 , -1, -1, -1, -1,   6,  7],
	"default_color": [0,0,0] 
      },
      { "Ibits": 12,
	"blue_pos":   [ -1, -1,   6,   7,  -1,  -1,  -1, -1,  6,  7,  -1, -1],
	"green_pos":    [  6,  7,  -1,  -1,  -1,  -1,  6,   7, -1,  -1, -1, -1],
	"red_pos":     [ -1, -1,  -1,  -1,   6,   7 , -1, -1, -1, -1,   6,  7],
	"default_color": [0,0,0] 
	  },
      { "Ibits": 12,
	"blue_pos":     [  -1,  -1,  5,  6,   7,  -1,  -1,  -1,  5,  6,  -1, -1],
	"green_pos":    [  -1,  -1, -1, -1,   5,   6,  -1,  -1, -1, -1,  5,  6],
	"red_pos":      [   6,  7,   6,  7,   6,   7,   6,   7,  6,  7,  6,  7],
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
	    "All Blue",
	    "All Purple",
	    "All Purple - purple center",
	    "All Maroon",
	    "All Orange",
	    "All Lime Green",
	    "All Gold",
	    "All Black",
	    "high resolution: blue, green",
	    "All Yellow",
	    "Reverse Yellow",
	    "All Magenta",
	    "Reverse Magenta",
	    "All Ocean",
	    "Reverse Ocean",
	    "All Pink",
	    "Reverse Pink",
	    "very bold: blue,green,red",
	    "very bold: green,red, blue",
	    "very bold: red,magenta,yellow"
	]`
)

// Field Names MUST start with a capital letter
type MandelColor struct {
	Ibits         int
	Blue_pos      []int
	Red_pos       []int
	Green_pos     []int
	Default_color []uint8
}

type Color struct {
	red   uint8
	green uint8
	blue  uint8
}

func NewColor(r, g, b uint8) Color {
	c := Color{
		red:   r,
		green: g,
		blue:  b,
	}
	return c
}
