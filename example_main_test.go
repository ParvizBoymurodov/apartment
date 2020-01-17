package main

import "fmt"

var apartaments = []apartment{
	{id: 1,
		price:                 25_000,
		title:                 "2-х комнатная квартира",
		district:              "Фирдоуси",
		distanceFromTheCenter: 5,
	},
	{id: 2,
		price:                 45_000,
		title:                 "3-х комнатная квартира",
		district:              "Сино",
		distanceFromTheCenter: 8,
	},
	{id: 3,
		price:                 55_000,
		title:                 "2-х комнатная квартира",
		district:              "Сино",
		distanceFromTheCenter: 15,
	},
}

func ExampleSortByPriceAsc() {
	sortedPriceAsc := sortByPriceAsc(apartaments)
	fmt.Println(sortedPriceAsc)
	//Output:[{1 25000 2-х комнатная квартира Фирдоуси 5} {2 45000 3-х комнатная квартира Сино 8} {3 55000 2-х комнатная квартира Сино 15}]
}
func ExampleSortByPriceDesc() {
	sortedPriceDesc := sortByPriceDesc(apartaments)
	fmt.Println(sortedPriceDesc)
	//Output:[{3 55000 2-х комнатная квартира Сино 15} {2 45000 3-х комнатная квартира Сино 8} {1 25000 2-х комнатная квартира Фирдоуси 5}]
}
func ExampleSortDistanceFromTheCenterAsc() {
	sortedByDistanceFromTheCenterAsc := sortDistanceAsc(apartaments)
	fmt.Println(sortedByDistanceFromTheCenterAsc)
	//Output:[{1 25000 2-х комнатная квартира Фирдоуси 5} {2 45000 3-х комнатная квартира Сино 8} {3 55000 2-х комнатная квартира Сино 15}]
}
func ExampleSortDistanceFromTheCenterDesc() {
	sortedByDistanceFromTheCenterDesc := sortDistanceDesc(apartaments)
	fmt.Println(sortedByDistanceFromTheCenterDesc)
	//Output:[{3 55000 2-х комнатная квартира Сино 15} {2 45000 3-х комнатная квартира Сино 8} {1 25000 2-х комнатная квартира Фирдоуси 5}]
}
func ExampleSearchByMaxPrice_NoResult() {
	searchMaxPrice := searchByMaxPrice(apartaments, 0)
	fmt.Println(searchMaxPrice)
	//Output:[]
}
func ExampleSearchByMaxPrice_OneResult() {
	searchMaxPrice := searchByMaxPrice(apartaments, 30_000)
	fmt.Println(searchMaxPrice)
	//Output:[{1 25000 2-х комнатная квартира Фирдоуси 5}]
}
func ExampleSearchByMaxPrice_ManyResult() {
	searchMaxPrice := searchByMaxPrice(apartaments, 200_000)
	fmt.Println(searchMaxPrice)
	//Output:[{1 25000 2-х комнатная квартира Фирдоуси 5} {2 45000 3-х комнатная квартира Сино 8} {3 55000 2-х комнатная квартира Сино 15}]
}
func ExampleSearchByMinPrice_NoResult() {
	searchMinPrice := searchByMinPrice(apartaments, 0)
	fmt.Println(searchMinPrice)
	//Output:[]
}

func ExampleSearchByMinPrice_OneResult() {
	searchMinPrice := searchByMinPrice(apartaments, 25_000)
	fmt.Println(searchMinPrice)
	//Output:[{1 25000 2-х комнатная квартира Фирдоуси 5}]
}
func ExampleSearchInIntervalPrice() {
	searchIntervalPrice := searchInIntervalPrice(apartaments, 30_000, 400_000)
	fmt.Println(searchIntervalPrice)
	//Output:[{2 45000 3-х комнатная квартира Сино 8} {3 55000 2-х комнатная квартира Сино 15}]
}
func ExampleSearchInIntervalPrice_OneResult() {
	searchIntervalPrice := searchInIntervalPrice(apartaments, 25_000, 40_000)
	fmt.Println(searchIntervalPrice)
	//Output:[{1 25000 2-х комнатная квартира Фирдоуси 5}]
}
func ExampleSearchInIntervalPrice_NoResult() {
	searchIntervalPrice := searchInIntervalPrice(apartaments, 0, 0)
	fmt.Println(searchIntervalPrice)
	//Output:[]
}
func ExampleSearchByDistrict_OneDistrict() {
	searchDistricts := searchByDistricts(apartaments, []string{"Сино"})
	fmt.Println(searchDistricts)
	//Output:[{2 45000 3-х комнатная квартира Сино 8} {3 55000 2-х комнатная квартира Сино 15}]
}
func ExampleSearchByDistrict_NoDistrict() {
	searchDistricts := searchByDistricts(apartaments, []string{"Рудаки"})
	fmt.Println(searchDistricts)
	//Output:[]
}
func ExampleSearchByDistrict_ManyDistrict() {
	searchDistricts := searchByDistricts(apartaments, []string{"Сино", "Фирдоуси"})
	fmt.Println(searchDistricts)
	//Output:[{1 25000 2-х комнатная квартира Фирдоуси 5} {2 45000 3-х комнатная квартира Сино 8} {3 55000 2-х комнатная квартира Сино 15}]
}
