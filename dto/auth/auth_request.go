package authdto

type AuthRequest struct {
	Username string "json:'name' gorm:'type: varchar(255)'"
	Password string "json:'password' gorm:'type: varchar(255)'"
}