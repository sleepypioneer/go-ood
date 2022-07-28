package pnp

import (
	_ "embed"
	"fmt"
	"sort"
)

type (
	// Gopher represents the P&P role of the Gopher
	Gopher struct {
		Character
	}
)

// NewGopher ...
func NewGopher() *Gopher {
	return &Gopher{Character{H: 100, Name: "Gopher"}}
}

// Skills returns the list of abilities the Gopher has
func (g Gopher) Skills() []Skill {
	abs := []Skill{TypeSafety}
	switch {
	case g.X >= 100:
		abs = append(abs, Generics)
		fallthrough
	case g.X >= 10:
		abs = append(abs, Interface)
	}
	sort.Slice(abs, func(i, j int) bool { return abs[i] < abs[j] })
	return abs
}

//go:embed resources/gopher.txt
var gopherArt string

// Art ...
func (g Gopher) Art() string {
	return fmt.Sprintf(gopherArt, g.H, g.X)
}
