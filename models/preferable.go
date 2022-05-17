package models

type UserPreferable struct {
	User_ID     uint     `bson:"user_id"`
	Preferable  []string `bson:"preferable"`
	Prohibition []string `bson:"prohibition"`
}

type JsonReciever struct {
	Preferable  []string `json:"preferables"`
	Prohibition []string `json:"prohibitions"`
}
