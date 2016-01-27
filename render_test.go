package engine

import "testing"

func TestRender(t *testing.T) {
    cases := []struct {
        input, expected string
    }{
        {"Something", "Rendered"},
        {"Another Thing", "Rendered"},
    }
    for _, c := range cases {
        got := Render(c.input)
        if got != c.expected {
            t.Errorf("Render(%q) == %q, want %q", c.input, got, c.expected)
        }
    }
}
