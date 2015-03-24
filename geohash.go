package geohash

import (
	"bytes"
	"strings"
)

var Base32Map = []byte("0123456789bcdefghjkmnpqrstuvwxyz")

type BoundingBox struct {
	MinLatitude  float64
	MaxLatitude  float64
	MinLongitude float64
	MaxLongitude float64
}

type Coords struct {
	Latitude  float64
	Longitude float64
}

func Decode(geohash string) *Coords {
	bbox := DecodeBoundingBox(geohash)

	return &Coords{
		Latitude:  (bbox.MinLatitude + bbox.MaxLatitude) / 2,
		Longitude: (bbox.MinLongitude + bbox.MaxLongitude) / 2,
	}
}

func DecodeBoundingBox(geohash string) *BoundingBox {
	var cd int
	geohash = strings.ToLower(geohash)

	isEven := true
	bits := []int{16, 8, 4, 2, 1}
	bbox := &BoundingBox{MinLatitude: -90.0, MaxLatitude: 90.0, MinLongitude: -180.0, MaxLongitude: 180.0}

	for _, code := range geohash {
		cd = bytes.Index(Base32Map, []byte(string(code)))

		for _, mask := range bits {
			if isEven {
				bbox.calcBboxRange(cd, mask, true)
			} else {
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
			geohash.WriteByte(Base32Map[hashValue])
			bit = 0
			hashValue = 0
		}
	}
	return geohash.String()
}

func Neighbors(geohash string) []string {
	v := []string{}
	directions := []string{"top", "right", "bottom", "left"}

	for _, direction := range directions {
		neighbor := Neighbor(geohash, direction)
		v = append(v, neighbor)

		if direction == "top" || direction == "bottom" {
			v = append(v, Neighbor(neighbor, "right"))
			v = append(v, Neighbor(neighbor, "left"))
		}
	}
	return v
}

func Neighbor(geohash string, dir string) string {
	neighbors := [][]string{
		[]string{
			"p0r21436x8zb9dcf5h7kjnmqesgutwvy",
			"bc01fg45238967deuvhjyznpkmstqrwx",
		},
		[]string{
			"bc01fg45238967deuvhjyznpkmstqrwx",
			"p0r21436x8zb9dcf5h7kjnmqesgutwvy",
		},
		[]string{
			"14365h7k9dcfesgujnmqp0r2twvyx8zb",
			"238967debc01fg45kmstqrwxuvhjyznp",
		},
		[]string{
			"238967debc01fg45kmstqrwxuvhjyznp",
			"14365h7k9dcfesgujnmqp0r2twvyx8zb",
		},
	}
	borders := [][]string{
		[]string{
			"prxz",
			"bcfguvyz",
		},
		[]string{
			"bcfguvyz",
			"prxz",
		},
		[]string{
			"028b",
			"0145hjnp",
		},
		[]string{
			"0145hjnp",
			"028b",
		},
	}

	var intDir int

	switch strings.ToLower(dir) {
	case "top":
	default:
		intDir = 0
	case "right":
		intDir = 1
	case "bottom":
		intDir = 2
	case "left":
		intDir = 3
	}

	hashLen := len(geohash)
	lastChar := geohash[hashLen-1:]
	isEven := hashLen % 2
	base := geohash[:hashLen-1]

	border := borders[intDir][isEven]
	neighbor := neighbors[intDir][isEven]

	if strings.Index(border, lastChar) != -1 {
		base = Neighbor(base, dir)
	}

	return base + string(Base32Map[strings.Index(neighbor, lastChar)])
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
