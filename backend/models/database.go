package models

import "gorm.io/gorm"

type Dbinstance struct {
	DB *gorm.DB
}
