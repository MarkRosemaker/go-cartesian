package points

import (
	"sort"

	"github.com/MarkRosemaker/go-cartesian/math/point"
)

// WithDistance is a slice of containers that each hold a point and its distance to the search origin point.
type WithDistance []point.WithDistance

// Points returns just the points, without the distance to the origin point.
func (pwd WithDistance) Points() Points {
	k := len(pwd)

	o := make(Points, k)
	for i := 0; i < k; i++ {
		o[i] = pwd[i].Point
	}
	return o
}

// less is defined to sort this slice
func (pwd WithDistance) less(i int, j int) bool {
	return pwd[i].Distance < pwd[j].Distance
}

// sortCut sorts the entire list and then cuts away all points that are not within the radius ("distance").
func (pwd WithDistance) sortCut(radius int) WithDistance {

	// sort the list
	sort.SliceStable(pwd, pwd.less)

	// cut off all points that are too far away
	k := len(pwd)
	for i := 0; i < k; i++ {
		if pwd[i].Distance > radius {
			return pwd[:i]
		}
	}

	return pwd
}

// collectSort first collects all points that are within the radius ("distance") and then sorts the collection.
func (pwd WithDistance) collectSort(radius int) WithDistance {

	k := len(pwd)
	list := make(WithDistance, 0) // todo find good capacity value for efficiency

	// collect all points within radius
	for i := 0; i < k; i++ {
		if pwd[i].Distance <= radius {
			list = append(list, pwd[i])
		}
	}

	// sort the list
	sort.SliceStable(list, list.less)

	return list
}
