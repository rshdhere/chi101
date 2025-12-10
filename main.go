package main

import "flag"

func main() {
	var PORT int

	flag.IntVar(&PORT, "port", 8080, "port for a go-backend")
	flag.Parse()
}
