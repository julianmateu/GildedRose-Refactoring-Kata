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
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
		}
		item.SellIn = item.SellIn - 1
		if item.SellIn < 0 {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}
		return
	}
	if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
			if item.SellIn < 11 {
				if item.Quality < 50 {
					item.Quality = item.Quality + 1
				}
			}
			if item.SellIn < 6 {
				if item.Quality < 50 {
					item.Quality = item.Quality + 1
				}
			}
		}
		item.SellIn = item.SellIn - 1
		if item.SellIn < 0 {
			item.Quality = item.Quality - item.Quality
		}
		return
	}
	if item.Quality > 0 {
		item.Quality = item.Quality - 1
	}

	item.SellIn = item.SellIn - 1

	if item.SellIn < 0 {
		if item.Quality > 0 {
			item.Quality = item.Quality - 1
		}

	}
}
