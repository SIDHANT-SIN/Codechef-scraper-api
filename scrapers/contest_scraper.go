package scraper

import (
    "Codechef-scraper-api/models"
    "fmt"
    "net/http"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

func ScrapeContests(username string) ([]models.UserContest, error) {
    url := fmt.Sprintf("https://www.codechef.com/users/%s", username)

    res, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    if res.StatusCode != 200 {
        return nil, fmt.Errorf("status code: %d", res.StatusCode)
    }

    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        return nil, err
    }

    contests := []models.UserContest{}

    doc.Find("table.rating-table tbody tr").Each(func(i int, row *goquery.Selection) {
        tds := row.Find("td")
        if tds.Length() >= 3 {
            contests = append(contests, models.UserContest{
                ContestCode: strings.TrimSpace(tds.Eq(0).Text()),
                ContestName: strings.TrimSpace(tds.Eq(1).Text()),
                Rank:        strings.TrimSpace(tds.Eq(2).Text()),
            })
        }
    })

    return contests, nil
}
