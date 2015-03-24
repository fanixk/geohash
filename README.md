[![Build Status](https://travis-ci.org/fanixk/geohash.svg?branch=travis)](https://travis-ci.org/fanixk/geohash)

# Geohash for Go

# Usage

geohash := geohash.Encode(37.941882, 23.653022)

bbox := geohash.DecodeBoundingBox("sw8zf5pe7r7w")

coords := geohash.Decode("sw8zf5pe7r7w")

neighbor := geohash.Neighbor("sw8zf5pe7r7w", "left")

neighbors := geohash.Neighbors("sw8zf5pe7r7w")

http://en.wikipedia.org/wiki/Geohash
