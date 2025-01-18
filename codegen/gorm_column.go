package codegen

import (
	_ "embed"
	"fmt"
	"github.com/dave/jennifer/jen"
	"strings"
	"time"
)

type GormColumn struct {
	EntityField string
	ColumnField string
	ColumnName  string
}

func GenerateGormColumn(filename string) error {
	res, err := Parse(filename)
	if err != nil {
		return err
	}

	data := make([]GormColumn, 0)
	for _, node := range res.StructNodes {
		for _, field := range node.Fields {
			value, ok := field.Tag.Lookup("gorm")
			if !ok {
				continue
			}
			properties := strings.Split(value, ",")
			for _, property := range properties {
				kv := strings.Split(property, ":")
				if len(kv) != 2 {
					continue
				}
				k := kv[0]
				if k != "column" {
					continue
				}
				data = append(data, GormColumn{
					EntityField: node.StructName,
					ColumnField: field.FieldName,
					ColumnName:  kv[1],
				})
				break
			}
		}
	}

	curTime := time.Now().Format("2006-01-02 15:04:05")

	f := jen.NewFile(res.PackageName)
	f.PackageComment("// Do not Edit Generated File.")
	f.PackageComment(fmt.Sprintf("// Generated at: %s\n", curTime))
	list := make([]jen.Code, 0)
	for _, e := range data {
		list = append(list, jen.Id(fmt.Sprintf("Column%s%s", e.EntityField, e.ColumnField)).Op("=").Lit(e.ColumnName))
	}
	f.Const().Defs(list...)

	if err = f.Save(fmt.Sprintf("%s_generated.go", strings.TrimRight(filename, ".go"))); err != nil {
		return err
	}
	return nil
}
