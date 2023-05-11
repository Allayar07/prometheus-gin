package models

type User struct {
	Name    string `json:"name"`
	Phone   int    `json:"phone"`
	Address string `json:"address"`
}
