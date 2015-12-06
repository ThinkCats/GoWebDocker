package main

import (
	"encoding/xml"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Greeting struct {
	XMLName xml.Name `xml:"greeting"`
	One     string   `xml:"one,attr"`
	Two     string   `xml:"two,attr"`
}

func main() {
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory:"templates",
		Layout:"layout",
		Extensions:[]string{".html"},
		Charset:"UTF-8",
	}))

	m.Get("/", func(r render.Render){
		r.HTML(200,"hello","wanglei")
	})

	m.Get("/api", func(r render.Render) {
		r.JSON(200, map[string]interface{}{"hello": "world"})
	})

	// This will set the Content-Type header to "text/xml; charset=UTF-8"
	m.Get("/xml", func(r render.Render) {
		r.XML(200, Greeting{One: "hello", Two: "world"})
	})

	// This will set the Content-Type header to "text/plain; charset=UTF-8"
	m.Get("/text", func(r render.Render) {
		r.Text(200, "hello, world")
	})

	m.Run()
}
