package model

type SessionAllowAttributes struct {
	SessionID             int64    `gorm:"column:session_id" json:"sessionId"`
	UserWithThisAttribute int64    `gorm:"column:user_with_this_attribute" json:"userWithThisAttribute"`
	Attributes            []string `gorm:"column:attributes;type:text[]" json:"attributes"`
}

func (SessionAllowAttributes) TableName() string {
	return "session_allow_attributes"
}
