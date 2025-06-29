package models

type UserContest struct {
    ContestCode   string `json:"contest_code"`
    ContestName   string `json:"contest_name"`
    Rank          string `json:"rank"`
}
