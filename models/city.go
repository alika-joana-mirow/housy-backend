package models

type City struct {
	ID   int    "json:'id' gorm:'primary_key:auto_increment'"
	Name string "json:'name'"
}
