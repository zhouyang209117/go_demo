package links

import (
	"golang.org/x/net/html"
	"net/http"
)

func Extract(url string) ([]string, error)  {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	node, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode || n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(attr.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(node, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node))  {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}