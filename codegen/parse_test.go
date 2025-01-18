package codegen

import "testing"

// Gorm gen
type Gorm struct {
	Id   string `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func TestParse(t *testing.T) {
	err := GenerateGormColumn("parse_test.go")
	if err != nil {
		t.Fatal(err)
	}
}
