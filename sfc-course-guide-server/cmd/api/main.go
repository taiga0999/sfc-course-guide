package main

import (
	"fmt"
	// "github.com/LuckyWindsck/sfc-course-guide/sfc-course-guide-server/internal/app/sfccoursesummary"
	"github.com/LuckyWindsck/sfc-course-guide/sfc-course-guide-server/internal/pkg/elasticclient"
	"io/ioutil"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("./assets/vu_sfc_keio_ac_jp/course_u/data/2019/")
	if err != nil {
		fmt.Println(err)
	}

	// sfccoursesummary.Download()
	// elasticclient.DeleteIndex("sfc")
	// elasticclient.CreateIndex("sfc")
	elasticclient.CountDocument("sfc")
	// elasticclient.DeleteAllCourse()
	// sfccoursesummary.ExtractFromAspect("./assets/vu_sfc_keio_ac_jp/course_u/data/2019/csec14_8_en.html")

	for _, f := range files {
		if strings.Contains(f.Name(), "csec14_") {
			// sfccoursesummary.ExtractFromAspect("./assets/vu_sfc_keio_ac_jp/course_u/data/2019/" + f.Name())
		}
	}
}
