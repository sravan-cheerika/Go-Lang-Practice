package subDirectory

import (
	"fmt"
	"net/http"
)

// PlayerServer currently returns Hello, world given _any_ request.
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}

func MethodHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Team")
}

func MethodHeader(w http.ResponseWriter, r *http.Request) {
	
	for name, headers := range r.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
	//fmt.Fprint(w, "Header code")
}
