package libs

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Recruit struct {
	href    string
	title   string
	endAt   time.Time
	startAt time.Time
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
					rec.href, _ = s.Find("a").Attr("href")
					rec.title = s.Find("a").Text()
				case 2:
					rec.endAt, err = time.Parse(pfmt, s.Text())
					if err != nil {
						log.Fatal(err)
					}
				case 3:
					rec.startAt, err = time.Parse(pfmt, s.Text())
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
	return
}
