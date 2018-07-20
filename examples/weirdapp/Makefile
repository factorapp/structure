.PHONY: wasm
wasm:
	GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o ./server/app.wasm ./client

.PHONY: server-app
server-app:
	GOOS=darwin GOARCH=amd64 ~/gowasm/bin/go build -o server-app ./server

.PHONY: run
run: wasm server-app
	./server-app
