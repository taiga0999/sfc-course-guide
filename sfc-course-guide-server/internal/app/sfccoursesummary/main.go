package sfccoursesummary

import (
	"encoding/json"
	"github.com/LuckyWindsck/sfc-course-guide/sfc-course-guide-server/internal/pkg/colorlog"
	"github.com/LuckyWindsck/sfc-course-guide/sfc-course-guide-server/internal/pkg/downloader"
	"github.com/LuckyWindsck/sfc-course-guide/sfc-course-guide-server/internal/pkg/elasticclient"
	"github.com/LuckyWindsck/sfc-course-guide/sfc-course-guide-server/internal/pkg/url2json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	index = uint8(iota)
	aspect
	syllabus
	downloadDirectory = "./assets"
)

var seen = map[string]map[string]bool{
	"Course": make(map[string]bool),
	"Other":  make(map[string]bool),
}

func http2https(url string) string {
	return strings.Replace(url, "http://", "https://", -1)
}

func createJSONPath(path string) (err error) {
	directory := path[:strings.LastIndex(path, "/")]

	_, err = os.Stat(directory)
	if os.IsNotExist(err) {
		colorlog.Warnln(directory, "does *not* exist")
		colorlog.Infoln("creating", directory, "...")
		err = os.MkdirAll(directory, os.ModePerm)
		if err != nil {
			colorlog.Errorln("[os.MkdirAll error]", err)
			return
		}
		colorlog.Infoln("successfully created", directory)
	} else if err != nil {
		// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go?answertab=votes#tab-top
		colorlog.Warnln("Schrodinger: file may or may not exist. See err for details.\n", err)
		return
	}

	return
}

func saveJSON(path string, j url2json.JSON) (err error) {
	err = createJSONPath(path)
	if err != nil {
		colorlog.Errorln("[createDownloadPath fail]")
		return
	}

	// 0644 : -rw-r--r-- (every one can read, only current user can write)
	err = ioutil.WriteFile(path, j.Indent(), 0644)
	if err != nil {
		colorlog.Errorln("[ioutil.WriteFile error]", err)
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

func isCourseSummaryPage(parsedLink *url.URL) bool {
	return strings.Contains(parsedLink.String(), "https://vu.sfc.keio.ac.jp/course_u/")
}

func getCourseSummaryIndex(currentYear int, lang string) (indexURLs []*url.URL, err error) {
	langs := []string{"Japanese", "English"}
	if lang != "" {
		langs = []string{lang}
	}

	titleURL := "https://vu.sfc.keio.ac.jp/course_u/data/" + strconv.Itoa(currentYear) + "/title.html"

	parsedTitleURL, err := url.Parse(titleURL)
	if err != nil {
		colorlog.Errorln("[url.Parse error]", err)
		return
	}

	var doc *goquery.Document
	doc, err = goquery.NewDocument(titleURL)
	if err != nil {
		colorlog.Errorln("[goquery.NewDocument error]", err)
		return
	}

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		html, err := s.Html()
		if err != nil {
			colorlog.Errorln("[goquery.Selection.html error]", err)
			return
		}

		for _, l := range langs {
			if html == l {
				val, exist := s.Attr("href")
				if exist {
					parsedLink, err := parsedTitleURL.Parse(val)
					if err != nil {
						colorlog.Errorln("[url.Parse error]", err)
						return
					}
					indexURLs = append(indexURLs, parsedLink)
				}
			}
		}
	})

	return
}

func saveIndexJSON(path string) (err error) {
	currentYear := time.Now().Year()

	indexURLs, err := getCourseSummaryIndex(currentYear, "English")
	if err != nil {
		colorlog.Fatalln("[getLatestCourseSummaryIndex fail]")
		return
	}

	j, err := url2json.URLs2JSON(indexURLs)
	if err != nil {
		colorlog.Fatalln("[url2json.URLs2JSON fail]")
		return
	}

	err = saveJSON(path, j)
	if err != nil {
		colorlog.Fatalln("[saveJSON fail]")
		return
	}
	colorlog.Infoln("created", filepath.Base(path))

	return
}

