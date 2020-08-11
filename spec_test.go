package builders

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypeSpec(t *testing.T) {
	i := Idents("name", "typ")
	spec := TypeSpec(i[0], i[1])

	require.Equal(t, i[0], spec.Name)
	require.Equal(t, i[1], spec.Type)
}
