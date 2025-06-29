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

    res, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    if res.StatusCode != 200 {
        return nil, fmt.Errorf("failed to fetch user: %s", username)
    }

    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        return nil, err
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
