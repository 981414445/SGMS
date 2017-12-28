package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	if "2.0" != Version {
		t.Error("version not 2.0")
	}
}
