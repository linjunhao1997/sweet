package codegen

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
)

type StructNode struct {
	StructName string
	Fields     []FieldNode
}

type FieldNode struct {
	TypeName  string
	FieldName string
	Tag       reflect.StructTag
}

type TagNode struct {
	Name  string
	Value string
}

func ParseStructNode(filename string) (res []StructNode, err error) {
	fileSet := token.NewFileSet()
	f, err := parser.ParseFile(fileSet, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			if d.Tok != token.TYPE {
				continue
			}
			if d.Doc == nil {
				continue
			}
			docText := strings.ToLower(d.Doc.Text())
			if !strings.HasPrefix(docText, "gen") {
				continue
			}

			for _, spec := range d.Specs {
				switch sp := spec.(type) {
				case *ast.TypeSpec:
					structNode := StructNode{
						StructName: sp.Name.Name,
						Fields:     make([]FieldNode, 0),
					}
					switch t := sp.Type.(type) {
					case *ast.StructType:
						if !t.Struct.IsValid() {
							continue
						}

						for _, field := range t.Fields.List {
							var fieldTypeName string
							switch ft := field.Type.(type) {
							case *ast.Ident:
								fieldTypeName = ft.Name
							}
							for _, name := range field.Names {
								structNode.Fields = append(structNode.Fields, FieldNode{
									TypeName:  fieldTypeName,
									FieldName: name.String(),
									Tag:       reflect.StructTag(strings.Trim(field.Tag.Value, "`")),
								})
								res = append(res, structNode)
							}

						}
					}

				}
			}
		}
	}
	return
}
