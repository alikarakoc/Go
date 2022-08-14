package models

type User struct {
	UserId   string `json:"userId"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Age      int64  `json:"age"`
}
