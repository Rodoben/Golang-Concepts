package main

import "testing"

func TestNewDeck(t *testing.T)
d:= newDeck()

if len(d)!= 16{
	t.Errorf("Expected deck length of 20, but go",len(d))
}