.PHONY: generate
generate:
	GOARCH=wasm GOOS=js ~/gowasm/bin/go generate

.PHONY: wasm
wasm:
	GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o ./app/app.wasm .

.PHONY: server-app
server-app:
	~/gowasm/bin/go build -o server-app ./server

.PHONY: run
run: generate wasm server-app
	./server-app
