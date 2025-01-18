package codegen

import (
	"fmt"
	"github.com/briandowns/spinner"
	"log"
	"strings"
	"testing"
	"time"
)

// Gorm gen
type Gorm struct {
	Id   string `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func TestParse(t *testing.T) {
	paths := strings.Split("parse_test.go other.go", " ")
	for _, path := range paths {
		split := strings.Split(path, "/")
		path = split[len(split)-1]
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Start()
		err := GenerateGormColumn(strings.TrimSpace(path))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(4 * time.Second)
		s.Stop()
		fmt.Println("generate gorm column file successful.")

	}
}
