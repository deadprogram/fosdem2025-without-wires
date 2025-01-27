.ONESHELL:

clean:
	rm -rf build

builddir:
	mkdir -p build

build/showvideo:
	go build -o ./build/showvideo ./demo/showvideo/main.go

showvideo: builddir build/showvideo
	./build/showvideo 2

advertising:
	tinygo flash -size short -target=pico-w -monitor ./demo/advertising

scanner:
	go run ./demo/scanner
