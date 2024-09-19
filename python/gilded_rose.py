# -*- coding: utf-8 -*-

def noop(item):
    return item

def decrease_sell_in(item):
    item.sell_in -= 1
    return item

def expired(item):
    return item.sell_in < 0

def quality_bounds(item):
    item.quality = max(0, min(50, item.quality))
    return item

def brie_quality_updater(item):
    if expired(item):
        item.quality += 2
    else:
        item.quality += 1
    return quality_bounds(item)

def backstage_quality_updater(item):
    if item.sell_in >= 10:
        item.quality += 1
    elif item.sell_in >= 5:
        item.quality += 2
    elif item.sell_in >= 0:
        item.quality += 3
    else:
        item.quality = 0
    return quality_bounds(item)

def conjured_quality_updater(item):
    if expired(item):
        item.quality -= 4
    else:
        item.quality -= 2
    return quality_bounds(item)

def default_quality_updater(item):
    if expired(item):
        item.quality -= 2
    else:
        item.quality -= 1
    return quality_bounds(item)

def exact_name_matcher(name):
    return lambda item: item.name == name

def starts_with_matcher(prefix):
    return lambda item: item.name.startswith(prefix)

class ItemUpdater:
    def __init__(self, sell_in_updater, quality_updater, name_matcher):
        self.sell_in_updater = sell_in_updater
        self.quality_updater = quality_updater
        self.name_matcher = name_matcher

    def update(self, item):
        if not self.name_matcher(item):
            return item
        item = self.sell_in_updater(item)
        item = self.quality_updater(item)
        return item
    
    def applies_to(self, item):
        return self.name_matcher(item)

    
ITEM_UPDATERS = [
    ItemUpdater(
        name_matcher=exact_name_matcher("Sulfuras, Hand of Ragnaros"),
        sell_in_updater=noop,
        quality_updater=noop,
    ),
    ItemUpdater(
        name_matcher=exact_name_matcher("Aged Brie"),
        sell_in_updater=decrease_sell_in,
        quality_updater=brie_quality_updater,
    ),
    ItemUpdater(
        name_matcher=starts_with_matcher("Backstage"),
        sell_in_updater=decrease_sell_in,
        quality_updater=backstage_quality_updater,
    ),
    ItemUpdater(
        name_matcher=starts_with_matcher("Conjured"),
        sell_in_updater=decrease_sell_in,
        quality_updater=conjured_quality_updater,
    ),
    ItemUpdater(
        name_matcher=lambda item: True,
        sell_in_updater=decrease_sell_in,
        quality_updater=default_quality_updater,
    ),
]

class GildedRose(object):

    def __init__(self, items):
        self.items = items

    def update_quality(self):
        for item in self.items:
            self.update_quality_for_item(item)
    
    def update_quality_for_item(self, item):
        for updater in ITEM_UPDATERS:
            if updater.applies_to(item):
                item = updater.update(item)
                break 

class Item:
    def __init__(self, name, sell_in, quality):
        self.name = name
        self.sell_in = sell_in
        self.quality = quality

    def __repr__(self):
        return "%s, %s, %s" % (self.name, self.sell_in, self.quality)
