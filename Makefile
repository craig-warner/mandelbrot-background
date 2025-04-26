# This is just a Makefile to simplify running some of the commands mentioned in the README.

build:
	flatpak-builder --user --force-clean build-dir io.github.craig-warner.mandelbrot-background.yml

install:
	flatpak-builder --user --install --force-clean build-dir io.github.craig-warner.mandelbrot-background.yml

run:
	flatpak run --user io.github.craig-warner.mandelbrot-background	