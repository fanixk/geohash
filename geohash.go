package geohash

import (
	"bytes"
	"strings"
)

var BASE32 = []byte("0123456789bcdefghjkmnpqrstuvwxyz")

type BoundingBox struct {
	MinLatitude  float64
	MaxLatitude  float64
	MinLongitude float64
	MaxLongitude float64
}

func (bbox *BoundingBox) calcBboxRange(cd int, mask int, isLon bool) {
	if isLon {
		lon := (bbox.MinLongitude + bbox.MaxLongitude) / 2
		if cd&mask > 0 {
			bbox.MinLongitude = lon
		} else {
			bbox.MaxLongitude = lon
		}
	} else {
		lat := (bbox.MinLatitude + bbox.MaxLatitude) / 2
		if cd&mask > 0 {
			bbox.MinLatitude = lat
		} else {
			bbox.MaxLatitude = lat
		}
	}
}

func DecodeBoundingBox(geohash string) *BoundingBox {
	var (
		code string
		cd   int
		mask int
	)
	hashLen := len(geohash)
	isEven := true
	latErr := 90.0
	lonErr := 180.0
	bits := []int{16, 8, 4, 2, 1}
	bbox := &BoundingBox{MinLatitude: -90.0, MaxLatitude: 90.0, MinLongitude: -180.0, MaxLongitude: 180.0}
	geohash = strings.ToLower(geohash)

	for i := 0; i < hashLen; i++ {
		code = geohash[i : i+1]
		cd = bytes.Index(BASE32, []byte(code))

		for j := 0; j < 5; j++ {
			mask = bits[j]
			if isEven {
				lonErr /= 2
				bbox.calcBboxRange(cd, mask, true)
			} else {
				latErr /= 2
				bbox.calcBboxRange(cd, mask, false)
			}
			isEven = !isEven
		}
	}

	return bbox
}

func Encode(lat float64, long float64) string {
	return PrecisionEncode(lat, long, 12)
}

func PrecisionEncode(latitude float64, longitude float64, precision int) string {
	var geohash bytes.Buffer
	var mid float64
	lat := []float64{-90.0, 90.0}
	long := []float64{-180.0, 180.0}
	hashValue := 0
	bit := 0
	isEven := true

	for geohash.Len() < precision {
		if isEven {
			mid = (long[0] + long[1]) / 2
			if longitude > mid {
				hashValue = (hashValue << 1) + 1
				long[0] = mid
			} else {
				hashValue = (hashValue << 1)
				long[1] = mid
			}
		} else {
			mid = (lat[0] + lat[1]) / 2
			if latitude > mid {
				hashValue = (hashValue << 1) + 1
				lat[0] = mid
			} else {
				hashValue = (hashValue << 1)
				lat[1] = mid
			}
		}

		isEven = !isEven
		if bit < 4 {
			bit++
		} else {
			geohash.WriteByte(BASE32[hashValue])
			bit = 0
			hashValue = 0
		}
	}
	return geohash.String()
}
