package models

type Pagination struct {
	Previous	int
	Next		int
	Last		int
	Count		int
	Middle		[5]int
}
