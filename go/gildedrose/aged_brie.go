package gildedrose

type AgedBrieUpdater struct{}

func (u *AgedBrieUpdater) Matches(item *Item) bool {
	return item.Name == "Aged Brie"
}

func (u *AgedBrieUpdater) Update(item *Item) {
	decreaseSellIn(item)
	if expired(item) {
		increaseQuality(item, 2)
	} else {
		increaseQuality(item, 1)
	}
}
