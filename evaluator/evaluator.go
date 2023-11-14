package evaluator

import (
	"github.com/goalm/intrpr/ast"
	"github.com/goalm/intrpr/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Boolean:
		return &object.Boolean{Value: node.Value}
	}
	return nil
}
