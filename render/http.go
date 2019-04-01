package render

import (
	"net/http"
	"net/url"
	"path"
)

func absURL(baseURL, resourceURL string) (string, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	resource, err := base.Parse(resourceURL)
	if err != nil {
		return "", err
	}
	return resource.String(), nil
}

type transport struct {
	siteURL         *url.URL
	localTransport  http.RoundTripper
	siteTransport   http.RoundTripper
	remoteTransport http.RoundTripper
}

func newTransport(site hugoSite) (*transport, error) {
	t := &transport{}

	siteURL, err := url.Parse(site.config.BaseURL)
	if err != nil {
		return nil, err
	}
	t.siteURL = siteURL

	t.localTransport = http.NewFileTransport(http.Dir("/"))
	publicDir := path.Join(site.dir, "public")
	t.siteTransport = http.NewFileTransport(http.Dir(publicDir))
	t.remoteTransport = &http.Transport{}

	return t, nil
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqURL := req.URL
	if reqURL.Scheme == "file" {
		return t.localTransport.RoundTrip(req)
	}
	if reqURL.Host == "" || reqURL.Host == t.siteURL.Host {
		return t.siteTransport.RoundTrip(req)
	}
	return t.remoteTransport.RoundTrip(req)
}

func newClient(site hugoSite) (*http.Client, error) {
	trans, err := newTransport(site)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Transport: trans,
	}
	return client, nil
}
