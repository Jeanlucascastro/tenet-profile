package model

import "reflect"

type Profile struct {
	GenericEntity
	Name       string `gorm:"column:name" json:"name"`
	LastName   string `gorm:"column:last_name" json:"lastName"`
	Nickname   string `gorm:"column:nickname" json:"nickname"`
	Age        int    `gorm:"column:age" json:"age"`
	Bio        string `gorm:"column:bio;type:text" json:"bio"`
	PictureUrl string `gorm:"column:picture_url" json:"pictureUrl"`
	UserID     int64  `gorm:"column:user_id" json:"userId"`
}

func (Profile) TableName() string {
	return "profile"
}

type ProfileDTO struct {
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	Nickname   string `json:"nickname"`
	Age        int    `json:"age"`
	Bio        string `json:"bio"`
	PictureUrl string `json:"pictureUrl"`
	UserID     int64  `json:"userId"`
}

func (p *ProfileDTO) ToEntity() *Profile {
	return &Profile{
		Name:       p.Name,
		LastName:   p.LastName,
		Nickname:   p.Nickname,
		Age:        p.Age,
		Bio:        p.Bio,
		PictureUrl: p.PictureUrl,
		UserID:     p.UserID,
	}
}

func (p *Profile) FilterByAttributes(allowed []string) map[string]interface{} {
	result := make(map[string]interface{})

	v := reflect.ValueOf(*p)
	t := reflect.TypeOf(*p)

	allowedSet := make(map[string]bool)
	for _, a := range allowed {
		allowedSet[a] = true
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			continue
		}

		if allowedSet[jsonTag] {
			result[jsonTag] = v.Field(i).Interface()
		}
	}

	return result
}
