package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting"
)

func main() {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			highlighting.Highlighting,
		),
	)
	source, err := ioutil.ReadFile("./test.md")
	if err != nil {
		panic(err.Error())
	}
	file, err := os.Create("test.html")
	if err != nil {
		panic(err.Error())
	}
	err = markdown.Convert(source, file)
	if err != nil {
		log.Printf("goldmark.Convert error:%s", err.Error())
	} else {
		fmt.Println("convert 2 html succes")
	}
}
