package main

import (
	"github.com/dylan-clarry/blogapi/app"
	"github.com/dylan-clarry/blogapi/config"
)

func main() {
	config := config.CreateConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":5000")
}