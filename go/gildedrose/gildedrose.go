package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		updateQualityForItem(item)
	}
}

func updateQualityForItem(item *Item) {
	if item.Name == "Sulfuras, Hand of Ragnaros" {
		return
	}
	if item.Name == "Aged Brie" {
		decreaseSellIn(item)
		if expired(item) {
			increaseQuality(item, 2)
		} else {
			increaseQuality(item, 1)
		}
		return
	}
	if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
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
		return
	}

	decreaseSellIn(item)
	if expired(item) {
		increaseQuality(item, -2)
	} else {
		increaseQuality(item, -1)
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
