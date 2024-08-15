package model

type Note struct {
	Id      int    `json:"id" gorm:"type:int;primaryKey"`
	Content string `json:"content,omitempty" gorm:"not null"`
}
