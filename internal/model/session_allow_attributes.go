package model

type SessionAllowAttributes struct {
	GenericEntity
	SessionID int64    `gorm:"column:session_id" json:"sessionId"`
	Attribute []string `gorm:"column:attribute;type:text[]" json:"attribute"`
}

func (SessionAllowAttributes) TableName() string {
	return "session_allow_attributes"
}
