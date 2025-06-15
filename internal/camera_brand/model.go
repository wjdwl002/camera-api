package camera_brand

import "gorm.io/gorm"

type CameraBrand struct {
	gorm.Model
	ID   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name" gorm:"unique;not null"`
}
