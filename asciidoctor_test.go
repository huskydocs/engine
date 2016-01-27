package engine

import "testing"

func TestInvoke(t *testing.T) {
    err := Invoke(Opts())
    if err != nil {
        t.Errorf("Unexpected Exec error %s", err)
    }
}
