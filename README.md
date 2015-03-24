[![Build Status](https://travis-ci.org/fanixk/geohash.svg?branch=travis)](https://travis-ci.org/fanixk/geohash)

## Geohash for Go

#### Usage
Encode:
```
geohash := geohash.Encode(37.941882, 23.653022)
```
Decode in BoundingBox
```
bbox := geohash.DecodeBoundingBox("sw8zf5pe7r7w")
```
Decode:
```
coords := geohash.Decode("sw8zf5pe7r7w")
```
Get Neighbor:
```
neighbor := geohash.Neighbor("sw8zf5pe7r7w", "left")
```
Get All Neighbors:
```
neighbors := geohash.Neighbors("sw8zf5pe7r7w")
```
http://en.wikipedia.org/wiki/Geohash
