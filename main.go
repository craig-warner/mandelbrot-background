package main

import (
	"github.com/craig-warner/mandelbrot-background/pkg/ctlprint"

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

const VERBOSE = false
const DEBUG = false

const WINDOW_SIZE = 512

// const WINDOW_SIZE = 1024
const (
	MAX_DISPLAY_SIZE = 10000
	MAX_IMAGES       = 256
)

/*
func DbgPrint(str ...interface{}) {
	if DEBUG {
		fmt.Println(str...)
		return
	}
}
*/

/*
type tappableRaster struct {
	fyne.CanvasObject
	OnTapped func()
}

func NewTappableRaster(raster fyne.CanvasObject, onTapped func()) *tappableRaster {
	return &tappableRaster{CanvasObject: raster, OnTapped: onTapped}
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

	cp := ctlprint.NewCtlPrint(VERBOSE, DEBUG)
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
	bg := NewBackground(cp)

	// Mandelbrot
	myMandel := NewMandel(-1.0, -1.5, 3.0, 256, 0, cp)
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
				cp,
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
		file_name_save.SetOnClosed(func() {
			//fmt.Println("Save Closed")
			cp.VerbosePrint("Save Closed")
		})
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
