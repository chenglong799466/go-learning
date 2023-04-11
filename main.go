package main

import (
	"fmt"
	"math"
)

func main() {

	n := 360
	year_month := 12
	p := 100
	r1 := float64(0.0490) / float64(year_month)
	r2 := float64(0.0325) / float64(year_month)
	a := get_pmt(r1, n, p)*float64(n) - float64(p)
	b := get_pmt(r2, n, p)*float64(n) - float64(p)
	fmt.Println(a, b)

}

func get_pmt(f_interest_rate float64, term_number int, principal int) float64 {
	compound_rate := math.Pow(1+f_interest_rate, float64(term_number))
	pmt := float64(principal) * f_interest_rate * compound_rate / (compound_rate - 1)
	return pmt
}
func get_month_provide(price int) {
	n := 360
	year_month := 12
	gongjijin_loan_limit := 900000
	shoufu_rate := 0.3
	r1 := float64(0.0490) / float64(year_month)
	r2 := float64(0.0325) / float64(year_month)
	month_provide := get_pmt(r2, n, gongjijin_loan_limit)
	month_provide += get_pmt(r1, n, int(float64(price)*(float64(1)-float64(shoufu_rate))-float64(gongjijin_loan_limit)))
	fmt.Println(month_provide)
}
