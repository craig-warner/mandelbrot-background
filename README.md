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

# Building .apk for Android Studio
## msys-x64
```
fyne package -os android -appID com.example.mandelbrotbackground -ic* Image Format Selection
```
 
# Installing .apk to Android Studio
Drag Mandelbrot_Background.apk to emulator ico* Image Format Selection
# Installaion on Ubuntu 
## - Assumptions
 * GOHOME = /home/craigwarner
```
%sudo apt install libgl1-mesa-dev xorg-dev
go run -v main.go
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
* FIXED: Generate Background Progress (image progress string)
* FIXED: Zoom / Pan Controll reset when selection popup is invoked

## Enhancements
* Image Format Selection
* Colors
  * Add More Colors
  * Color Preview (DONE)
  * Yellow, Orange, Pink, Iowa State Red
* Math
   * Output json of span, x, y
   * Load pathfile
   * Color for each image
   * Threshold control (DONE)
* Templates
    * Preview (DONE)
    * Numbers (DONE)
    * More Templates
* Stretch Progress Bar
* Documentation
    * Pop Up Window
    * Pointer to PBS
    * Mandelbrot math
* Image Growth with Window Resize
* Show Intermediate images
* Allow color changes between images
* Android
   * Test
   * Save Test

## Spin-offs
* Controlled Info, Error Debug Print 

## Go Code Enhacements
* Multiple files 

## Test Work
* Test on Android 
