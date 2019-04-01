package render

import (
	"io/ioutil"
	"os"

	"github.com/resumic/schema/schema"
)

func RenderHTML(resumeJSON []byte, themeDir string) ([]byte, error) {
	if err := schema.ValidateResume(resumeJSON); err != nil {
		return nil, err
	}

	siteDir, err := ioutil.TempDir(os.TempDir(), "resumic")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(siteDir)
	site, err := initHugoSite(siteDir)
	if err != nil {
		return nil, err
	}

	resumeName := "resume"
	if err := site.writeResumeJSON(resumeJSON, resumeName); err != nil {
		return nil, err
	}
	resumeURL, err := site.getResumeURL(resumeName)
	if err != nil {
		return nil, err
	}

	if err := site.build(themeDir); err != nil {
		return nil, err
	}

	client, err := newClient(site)
	if err != nil {
		return nil, err
	}
	response, err := client.Get(resumeURL)
	if err != nil {
		return nil, err
	}
	resumeHTML, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return embedHTML(resumeHTML, resumeURL, client)
}
