// cmd/marketquantify/main.go
package main

import (
"flag"
"log"
"os"

"marketquantify/internal/marketquantify"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := marketquantify.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
