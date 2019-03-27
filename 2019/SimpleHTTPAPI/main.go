package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	name    string
	version string
	gitSHA  string
)

const usage = `version: %s - git: %s
Usage: %s [-h] [-v]
Options:
  -h            this help
  -v            show version and exit

Examples: 
  %[3]s -t localhost:8888       run the server
`

func main() {
	var vers bool

	flag.Usage = func() {
		w := os.Stderr
		for _, arg := range os.Args {
			if arg == "-h" {
				w = os.Stdout
				break
			}
		}
		fmt.Fprintf(w, usage, version, gitSHA, name)
	}

	flag.BoolVar(&vers, "v", false, "")
	flag.Parse()

	if vers {
		fmt.Fprintf(os.Stdout, "version: %s\n", version)
		return
	}

	router := mux.NewRouter()

	router.Handle("/some/path", xHandler).Methods(http.MethodGet)

}

func xHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
