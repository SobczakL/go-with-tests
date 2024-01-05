package main

import "testing"

func TestHello()  {
    got := Hello()
    want := "Hello, world"

    if got != want{
        t.Errorf("got %q want %q", got, want)
    }
}
