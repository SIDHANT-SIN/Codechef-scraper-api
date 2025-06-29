package models

type UserData struct {
    Username       string `json:"username"`
    Name           string `json:"name"`
    Stars          string `json:"stars"`
    Rating         string `json:"rating"`
    HighestRating  string `json:"highest_rating"`
    GlobalRank     string `json:"global_rank"`
    CountryRank    string `json:"country_rank"`
    Country    string `json:"country"`
   Institution string `json:"institution"`
}

