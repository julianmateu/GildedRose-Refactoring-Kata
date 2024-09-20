package gildedrose

import "strings"

type BackstagePassUpdater struct{}

func (u *BackstagePassUpdater) Matches(item *Item) bool {
	return strings.HasPrefix(item.Name, "Backstage passes")
}

func (u *BackstagePassUpdater) Update(item *Item) {
	decreaseSellIn(item)
	if expired(item) {
		item.Quality = 0
	} else if item.SellIn < 5 {
		increaseQuality(item, 3)
	} else if item.SellIn < 10 {
		increaseQuality(item, 2)
	} else {
		increaseQuality(item, 1)
	}
}
