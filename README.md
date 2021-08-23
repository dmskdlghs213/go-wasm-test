# go-wasm-test

Go WebAssembly Hello World

# How to Started

1. create simple main.go

```Go
package main

import "fmt"

func main() {
	fmt.Println("Hello Gopher !!!!")
}

```

2. build from main.go to main.wasm

```Shell
$ GOOS=js GOARCH=wasm go build -o main.wasm
```

3. copy wasm.exec.js for usage go version

```Shell
$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

4. create simple go server for Output html

```Go
package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("listen on http://localhost:8080")
	http.Handle("/", http.FileServer(http.Dir("app")))
	http.ListenAndServe(":"+"8080", nil)
}

```

5. create simple html for read wasm

```html
<html>
    <head>
       <meta charset="utf-8">
       <script src="wasm_exec.js"></script>
       <script>
           const go = new Go();
           WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
               go.run(result.instance);
           });
       </script>
   </head>
    <body>
        <h2>Hello Wasm</h2>
    </body>
</html>
```

6. check read wasm file
![hogehgoe](https://user-images.githubusercontent.com/42029907/130448499-bae03baf-b761-4836-91fd-5f2c1b051193.png)

