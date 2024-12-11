package tests

import (
	"github.com/bwafi/trendora-backend/internal/entity"
)

func ClearAll() {
	ClearUser()
}

func ClearUser() {
	err := db.Where("id IS NOT NULL").Unscoped().Delete(&entity.Customers{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}

func StrPointer(s string) *string {
	return &s
}
