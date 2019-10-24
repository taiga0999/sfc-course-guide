package url2json

import (
	"bytes"
	"encoding/json"
	"github.com/LuckyWindsck/sfc-course-guide/sfc-course-guide-server/internal/pkg/colorlog"
	"net/url"
)

// JSON ...
type JSON []byte

// URL2JSON ...
func URL2JSON(parsedLink *url.URL) (j JSON, err error) {
	j, err = URLs2JSON([]*url.URL{parsedLink})
	if err != nil {
		colorlog.Error(err)
		return
	}

	return
}

// URLs2JSON ...
func URLs2JSON(parsedLinks []*url.URL) (j JSON, err error) {
	var links []string
	for _, parsedLink := range parsedLinks {
		links = append(links, parsedLink.String())
	}

	j, err = json.Marshal(links)
	if err != nil {
		colorlog.Error(err)
		return
	}

	return
}

// Indent ...
func (j JSON) Indent() []byte {
	var out bytes.Buffer
	json.Indent(&out, j, "", "\t")
	return out.Bytes()
}

// JSON2str ...
func (j JSON) String() string {
	return string(j.Indent()[:])
}
