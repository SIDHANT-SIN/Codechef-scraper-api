package models

type Rating struct {
	Code        string `json:"code"`
	GetYear     string `json:"getyear"`
	GetMonth    string `json:"getmonth"`
	GetDay      string `json:"getday"`
	Rating      string `json:"rating"` 
	Rank        string `json:"rank"`   
	Name        string `json:"name"`
	EndDate     string `json:"end_date"`
	
}

