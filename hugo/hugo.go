package hugo

import (
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
)

const RESUME_CODE_LENGTH = 50

func extractSite(root string) error {
	box := packr.New("site", "./hugo/site")
	return box.Walk(func(name string, file packd.File) error {
		p := path.Join(root, name)
		os.MkdirAll(path.Dir(p), 0777)
		dst, err := os.Create(p)
		defer dst.Close()
		if err != nil {
			return err
		}
		_, err = io.Copy(dst, file)
		return err
	})
}

type Hugo struct {
	root string
}

func New(root string) (*Hugo, error) {
	err := extractSite(root)
	if err != nil {
		return nil, err
	}
	return &Hugo{root: root}, nil
}

func (h *Hugo) resumeContentPath(code string) string {
	return path.Join(h.root, "content", "resume", code+".md")
}

func (h *Hugo) resumeJSONPath(code string) string {
	return path.Join(h.root, "data", "resume", code+".json")
}

func (h *Hugo) resumeExist(code string) bool {
	p := h.resumeContentPath(code)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return false
	}
	return true
}

func (h *Hugo) NewResume(resume []byte) (string, error) {
	code := randomString(RESUME_CODE_LENGTH)
	for h.resumeExist(code) {
		code = randomString(RESUME_CODE_LENGTH)
	}
	err := ioutil.WriteFile(h.resumeContentPath(code), []byte{}, 0600)
	if err != nil {
		os.Remove(h.resumeContentPath(code))
		return "", err
	}
	err = ioutil.WriteFile(h.resumeJSONPath(code), resume, 0600)
	if err != nil {
		os.Remove(h.resumeJSONPath(code))
		return "", err
	}
	return code, nil
}
