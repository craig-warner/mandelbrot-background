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
fyne package -os android -appID com.example.mandelbrot-background-icon assets/mandelbrot-background.png
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
go run -v main.go dict.go
```

# Building .apk for Android Studio
## msys-x64
```
fyne package -os android -appID com.example.mandelbrotbackground -icon assets/mandelbrot-background.png
```
```
fyne package -os windows -icon assets/mandelbrot-back.png
```
 
# Installing .apk to Android Studio
Drag Mandelbrot_Background.apk to emulator icon

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

## Bub Fixes

## Enhancements
* Image Format Selection
* Add More Colors
* Stretch Progress Bar
* Bigger Window
* Remove Debug Print 
* Add Documentation
* Image Growth with Window Resize
* Show Intermediate images
* Allow color changes between images

## Go Code Enhacements
* Multiple files 

## Test Work
* Test on Android 
