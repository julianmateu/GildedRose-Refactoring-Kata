# -*- coding: utf-8 -*-
import pytest

from gilded_rose import Item, GildedRose


def test_sulfuras_should_not_change():
    items = [Item("Sulfuras, Hand of Ragnaros", 60, 80)]
    gilded_rose = GildedRose(items)
    gilded_rose.update_quality()
    assert 60 == items[0].sell_in
    assert 80 == items[0].quality

def test_regular_item_quality_and_sell_in_decreases_by_1_if_not_expired():
    items = [Item("Regular Item", 10, 20)]
    gilded_rose = GildedRose(items)
    gilded_rose.update_quality()
    assert 9 == items[0].sell_in
    assert 19 == items[0].quality

@pytest.mark.parametrize("sell_in", [
    -2, 0
])
def test_regular_item_quality_decreases_by_2_if_expired(sell_in):
    items = [Item("Regular Item", sell_in, 20)]
    gilded_rose = GildedRose(items)
    gilded_rose.update_quality()
    assert sell_in-1 == items[0].sell_in
    assert 18 == items[0].quality

@pytest.mark.parametrize("sell_in", [
    2, -1, 0
])
def test_regular_item_quality_never_negative(sell_in):
    items = [Item("Regular Item", sell_in, 0)]
    gilded_rose = GildedRose(items)
    gilded_rose.update_quality()
    assert sell_in - 1 == items[0].sell_in
    assert 0 == items[0].quality

def test_aged_brie_quality_increases_by_1_if_not_expired():
    items = [Item("Aged Brie", 10, 20)]
    gilded_rose = GildedRose(items)
    gilded_rose.update_quality()
    assert 9 == items[0].sell_in
    assert 21 == items[0].quality
    
def test_aged_brie_quality_increases_by_2_if_expired():
    items = [Item("Aged Brie", -2, 20)]
    gilded_rose = GildedRose(items)
    gilded_rose.update_quality()
    assert -3 == items[0].sell_in
    assert 22 == items[0].quality

@pytest.mark.parametrize("sell_in,quality", [
    (2, 50),
    (-1, 50),
    (0, 50),
    (2, 49),
    (-1, 49),
    (0, 49),
])
def test_aged_brie_quality_stays_at_fifty(sell_in, quality):
    items = [Item("Aged Brie", sell_in, quality)]
    gilded_rose = GildedRose(items)
    gilded_rose.update_quality()
    assert sell_in-1 == items[0].sell_in
    assert 50 == items[0].quality

@pytest.mark.parametrize("sell_in", [
    12, 11
])
def test_tickets_increase_quality_by_one_if_gte_11_days(sell_in):
    items = [Item("Backstage passes to a TAFKAL80ETC concert", sell_in, 20)]
    gilded_rose = GildedRose(items)
    gilded_rose.update_quality()
    assert sell_in - 1 == items[0].sell_in
    assert 21 == items[0].quality

@pytest.mark.parametrize("sell_in", [
    6, 7, 10
])
def test_tickets_increase_quality_by_2_if_between_6_and_11_days(sell_in):
    items = [Item("Backstage passes to a TAFKAL80ETC concert", sell_in, 20)]
    gilded_rose = GildedRose(items)
    gilded_rose.update_quality()
    assert sell_in - 1 == items[0].sell_in
    assert 22 == items[0].quality

@pytest.mark.parametrize("sell_in", [
    5, 1
])
def test_tickets_increase_quality_by_3_if_lt_6_days(sell_in):
    items = [Item("Backstage passes to a TAFKAL80ETC concert", sell_in, 20)]
    gilded_rose = GildedRose(items)
    gilded_rose.update_quality()
    assert sell_in - 1 == items[0].sell_in
    assert 23 == items[0].quality

@pytest.mark.parametrize("sell_in", [
    0, -2
])
def test_tickets_quality_is_0_if_expired(sell_in):
    items = [Item("Backstage passes to a TAFKAL80ETC concert", sell_in, 20)]
    gilded_rose = GildedRose(items)
    gilded_rose.update_quality()
    assert sell_in - 1 == items[0].sell_in
    assert 0 == items[0].quality

def test_conjured_item_quality_decreases_by_2_if_not_expired():
    items = [Item("Conjured Item", 10, 20)]
    gilded_rose = GildedRose(items)
    gilded_rose.update_quality()
    assert 9 == items[0].sell_in
    assert 18 == items[0].quality

def test_conjured_item_quality_decreases_by_4_if_expired():
    items = [Item("Conjured Item", -1, 20)]
    gilded_rose = GildedRose(items)
    gilded_rose.update_quality()
    assert -2 == items[0].sell_in
    assert 16 == items[0].quality

@pytest.mark.parametrize("sell_in", [
    2, -1, 0
])
def test_conjured_item_quality_never_negative(sell_in):
    items = [Item("Regular Item", sell_in, 1)]
    gilded_rose = GildedRose(items)
    gilded_rose.update_quality()
    assert sell_in - 1 == items[0].sell_in
    assert 0 == items[0].quality