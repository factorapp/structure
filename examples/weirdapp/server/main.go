// build !js,wasm
package main

import (
	"log"
	"net/http"
)

func wasmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	http.ServeFile(w, r, "./server/app.wasm")
}
func main() {
	http.HandleFunc("/app/app.wasm", wasmHandler)
	/*	cwd, err := os.Getcwd()
		if err != nil {
			panic(err)
		}
		app := filepath.Join(cwd, "app")
	*/
	http.HandleFunc("/wasm_exec.js", jsHandler)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", indexHandler)
	log.Println("serving on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func jsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./server/wasm_exec.js")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./components/index.html")
}
