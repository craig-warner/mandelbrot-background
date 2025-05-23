package main

import (
	"math"
	"strconv"

	"github.com/craig-warner/mandelbrot-background/pkg/ctlprint"

	"bytes"
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
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const VERBOSE = false
const DEBUG = false

const WINDOW_SIZE = 512
const COLOR_PREVIEW_SIZE = 64
const TEMPLATE_PREVIEW_SIZE_X = 256
const TEMPLATE_PREVIEW_SIZE_Y = 180
const PREVIEW_SIZE = 256

// const WINDOW_SIZE = 1024
const (
	MAX_DISPLAY_SIZE = 10000
	MAX_IMAGES       = 256
)

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

	// Background
	bg := NewBackground(cp)
	// Mandelbrot
	myMandel := NewMandel(-1.0, -1.5, 3.0, PREVIEW_SIZE, 0, bg.threshold, cp)
	myMandel.ResetSpan()
	myMandel.ResetWindow(PREVIEW_SIZE, PREVIEW_SIZE)
	// Raster
	myRaster := canvas.NewRasterWithPixels(myMandel.DrawOneDot)

	// Resize ignored by Mobile Platforms
	// - Mobile platforms are always full screen
	// - 27 is a hack determined by Ubuntu/Gnome
	//myWindow.Resize(fyne.NewSize(256, (256 + 27)))
	myWindow.Resize(fyne.NewSize(WINDOW_SIZE, (WINDOW_SIZE + 27)))

	// Control Menu Set up
	menuItemTemplate := fyne.NewMenuItem("Template Settings", func() {
		cp.InfoPrint("In Template Settings")
		var popup *widget.PopUp
		new_template_num := bg.template_num
		selectBackgroundTemplateText := widget.NewLabel("Select a background temple")
		selectBackgroundTemplateChoicesStrings := bg.GetTemplateChoicesStrings()
		selectBackgroundTemplateChoices := widget.NewSelect(selectBackgroundTemplateChoicesStrings, func(s string) {
			cp.DbgPrint("Select Background Template Callback:", s)
			for i := 0; i < len(selectBackgroundTemplateChoicesStrings); i++ {
				if selectBackgroundTemplateChoicesStrings[i] == s {
					// New Template
					new_template_num = i
					break
				}
			}
		})
		/*
			// Template Preview
			template_preview := container.NewWithoutLayout()
			// Big Rectangle
			rect := canvas.NewRectangle(color.Black)
			rect.SetMinSize(fyne.NewSize(TEMPLATE_PREVIEW_SIZE, TEMPLATE_PREVIEW_SIZE))
			template_preview.Add(rect)
			cp.DbgPrint("TemplateNum:", bg.template_num)
			// Image Numbers
			units := TEMPLATE_PREVIEW_SIZE / bg.templates[bg.template_num].X_units
			units_half := units / 2
			cp.DbgPrint("Units:", units)
			num_images := bg.templates[bg.template_num].Num_images
			cp.DbgPrint("NumImages:", num_images)
			for i := 0; i < num_images; i++ {
				image_size := bg.templates[bg.template_num].Images[i].Side_size
				// Calculate the x and y
				pos_x := bg.templates[bg.template_num].Images[i].Bg_x*units + units_half*image_size
				pos_y := bg.templates[bg.template_num].Images[i].Bg_y*units + units_half*image_size
				image_num_str := strconv.Itoa(i)
				image_num_text := canvas.NewText(image_num_str, color.Black)
				template_preview.Add(image_num_text)
				image_num_text.Move(fyne.NewPos(float32(pos_x), float32(pos_y)))
				cp.DbgPrint("id:posx:posy:", i, pos_x, pos_y)
			}
			//MinSize(fyne.NewSize(TEMPLATE_PREVIEW_SIZE, TEMPLATE_PREVIEW_SIZE))
			//template_preview := container.New(layout.NewAdaptiveGridLayout(3))
			//template_preview.Add(rect)
			//template_preview.Add(text1)
			//template_preview.Add(text2)
			//template_preview.Add(text3)
			//template_preview.Add(text4)
			//text1.Move(fyne.NewPos(10, 10))
			//text2.Move(fyne.NewPos(20, 20))
		*/
		selectBackgroundTemplateChoices.SetSelectedIndex(bg.template_num)
		popUpContent := container.NewVBox(
			selectBackgroundTemplateText,
			selectBackgroundTemplateChoices,
			//template_preview,
			widget.NewButton("See Template Image Order", func() {
				var nested_popup *widget.PopUp
				// Template Preview
				template_preview := container.NewWithoutLayout()
				// Big Rectangle
				rect := canvas.NewRectangle(theme.Color(theme.ColorNameBackground))
				rect.SetMinSize(fyne.NewSize(TEMPLATE_PREVIEW_SIZE_X, TEMPLATE_PREVIEW_SIZE_Y))
				template_preview.Add(rect)
				cp.DbgPrint("TemplateNum:", new_template_num)
				// Image Numbers
				units := TEMPLATE_PREVIEW_SIZE_X / bg.templates[new_template_num].X_units
				units_half := units / 2
				cp.DbgPrint("Units:", units)
				num_images := bg.templates[new_template_num].Num_images
				cp.DbgPrint("NumImages:", num_images)
				for i := 0; i < num_images; i++ {
					image_size := bg.templates[new_template_num].Images[i].Side_size
					// Calculate the x and y
					pos_x := bg.templates[new_template_num].Images[i].Bg_x*units + units_half*image_size
					pos_y := bg.templates[new_template_num].Images[i].Bg_y*units + units_half*image_size
					// Image Numbering starts with 1 (not 0)
					image_num_str := strconv.Itoa(i + 1)
					image_num_text := canvas.NewText(image_num_str, theme.Color(theme.ColorNameForeground))
					template_preview.Add(image_num_text)
					image_num_text.Move(fyne.NewPos(float32(pos_x), float32(pos_y)))
					cp.DbgPrint("id:posx:posy:", i, pos_x, pos_y)
					/*
						size := image_size * units
						image_num_rect := canvas.NewRectangle(color.RGBA{R: 173, G: 100, B: 156, A: 200})
						template_preview.Add(image_num_rect)
						cp.DbgPrint("size:", i, size)
						image_num_rect.Resize(fyne.NewSize(float32(size), float32(size)))
						image_num_rect.Move(fyne.NewPos(float32(pos_x-(size/2)), float32(pos_y-(size/2))))
						image_num_rect.Show()
					*/
				}
				popUpContent := container.NewVBox(
					template_preview,
					widget.NewButton("Ok", func() {
						nested_popup.Hide()
					}),
				)
				nested_popup = widget.NewModalPopUp(popUpContent, myWindow.Canvas())
				nested_popup.Show()
			}),
			container.NewHBox(
				layout.NewSpacer(),
				widget.NewButton("Ok", func() {
					bg.template_num = new_template_num
					bg.image_defined = 0
					zoomPathString = bg.PathImageString()
					zoomPathLabel.SetText(zoomPathString)
					popup.Hide()

				}),
				widget.NewButton("Cancel", func() {
					popup.Hide()
				}),
				layout.NewSpacer(),
			),
		)
		popup = widget.NewModalPopUp(popUpContent, myWindow.Canvas())
		popup.Show()
	})
	menuItemDesktop := fyne.NewMenuItem("Desktop Settings", func() {
		cp.InfoPrint("In Desktop Settings")
		var popup *widget.PopUp
		new_desktop_num := 0
		selectDesktopSizeText := widget.NewLabel("Select a background temple")
		selectDesktopSizeChoicesStrings := bg.GetDesktopChiocesStrings()
		selectDesktopSizeChoices := widget.NewSelect(selectDesktopSizeChoicesStrings, func(s string) {
			cp.DbgPrint("Select Desktop Size Callback:", s)
			for i := 0; i < len(selectDesktopSizeChoicesStrings); i++ {
				if selectDesktopSizeChoicesStrings[i] == s {
					new_desktop_num = i
					break
				}
			}
		})
		selectDesktopSizeChoices.SetSelectedIndex(bg.desktop_num)
		popUpContent := container.NewVBox(
			selectDesktopSizeText,
			selectDesktopSizeChoices,
			container.NewHBox(
				layout.NewSpacer(),
				widget.NewButton("Ok", func() {
					bg.desktop_num = new_desktop_num
					popup.Hide()

				}),
				widget.NewButton("Cancel", func() {
					popup.Hide()
				}),
				layout.NewSpacer(),
			),
		)
		popup = widget.NewModalPopUp(popUpContent, myWindow.Canvas())
		popup.Show()
	})
	// Color Menu Set up
	menuItemColor := fyne.NewMenuItem("Color Settings", func() {
		cp.InfoPrint("In Color Settings")
		var popup *widget.PopUp
		new_color_num := 0
		// Color Preview Mandel
		selectColorPreviewMandel := NewMandel(-1.0, -1.5, 3.0, COLOR_PREVIEW_SIZE, 0, bg.threshold, cp)
		selectColorPreviewMandel.ResetSpan()
		selectColorPreviewMandel.ResetWindow(COLOR_PREVIEW_SIZE, COLOR_PREVIEW_SIZE)
		selectColorPreviewMandel.SetColorTheme(bg.color_theme_num)
		selectColorPreviewMandel.UpdateAll()
		// Color Preview Raster
		selectColorPreviewRaster := canvas.NewRasterWithPixels(selectColorPreviewMandel.DrawOneDot)
		selectColorPreviewRaster.SetMinSize(fyne.NewSize(COLOR_PREVIEW_SIZE, COLOR_PREVIEW_SIZE))
		selectColorPreferenceText := widget.NewLabel("Select your color preference")
		selectColorPreferenceChoicesStrings := bg.GetColorChiocesStrings()
		selectColorPreferenceChoices := widget.NewSelect(selectColorPreferenceChoicesStrings, func(s string) {
			cp.DbgPrint("Select Color Preference Callback:", s)
			for i := 0; i < len(selectColorPreferenceChoicesStrings); i++ {
				if selectColorPreferenceChoicesStrings[i] == s {
					new_color_num = i
					selectColorPreviewMandel.SetColorTheme(new_color_num)
					selectColorPreviewMandel.UpdateAll()
					selectColorPreviewRaster.Refresh()
					break
				}
			}
		})
		selectColorPreferenceChoices.SetSelectedIndex(bg.color_theme_num)
		colorPreviewCenter := container.NewHBox(container.New(layout.NewHBoxLayout()))
		colorPreviewCenter.Add(layout.NewSpacer())
		colorPreviewCenter.Add(selectColorPreviewRaster)
		colorPreviewCenter.Add(layout.NewSpacer())
		popUpContent := container.NewVBox(
			selectColorPreferenceText,
			selectColorPreferenceChoices,
			colorPreviewCenter,
			container.NewHBox(
				layout.NewSpacer(),
				widget.NewButton("Ok", func() {
					bg.color_theme_num = new_color_num
					myMandel.SetColorTheme(bg.color_theme_num)
					myRaster.Refresh()
					popup.Hide()

				}),
				widget.NewButton("Cancel", func() {
					popup.Hide()
				}),
				layout.NewSpacer(),
			),
		)
		popup = widget.NewModalPopUp(popUpContent, myWindow.Canvas())
		popup.Show()
		selectColorPreviewRaster.Refresh()
	})
	menuItemPanZoom := fyne.NewMenuItem("Pan and Zoom Settings", func() {
		cp.InfoPrint("In Pan and Zoom Settings")
		var popup *widget.PopUp
		new_fine_grain_pan_setting := false
		new_zoom := float64(bg.zoom_magnification)
		zoomMagnificationText := widget.NewLabel("Zoom in Magnification (1x to 2x)")
		zoomMagnificationSlider := widget.NewSlider(1.0, 10.0)
		zoomMagnificationSlider.SetValue(float64(bg.zoom_magnification))
		zoomMagnificationSlider.OnChanged = func(f float64) {
			cp.DbgPrint("Zoom Magnification Callback:", f)
			new_zoom = f
		}
		panCheckBox := widget.NewCheck("Fine Grained Pan", func(b bool) {
			cp.DbgPrint("Zoom In Callback:", b)
			new_fine_grain_pan_setting = b
		})
		cp.DbgPrint("Fine Grain Pan Setting:", bg.fine_grain_pan)
		panCheckBox.SetChecked(bg.fine_grain_pan)
		popUpContent := container.NewVBox(
			zoomMagnificationText,
			zoomMagnificationSlider,
			panCheckBox,
			container.NewHBox(
				layout.NewSpacer(),
				widget.NewButton("Ok", func() {
					bg.zoom_magnification = int(new_zoom)
					// fine grain pan
					bg.fine_grain_pan = new_fine_grain_pan_setting
					if new_fine_grain_pan_setting {
						bg.pan_speed = 0.01
					} else {
						bg.pan_speed = 0.1
					}
					popup.Hide()
				}),
				widget.NewButton("Cancel", func() {
					popup.Hide()
				}),
				layout.NewSpacer(),
			),
		)
		popup = widget.NewModalPopUp(popUpContent, myWindow.Canvas())
		popup.Show()
	})
	menuItemThreshold := fyne.NewMenuItem("Threshold Settings", func() {
		cp.InfoPrint("In Threshold Settings")
		var popup *widget.PopUp
		thresholdText := widget.NewLabel("Mandebrot Color Calculation Threshold (1.0 to 10,0000.0)")
		thresholdEntry := widget.NewEntry()
		thresholdEntry.SetText(bg.GetThresholdString())
		//thresholdEntry.OnChanged = func(s string) {
		//	cp.DbgPrint("Threshold Entry Callback:", s)
		//	new_threshold = true
		//	new_theshold_value, _ = strconv.ParseFloat(s, 64)
		//}
		popUpContent := container.NewVBox(
			thresholdText,
			thresholdEntry,
			container.NewHBox(
				layout.NewSpacer(),
				widget.NewButton("Ok", func() {
					//new_threshold_string := thresholdEntry.SelectedText()
					new_threshold_string := thresholdEntry.Text
					cp.DbgPrint("Threshold Value String:", new_threshold_string)
					new_theshold_value, err := strconv.ParseFloat(new_threshold_string, 64)
					if err != nil {
						cp.DbgPrint("Threshold Value Error")
						bg.SetThreshold(10.0)
						myMandel.SetThreshold(10.0)
					} else {
						// Check: One is the smallest threshold
						if math.Abs(new_theshold_value) < 1.0 {
							new_theshold_value = 1.0
						}
						bg.SetThreshold(math.Abs(new_theshold_value))
						myMandel.SetThreshold(bg.threshold)
					}
					popup.Hide()
				}),
				widget.NewButton("Cancel", func() {
					popup.Hide()
				}),
				layout.NewSpacer(),
			),
		)
		popup = widget.NewModalPopUp(popUpContent, myWindow.Canvas())
		popup.Show()
	})
	menuItemQuit := fyne.NewMenuItem("Quit", func() {
		cp.InfoPrint("In DoQuit:")
		os.Exit(0)
	})
	//	menuControl:= fyne.NewMenu("Control", menuItemColor, menuItemZoom, menuItemQuit);
	//menuControl := fyne.NewMenu("Control", menuItemGenerate, menuItemQuit)
	menuControl := fyne.NewMenu("Control", menuItemTemplate, menuItemDesktop, menuItemColor,
		menuItemPanZoom, menuItemThreshold, menuItemQuit)
	// About Menu Set up
	menuItemAbout := fyne.NewMenuItem("About...", func() {
		dialog.ShowInformation("About Mandelbrot Background v1.1.0", "Author: Craig Warner \n\ngithub.com/craig-warner/mandelbrot-background", myWindow)
	})
	menuHelp := fyne.NewMenu("Help ", menuItemAbout)
	mainMenu := fyne.NewMainMenu(menuControl, menuHelp)
	myWindow.SetMainMenu(mainMenu)

	// Content

	addResetContent := container.New(layout.NewHBoxLayout())
	addImageBtn := widget.NewButton("Add Image", func() {
		// Check
		if bg.image_defined >= bg.TotalImages() {
			cp.ErrorPrint("Too many images defined")
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
		cp.InfoPrint("Add Image")
		bg.all_min_x[bg.image_defined] = myMandel.min_x
		bg.all_min_y[bg.image_defined] = myMandel.min_y
		bg.all_span[bg.image_defined] = myMandel.span
		bg.all_color_nums[bg.image_defined] = bg.color_theme_num
		bg.image_defined++
		zoomPathString = bg.PathImageString()
		zoomPathLabel.SetText(zoomPathString)
		zoomPathLabel.Refresh()
		cp.DbgPrint("zoomPathString: %d", bg.image_defined)
	})
	resetPathBtn := widget.NewButton("Reset", func() {
		cp.InfoPrint("Reset")
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
	//colOneContent.Add(selectBackgroundTemplateText)
	//colOneContent.Add(selectBackgroundTemplateChoices)
	//colOneContent.Add(selectDesktopSizeText)
	//colOneContent.Add(selectDesktopSizeChoices)
	//colOneContent.Add(selectColorPreferenceText)
	//colOneContent.Add(selectColorPreferenceChoices)
	//colOneContent.Add(zoomContent)
	//colOneContent.Add(panCheckBox)
	colOneContent.Add(addResetContent)
	colOneContent.Add(zoomPathLabel)

	//	previewText := canvas.NewText("Preview", color.Black)
	myRaster.SetMinSize(fyne.NewSize(256, 256))
	previewContent := container.New(layout.NewHBoxLayout())
	previewContent.Add(layout.NewSpacer())
	previewContent.Add(myRaster)
	previewContent.Add(layout.NewSpacer())

	panControlContent := container.New(layout.NewHBoxLayout())
	panUpBtn := widget.NewButton("Up", func() {
		cp.DbgPrint("Up")
		bg.PanUp(&myMandel)
	})
	panDownBtn := widget.NewButton("Down", func() {
		cp.DbgPrint("Down")
		bg.PanDown(&myMandel)
	})
	panLeftBtn := widget.NewButton("Left", func() {
		cp.DbgPrint("Left")
		bg.PanLeft(&myMandel)
	})
	panRightBtn := widget.NewButton("Right", func() {
		cp.DbgPrint("Right")
		bg.PanRight(&myMandel)
	})
	panZoomInBtn := widget.NewButton("Zoom In", func() {
		cp.DbgPrint("Zoom In")
		bg.PanZoomIn(&myMandel)
	})
	panZoomOutBtn := widget.NewButton("Zoom Out", func() {
		cp.DbgPrint("Zoom Out")
		bg.PanZoomOut(&myMandel)
	})
	panControlContent.Add(panUpBtn)
	panControlContent.Add(panDownBtn)
	panControlContent.Add(panLeftBtn)
	panControlContent.Add(panRightBtn)
	panControlContent.Add(panZoomInBtn)
	panControlContent.Add(panZoomOutBtn)

	// Botton Content Creation
	imageGenerationProgressBar := widget.NewProgressBar()
	backgroundGenerationProgressBar := widget.NewProgressBar()

	//bottomContent := container.New(layout.NewVBoxLayout())
	generateBtn := widget.NewButton("Generate Background", func() {
		// Check
		if bg.image_defined != bg.TotalImages() {
			cp.InfoPrint("Location for all images is not defined")
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
		cp.DonePrint("Generate Background")
		mbg_image := image.NewRGBA(image.Rect(0, 0, bg.desktop_x_dots[bg.desktop_num], bg.desktop_y_dots[bg.desktop_num]))
		cp.DbgPrint("Making Black")
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
				bg.all_color_nums[i_num],
				bg.threshold,
				cp,
			)
			// update background progress bar
			backgroundGenerationProgressBar.SetValue(float64(i_num) / float64(bg.TotalImages()))
			for {
				imageGenerationProgressBar.SetValue(new_mandelbrot.PercentCalced())
				if new_mandelbrot.up_to_date {
					break
				} else {
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
		}
		backgroundGenerationProgressBar.SetValue(float64(1.0))
		// Save the image
		//filename_save := bg.templates[bg.template_num].Name + ".bmp"
		//default_file_name_save := "mbg"+ ".png"
		file_name_save := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
			if err != nil {
				cp.ErrorPrint("Error in Save")
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
				cp.ErrorPrint("Error in Save: Converting")
				return
			}

			// save file
			//os.WriteFile(uc.URI().Path(), bmp.Encode(mbg_image), 0644)
			//os.WriteFile(uc.URI().Path(), []byte{0xff}, 0644)
			save_err := os.WriteFile(uc.URI().Path(), png_bytes, 0644)
			if save_err != nil {
				cp.ErrorPrint("Error in Save: Writing")
			}

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
			cp.InfoPrint("Save Closed")
		})
		file_name_save.Show()
	})
	backgroundGenerationProgressText := canvas.NewText("Background Generation Progress", theme.Color(theme.ColorNameForeground))
	backgroundProgrogressContent := container.New(layout.NewHBoxLayout(), backgroundGenerationProgressText, backgroundGenerationProgressBar)
	// Center Layout
	/*
		backgroundProgrogressContentHBox := container.New(layout.NewHBoxLayout())
		backgroundProgrogressContent := container.New(layout.NewCenterLayout())
		backgroundProgrogressContentHBox.Add(backgroundGenerationProgressText)
		backgroundProgrogressContentHBox.Add(backgroundGenerationProgressBar)
		backgroundProgrogressContent.Add(backgroundProgrogressContentHBox)
	*/
	// Center Layout - doesn't make a difference
	/*
		backgroundProgrogressContentHBox := container.New(layout.NewHBoxLayout())
		backgroundProgrogressContent := container.New(layout.NewCenterLayout())
		backgroundProgrogressContentHBox.Add(layout.NewSpacer())
		backgroundProgrogressContentHBox.Add(backgroundGenerationProgressText)
		backgroundProgrogressContentHBox.Add(backgroundGenerationProgressBar)
		backgroundProgrogressContentHBox.Add(layout.NewSpacer())
		backgroundProgrogressContent.Add(backgroundProgrogressContentHBox)
	*/

	imageGenerationProgressText := canvas.NewText("Image Generation Progress", theme.Color(theme.ColorNameForeground))
	imageGenerationProgrogressContent := container.New(layout.NewHBoxLayout())
	imageGenerationProgrogressContent.Add(imageGenerationProgressText)
	imageGenerationProgrogressContent.Add(imageGenerationProgressBar)

	colTwoContent := container.New(layout.NewVBoxLayout())
	colTwoContent.Add(layout.NewSpacer())
	//	colTwoContent.Add(previewText)
	colTwoContent.Add(previewContent)
	colTwoContent.Add(panControlContent)
	colTwoContent.Add(addResetContent)
	colTwoContent.Add(zoomPathLabel)
	colTwoContent.Add(generateBtn)
	colTwoContent.Add(backgroundProgrogressContent)
	colTwoContent.Add(imageGenerationProgrogressContent)
	colTwoContent.Add(layout.NewSpacer())

	topContent := container.New(layout.NewHBoxLayout())
	topContent.Add(layout.NewSpacer())
	//topContent.Add(colOneContent)
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
