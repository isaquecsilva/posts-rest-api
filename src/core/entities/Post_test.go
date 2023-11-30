package entities

import (
	"testing"
)

func TestNewPost(test *testing.T) {
	test.Run("Instantiate Post structure", func(t *testing.T) {
		post := NewPost("gopher", "Testing post")

		if post.NickName != "gopher" || post.Content != "Testing post" {
			t.FailNow()
		}
	})
}
