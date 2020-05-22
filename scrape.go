package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

type Recruit struct {
	Href    string
	Title   string
	EndAt   time.Time
	StartAt time.Time
}

type Recruits []Recruit

func (r Recruits) Len() int {
	return len(r)
}

func (r Recruits) Less(i, j int) bool {
	return r[j].StartAt.Before(r[i].StartAt)
}

func (r Recruits) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func PostToDoc(Url string, data url.Values) (doc *goquery.Document, err error) {
	resp, err := http.PostForm(Url, data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	return
}

func ParseRecruits(Url string, size int) (recruits []Recruit) {
	const pfmt = "2006-01-02"
	for i := 1; i <= size; i++ {
		doc, err := PostToDoc(Url, url.Values{
			"ar_eopjong_gbcd":   {"11111"},
			"eopjong_gbcd_list": {"11111"},
			"eopjong_gbcd":      {"1"},
			"pageUnit":          {"10"},
			"pageIndex":         {strconv.Itoa(i)},
		})
		if err != nil {
			log.Fatal(err)
		}

		doc.Find("table.brd_list_n tbody tr").Each(func(_ int, s *goquery.Selection) {
			var rec Recruit
			s.Find("td").Each(func(i int, s *goquery.Selection) {
				switch i {
				case 0:
					rec.Href, _ = s.Find("a").Attr("href")
					rec.Title = s.Find("a").Text()
				case 2:
					rec.EndAt, err = time.Parse(pfmt, s.Text())
					if err != nil {
						log.Fatal(err)
					}
				case 3:
					rec.StartAt, err = time.Parse(pfmt, s.Text())
					if err != nil {
						log.Fatal(err)
					}
				}
			})
			recruits = append(recruits, rec)
		})
	}
	return
}

func MMAScrape() (recruits []Recruit) {
	Url := "https://work.mma.go.kr/caisBYIS/search/cygonggogeomsaek.do"
	doc, err := PostToDoc(Url, url.Values{
		"ar_eopjong_gbcd":   {"11111"},
		"eopjong_gbcd_list": {"11111"},
		"eopjong_gbcd":      {"1"},
		"eopjong_cd":        {"11111"},
	})
	if err != nil {
		log.Fatal(err)
	}

	size, err := strconv.Atoi(doc.Find("div.page_move_n span").Last().Text())
	if err != nil {
		log.Fatal(err)
	}

	recruits = ParseRecruits(Url, size)
	sort.Sort(Recruits(recruits))
	return
}
