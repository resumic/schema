package render

import (
	"bytes"
	"encoding/base64"
	"mime"
	"net/url"
	"path"

	"golang.org/x/net/html"
)

func makeDataURL(resourceURL *url.URL, site hugoSite) (string, error) {
	resourceContent, err := site.readPublic(resourceURL)
	if err != nil {
		return "", err
	}
	resourceBase64 := base64.StdEncoding.EncodeToString(resourceContent)
	resourceExt := path.Ext(resourceURL.Path)
	resourceMime := mime.TypeByExtension(resourceExt)
	return "data:" + resourceMime + ";base64," + resourceBase64, nil
}

func embedHTMLNode(node *html.Node, baseURL *url.URL, site hugoSite) error {
	if node.Type == html.ElementNode {
		for index, attr := range node.Attr {
			if attr.Key == "href" || attr.Key == "src" {
				resourceURL, err := baseURL.Parse(attr.Val)
				if err != nil {
					return err
				}
				if baseURL.Host == resourceURL.Host {
					dataURL, err := makeDataURL(resourceURL, site)
					if err != nil {
						return err
					}
					node.Attr[index].Val = dataURL
				}
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if err := embedHTMLNode(child, baseURL, site); err != nil {
			return err
		}
	}

	return nil
}

func embedHTML(pageHTML []byte, pageURL *url.URL, site hugoSite) ([]byte, error) {
	node, err := html.Parse(bytes.NewReader(pageHTML))
	if err != nil {
		return nil, err
	}

	if err := embedHTMLNode(node, pageURL, site); err != nil {
		return nil, err
	}

	embeddedHTML := &bytes.Buffer{}
	if err := html.Render(embeddedHTML, node); err != nil {
		return nil, err
	}

	return embeddedHTML.Bytes(), nil
}
