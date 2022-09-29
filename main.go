package main

import (
	a "github.com/jdboisvert/quotes-service-go/app"
)

func main() {
	app := a.App{}

	app.Initialize()
	app.Run()
}
