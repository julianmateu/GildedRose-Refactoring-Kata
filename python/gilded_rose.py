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
        
        if item.name != "Aged Brie" and item.name != "Backstage passes to a TAFKAL80ETC concert":
            ## Decreasing quality by 1 except brie, concert and sulfuras
            if item.quality > 0:
                item.quality = item.quality - 1
        else:
            ## Either brie or concert
            if item.quality < 50:
                item.quality = item.quality + 1
                if item.name == "Backstage passes to a TAFKAL80ETC concert":
                    ## concert will potentially enter both if conditions
                    if item.sell_in < 11:
                        if item.quality < 50:
                            item.quality = item.quality + 1
                    if item.sell_in < 6:
                        if item.quality < 50:
                            item.quality = item.quality + 1


        ## This will always be executed
        item.sell_in = item.sell_in - 1
        if item.sell_in < 0:
            ## expired
            if item.name != "Aged Brie":
                if item.name != "Backstage passes to a TAFKAL80ETC concert":
                    if item.quality > 0:
                        ## if not brie, concert or sulfuras, decrease quality by 1 more
                        item.quality = item.quality - 1
                else:
                    # concert
                    item.quality = item.quality - item.quality ## set quality to 0
            else:
                ## brie
                if item.quality < 50:
                    item.quality = item.quality + 1


class Item:
    def __init__(self, name, sell_in, quality):
        self.name = name
        self.sell_in = sell_in
        self.quality = quality

    def __repr__(self):
        return "%s, %s, %s" % (self.name, self.sell_in, self.quality)
