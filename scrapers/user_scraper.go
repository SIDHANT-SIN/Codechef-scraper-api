package scraper

import (
	"Codechef-scraper-api/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeUser(username string) (*models.UserData, error) {
url := fmt.Sprintf("https://www.codechef.com/users/%s", username)

req, err := http.NewRequest("GET", url, nil)
if err != nil {
	return nil, fmt.Errorf("failed to create HTTP request: %w", err)
}
req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
	"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0 Safari/537.36")
req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
req.Header.Set("Accept-Language", "en-US,en;q=0.9")

client := &http.Client{}
res, err := client.Do(req)
if err != nil {
	return nil, fmt.Errorf("failed to make HTTP request to %s: %w", url, err)
}
defer res.Body.Close()

if res.StatusCode != http.StatusOK {
	return nil, fmt.Errorf("received non-OK status code %d from %s", res.StatusCode, url)
}

doc, err := goquery.NewDocumentFromReader(res.Body)
if err != nil {
	return nil, fmt.Errorf("failed to parse HTML document: %w", err)
}

	raw := doc.Find(".rating-header small").Text()

	start := strings.LastIndex(raw, " ") + 1
	end := strings.LastIndex(raw, ")")
	
	highestRating := ""
	if start > 0 && end > start {
		highestRating = raw[start:end]
	}

    raw_name := doc.Find(".content .breadcrumbs .breadcrumb").Text();
	parts_name := strings.Split(raw_name, "»");
	test_name := strings.TrimSpace(parts_name[len(parts_name)-1])
   // fmt.Println("Item");
	// doc.Find(".rating-ranks li").Each(func(i int, s *goquery.Selection) {
	// 	text := strings.TrimSpace(s.Text())
	// 	fmt.Println("Item", i, ":", text)
	// })

	institution := doc.Find("div.user-details-container.plr10 section.user-details ul.side-nav li:contains('Institution:') span").Text()
	
	//fmt.Println( institution)

    user := &models.UserData{
        Username: username,
        Name:     test_name,
        Stars:    doc.Find(".rating-star").Text(),
        Rating:   doc.Find(".rating-number").Text(),
        HighestRating: highestRating,
		

        GlobalRank:   doc.Find(".rating-ranks ul.inline-list li:contains('Global Rank') strong").Text(),
        CountryRank:    doc.Find(".rating-ranks ul.inline-list li:contains('Country Rank') strong").Text(),
        Country:   doc.Find(".user-details-container .user-country-name").Text(),
          
        Institution:  institution,
    }

    return user, nil
}
