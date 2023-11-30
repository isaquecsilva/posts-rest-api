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

func TestGetNickLength(test *testing.T) {
	test.Run("Should return valid NickName field length", func(t *testing.T) {
		var name string = "Alice"
		post := NewPost(name, "Sunny day.")

		if post.GetNickLength() != len(name) {
			t.FailNow()
		}
	})
}

func TestGetContentLength(test *testing.T) {
	test.Run("Should return valid Content field length", func(t *testing.T) {
		var content string = "Sunny day."
		post := NewPost("Alice", content)

		if post.GetContentLength() != len(content) {
			t.FailNow()
		}
	})
}
