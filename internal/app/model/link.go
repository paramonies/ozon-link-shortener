package model

type ClientLink struct {
	Url string `json:"url" binding:"required"`
}
