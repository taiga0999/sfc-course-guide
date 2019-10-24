package elasticclient

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"reflect"
)

type Course struct { // Example Course
	Link              string
	Subject_sort      string   // B6117
	Title_en          string   // FUNDAMENTALS OF LOGIC
	Title_memo_en     string   // (GIGA/GG/GI)
	Faculty_in_charge []string // [Tatsuya Hagino]
	Semester_en       string   // Spring
	Days_en           []string // Tuesday 1st Period
	Language_en       string   // English
	Description       string
}

// Search Result
type Hit interface {
}

type CourseHit struct {
	Course    Course
	Highlight map[string][]string `json:",omitempty"`
}

type HitStat struct {
	Total   int64
	Latency int64 // Unit: milliseconds
}

type ClientSearchResult struct {
	Query string `json:",omitempty"`
	Stat  HitStat
	Hits  []Hit
}

func initClient() (client *elastic.Client) {
	client, err := elastic.NewClient()
	if err != nil {
		fmt.Println(err)
		fmt.Println("connect es error")
	}

	return
}

var client = initClient()
var ctx = context.Background()
var matchAllQuery = elastic.NewMatchAllQuery()
var course Course

func CreateIndex(index string) {
	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	if exists {
		return
	}

	createIndex, err := client.CreateIndex(index).Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !createIndex.Acknowledged {
		fmt.Println("Not acknowledged")
	}

	fmt.Printf("Successfully create index: %s\n", index)
}

func DeleteIndex(index string) {
	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !exists {
		return
	}

	deleteIndex, err := client.DeleteIndex(index).Do(ctx)
	if err != nil {
		fmt.Println(err)
	}
	if !deleteIndex.Acknowledged {
		fmt.Println("Not acknowledged")
	}

	fmt.Printf("Successfully delete index: %s\n", index)
}

func CountDocument(index string) (count int64, err error) {
	count, err = client.Count(index).Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Printf("Total number of documents in index %s: %d\n", index, count)
	return
}

func CreateCourse(course Course) (err error) {
	_, err = client.Index().
		Index("sfc").
		Type("course").
		BodyJson(course).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func GetAllCourse() (clientSearchResult ClientSearchResult, err error) {
	count, err := CountDocument("sfc")
	if err != nil {
		fmt.Println(err)
		return
	}

	searchResult, err := client.Search("sfc").
		Type("course").
		Sort("Subject_sort.keyword", true).
		From(0).Size(int(count)).
		// Pretty(true).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	clientSearchResult.Stat.Latency = searchResult.TookInMillis
	clientSearchResult.Stat.Total = searchResult.TotalHits()

	for _, item := range searchResult.Each(reflect.TypeOf(course)) {
		c, ok := item.(Course)
		courseHit := CourseHit{Course: c}
		if ok {
			clientSearchResult.Hits = append(clientSearchResult.Hits, courseHit)
		}
	}

	return
}

func DeleteAllCourse() {
	_, err := client.DeleteByQuery("sfc").Query(matchAllQuery).Do(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully delete all course in sfc")
}

func SearchCourse(query string) (clientSearchResult ClientSearchResult, err error) {
	count, err := CountDocument("sfc")
	if err != nil {
		fmt.Println(err)
		return
	}

	multiMatchQuery := elastic.
		NewMultiMatchQuery(
			query,
			"Subject_sort",
			"Title_en",
			"Title_memo_en",
			"Faculty_in_charge",
			"Semester_en",
			"Days_en",
			"Language_en").
		Type("cross_fields").
		Operator("And")

	// highLight := elastic.NewHighlight().
	// 	Field("Subject_sort").
	// 	Field("Title_en").
	// 	Field("Title_memo_en").
	// 	Field("Faculty_in_charge").
	// 	Field("Semester_en").
	// 	Field("Days_en").
	// 	Field("Language_en").
	// 	PreTags("<highlight>").
	// 	PostTags("</highlight>")

	searchResult, err := client.Search("sfc").
		Type("course").
		Query(multiMatchQuery).
		From(0).Size(int(count)).
		// Pretty(true).
		// Highlight(highLight).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
	}

	clientSearchResult.Query = query
	clientSearchResult.Stat.Latency = searchResult.TookInMillis
	clientSearchResult.Stat.Total = searchResult.TotalHits()

	for i, item := range searchResult.Each(reflect.TypeOf(course)) {
		c, ok := item.(Course)
		courseHit := CourseHit{Course: c, Highlight: searchResult.Hits.Hits[i].Highlight}
		if ok {
			clientSearchResult.Hits = append(clientSearchResult.Hits, courseHit)
		}
	}

	return
}
