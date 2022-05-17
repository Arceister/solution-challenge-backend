package models

type UserPreferable struct {
	User_ID    uint     `bson:"user_id"`
	Preferable []string `bson:"preferable"`
}

type JsonReciever struct {
	Preferable []string `json:"preferables"`
}
