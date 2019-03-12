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

func hugoInitSite(siteDir string) error {
	config := siteConfig{}
	config.DisableKinds = []string{"taxonomy", "taxonomyTerm", "category", "sitemap", "RSS", "404", "robotsTXT", "home", "section"}
	configJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	configPath := path.Join(siteDir, "config.json")
	return ioutil.WriteFile(configPath, configJSON, 0600)
}

func hugoWriteResumeJSON(resumeJSON []byte, resumeName, siteDir string) error {
	dataPath := path.Join(siteDir, "data", "resumic", resumeName+".json")
	if err := os.MkdirAll(path.Dir(dataPath), 0700); err != nil {
		return err
	}
	if err := ioutil.WriteFile(dataPath, resumeJSON, 0600); err != nil {
		return err
	}

	content := frontmatter{}
	content.Layout = "resumic"
	contentJSON, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return err
	}
	contentPath := path.Join(siteDir, "content", "resumic", resumeName+".md")
	if err := os.MkdirAll(path.Dir(contentPath), 0700); err != nil {
		return err
	}
	return ioutil.WriteFile(contentPath, contentJSON, 0600)
}

func hugoBuild(themeDir, siteDir string) error {
	resp := commands.Execute([]string{"--quiet", "--themesDir", themeDir, "--source", siteDir})
	return resp.Err
}

func hugoReadResumeHTML(name, siteDir string) ([]byte, error) {
	htmlPath := path.Join(siteDir, "public", "resumic", name, "index.html")
	return ioutil.ReadFile(htmlPath)
}
