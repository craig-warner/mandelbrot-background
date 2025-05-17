# This is just a Makefile to simplify running some of the commands mentioned in the README.

build:
	flatpak-builder --user --force-clean build-dir io.github.craigwarner.mandelbrot-background.yml

install:
	flatpak-builder --user --install --force-clean build-dir io.github.craigwarner.mandelbrot-background.yml

run:
	flatpak run --user io.github.craigwarner.mandelbrot-background	

clean:
	rm -rf build-dir
	flatpak-builder --user --force-clean --delete-build-dir build-dir io.github.craigwarner.mandelbrot-background.yml
	flatpak-builder --user --force-clean --delete-cache build-dir io.github.craigwarner.mandelbrot-background.yml
	flatpak uninstall --unused