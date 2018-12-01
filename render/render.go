package render

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/gohugoio/hugo/commands"
)

type siteConfig struct {
	DisableKinds []string `json:"disableKinds"`
	Theme        string   `json:"theme"`
}

type frontmatter struct {
	Layout string `json:"layout"`
}

func build(root string) error {
	resp := commands.Execute([]string{"--quiet", "-s", root})
	return resp.Err
}

func RenderHTML(resume []byte, theme string) ([]byte, error) {
	sitePath, err := ioutil.TempDir(os.TempDir(), "resumic")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(sitePath)

	themePath := path.Join(sitePath, "themes")
	err = extractTheme(themePath, theme)
	if err != nil {
		return nil, err
	}

	config := siteConfig{}
	config.Theme = theme
	config.DisableKinds = []string{"taxonomy", "taxonomyTerm", "category", "sitemap", "RSS", "404", "robotsTXT", "home", "section"}

	configJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return nil, err
	}
	configPath := path.Join(sitePath, "config.json")
	err = ioutil.WriteFile(configPath, configJSON, 0600)
	if err != nil {
		return nil, err
	}

	dataPath := path.Join(sitePath, "data", "resumic", "resume.json")
	err = os.MkdirAll(path.Dir(dataPath), 0700)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(dataPath, resume, 0600)
	if err != nil {
		return nil, err
	}

	content := frontmatter{}
	content.Layout = "resumic"

	contentJSON, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return nil, err
	}
	contentPath := path.Join(sitePath, "content", "resumic", "resume.md")
	err = os.MkdirAll(path.Dir(contentPath), 0700)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(contentPath, contentJSON, 0600)
	if err != nil {
		return nil, err
	}

	err = build(sitePath)
	if err != nil {
		return nil, err
	}

	htmlPath := path.Join(sitePath, "public", "resumic", "resume", "index.html")
	return ioutil.ReadFile(htmlPath)
}
