package main

import (
	"github.com/cagiti/go-tawerin/pkg/app"
)

func main() {
	a := app.App{}
	a.Initialize()
	a.Run(":8080")
}
