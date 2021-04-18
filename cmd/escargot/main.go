package main

import (
	"github.com/oryanmoshe/escargot/internal/store"
)

func main() {
	s := store.FromID("oryan-store")
	s.Whoop("A log message")
}
