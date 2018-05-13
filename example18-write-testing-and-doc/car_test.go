package car

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simple testing what different between Fatal and Error
func TestNew(t *testing.T) {
	c, err := New("", 100)
	if err != nil {
		t.Fatal("got errors:", err)
	}

	if c == nil {
		t.Error("car should be nil")
	}
}

// Simple testing with testify tool
func TestNewWithAssert(t *testing.T) {
	c, err := New("", 100)
	assert.NotNil(t, err)
	assert.Nil(t, c)
}
