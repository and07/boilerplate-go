package handlers

import (
	"testing"
)

func TestRandom(t *testing.T) {
	var v int
	v = random(0, 5)
	t.Error(v)
}
