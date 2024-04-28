package model

type UserData struct {
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	Number string `json:"number,omitempty" bson:"number,omitempty"`
}
type ResponseData struct {
	Code       string    `json:"code,omitempty" bson:"code,omitempty"`
	DataCookie *UserData `json:"DataCookie,omitempty" bson:"DataCookie,omitempty"`
}
