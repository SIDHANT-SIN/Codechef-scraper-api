package models

type UserSolved struct {
    FullySolvedCount     string   `json:"fully_solved_count"`
    FullySolvedList      []string `json:"fully_solved_list,omitempty"`  
   
}