package render

import (
	"bytes"
	"encoding/base64"
	"mime"
	"net/url"
	"path"

	"golang.org/x/net/html/atom"

	"golang.org/x/net/html"
)

func findAttributeIndex(attrs []html.Attribute, key string) int {
	index := -1
	for i := 0; i < len(attrs); i++ {
		if attrs[i].Key == key {
			index = i
			break
		}
	}
	return index
}

func embedStyleNode(node *html.Node, baseURL *url.URL, site hugoSite) error {
	hrefAttrIndex := findAttributeIndex(node.Attr, "href")
	if hrefAttrIndex == -1 {
		return nil
	}
	hrefAttr := node.Attr[hrefAttrIndex]
	contentURL, err := baseURL.Parse(hrefAttr.Val)
	if err != nil {
		return err
	}
	if contentURL.Host != baseURL.Host {
		return nil
	}

	content, err := site.readPublic(contentURL)
	if err != nil {
		return err
	}
	contentNode := &html.Node{}
	contentNode.Type = html.TextNode
	contentNode.Data = string(content)

	styleNode := &html.Node{}
	styleNode.Type = html.ElementNode
	styleNode.DataAtom = atom.Style
	styleNode.Data = "style"
	if index := findAttributeIndex(node.Attr, "media"); index != -1 {
		attr := styleNode.Attr[index]
		styleNode.Attr = append(styleNode.Attr, attr)
	}
	if index := findAttributeIndex(node.Attr, "title"); index != -1 {
		attr := styleNode.Attr[index]
		styleNode.Attr = append(styleNode.Attr, attr)
	}

	styleNode.AppendChild(contentNode)
	node.Parent.InsertBefore(styleNode, node)
	node.Parent.RemoveChild(node)
	return nil
}

func embedScriptNode(node *html.Node, baseURL *url.URL, site hugoSite) error {
	srcAttrIndex := findAttributeIndex(node.Attr, "src")
	if srcAttrIndex == -1 {
		return nil
	}
	srcAttr := node.Attr[srcAttrIndex]
	contentURL, err := baseURL.Parse(srcAttr.Val)
	if err != nil {
		return err
	}
	if contentURL.Host != baseURL.Host {
		return nil
	}

	content, err := site.readPublic(contentURL)
	if err != nil {
		return err
	}
	contentNode := &html.Node{}
	contentNode.Type = html.TextNode
	contentNode.Data = string(content)

	node.Attr = append(node.Attr[:srcAttrIndex], node.Attr[srcAttrIndex+1:]...)
	node.AppendChild(contentNode)
	return nil
}

func embedImgNode(node *html.Node, baseURL *url.URL, site hugoSite) error {
	srcAttrIndex := findAttributeIndex(node.Attr, "src")
	if srcAttrIndex == -1 {
		return nil
	}
	srcAttr := node.Attr[srcAttrIndex]
	imageURL, err := baseURL.Parse(srcAttr.Val)
	if err != nil {
		return err
	}
	if imageURL.Host != baseURL.Host {
		return nil
	}

	imageExt := path.Ext(imageURL.Path)
	imageMime := mime.TypeByExtension(imageExt)

	imageData, err := site.readPublic(imageURL)
	if err != nil {
		return err
	}
	imageBase64 := base64.StdEncoding.EncodeToString(imageData)

	node.Attr[srcAttrIndex].Val = "data:" + imageMime + ";base64," + imageBase64
	return nil
}

func embedNode(node *html.Node, baseURL *url.URL, site hugoSite) error {
	var err error
	switch node.DataAtom {
	case atom.Link:
		relIndex := findAttributeIndex(node.Attr, "rel")
		if relIndex == -1 {
			break
		}
		relValue := node.Attr[relIndex].Val

		switch relValue {
		case "stylesheet":
			err = embedStyleNode(node, baseURL, site)
		}

	case atom.Script:
		err = embedScriptNode(node, baseURL, site)

	case atom.Img:
		err = embedImgNode(node, baseURL, site)
	}

	if err != nil {
		return err
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if err := embedNode(child, baseURL, site); err != nil {
			return err
		}
	}
	return nil
}

func embedHTML(pageURL *url.URL, site hugoSite) ([]byte, error) {
	rawHTML, err := site.readPublic(pageURL)
	if err != nil {
		return nil, err
	}

	node, err := html.Parse(bytes.NewReader(rawHTML))
	if err != nil {
		return nil, err
	}

	if err := embedNode(node, pageURL, site); err != nil {
		return nil, err
	}

	embeddedHTML := &bytes.Buffer{}
	if err := html.Render(embeddedHTML, node); err != nil {
		return nil, err
	}

	return embeddedHTML.Bytes(), nil
}
