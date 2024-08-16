
all: demo.linux

demo.linux: main.go
	go build -o demo.linux main.go

demo.wasm: main.go
	GOOS=js GOARCH=wasm go build -o  demo.wasm