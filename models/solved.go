package models


type Solved struct {
    
    Name      string `json:"name"`        
	Queslink  string `json:"ques_link"`   
	Sollink  string `json:"sol_link"`  
   
}



type UserSolved struct {

    FullySolvedList      []Solved `json:"fully_solved_list,omitempty"`  
   
}