package filter

import (
	"github.com/bio-routing/bio-rd/net"
	"github.com/bio-routing/bio-rd/route"
)

type From struct {
	prefixLists []*PrefixList
}

func (f *From) Matches(p net.Prefix, pa *route.Path) bool {
	return f.matchesAnyPrefixList(p)
}

func (t *From) matchesAnyPrefixList(p net.Prefix) bool {
	for _, l := range t.prefixLists {
		if l.Matches(p) {
			return true
		}
	}

	return false
}