func extractFromIndex(indexURL *url.URL, indexJSONPath string) (JSONPaths []string, err error) {
	f, err := os.Open(indexJSONPath)
	defer f.Close()
	if err != nil {
		colorlog.Errorln(err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		colorlog.Errorln("[goquery.NewDocumentFromReader error]", err)
		return
	}

	// Course Page By Academic degree
	// doc.Find("#toc > .np a").Each(func(_ int, s *goquery.Selection) {
	// 	val, exist := s.Attr("href")
	// 	if exist {
	// 		parsedLink, err := indexURL.Parse(http2https(val))
	// 		if err != nil {
	// 			colorlog.Errorln("[url.Parse error]", err)
	// 			return
	// 		}
	// 		colorlog.Debugln(parsedLink)
	// 	}
	// })

	// Other pages
	// doc.Find("#main a").Each(func(_ int, s *goquery.Selection) {
	// 	val, exist := s.Attr("href")
	// 	if exist {
	// 		parsedLink, err := indexURL.Parse(http2https(val))
	// 		if err != nil {
	// 			colorlog.Errorln("[url.Parse error]", err)
	// 			return
	// 		}
	// 		colorlog.Println(parsedLink)
	// 	}
	// })

	// Toc pages (toc from <div id="toc"/>)
	// 	PEMI
	// 	Perspectives
	// 	Course List by Aspect
	PMEI := doc.Find("h2 + ol").Eq(0)
	var PEMIURLs []*url.URL

	PMEI.Find("a").Each(func(_ int, s *goquery.Selection) {
		// name, err := s.Html()
		// if err != nil {
		// 	colorlog.Errorln("[goquery.Selection.html error]", err)
		// 	return
		// }
		// colorlog.Println(name)

		val, exist := s.Attr("href")
		if exist {
			parsedLink, err := indexURL.Parse(http2https(val))
			if err != nil {
				colorlog.Errorln("[url.Parse error]", err)
				return
			}
			PEMIURLs = append(PEMIURLs, parsedLink)
		}
	})

	j, err := url2json.URLs2JSON(PEMIURLs)
	if err != nil {
		colorlog.Fatalln("[url2json.URLs2JSON fail]")
		return
	}

	PEMIJSONPath := "PEMI.json"
	if strings.Contains(indexURL.String(), "en") {
		PEMIJSONPath = "PEMI_en.json"
	}

	PEMIJSONPath = filepath.Join(filepath.Dir(indexJSONPath), PEMIJSONPath)
	err = saveJSON(PEMIJSONPath, j)
	if err != nil {
		colorlog.Fatalln("[saveJSON fail]")
		return
	}
	colorlog.Infoln("created", PEMIJSONPath)
	JSONPaths = append(JSONPaths, PEMIJSONPath)

	return
}

func Download() {
	IndexJSONPath := "./assets/index.json"

	err := saveIndexJSON(IndexJSONPath)
	if err != nil {
		colorlog.Fatalln("[saveIndexJSON fail]")
		os.Exit(1)
	}

	downloads, err := downloader.DownloadFromFile(IndexJSONPath, downloadDirectory)
	if err != nil {
		colorlog.Fatalln("[downloader.DownloadFromFile fail]")
		os.Exit(1)
	}

	var PEMIJSONPaths []string

	for _, download := range downloads {
		JSONPaths, err := extractFromIndex(download.Link, download.Path)
		if err != nil {
			colorlog.Errorln("[extractFromIndex fail]")
			return
		}

		PEMIJSONPaths = append(PEMIJSONPaths, JSONPaths...)
	}

	for _, PEMIJSONPath := range PEMIJSONPaths {
		_, err := downloader.DownloadFromFile(PEMIJSONPath, downloadDirectory)
		if err != nil {
			colorlog.Fatalln("[downloader.DownloadFromFile fail]")
			os.Exit(1)
		}
	}
}

func ExtractFromAspect(path string) (JSONPaths []string, err error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		colorlog.Errorln(err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		colorlog.Errorln("[goquery.NewDocumentFromReader error]", err)
		return
	}

	var courses []elasticclient.Course

	doc.Find(".course").Each(func(_ int, s *goquery.Selection) {
		title_memo_en := s.Find(".course_title > .memo").Text()

		s.Find(".course_title").Children().Remove()
		course_title := s.Find(".course_title").Eq(0).Text()

		subject_sort := course_title[0:5]
		title_en := strings.Split(course_title[6:], "\n")[0]
		semester_en := strings.Trim(regexp.MustCompile(`[a-zA-Z0-9\s]*`).FindAllString(strings.Split(course_title[6:], "\n")[1], -1)[1], " ")

		var faculty_in_charge []string
		var days_en []string
		var language_en string

		s.Find("table tr").Each(func(_ int, tr *goquery.Selection) {
			tag := tr.Find(".tag").Text()

			if strings.Contains(tag, "Faculty-in-charge") {
				tr.Find(".value > a").Each(func(_ int, a *goquery.Selection) {
					faculty_in_charge = append(faculty_in_charge, a.Text())
				})
			}

			if strings.Contains(tag, "Semester, Day and Period") {
				for _, char := range regexp.MustCompile(`[,()]`).Split(strings.Split(tr.Children().Eq(1).Text(), "\n")[0], -1) {
					if char != "" {
						days_en = append(days_en, char)
					}
				}
			}

			if strings.Contains(tag, "Language used in the class") {
				language_en = strings.Split(tr.Children().Eq(1).Text(), "\n")[0]
			}
		})

		link := strings.Replace(path, "./assets/vu_sfc_keio_ac_jp/", "https://vu.sfc.keio.ac.jp/", -1) + "#ks" + subject_sort

		s.Find(".course_title").Remove()
		s.Find("table").Remove()
		description := s.Text()

		courses = append(courses, elasticclient.Course{
			Link:              link,
			Subject_sort:      subject_sort,
			Title_en:          title_en,
			Title_memo_en:     title_memo_en,
			Faculty_in_charge: faculty_in_charge,
			Semester_en:       semester_en,
			Days_en:           days_en,
			Language_en:       language_en,
			Description:       description,
		})
	})

	colorlog.Debugln(len(courses))

	for _, course := range courses {
		elasticclient.CreateCourse(course)
	}

	// Get all urls from main div
	// Remove urls link to same page

	// Get all professor pages url

	// Get all syllabus pages url

	// Save those urls to json file
	return
}
