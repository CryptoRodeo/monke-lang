package object

import (
	"fmt"
)

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

/**
Dev notes:
- every value we encounter and evaluate will be represented using an Object interace
- the reason for the interface is that every value needs a different internal representation
  and having different struct types makes it easier to define different values instead of trying to fit them
  all in the same struct field.

**/