package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph_Union(t *testing.T) {
	uf := New(10)

	assert.True(t, uf.Empty())

	uf.Union(2, 3) // [0, 1, 2, 2, 4, 5, 6, 7, 8, 9]
	uf.Union(2, 9) // [0, 1, 2, 2, 4, 5, 6, 7, 8, 2]
	uf.Union(8, 9) // [0, 1, 2, 2, 4, 5, 6, 7, 2, 2]

	assert.Equal(t, &Graph{
		count:   7,
		parents: []int{0, 1, 2, 2, 4, 5, 6, 7, 2, 2},
		sizes:   []int{1, 1, 4, 1, 1, 1, 1, 1, 1, 1},
	}, uf)
}

func TestGraph_Connected(t *testing.T) {
	uf := New(10)

	assert.True(t, uf.Empty())

	assert.False(t, uf.Connected(2, 3))

	uf.Union(2, 3)
	assert.True(t, uf.Connected(2, 3))

	assert.False(t, uf.Connected(2, 8))

	uf.Union(2, 9)
	uf.Union(8, 9)
	assert.True(t, uf.Connected(2, 8))
	assert.True(t, uf.Connected(3, 8))
}

func TestGraph_Count(t *testing.T) {
	uf := New(10)

	assert.True(t, uf.Empty())

	assert.Equal(t, 10, uf.Count())

	uf.Union(2, 3)

	assert.Equal(t, 9, uf.Count())

	uf.Union(2, 3)

	assert.Equal(t, 9, uf.Count())

	uf.Union(2, 9)
	uf.Union(8, 9)

	assert.Equal(t, 7, uf.Count())
}

func TestGraphSerialization(t *testing.T) {
	uf := New(10)

	uf.Union(2, 3)
	uf.Union(3, 4)
	uf.Union(8, 9)

	assert.Equal(t, 3, uf.Size())
	assert.Equal(t, []interface{}{0, 1, 2, 2, 2, 5, 6, 7, 8, 8}, uf.Values())

	data, err := uf.ToJSON()
	assert.NoError(t, err)

	newUF := &Graph{}
	err = newUF.FromJSON(data)
	assert.NoError(t, err)
	assert.Equal(t, uf, newUF)
}
