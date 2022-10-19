package butter

import "testing"

type user struct {
	Name string
	Msg  string
}

func TestMergeStructsToMap(t *testing.T) {
	user1 := user{
		Name: "joe",
	}
	user2 := user{
		Msg: "hello!",
	}
	result := MergeStructsToMap(user1, user2)
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
}
