package time

import "testing"

func TestGetInfo(t *testing.T) {
    want := "gorb-util/time  version 1"
    if got := GetInfo(); got != want {
        t.Errorf("TestGetInfo () = %q, want %q",got,want)
    }
}
