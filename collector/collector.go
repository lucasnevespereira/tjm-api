package collector

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
)

// getInfo searches Job Position, Salary (TJM) and check for an err bool
// prepares scrapdata from freelance-info.fr
func getInfo(e *colly.HTMLElement) (string, int, bool) {
	var title string
	var salary int
	var err bool

	// Remove certain unwanted fields
	// MatchString reports if e.Text contains any match of the regular expression pattern.
	match, _ := regexp.MatchString("Freelance|Réflexion|Travail|Formation|Typologie|astreinte", e.Text)
	if !match {
		// to understand this regexpression go to (https://regexr.com/) :)
		var re = regexp.MustCompile(`\s*(Maj|Lu)?\s*(.*)\s+(\d*) €\/j`)
		title = re.ReplaceAllString(e.Text, `$2`)
		sal := re.ReplaceAllString(e.Text, `$3`)
		// converts to int
		salary, _ = strconv.Atoi(sal)
		if salary == 0 {
			err = true
		}
	} else {
		err = true
	}

	fmt.Println("here")
	fmt.Println(title, salary, err)

	return title, salary, err
}

// Start func inits a gocolly collector
func Start() {

	info := map[string]int{}

	// Inits a collector
	c := colly.NewCollector(
		// Visit only this domains
		colly.AllowedDomains("freelance-info.fr", "www.freelance-info.fr"),
	)

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Runned into some issue :", err)
	})

	// On every li element with class vi call callback
	c.OnHTML("li.vi", func(e *colly.HTMLElement) {
		t, s, err := getInfo(e)
		if !err {
			info[t] = s
		}
	})

	// Before making a request prints requested website
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Requesting this website : ", r.URL.String())
	})

	// Action after scrapping is done
	c.OnScraped(func(s *colly.Response) {
		jsonData, _ := json.Marshal(info)
		fmt.Printf("%s", jsonData)
	})

	// Visit starts the actual scrapping request by calling all the previous funcs
	c.Visit("https://www.freelance-info.fr/tarifs/")
}
