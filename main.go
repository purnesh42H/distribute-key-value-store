package main

import (
  "log"
  "net/http"
	"time"
  "os"
  "strconv"
)

var store = map[string]*Data{}
var prefix  = ""
var serverId = ""
var n uint32

func main() {
	router := NewRouter()
  prefix = os.Args[1]
  serverId = os.Args[2]
  conv, err := strconv.ParseInt(os.Args[3], 10, 32)
  n = uint32(conv)
  if err != nil {
        panic(err)
  }
  log.Fatal(http.ListenAndServe(":" + prefix + serverId, router))
}

func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
    })
}
