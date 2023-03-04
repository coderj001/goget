package src

import (
	"testing"
)

func TestModuleName(t *testing.T) {
	if ProjectName() != "goget" {
		t.Errorf("Project name `%s` incorrect", ProjectName())
	}
}
