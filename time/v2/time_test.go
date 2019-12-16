package time

import "testing"

func TestGetInfo(t *testing.T) {
    want := "gorb-util/time  version 2"
    if got := GetInfoV2(); got != want {
        t.Errorf("TestGetInfo () = %q, want %q",got, want)
    }
}
