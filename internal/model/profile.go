package model

type Profile struct {
	GenericEntity
	Name       string `gorm:"column:name" json:"name"`
	LastName   string `gorm:"column:last_name" json:"lastName"`
	Nickname   string `gorm:"column:nickname" json:"nickname"`
	Age        int    `gorm:"column:age" json:"age"`
	Bio        string `gorm:"column:bio;type:text" json:"bio"`
	PictureUrl string `gorm:"column:picture_url" json:"pictureUrl"`
}

func (Profile) TableName() string {
	return "profile"
}
