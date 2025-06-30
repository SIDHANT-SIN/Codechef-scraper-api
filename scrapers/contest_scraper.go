package scraper

import (
	"Codechef-scraper-api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeContests(username string) ([]models.Rating, error) {
	url := fmt.Sprintf("https://www.codechef.com/users/%s", username)

	res, err := http.Get(url)
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

	var allRatings []models.Rating
	foundAndParsed := false

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		if foundAndParsed {
			return
		}

		scriptContent := s.Text()
		if strings.Contains(scriptContent, "var all_rating") {
			re := regexp.MustCompile(`var all_rating = (\[.*?\]);`)
			matches := re.FindStringSubmatch(scriptContent)

			if len(matches) > 1 {
				jsonString := matches[1]
				
				var ratingsFromScript []models.Rating
				err := json.Unmarshal([]byte(jsonString), &ratingsFromScript)
				if err != nil {
					fmt.Printf("Error unmarshaling JSON from script %d: %v\n", i, err)
					return
				}

                

				allRatings = append(allRatings, ratingsFromScript...)
				foundAndParsed = true
			}
		}
	})

	if !foundAndParsed && len(allRatings) == 0 {
		return nil, fmt.Errorf("no 'var all_rating' script content found or parsed for user %s", username)
	}

	return allRatings, nil
}