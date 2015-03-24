package geohash

import (
	"reflect"
	"sort"
	"testing"
)

type encodeTest struct {
	coords  Coords
	geohash string
}

func TestEncode(t *testing.T) {
	var tests = []encodeTest{
		encodeTest{
			Coords{45.37, -121.7},
			"c216nekg2kyz",
		},
		encodeTest{
			Coords{47.6062095, -122.3320708},
			"c23nb62w20st",
		},
		encodeTest{
			Coords{35.6894875, 139.6917064},
			"xn774c06kdtv",
		},
		encodeTest{
			Coords{-33.8671390, 151.2071140},
			"r3gx2f9tt5sn",
		},
		encodeTest{
			Coords{51.5001524, -0.1262362},
			"gcpuvpk44kpr",
		},
	}

	prGeohash := PrecisionEncode(45.37, -121.7, 6)
	if prGeohash != "c216ne" {
		t.Errorf("Expected %s, got %s", "c216ne", prGeohash)
	}

	for _, test := range tests {
		geohash := Encode(test.coords.Latitude, test.coords.Longitude)
		if test.geohash != geohash {
			t.Errorf("Expected %s, got %s", test.geohash, geohash)
		}
	}
}

type decodeTest struct {
	bbox    *BoundingBox
	geohash string
}

func TestDecode(t *testing.T) {
	var tests = []decodeTest{
		decodeTest{
			&BoundingBox{
				MinLatitude:  45.3680419921875,
				MinLongitude: -121.70654296875,
				MaxLatitude:  45.37353515625,
				MaxLongitude: -121.695556640625,
			},
			"c216ne",
		},
		decodeTest{
			&BoundingBox{
				MinLatitude:  45.3680419921875,
				MinLongitude: -121.70654296875,
				MaxLatitude:  45.37353515625,
				MaxLongitude: -121.695556640625,
			},
			"C216Ne",
		},
		decodeTest{
			&BoundingBox{
				MinLatitude:  39.0234375,
				MinLongitude: -76.552734375,
				MaxLatitude:  39.0673828125,
				MaxLongitude: -76.5087890625,
			},
			"dqcw4",
		},
		decodeTest{
			&BoundingBox{
				MinLatitude:  39.0234375,
				MinLongitude: -76.552734375,
				MaxLatitude:  39.0673828125,
				MaxLongitude: -76.5087890625,
			},
			"DQCW4",
		},
	}

	for _, test := range tests {
		bbox := DecodeBoundingBox(test.geohash)
		if !helperEqualBbox(test.bbox, bbox) {
			t.Errorf("Expected %s, got %s", test.bbox, bbox)
		}
	}
}

func helperEqualBbox(bbox1, bbox2 *BoundingBox) bool {
	isEqual := bbox1.MinLatitude == bbox2.MinLatitude &&
		bbox1.MaxLatitude == bbox2.MaxLatitude &&
		bbox1.MinLongitude == bbox2.MinLongitude &&
		bbox1.MaxLongitude == bbox2.MaxLongitude
	return isEqual
}

type neighborTest struct {
	dir     string
	geohash string
}

func TestNeighbor(t *testing.T) {
	var tests = []neighborTest{
		neighborTest{
			"top",
			"dqcjw",
		},
		neighborTest{
			"bottom",
			"dqcjn",
		},
		neighborTest{
			"left",
			"dqcjm",
		},
		neighborTest{
			"right",
			"dqcjr",
		},
	}

	for _, test := range tests {
		neighbor := Neighbor("dqcjq", test.dir)
		if neighbor != test.geohash {
			t.Errorf("Expected %s, got %s", test.geohash, neighbor)
		}
	}
}

type neighborsTest struct {
	geohash   string
	neighbors []string
}

func TestNeighbors(t *testing.T) {
	var tests = []neighborsTest{
		neighborsTest{
			"dqcw5",
			[]string{"dqcw7", "dqctg", "dqcw4", "dqcwh", "dqcw6", "dqcwk", "dqctf", "dqctu"},
		},
		neighborsTest{
			"xn774c",
			[]string{"xn774f", "xn774b", "xn7751", "xn7749", "xn774d", "xn7754", "xn7750", "xn7748"},
		},
		neighborsTest{
			"gcpuvpk",
			[]string{"gcpuvps", "gcpuvph", "gcpuvpm", "gcpuvp7", "gcpuvpe", "gcpuvpt", "gcpuvpj", "gcpuvp5"},
		},
		neighborsTest{
			"c23nb62w",
			[]string{"c23nb62x", "c23nb62t", "c23nb62y", "c23nb62q", "c23nb62r", "c23nb62z", "c23nb62v", "c23nb62m"},
		},
	}

	for _, test := range tests {
		neighbors := Neighbors(test.geohash)
		sort.Strings(neighbors)
		sort.Strings(test.neighbors)

		if !reflect.DeepEqual(neighbors, test.neighbors) {
			t.Errorf("Expected %s, got %s", test.neighbors, neighbors)
		}
	}
}
