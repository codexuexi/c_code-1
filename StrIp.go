package c_code

import (
	"bytes"
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

//Nodes structure with html.Node elements
type Nodes struct {
	Elements *html.Node
}

//Strip HTML tags from a string. This function allows you to provide an array of allowable tags which will be skipped
//from removing. Also, you can strip the HTML tag attributes (e.g. style, class, id ...)
func Strip(content string, allowableTags []string, stripInlineAttributes bool) (Nodes, error) {
	document, err := toNodes(content)
	handleError(err)
	var nodeTree html.Node

	var output func(document *html.Node, nt *html.Node)
	output = func(document *html.Node, nt *html.Node) {
		for c := document.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode || (c.Type == html.ElementNode && inArray(c.Data, allowableTags)) {
				var childNode html.Node
				childNode.Type = c.Type
				childNode.Data = c.Data
				if stripInlineAttributes == true {
					childNode.Attr = []html.Attribute{}
				} else {
					childNode.Attr = c.Attr
				}
				nt.AppendChild(&childNode)
				output(c, nt.LastChild)
			} else {
				output(c, nt)
			}
		}
	}
	output(document, &nodeTree)
	return Nodes{Elements: &nodeTree}, nil
}

func StripRemoveAttr(content string) string {
	allowableTags := []string{"img", "stop_node_tag"}
	document, err := toNodes(content)
	handleError(err)
	var nodeTree html.Node

	var output func(document *html.Node, nt *html.Node)
	output = func(document *html.Node, nt *html.Node) {
		for c := document.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode || c.Type == html.ElementNode {
				var childNode html.Node
				childNode.Type = c.Type
				childNode.Data = c.Data
				if inArray(c.Data, allowableTags) {
					childNode.Attr = c.Attr
				} else {
					childNode.Attr = []html.Attribute{}
				}

				nt.AppendChild(&childNode)
				output(c, nt.LastChild)
			} else {
				output(c, nt)
			}
		}
	}
	output(document, &nodeTree)

	nodes := Nodes{Elements: &nodeTree}

	findString := regexp.MustCompile(`<body>(.*?)</body>`).FindStringSubmatch(nodes.ToString())
	if len(findString) >= 2 {
		return findString[1]
	}
	return content
}

//func HtmlRemoveAttr(content string) (string) {
//	document, err := toNodes(content)
//	handleError(err)
//	var nodeTree html.Node
//
//	var output func(document *html.Node, nt *html.Node)
//	output = func(document *html.Node, nt *html.Node) {
//		for c := document.FirstChild; c != nil; c = c.NextSibling {
//			if c.Type == html.TextNode {
//				var childNode html.Node
//				childNode.Type = c.Type
//				childNode.Data = c.Data
//				childNode.Attr = []html.Attribute{}
//				nt.AppendChild(&childNode)
//				output(c, nt.LastChild)
//			} else {
//				output(c, nt)
//			}
//		}
//	}
//	output(document, &nodeTree)
//	toString := Nodes{Elements: &nodeTree}.ToString()
//	return toString
//}

//String to nodes helper.
func toNodes(document string) (*html.Node, error) {
	nodes, err := html.Parse(strings.NewReader(html.UnescapeString(document)))
	handleError(err)
	return nodes, nil
}

//ToString is a Nodes method. Converts Nodes.Elements to string
func (nodes *Nodes) ToString() string {
	var buf bytes.Buffer
	for n := nodes.Elements.FirstChild; n != nil; n = n.NextSibling {
		html.Render(&buf, n)
	}
	return html.UnescapeString(buf.String())
}

//Check if needle is in the array
func inArray(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//Show error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
