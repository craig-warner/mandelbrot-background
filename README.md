# mandelbrot-background
Mandelbrot background drawing program written in go using fyne.io 

# Running on Windows
## msys-x64
% go build -v main.go
% go run -v main.go

# Building .apk for Android Studio
## msys-x64 
% fyne package -os android -appID com.example.mandelbrot-background-icon assets/mandelbrot-background.png
## Installing .apk to Android Studio
Drag Mandelbrot_Background.apk to emulator icon

# Building Windows 
% fyne package -os windows -icon assets/mandelbrot-background.png

# Building Linux 
% fyne package -os linux -icon assets/mandelbrot-background.png
## Installaion on Ubuntu 
%tar xf 'Mandelbrot Background.tar.xz'
%sudo make install
## Assumptions
 * GOHOME = /home/craigwarner
%sudo apt install libgl1-mesa-dev xorg-dev
%go run -v main.go dict.go

# Building .apk for Android Studio
## msys-x64 
% fyne package -os android -appID com.example.mandelbrotbackground -icon assets/mandelbrot-background.png
% fyne package -os windows -icon assets/mandelbrot-back.png
 
# Installing .apk to Android Studio
Drag Mandelbrot_Background.apk to emulator icon

# Installaion on Ubuntu 
## - Assumptions
 * GOHOME = /home/craigwarner
%sudo apt install libgl1-mesa-dev xorg-dev
%go run -v main.go

# Intalling  
1) sudo /home/craigwarner/go/bin/fyne get github.com/craig-warner/mandelbrot-background

# Constuction Notes
1) go mod init
2) go mod tidy


