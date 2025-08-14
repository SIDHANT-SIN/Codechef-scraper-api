package scraper

import (
	"Codechef-scraper-api/models" // Assuming this is the correct path to your models package
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Remove the redundant Solved struct definition here.
// type Solved struct {
//     Name      string
//     Queslink  string
//     Sollink  string
// }

func ScrapeSolved(username string) (*models.UserSolved, error) {
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


    //fmt.Println(doc);

	var questions []models.Solved 

	doc.Find("div#rankContentDiv.table-questions table.dataTable tbody tr").Each(func(i int, row *goquery.Selection) {
		var p models.Solved 
		secondTd := row.Find("td:nth-child(2)")
		if secondTd.Length() > 0 {
			if tdTitle, exists := secondTd.Attr("title"); exists {
				p.Name = tdTitle
			}

			linkATag := secondTd.Find("a")
			if linkATag.Length() > 0 {
				if href, exists := linkATag.Attr("href"); exists {
					p.Queslink = href
				}
			}
		}

		FifthTd := row.Find("td:nth-child(5)")
		if FifthTd.Length() > 0 {
			linkATag := FifthTd.Find("a")
			if linkATag.Length() > 0 {
				if href, exists := linkATag.Attr("href"); exists {
					p.Sollink = href
				}
			}
		}

		if p.Name != "" || p.Queslink != "" || p.Sollink != "" {

            fmt.Println("name-",p.Name," questionlink-",p.Queslink," sol-",p.Sollink);
			questions = append(questions, p)
		}
	})



	return &models.UserSolved{
		FullySolvedList: questions,
	}, nil
}