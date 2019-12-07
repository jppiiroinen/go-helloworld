/*
#############################################################################
# License: MIT
#
# (C) 2019 Juhapekka Piiroinen <juhapekka.piiroinen@1337.fi>
# All Rights Reserved.
#############################################################################
*/
package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
