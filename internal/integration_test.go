package internal

import "testing"

func Test(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		t.Fatal("test")
	})
}
