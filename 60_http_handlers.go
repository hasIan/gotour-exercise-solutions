package main

import (
    "fmt"
    "net/http"
)

type String string

type Greeting struct {
    Greeting    string
    Punctuation string
    Who         string
}

func (s String) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, s)
}

func (g *Greeting) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprintf(w, "%s%s %s",
      g.Greeting, g.Punctuation, g.Who)
}

func main() {
    http.Handle("/string", String("Wh00p Wh00p"))
    http.Handle("/struct", &Greeting{
        "Hello", ":", "Gophers!"
    })
    http.ListenAndServe("localhost:4000", nil)
}
