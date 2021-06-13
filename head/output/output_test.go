package output

import "testing"

func TestOutput_SprintKV(t *testing.T) {
	o := New()

	o.SprintKV("TestKey", map[string]string{
		"Some": "Key",
	})

	o.SprintKV("TestKey", "Another")
}
