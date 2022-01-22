package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList_ToFront(t *testing.T) {
	t.Parallel()

	t.Run("simple positive", func(t *testing.T) {
		r := require.New(t)
		t.Parallel()

		list := NewList(nil)

		front, err := list.ToFront(nil)

		r.NoError(err, "err must be nil")
		r.Nil(front.prev, "prev is not nil")
		r.Equal(front, list.root, "root != front")
	})

	t.Run("simple negative", func(t *testing.T) {
		r := require.New(t)
		t.Parallel()

		list := List{len: 0}

		front, err := list.ToFront(nil)

		r.Error(err, "err is nil")
		r.Nil(front, "new elem must be nil")
	})
}

func TestList_ToBack(t *testing.T) {
	t.Parallel()

	t.Run("simple positive", func(t *testing.T) {
		r := require.New(t)
		t.Parallel()

		list := NewList(nil)

		back := list.ToBack(nil)

		r.Nil(back.next, "back is not nil")
		r.Equal(back, list.head, "head != back")
	})
}
