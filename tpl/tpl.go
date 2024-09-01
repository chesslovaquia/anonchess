// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"html/template"
	"log"
	"os"
)

type page struct {
	Root string
}

var tplbase string = "tpl/base.html"

func render(dst, src string, data page) {
	t, err := template.ParseFiles(tplbase, src)
	if err != nil {
		log.Fatalf("%s - parse error: %v", src, err)
	}

	fh, err := os.Create(dst)
	if err != nil {
		log.Fatalf("%s: %v", dst, err)
	}
	defer fh.Close()

	err = t.Execute(fh, &data)
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Println(dst)
}

func main() {
	render("static/index.html", "tpl/home.html", page{
		Root: ".",
	})
	render("static/play/index.html", "tpl/play.html", page{
		Root: "..",
	})
}
