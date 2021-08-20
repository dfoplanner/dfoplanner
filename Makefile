all: static/main.wasm static/wasm_exec.js
	goexec 'http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`./tests/web/static/`)))'

static/wasm_exec.js:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" tests/web/static

static/main.wasm : tests/web/main.go
	GO111MODULE=auto GOOS=js GOARCH=wasm go build -o tests/web/static/main.wasm .