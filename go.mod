module MyCache

require geecache v0.0.0
replace lru => ./geecache/lru
replace geecache => ./geecache

go 1.16
