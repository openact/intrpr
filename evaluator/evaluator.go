package evaluator

import (
	"intrpr/ast"
	"intrpr/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Boolean:
		return &object.Boolean{Value: node.Value}
	}
	return nil
}
