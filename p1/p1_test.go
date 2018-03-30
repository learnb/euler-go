package main

import "testing"

func TestP1(t *testing.T) {
    cases := []struct {
        in, want int
    }{
        {10, 23},
    }
    for _, c := range cases {
        got := p1(c.in)
        if got != c.want {
            t.Errorf("p1(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
