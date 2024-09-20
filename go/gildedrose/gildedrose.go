package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

type ItemUpdater interface {
	Update(item *Item)
	Matches(item *Item) bool
}

var itemUpdaters = []ItemUpdater{
	&SulfurasUpdater{},
	&AgedBrieUpdater{},
	&BackstagePassUpdater{},
	&ConjuredUpdater{},
	&DefaultUpdater{},
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		for _, updater := range itemUpdaters {
			if updater.Matches(item) {
				updater.Update(item)
				break
			}
		}
	}
}
