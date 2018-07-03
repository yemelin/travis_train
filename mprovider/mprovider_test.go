package mprovider

import (
	"testing"
)

func TestMessage(t *testing.T) {
	m := New()
	if m.Message() != "burn in hell!" {
		t.Error("message incorrect")
	}
}
