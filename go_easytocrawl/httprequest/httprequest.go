package httprequest

import (
	"fmt"
	"net/http"
	"bytes"
	"os"
	"github.com/axgle/mahonia"
	"golang.org/x/net/html"
	"strings"
)

var hasfind bool
var buf bytes.Buffer

func start() {
	resp, err := http.Get("https://www.biqiuge.com/book/4772/2946389.html")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", *doc)
	os.Exit(0)
	parse(doc)
	text, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("open: ", err)
	}

	defer text.Close()

	str := utf8(buf.String())
	file := strings.NewReader(str)
	file.WriteTo(text)
}

func parse(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == "content" {
				hasfind = true
				parseTxt(&buf, n)
				break
			}
		}
	}
	if !hasfind {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parse(c)
		}
	}
}

func parseTxt(buf *bytes.Buffer, n *html.Node) {
	for c:= n.FirstChild; c != nil; c=c.NextSibling {
		if c.Data != "br" {
			buf.WriteString(c.Data + "\n")
		}
	}
}

func utf8(str string) string {
	decoder := mahonia.NewDecoder("gbk")
	s := decoder.ConvertString(str)
	return strings.Replace(s, "聽", " ", -1)
}