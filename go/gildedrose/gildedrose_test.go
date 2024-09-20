package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Foo(t *testing.T) {

	testCases := []struct {
		testName        string
		name            string
		sellIn          int
		quality         int
		expectedSellIn  int
		expectedQuality int
	}{
		{
			"Sulfuras does not change if sellIn = 0",
			"Sulfuras, Hand of Ragnaros",
			0,
			80,
			0,
			80,
		},
		{
			"Sulfuras does not change if not expired",
			"Sulfuras, Hand of Ragnaros",
			1,
			80,
			1,
			80,
		},
		{
			"Sulfuras does not change if expired",
			"Sulfuras, Hand of Ragnaros",
			-1,
			80,
			-1,
			80,
		},
		{
			"Normal item decreases in quality and sellIn by 1 if not expired",
			"+5 Dexterity Vest",
			10,
			20,
			9,
			19,
		},
		{
			"Normal item decreases in quality and sellIn by 1 if not expired",
			"+5 Dexterity Vest",
			1,
			20,
			0,
			19,
		},
		{
			"Normal item decreases in quality by 2 if sellIn = 0",
			"+5 Dexterity Vest",
			0,
			20,
			-1,
			18,
		},
		{
			"Normal item decreases in quality by 2 if expired",
			"+5 Dexterity Vest",
			-1,
			20,
			-2,
			18,
		},
		{
			"Normal Item quality does not go below 0",
			"+5 Dexterity Vest",
			0,
			0,
			-1,
			0,
		},
		{
			"Aged Brie quality does not go above 50",
			"Aged Brie",
			2,
			50,
			1,
			50,
		},
		{
			"Aged Brie quality does not go above 50",
			"Aged Brie",
			-1,
			50,
			-2,
			50,
		},
		{
			"Aged Brie quality does not go above 50",
			"Aged Brie",
			-1,
			49,
			-2,
			50,
		},
		{
			"Aged Brie increases in quality by 1 if not expired",
			"Aged Brie",
			2,
			0,
			1,
			1,
		},
		{
			"Aged Brie increases in quality by 1 if not expired",
			"Aged Brie",
			1,
			0,
			0,
			1,
		},
		{
			"Aged Brie increases in quality by 2 if expired",
			"Aged Brie",
			0,
			0,
			-1,
			2,
		},
		{
			"Aged Brie increases in quality by 2 if expired",
			"Aged Brie",
			-1,
			0,
			-2,
			2,
		},
		{
			"Aged Brie quality does not go above 50",
			"Aged Brie",
			2,
			50,
			1,
			50,
		},
		{
			"Aged Brie quality does not go above 50",
			"Aged Brie",
			-2,
			50,
			-3,
			50,
		},
		{
			"Backstage passes increases in quality by 1 if sellIn > 10",
			"Backstage passes to a TAFKAL80ETC concert",
			15,
			20,
			14,
			21,
		},
		{
			"Backstage passes increases in quality by 1 if sellIn > 10",
			"Backstage passes to a TAFKAL80ETC concert",
			11,
			20,
			10,
			21,
		},
		{
			"Backstage passes increases in quality by 2 if 5 <= sellIn <= 10",
			"Backstage passes to a TAFKAL80ETC concert",
			10,
			20,
			9,
			22,
		},
		{
			"Backstage passes increases in quality by 2 if 5 < sellIn <= 10",
			"Backstage passes to a TAFKAL80ETC concert",
			6,
			20,
			5,
			22,
		},
		{
			"Backstage passes increases in quality by 3 if 0 < sellIn <= 5",
			"Backstage passes to a TAFKAL80ETC concert",
			5,
			20,
			4,
			23,
		},
		{
			"Backstage passes increases in quality by 3 if 0 < sellIn <= 5",
			"Backstage passes to a TAFKAL80ETC concert",
			1,
			20,
			0,
			23,
		},
		{
			"Backstage passes quality drops to 0 if expired",
			"Backstage passes to a TAFKAL80ETC concert",
			0,
			20,
			-1,
			0,
		},
		{
			"Backstage passes quality drops to 0 if expired",
			"Backstage passes to a TAFKAL80ETC concert",
			-1,
			20,
			-2,
			0,
		},
		{
			"Backstage passes quality does not go above 50",
			"Backstage passes to a TAFKAL80ETC concert",
			15,
			50,
			14,
			50,
		},
		{
			"Backstage passes quality does not go above 50",
			"Backstage passes to a TAFKAL80ETC concert",
			8,
			49,
			7,
			50,
		},
		{
			"Backstage passes quality does not go above 50",
			"Backstage passes to a TAFKAL80ETC concert",
			1,
			49,
			0,
			50,
		},
		{
			"Conjured decreases in quality by 2 if not expired",
			"Conjured Mana Cake",
			3,
			6,
			2,
			4,
		},
		{
			"Conjured decreases in quality by 4 if expired",
			"Conjured Mana Cake",
			0,
			6,
			-1,
			2,
		},
		{
			"Conjured quality does not go below 0",
			"Conjured Mana Cake",
			0,
			0,
			-1,
			0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			var items = []*gildedrose.Item{
				{tc.name, tc.sellIn, tc.quality},
			}
			gildedrose.UpdateQuality(items)
			if items[0].SellIn != tc.expectedSellIn {
				t.Errorf("SellIn: Expected %d but got %d ", tc.expectedSellIn, items[0].SellIn)
			}
			if items[0].Quality != tc.expectedQuality {
				t.Errorf("Quality: Expected %d but got %d ", tc.expectedQuality, items[0].Quality)
			}
			if items[0].Name != tc.name {
				t.Errorf("Name: Expected %s but got %s ", tc.name, items[0].Name)
			}
		})
	}
}
