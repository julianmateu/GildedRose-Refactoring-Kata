package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

type ItemUpdater interface {
	UpdateQuality(item *Item)
	Matches(item *Item) bool
}

type AgedBrieUpdater struct{}

func (u *AgedBrieUpdater) Matches(item *Item) bool {
	return item.Name == "Aged Brie"
}

func (u *AgedBrieUpdater) UpdateQuality(item *Item) {
	decreaseSellIn(item)
	if expired(item) {
		increaseQuality(item, 2)
	} else {
		increaseQuality(item, 1)
	}
}

type BackstagePassUpdater struct{}

func (u *BackstagePassUpdater) Matches(item *Item) bool {
	return item.Name == "Backstage passes to a TAFKAL80ETC concert"
}

func (u *BackstagePassUpdater) UpdateQuality(item *Item) {
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

type SulfurasUpdater struct{}

func (u *SulfurasUpdater) Matches(item *Item) bool {
	return item.Name == "Sulfuras, Hand of Ragnaros"
}

func (u *SulfurasUpdater) UpdateQuality(item *Item) {
}

type DefaultUpdater struct{}

func (u *DefaultUpdater) Matches(item *Item) bool {
	return true
}

func (u *DefaultUpdater) UpdateQuality(item *Item) {
	decreaseSellIn(item)
	if expired(item) {
		increaseQuality(item, -2)
	} else {
		increaseQuality(item, -1)
	}
}

var itemUpdaters = []ItemUpdater{
	&SulfurasUpdater{},
	&AgedBrieUpdater{},
	&BackstagePassUpdater{},
	&DefaultUpdater{},
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		for _, updater := range itemUpdaters {
			if updater.Matches(item) {
				updater.UpdateQuality(item)
				break
			}
		}
	}
}

func increaseQuality(item *Item, increase int) {
	item.Quality = item.Quality + increase
	item.Quality = max(0, min(item.Quality, 50))
}

func expired(item *Item) bool {
	return item.SellIn < 0
}

func decreaseSellIn(item *Item) {
	item.SellIn = item.SellIn - 1
}
