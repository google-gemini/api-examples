package examples

import (
	"testing"
)

func TestCodeExecutionBasic(t *testing.T) {
	_, err := CodeExecutionBasic()
	if err != nil {
		t.Errorf("CodeExecutionBasic returned an error.")
	}
}
