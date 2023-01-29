package models

type User struct {
	ID       int      "json:'id' gorm:'primary_key:auto_increment'"
	Username string   "json:'username' gorm:'type: varchar(255)'"
	Email    string   "json:'email' gorm:'type: varchar(255)'"
	Password string   "json:'-' gorm:'type: varchar(255)'"
	ListAs   []ListAs "json:'list_as' gorm:'hasone:user_as"
	ListAsID []int    "json:'list_as_id' form:'list_as_id' gorm:'-'"
}

type UserProfileResponse struct {
	ID       int    "json:'id'"
	Username string "json:'username'"
}

func (UserProfileResponse) TableName() string {
	return "users"
}
