package main

import (
	"github.com/screw-coding/http"
	"github.com/screw-coding/index"
)

func main() {
	println(http.NewHttpServer())
	println(index.InvertedIndex())
}
