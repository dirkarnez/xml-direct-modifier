package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/antchfx/xmlquery"
)

var (
	source string
	target string
	mode   string
	xpath  string
	value  string
)

func main() {
	flag.StringVar(&source, "source", "", "Absolute path for source .xml file")
	flag.StringVar(&target, "target", "", "Absolute path for target .xml file")
	flag.StringVar(&mode, "mode", "innertext", `Mode: "innertext" is only supported. Default "innertext" mode`)
	flag.StringVar(&xpath, "xpath", "", "XPath for manipulate the source .xml file")
	flag.StringVar(&value, "value", "", "the updated value")
	flag.Parse()

	if len(source) < 1 {
		log.Fatal(`"source" must be pprovided!`)
	}

	if len(target) < 1 {
		log.Fatal(`"target" must be pprovided!`)
	}

	if mode != "innertext" {
		log.Fatal(`"mode" must be "innertext" in this version`)
	}

	file, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	doc, err := xmlquery.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	// "Project/PropertyGroup/ConfigurationType"
	for _, n := range xmlquery.Find(doc, xpath) {
		switch n.Type {
		case xmlquery.ElementNode:
			if n.FirstChild.Type == xmlquery.TextNode {
				n.FirstChild.Data = value
			}
		}
	}

	root := xmlquery.FindOne(doc, "/") // root node
	newFile, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
	}

	defer newFile.Close()
	w := bufio.NewWriter(newFile)
	fmt.Fprintln(w, root.OutputXML(true))
	w.Flush()
}
