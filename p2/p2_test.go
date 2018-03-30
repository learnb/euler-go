package main

import "testing"

func TestP2(t *testing.T) {
    cases := []struct {
        in, want int
    }{
        {89, 44},
    }
    for _, c := range cases {
        got := p2(c.in)
        if got != c.want {
            t.Errorf("p2(%q) == %q, want %q", c.in, got, c.want)
        }
    } 
}
