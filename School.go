package main

import "fmt"

// kamus global
const NMAX = 1001

type tabMurid [NMAX]string

func main() {
	// kamus lokal
	var data tabMurid
	var n, x int
	var dicari string
	var ketemu bool

	// algoritma
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&data[i])
	}
	fmt.Scanln(&dicari)
	absen(data, n, dicari, &x, &ketemu)

	if ketemu {
		fmt.Printf("Murid terdaftar dan berada di urutan absen ke-%d", x)
	} else {
		fmt.Println("Murid tidak terdaftar")
	}
}

func absen(data tabMurid, n int, dicari string, x *int, ketemu *bool) {
	// kamus lokal
	var left, right int

	// algoritma
	left = 0
	right = n - 1
	*x = (left + right) / 2
	for left <= right && data[*x] != dicari {
		if dicari < data[*x] {
			right = *x - 1
		} else {
			left = *x + 1
		}
		*x = (left + right) / 2
	}
	*ketemu = *x > 0 && data[*x] == dicari
}
