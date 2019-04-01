package render

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"os"
	"path"

	"github.com/gohugoio/hugo/commands"
)

type CouldNotReadPublicError struct {
	url *url.URL
	dir string
}

func (e *CouldNotReadPublicError) Error() string {
	return "Could not find " + e.url.String() + " in " + e.dir
}

type contentFrontmatter struct {
	Layout string `json:"layout"`
}

type hugoConfig struct {
	DisableKinds []string `json:"disableKinds"`
	Theme        string   `json:"theme"`
	BaseURL      string   `json:"baseURL"`
	Title        string   `json:"title"`
}

func newHugoConfig() hugoConfig {
	config := hugoConfig{}
	config.DisableKinds = []string{"taxonomy", "taxonomyTerm", "category", "sitemap", "RSS", "404", "robotsTXT", "home", "section"}
	config.BaseURL = "http://example.org/"
	config.Title = "resumic"
	return config
}

type hugoSite struct {
	dir    string
	config hugoConfig
}

func initHugoSite(siteDir string) (hugoSite, error) {
	site := hugoSite{
		dir: siteDir,
	}
	if err := os.MkdirAll(site.dir, 0700); err != nil {
		return site, err
	}

	site.config = newHugoConfig()
	configJSON, err := json.MarshalIndent(site.config, "", "  ")
	if err != nil {
		return site, err
	}
	configPath := path.Join(site.dir, "config.json")
	if err := ioutil.WriteFile(configPath, configJSON, 0600); err != nil {
		return site, err
	}

	return site, nil
}

func (s hugoSite) writeResumeJSON(resumeJSON []byte, resumeName string) error {
	dataDir := path.Join(s.dir, "data", "resumic")
	if err := os.MkdirAll(dataDir, 0700); err != nil {
		return err
	}
	dataPath := path.Join(dataDir, resumeName+".json")
	if err := ioutil.WriteFile(dataPath, resumeJSON, 0600); err != nil {
		return err
	}

	content := contentFrontmatter{}
	content.Layout = "resumic"
	contentJSON, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return err
	}
	contentDir := path.Join(s.dir, "content", "resumic")
	if err := os.MkdirAll(contentDir, 0700); err != nil {
		return err
	}
	contentPath := path.Join(contentDir, resumeName+".md")
	return ioutil.WriteFile(contentPath, contentJSON, 0600)
}

func (s hugoSite) getResumeURL(resumeName string) (string, error) {
	return absURL(s.config.BaseURL, "/resumic/"+resumeName)
}

func (s hugoSite) build(themeDir string) error {
	resp := commands.Execute([]string{"--quiet", "--source", s.dir, "--themesDir", themeDir})
	return resp.Err
}
