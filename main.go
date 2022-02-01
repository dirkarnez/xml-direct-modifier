package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/antchfx/xmlquery"
)

var (
	source string
	target string
	xpath  string
	value  string
)

func main() {
	file, err := os.Open("WinReg.vcxproj")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// sp, err := xmlquery.CreateStreamParser(file, "Project/PropertyGroup/ConfigurationType")
	// if err != nil {
	// 	panic(err)
	// }
	// for {
	// 	n, err := sp.Read()
	// 	if err != nil {
	// 		break
	// 	}

	// 	fmt.Println(n.InnerText())
	// 	fmt.Println(n.OutputXML(true))
	// }

	doc, err := xmlquery.Parse(file)
	if err != nil {
		panic(err)
	}
	for _, n := range xmlquery.Find(doc, "Project/PropertyGroup/ConfigurationType") {
		switch n.Type {
		case xmlquery.ElementNode:
			if n.FirstChild.Type == xmlquery.TextNode {
				n.FirstChild.Data = "Hahaha"
			}
		}
	}

	root := xmlquery.FindOne(doc, "/")

	sss, _ := os.Create("exported")

	defer sss.Close()

	w := bufio.NewWriter(sss)

	fmt.Fprintln(w, root.OutputXML(true))

	w.Flush()
}
