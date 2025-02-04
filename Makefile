.ONESHELL:

clean:
	rm -rf build

builddir:
	mkdir -p build

showvideo:
	wasmvision run -p hello -logging=error --source 2

advertising:
	tinygo flash -size short -target=pico-w -monitor ./demo/advertising

scanner:
	go run ./demo/scanner
