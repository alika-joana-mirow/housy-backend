package models

type House struct {
	ID        int                 "json:'id' gorm:'primary_key:auto_increment'"
	Name      string              "json:'name' form:'name' gorm:'type: varchar(255)'"
	Address   string              "json:'address' form:'address' gorm:'type: text'"
	Price     int                 "json:'price' form:'price' gorm:'type: int'"
	TypeRent  string              "json:'type_rent' form:'type_rent' gorm:'type: varchar(255)'"
	// Ameneties []Ameneties         "json:'ameneties' form:'amenities"
	// Image     string              "json:'image' form:'image' gorm:'type: varchar(255)'"
	UserID    int                 "json:'user_id' form:'user_id'"
	User      UserProfileResponse "json:'user'"
	City      []City              "json:'city' gorm:'hasone:house_city"
	CityID    []int               "json:'city_id' form:'city_id' gorm:'-'"
}
