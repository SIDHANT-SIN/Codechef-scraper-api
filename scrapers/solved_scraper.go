package scraper

import (
	"Codechef-scraper-api/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeSolved(username string) (*models.UserSolved, error) {
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
	//var contests []models.UserSolved

	fullSolved := []string{}
    fullCount := doc.Find("section.problems-solved h5").First().Text()

    doc.Find("section.problems-solved .content ul").First().Find("li a").Each(func(i int, s *goquery.Selection) {
        fullSolved = append(fullSolved, strings.TrimSpace(s.Text()))
    })

    return &models.UserSolved{
        FullySolvedCount: fullCount,
        FullySolvedList:  fullSolved,
    }, nil




}

//  ContestHistory:  doc.Find(".rating-data-section.problems-solved h5").Text(),