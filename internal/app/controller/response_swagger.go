package controller

//types for swagger
type InputShortLink struct {
	Url string `json:"url" example:"https://test.ru"`
}

type InputLongLink struct {
	Url string `json:"url" example:"8wSnscuTr6"`
}

type GetShortLinkMessage400 struct {
	Message string `json:"error" example:"invalid response body"`
}

type GetShortLinkMessage500 struct {
	Message string `json:"error" example:"internal server error"`
}
