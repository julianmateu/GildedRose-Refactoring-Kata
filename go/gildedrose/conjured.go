package gildedrose

import "strings"

type ConjuredUpdater struct{}

func (u *ConjuredUpdater) Matches(item *Item) bool {
	return strings.HasPrefix(item.Name, "Conjured")
}

func (u *ConjuredUpdater) Update(item *Item) {
	decreaseSellIn(item)
	if expired(item) {
		increaseQuality(item, -4)
	} else {
		increaseQuality(item, -2)
	}
}
