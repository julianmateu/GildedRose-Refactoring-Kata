package gildedrose

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
