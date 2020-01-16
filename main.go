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

func sortByPriceAsc(apartments []apartment) []apartment {
	result := make([]apartment, len(apartments))
	copy(result, apartments)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price < result[j].price
	})
	return result
}
func sortByPriceDesc(apartments []apartment) []apartment {
	result := make([]apartment, len(apartments))
	copy(result, apartments)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price > result[j].price
	})
	return result
}
func sortDistanceAsc(apartments []apartment) []apartment {
	result := make([]apartment, len(apartments))
	copy(result, apartments)
	sort.Slice(result, func(i, j int) bool {
		return result[i].distanceFromTheCenter < result[j].distanceFromTheCenter
	})
	return result
}
func sortDistanceDesc(apartments []apartment) []apartment {
	result := make([]apartment, len(apartments))
	copy(result, apartments)
	sort.Slice(result, func(i, j int) bool {
		return result[i].distanceFromTheCenter > result[j].distanceFromTheCenter
	})
	return result
}
func searchByMaxPrice(apartments []apartment, limit int64) []apartment {
	result := make([]apartment, 0)
	for _, apartment := range apartments {
		if apartment.price <= limit {
			result = append(result, apartment)
		}
	}
	return result
}
func searchByMinPrice(apartments []apartment, limit int64) []apartment {
	result := make([]apartment, 0)
	for _, apartment := range apartments {
		if apartment.price <= limit {
			result = append(result, apartment)
		}
	}
	return result
}
func searchInIntervalPrice(apartments []apartment, MinInterval int64, MaxInterval int64) []apartment {
	result := make([]apartment, 0)
	for _, apartment := range apartments {
		if apartment.price >= MinInterval && apartment.price <= MaxInterval {
			result = append(result, apartment)
		}
	}
	return result
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
