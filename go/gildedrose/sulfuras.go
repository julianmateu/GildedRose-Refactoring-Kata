package gildedrose

type SulfurasUpdater struct{}

func (u *SulfurasUpdater) Matches(item *Item) bool {
	return item.Name == "Sulfuras, Hand of Ragnaros"
}

func (u *SulfurasUpdater) Update(item *Item) {
}
