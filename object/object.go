package object

import "fmt"

type ObjectType string

const (
	// Object types
	NULL_OBJ    = "NULL"
	BOOLEAN_OBJ = "BOOLEAN"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }

func (n *Null) Inspect() string { return "null" }
