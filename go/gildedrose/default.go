package gildedrose

type DefaultUpdater struct{}

func (u *DefaultUpdater) Matches(item *Item) bool {
	return true
}

func (u *DefaultUpdater) Update(item *Item) {
	decreaseSellIn(item)
	if expired(item) {
		increaseQuality(item, -2)
	} else {
		increaseQuality(item, -1)
	}
}
