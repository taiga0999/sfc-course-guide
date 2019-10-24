package downloader

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/LuckyWindsck/sfc-course-guide/sfc-course-guide-server/internal/pkg/colorlog"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Download ...
type Download struct {
	Link *url.URL
	Path string
}

// createDownloadPath mkdir the directory for saving file and then return the path to save the file
func createDownloadPath(parsedLink *url.URL, downloadDirectory string) (path string, err error) {
	hostname := strings.Replace(parsedLink.Hostname(), ".", "_", -1)
	path = filepath.Join(downloadDirectory, hostname, parsedLink.Path)
	downloadDirectory = path[:strings.LastIndex(path, "/")]

	_, err = os.Stat(downloadDirectory)
	if os.IsNotExist(err) {
		colorlog.Warnln(downloadDirectory, "does *not* exist")
		colorlog.Infoln("creating", downloadDirectory, "...")
		err = os.MkdirAll(downloadDirectory, os.ModePerm)
		if err != nil {
			colorlog.Errorln("[os.MkdirAll error]", err)
			return
		}
		colorlog.Infoln("successfully created", downloadDirectory)
	} else if err != nil {
		// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go?answertab=votes#tab-top
		colorlog.Warnln("Schrodinger: file may or may not exist. See err for details.\n", err)
		return
	}

	return
}

// httpGet use http get request to obtain the html document and then return it
func httpGet(parsedLink *url.URL) (body []byte, err error) {
	resp, err := http.Get(parsedLink.String())
	if err != nil {
		colorlog.Errorln("[http.Get error]", err)
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		colorlog.Errorln("[ioutil.ReadAll error]", err)
		return
	}

	return
}

func detectCharsetFromMeta(r io.Reader) (htmlCharset string) {
	htmlCharset = "utf-8"

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		colorlog.Warnln("[goquery.NewDocumentFromReader error]", err)
		return
	}

	doc.Find("meta").Each(func(_ int, s *goquery.Selection) {
		val, exist := s.Attr("content")
		if exist {
			htmlCharset = regexp.MustCompile(`charset\s*=\s*(.*?)\s*;*`).Split(val, -1)[1]
		}

		val, exist = s.Attr("charset")
		if exist {
			htmlCharset = val
		}
	})

	return
}

func detectCharsetFromContent(r io.Reader) (htmlCharset string) {
	htmlCharset = "utf-8"

	data, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		colorlog.Warnln("[(*bufio.Reader).Peek error]", err)
		return
	}

	_, name, ok := charset.DetermineEncoding(data, "")
	if ok {
		htmlCharset = name
	}

	return
}

func detectCharset(r io.Reader) (htmlCharset string, err error) {
	htmlCharset = detectCharsetFromMeta(r)

	if htmlCharset != "utf-8" {
		return
	}

	htmlCharset = detectCharsetFromContent(r)
	return
}

// DecodeHTMLBody ...
func DecodeHTMLBody(parsedLink *url.URL, charset string) (body []byte, err error) {
	body, err = httpGet(parsedLink)
	if err != nil {
		colorlog.Errorln("[httpGet fail]", err)
		return
	}

	if charset == "" {
		// bytes.NewReader(body) will be read so cannot use a variable to replace it
		charset, err = detectCharset(bytes.NewReader(body))

		if err != nil {
			colorlog.Errorln("[detectCharset fail]", err)
			return
		}
	}

	encode, err := htmlindex.Get(charset)
	if err != nil {
		colorlog.Errorln("[htmlindex.Get error]", err)
		return
	}

	name, err := htmlindex.Name(encode)
	if err != nil {
		colorlog.Errorln("[htmlindex.Name error]", err)
		return
	}

	if name != "utf-8" {
		buf := new(bytes.Buffer)
		// [] byte => io.reader => decode => [] byte
		buf.ReadFrom(encode.NewDecoder().Reader(bytes.NewReader(body)))
		body = buf.Bytes()
	}

	return
}

// saveHTMLDocument
func saveHTMLDocument(path string, body []byte) (err error) {
	// 0644 : -rw-r--r-- (every one can read, only current user can write)
	err = ioutil.WriteFile(path, body, 0644)
	if err != nil {
		colorlog.Errorln("[ioutil.WriteFile error]", err)
		return
	}

	return
}

// DownloadFromLink download the html document of given url
func DownloadFromLink(parsedLink *url.URL, downloadDirectory string) (dl Download, err error) {
	dl.Link = parsedLink

	body, err := DecodeHTMLBody(parsedLink, "")
	if err != nil {
		colorlog.Errorln("[decodeHTMLBody fail]", err)
		return
	}

	dl.Path, err = createDownloadPath(parsedLink, downloadDirectory)
	if err != nil {
		colorlog.Errorln("[createDownloadPath fail]")
		return
	}

	colorlog.Infoln("downloading", parsedLink.String())

	err = saveHTMLDocument(dl.Path, body)
	if err != nil {
		colorlog.Errorln("[saveHTMLDocument failed]")
		return
	}

	return
}

func getURLs(path string) (parsedLinks []*url.URL, err error) {
	lines, err := ioutil.ReadFile(path)
	if err != nil {
		colorlog.Errorln("[ioutil.ReadFile error]", err)
		return
	}

	var links []string
	err = json.Unmarshal(lines, &links)
	if err != nil {
		colorlog.Errorln("[json.Unmarshal error]", err)
		return
	}

	for _, link := range links {
		parsedLink, err := url.Parse(link)
		if err != nil {
			colorlog.Errorln("[url.Parse error]", err)
		}
		parsedLinks = append(parsedLinks, parsedLink)
	}

	return
}

// DownloadFromFile reads the file named by filename, and download the html
// document according to the url in each line of the file
func DownloadFromFile(filename string, downloadDirectory string) (downloads []Download, err error) {
	links, err := getURLs(filename)
	if err != nil {
		colorlog.Errorln("[getURLs fail]")
		return
	}

	for _, link := range links {
		var dl Download
		dl, err = DownloadFromLink(link, downloadDirectory)
		if err != nil {
			colorlog.Errorln("[downloader.DownloadFromLink fail]")
			return
		}

		downloads = append(downloads, dl)
	}

	return
}
