package main

import (
	"sort"
)

type apartment struct {
	id                    int64
	price                 int64
	title                 string
	district              string
	distanceFromTheCenter int64
}

//var result []apartment
func sortBy(apartments []apartment, compare func(apartment, apartment) bool) []apartment {
	result := make([]apartment, len(apartments))
	copy(result, apartments)
	sort.Slice(result, func(i, j int) bool {
		return compare(result[i], result[j])
	})
	return result
}

func sortByPriceAsc(apartments []apartment) []apartment {
	return sortBy(apartments, func(a, b apartment) bool {
		return a.price < b.price
	})
}

func sortByPriceDesc(apartments []apartment) []apartment {
	return sortBy(apartments, func(a, b apartment) bool {
		return a.price > b.price
	})
}
func sortDistanceAsc(apartments []apartment) []apartment {
	return sortBy(apartments, func(a, b apartment) bool {
		return a.distanceFromTheCenter < b.distanceFromTheCenter
	})
}
func sortDistanceDesc(apartments []apartment) []apartment {
	return sortBy(apartments, func(a, b apartment) bool {
		return a.distanceFromTheCenter > b.distanceFromTheCenter
	})
}

func searchBy(apatments []apartment, find func(apartment) bool) []apartment {
	result := make([]apartment, 0)
	for _, apartment := range apatments {
		if find(apartment) {
			result = append(result, apartment)
		}
	}
	return result
}
func searchByMaxPrice(apartments []apartment, limit int64) []apartment {
	return searchBy(apartments, func(a apartment) bool {
		return a.price <= limit
	})
}
func searchByMinPrice(apartments []apartment, limit int64) []apartment {
	return searchBy(apartments, func(a apartment) bool {
		return a.price >= limit
	})
}
func searchInIntervalPrice(apartments []apartment, MinInterval int64, MaxInterval int64) []apartment {
	return searchBy(apartments, func(a apartment) bool {
		return a.price >= MinInterval && a.price <= MaxInterval
	})
}
func searchByDistricts(apartments []apartment, districts []string) []apartment {
	result := make([]apartment, 0)
	for _, apartment := range apartments {
		for _, didstrict := range districts {
			if apartment.district == didstrict {
				result = append(result, apartment)
				break
			}
		}

	}
	return result
}

func main() {}
