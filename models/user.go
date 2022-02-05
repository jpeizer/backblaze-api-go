package models

type User struct {
	Id               int64  `json:"id"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	ApplicationKeyId string `json:"applicationKeyId"`
	ApplicationKey   string `json:"applicationKey"`
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
