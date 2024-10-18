// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"html/template"
	"fmt"
	"log"
	"os"
)

type js struct {
	Root string
}

type doc struct {
	Root string
}

func gen() {
	render("static/index.html", "tpl/home.html", doc{
		Root: ".",
	})
	renderJS("static/loader.js", "tpl/loader.js", js{
		Root: ".",
	})
	render("static/about.html", "tpl/about.html", doc{
		Root: ".",
	})
	render("static/play/index.html", "tpl/play.html", doc{
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

var tplbase string = "tpl/base.html"

func render(dst, src string, data doc) {
	// clean
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
	// gen
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

func renderJS(dst, src string, data js) {
	// clean
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
	// gen
	t, err := template.ParseFiles(src)
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
