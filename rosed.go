package rosed

import "github.com/dekarrin/rosed/internal/gem"

func _g(s string) gem.String {
	return gem.New(s)
}
