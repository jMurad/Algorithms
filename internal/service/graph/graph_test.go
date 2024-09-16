package graph

import "testing"

func TestAdd(t *testing.T) {
	gg := New()

	gg.AddNode("Saf", Neighbor{"Mur", 1}, Neighbor{"Pat", 2})
	gg.AddNode("Mur", Neighbor{"Lim", 1}, Neighbor{"Zub", 2})
	gg.AddNode("Pat", Neighbor{"Mag", 1}, Neighbor{"Sar", 2})

	gg.Neighbors()

	t.Error("")

}
