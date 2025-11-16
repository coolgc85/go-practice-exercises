package main

import (
	"fmt"
	"slices"
)

type Book struct {
	title    string
	author   string
	price    float64
	category string
}

const extraDiscount float64 = 5.0

func (book *Book) Discount(percent float64) {

	discountedCategories := getDiscountedCategories()
	hasExtraDiscount := slices.Contains(discountedCategories, book.category)

	if hasExtraDiscount {
		percent += extraDiscount
		fmt.Printf("Applying extra discount of %.2f%% for the category %s \n", extraDiscount, book.category)
	}
	discountAmount := book.price * (percent / 100)
	book.price -= discountAmount
}

func (b Book) String() string {
	return fmt.Sprintf("Title: %s, Author: %s, Category: %s, Price: %.2f", b.title, b.author, b.category, b.price)
}

func getDiscountedCategories() []string {
	return []string{"fictions", "fantasy", "sci-fi", "computer-science"}
}

func main() {

	const usualDiscount = 30.0

	favorite_book := Book{
		title:    "Song of fire and ice",
		author:   "RR Martin",
		price:    40.30,
		category: "fantasy",
	}

	fmt.Println("================================")
	fmt.Println("Original book details:", favorite_book)
	fmt.Printf("Applying a %.2f%% usual discount...\n", usualDiscount)
	favorite_book.Discount(usualDiscount)
	fmt.Println("Book details after discount:", favorite_book)

	new_book := Book{
		title:    "Time Management",
		author:   "John Doe",
		price:    40.30,
		category: "Personal Development",
	}

	fmt.Println("================================")
	fmt.Println("Original book details:", new_book)
	fmt.Printf("Applying a %.2f%% usual discount...\n", usualDiscount)
	new_book.Discount(usualDiscount)
	fmt.Println("Book details after discount:", new_book)
}
