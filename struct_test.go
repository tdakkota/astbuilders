package builders

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStructBuilder(t *testing.T) {
	a := require.New(t)
	s := NewStructBuilder()

	i := Idents("field", "typ")
	tag := StringLit("tag")
	field := Field(i[0])(i[1], tag)

	s = s.AddFields(field)
	typ := s.Complete()
	fields := typ.Fields

	a.Len(fields.List, 1)
	a.Equal(field, fields.List[0])
	a.Equal(typ, s.CompleteAsLit().Type)
}
