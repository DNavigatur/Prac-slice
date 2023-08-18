package main

import "fmt"

/*
func main() {

	// literal array
	a := [3]int{45, 89, 56}
	fmt.Println(a)

	//another way of assinging values to an array
	b := [...]string{"Henry", "Samuel"}
	fmt.Println(b)

	//aonther way
	var c [2]int
	c[0] = 45
	c[1] = 93
	fmt.Println(c)

	// we can use len to check the lenght of the array
	fmt.Println(len(c))

}


//array execise
func main() {
	a := [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	fmt.Println(a)
	fmt.Println("The number of Icecream flavour", len(a))
}

func main() {
	xs := []string{"hello", "Henry"}
	fmt.Println(xs)
}

//slice data structure
func main() {
	xf := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}

	for index, value := range xf {
		fmt.Printf("key is %v\t value is %v\n", index, value)
	}
	fmt.Println(len(xf))
}

func main() {
	xf := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	for _, v := range xf {
		fmt.Printf("%v\n", v)
	}
}

//append values
func main() {
	xi := []int{23, 86, 90}
	fmt.Println(xi)

	//variadic parameter
	xi = append(xi, 87, 90, 46)
	fmt.Println(xi)
}

//slicing a slice
func main() {
	xi := []int{23, 86, 96, 41, 30, 62, 90}
	fmt.Printf("%v ", xi[1:4])
}

func main() {
	xi := []int{23, 86, 96, 41, 30, 62, 90}
	fmt.Printf("%v", xi)

	// this how we deleted 41
	// we append a new version of the slice with a slice r from the begining and the ending
	xi = append(xi[:3], xi[4:]...)
	fmt.Printf("%v", xi)

}

// making a slice
func main() {
	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	xi = append(xi, 0, 1, 2, 3, 4, 5)
	fmt.Println(xi)
}

//multi-dimensional slice
func main() {
	eh := []string{"Eze", "Henry", "Puff Puff"}
	hp := []string{"Hewlet", "Pakard", "Akara"}
	fmt.Println(eh)
	fmt.Println(hp)

	xxs := [][]string{eh, hp}
	fmt.Println(xxs)
}

//copy in slice and how to manipulate slice
func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := make([]int, 6)

	copy(b, a)

	fmt.Println(a)
	fmt.Println(b)

}

//Hand-exercise 1
func main() {
	var a [5]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5

	for i, v := range a {
		fmt.Printf("Index %v\t Value %v\n", i, v)
	}
}

//Hand-exercise 2
func main() {
	xa := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, v := range xa {
		fmt.Printf("Value %v\t Type %T\n", v, v)
	}
}

//Hand-exercise 3
func main() {
	xa := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("1st one %v\t\n", xa[:5])
	fmt.Printf("2nd one %v\t\n", xa[5:])
	fmt.Printf("3rd one %v\t\n", xa[2:7])
	fmt.Printf("4th one %v\t\n", xa[1:6])

}

func main() {
	xa := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	xa = append(xa, 52, 53, 54, 55)
	fmt.Printf("%v", xa)
	xy := []int{56, 57, 58, 59, 60}

	xxay := [][]int{xa, xy}

	fmt.Printf("%v", xxay)
}

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	xa := append(x[:3], x[6:]...)

	fmt.Printf("%v", xa)
}

//Hands-exercise
func main() {

	//xs := [50]string{}
	xsn := make([]string, 0, 50)
	xsn = append(xsn, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Println(len(xsn))
	fmt.Println(cap(xsn))
	for i := 0; i < len(xsn); i++ {
		fmt.Println(xsn[i], i)
	}

}
*/

func main() {
	x1 := []string{"James", "Bond", "Shaken", "not stirred"}
	x2 := []string{"Miss", "Moneypenny", "I'm 008"}

	xa := [][]string{x1, x2}

	for i, v := range xa {
		fmt.Printf("%v, %v\n", i, v)
	}
}
