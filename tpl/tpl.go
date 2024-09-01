// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"html/template"
	"fmt"
	"log"
	"os"
)

func gen() {
	render("static/index.html", "tpl/home.html", page{
		Root: ".",
	})
	render("static/play/index.html", "tpl/play.html", page{
		Root: "..",
	})
}

var cleanup bool = false

func main() {
	if os.Getenv("TPL") == "clean" {
		cleanup = true
	}
	gen()
}

type page struct {
	Root string
}

var tplbase string = "tpl/base.html"

func render(dst, src string, data page) {
	if cleanup {
		if err := os.Remove(dst); err != nil {
			if !os.IsNotExist(err) {
				log.Fatal(err)
			}
		} else {
			fmt.Println(dst, "removed")
		}
		return
	}

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

	fmt.Println(dst)
}
