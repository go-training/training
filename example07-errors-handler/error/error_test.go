package error

import "testing"

func TestIsMyError(t *testing.T) {
	err := MyError{"title", "message"}

	ok := IsMyError(err)

	if !ok {
		t.Fatal("testing error")
	}

	if err.Error() != "title: message" {
		t.Fatal("message error")
	}
}
