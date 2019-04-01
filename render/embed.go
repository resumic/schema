package render

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

func makeDataURL(resourceURL string, client *http.Client) (string, error) {
	response, err := client.Get(resourceURL)
	if err != nil {
		return "", err
	}

	contentType := response.Header.Get("Content-type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	contentBase64 := base64.StdEncoding.EncodeToString(content)

	dataURL := "data:" + contentType + ";base64," + contentBase64
	return dataURL, nil
}

func embedHTMLNode(node *html.Node, baseURL string, client *http.Client) error {
	if node.Type == html.ElementNode {
		for index, attr := range node.Attr {
			if attr.Key == "href" || attr.Key == "src" {
				resourceURL, err := absURL(baseURL, attr.Val)
				if err != nil {
					return nil
				}
				dataURL, err := makeDataURL(resourceURL, client)
				if err != nil {
					return err
				}
				node.Attr[index].Val = dataURL
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if err := embedHTMLNode(child, baseURL, client); err != nil {
			return err
		}
	}

	return nil
}

func embedHTML(pageHTML []byte, pageURL string, client *http.Client) ([]byte, error) {
	node, err := html.Parse(bytes.NewReader(pageHTML))
	if err != nil {
		return nil, err
	}

	if err := embedHTMLNode(node, pageURL, client); err != nil {
		return nil, err
	}

	embeddedHTML := &bytes.Buffer{}
	if err := html.Render(embeddedHTML, node); err != nil {
		return nil, err
	}

	return embeddedHTML.Bytes(), nil
}
