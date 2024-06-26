package dto

import "time"

type ToResponseGetCat struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IssuedBy struct {
	Name string	`json:"name"`
	Email string	`json:"email"`
	CreatedAt string	`json:"createdAt"`
}

type CatDetail struct {
	Id int `json:"id"`
	UserId int `json:"userId"`
	Name string `json:"name"`
	Race string `json:"race"`
	Sex string `json:"sex"`
	Description string `json:"description"`
	AgeInMonth int `json:"ageInMonth"`
	ImageUrls []string `json:"imageUrls"` 
	HasMatched bool `json:"hasMatched"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
type ResponseGetMatch struct {
	Id int		`json:"id"`
	IssuedBy IssuedBy `json:"issuedBy"`
	Message	string	`json:"message"`
	CreatedAt	string	`json:"createdAt"`
	MatchCatDetail CatDetail `json:"matchCatDetail"`
	UserCatDetail CatDetail `json:"userCatDetail"`
}