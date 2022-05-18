module MyCache

require lru v0.0.0
replace lru => ./geecache/lru
require geecache v0.0.0
replace geecache => ./geecache

go 1.16
