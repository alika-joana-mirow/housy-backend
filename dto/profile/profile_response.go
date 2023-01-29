package profile

import "housy/models"

type ProfileResponse struct {
	ID       int                        "json:'id' gorm:'primary_key:auto_increment'"
	Fullname string                     "json:'fullname' gorm:'type: varchar(255)'"
	Gender   string                     "json:'gender' gorm:'type: varchar(255)'"
	Image    string                     "json:'image' form:'image' gorm:'type: varchar(255)'"
	UserID   int                        "json:'user_id'"
	User     models.UserProfileResponse "json:'user'"
}
