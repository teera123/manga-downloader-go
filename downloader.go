package manga_downloader

import (
	"io/ioutil"
	"net/http"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/html"
)

type MangaPage struct {
	URL string
}

func (mp MangaPage) GetTitle() (string, error) {
	doc, err := parseHTML(mp.URL)
	if err != nil {
		return "", err
	}
	defer doc.Free()

	nodes, err := doc.Search("//title")
	if err != nil {
		return "", err
	}

	if len(nodes) == 0 {
		return "", nil
	}

	return nodes[0].String(), nil
}

func parseHTML(url string) (*html.HtmlDocument, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return gokogiri.ParseHtml(page)
}
