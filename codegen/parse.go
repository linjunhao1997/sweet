package codegen

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
)

type GoFile struct {
	PackageName string
	StructNodes []StructNode
}

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

func Parse(filename string) (res *GoFile, err error) {

	fileSet := token.NewFileSet()
	f, err := parser.ParseFile(fileSet, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	res = &GoFile{
		PackageName: f.Name.Name,
		StructNodes: make([]StructNode, 0),
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
			// docText := strings.ToLower(d.Doc.Text())

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
							}
						}
						res.StructNodes = append(res.StructNodes, structNode)
					}
				}
			}
		}
	}
	return res, nil
}
