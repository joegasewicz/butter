package butter

import "testing"

type user struct {
	ID   uint
	Name string
	Msg  string
}

func TestMergeStructsToMap(t *testing.T) {
	incoming := user{
		Name: "joe",
	}
	original := user{
		ID:  1,
		Msg: "hello!",
	}
	result := MergeStructsToMap(original, incoming)
	if result == nil {
		t.Logf("expected result to not be null\n")
		t.Fail()
	}
	if result["Name"] != "joe" {
		t.Logf("expected name to be joe but got %s", result["Name"])
		t.Fail()
	}
	if result["Msg"] != "hello!" {
		t.Logf("expected name to be joe but got %s", result["Msg"])
		t.Fail()
	}
	if result["ID"] != uint(1) {
		t.Logf("expected id to be 1 but got %d", result["ID"])
		t.Fail()
	}
}
