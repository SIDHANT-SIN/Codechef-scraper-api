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

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch user: %s (status code: %d)", username, res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
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