# mandelbrot-background
Mandelbrot background drawing program written in go using fyne.io 

# Running on Windows
## msys-x64
```
go build -v *.go
```
```
go run -v *.go
```

# Building .apk for Android Studio
## msys-x64 
```
fyne package -os android -appID com.example.mandelbrotbackground -icon assets/mandelbrot-background.png
```
## Installing .apk to Android Studio
Drag Mandelbrot_Background.apk to emulator icon

# Building Windows 
```
fyne package -os windows -icon assets/mandelbrot-background.png
```

# Building Linux 
```
fyne package -os linux -icon assets/mandelbrot-background.png
```
## Installaion on Ubuntu 
```
tar xf 'Mandelbrot Background.tar.xz'
sudo make install
```
## Assumptions
 * GOHOME = /home/craigwarner
```
%sudo apt install libgl1-mesa-dev xorg-dev
```

# Intalling  
```
sudo /home/craigwarner/go/bin/fyne get github.com/craig-warner/mandelbrot-background
```

# Constuction Notes
1) go mod init
2) go mod tidy


# Development To Do

## Basic

## Bug 
* Android Save only generates zero byte files
* Small images on preview at first
* FIXED: Generate Background Progress (image progress string)
* FIXED: Zoom / Pan Controll reset when selection popup is invoked
* Pink needs to be changed to be mroe like what people think of as pink
* Figure out why the images get less well formed the more you zoom in
  * Releation to Threshold?
  * Releation to bits per color?

## Enhancements
* Image Format Selection
* Colors
  * Add More Colors
  * Color Preview (DONE)
  * Yellow, Orange, Pink, Iowa State Red
* Math
   * Output json of span, x, y
   * Load pathfile
   * Threshold control (DONE)
* Templates
    * Preview (DONE)
    * Numbers (DONE)
    * More Templates
* Stretch Progress Bar
* Documentation
    * Pop Up Window
    * Pointer to PBS
        https://www.youtube.com/watch?v=qABFYiYqXSU
    * Mandelbrot math
* Image Growth with Window Resize
* Show Intermediate images
* Allow color changes between images (DONE)
* Android
   * Test

## Spin-offs
* Controlled Info, Error Debug Print 

## Go Code Enhacements
* Multiple files 

## Test Work
* Test on Android 
