# -*- coding: utf-8 -*-

class GildedRose(object):

    def __init__(self, items):
        self.items = items

    def update_quality(self):
        for item in self.items:
            self.update_quality_for_item(item)
    
    def update_quality_for_item(self, item):
        if item.name == "Sulfuras, Hand of Ragnaros":
            return

        expired = item.sell_in <= 0

        if item.name == "Aged Brie":
            if not expired:
                item.quality += 1
            else:
                item.quality += 2
            item.quality = min(item.quality, 50)
            ## This will always be executed
            item.sell_in = item.sell_in - 1
            return
    
        if item.name == "Backstage passes to a TAFKAL80ETC concert":
            if expired:
                item.quality = 0
                item.sell_in = item.sell_in - 1
                return
            if item.sell_in >= 11:
                item.quality += 1
            elif item.sell_in >= 6:
                item.quality += 2
            elif item.sell_in >= 0:
                item.quality += 3
            item.quality = min(item.quality, 50)
            item.sell_in = item.sell_in - 1
            return
    
        ## Regular item
        if not expired:
            item.quality -= 1
        else:
            item.quality -= 2
        item.quality = max(item.quality, 0)
        item.sell_in = item.sell_in - 1
        return

class Item:
    def __init__(self, name, sell_in, quality):
        self.name = name
        self.sell_in = sell_in
        self.quality = quality

    def __repr__(self):
        return "%s, %s, %s" % (self.name, self.sell_in, self.quality)
