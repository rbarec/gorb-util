package str

import "testing"

func TestGetInfo(t *testing.T) {
    want := "gorb-util/str"
    if got := GetInfo(); got != want {
        t.Errorf("TestGetInfo () = %q, want %q")
    }
}
