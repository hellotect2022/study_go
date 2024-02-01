package scrapper_goroutine

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"

type JobExtracted struct {
	id    string
	title string
	url   string
	date  string
	// city        string
	// location    string
	condition   string
	jobSelector string
}

func ScrapperMain() {
	var allJobs []JobExtracted
	pages := getPages()
	//fmt.Println(pages)
	for i := 0; i < pages; i++ {
		extractedJobs := getPage(i + 1)
		allJobs = append(allJobs, extractedJobs...)
	}
	writeJobs(allJobs)
	fmt.Println(allJobs)
}

func getPage(page int) []JobExtracted {
	var jobsForPage = []JobExtracted{}
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("request pag : " + pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	recruit := doc.Find(".item_recruit")
	//fmt.Println(recruit.Html())
	recruit.Each(func(i int, card *goquery.Selection) {
		id, _ := card.Attr("value")
		title := cleanString(card.Find(".job_tit").Find("span").Text())
		url, _ := card.Find(".job_tit").Find("a").Attr("href")
		date := cleanString(card.Find(".job_date").Find("span").Text())
		condition := cleanString(card.Find(".job_condition").Find("span").Text())
		jobSelector := cleanString(card.Find(".job_sector").Find("a").Text())
		jobExtracted := JobExtracted{id: id, title: title, url: "https://www.saramin.co.kr" + url, date: date, condition: condition, jobSelector: jobSelector}
		// fmt.Println(jobExtracted)
		// fmt.Println("++++")
		jobsForPage = append(jobsForPage, jobExtracted)
	})
	return jobsForPage
}

func writeJobs(jobs []JobExtracted) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "TITLE", "URL", "DATE", "CONDITION", "JOBSELECTOR"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSliece := []string{job.id, job.title, job.url, job.date, job.condition, job.jobSelector}
		wErr := w.Write(jobSliece)
		checkErr(wErr)
	}
}

func cleanString(str string) string {
	return strings.TrimSpace(str)
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
		//fmt.Println(s.Find("a").Length())
		// s.Find("a").Each(func(i int, s *goquery.Selection) {
		// 	res, _ := s.Html()
		// 	fmt.Println(i, res)
		// })
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Failed with StatusCode : ", res.StatusCode)
	}
}
