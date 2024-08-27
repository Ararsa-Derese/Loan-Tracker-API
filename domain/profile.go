package domain

import "mime/multipart"

type Profile struct {
	First_Name      string                `bson:"first_name" form:"first_name"`
	Last_Name       string                `bson:"last_name" form:"last_name" `
	Bio             string                `bson:"bio" form:"bio"`
	Profile_Picture *multipart.FileHeader `bson:"profile_picture" form:"profile_picture"`
	Contact_Info    []ContactInfo         `bson:"contact_info" form:"contact_info"`
}

type ProfileResponse struct {
	First_Name      string        `bson:"first_name" json:"first_name"`
	Last_Name       string        `bson:"last_name" json:"last_name" `
	Bio             string        `bson:"bio" json:"bio"`
	Profile_Picture string        `bson:"profile_picture" json:"profile_picture"`
	Contact_Info    []ContactInfo `bson:"contact_info" json:"contact_info"`
}

type ContactInfo struct {
	Address      string `bson:"address"`
	Phone_number string `bson:"phone_number" json:"phone_number"`
}
