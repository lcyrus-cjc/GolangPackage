package main

//var x int = 10
//const y int = 20
/*
type person struct {
	name string
}

func main() {

	p1 := person{
		name: "Bob",
	}

	p2 := person{
		name: "Tim",
	}

	p1.people()
	p2.people()

}

func (n person) people() {
	fmt.Println("My name is", n.name)
}


func main() {

	fmt.Println("Testing")
	test(10, 20)

		lads := struct {  //anonymous Structs
			first     string
			friends   map[string]int
			favDrinks []string
		}{
			first: "Glen",
			friends: map[string]int{
				"Todd": 28,
				"Tim":  27,
				"Ted":  19,
			},
			favDrinks: []string{"Cola", "Coffee"},
		}
		fmt.Println(lads)

		type engine struct { // Embeded Structs
			electric bool
		}

		type vehicle struct {
			engine
			make   string
			model  string
			doors  int
			colour string
		}

		p1 := vehicle{
			engine: engine{
				electric: true,
			},
			make:   "Hyundi",
			model:  "X21",
			doors:  4,
			colour: "Red",
		}
		p2 := vehicle{
			engine: engine{
				electric: false,
			},
			make:   "Ford",
			model:  "S23",
			doors:  2,
			colour: "Black",
		}

		fmt.Println(p1, p2)

		fmt.Println(p1.colour)
		fmt.Println(p2.colour)


		type person struct {
			first string
			last  string
			favIc []string //Slice of string
		}
		// Values of Type Person (p1 and p2)

		p1 := person{
			first: "Peter",
			last:  "Price",
			favIc: []string{"Vanilla", "Chocolate", "Cookie Dough"},
		}

		p2 := person{
			first: "Harry",
			last:  "Bland",
			favIc: []string{"Strawberry", "Hazel Nut", "Mint"},
		}

		fmt.Println(p1)
		fmt.Println(p2)

		// Print out the values, ranging over the elements in the slice which stores the favorite flavors.

		for _, v := range p1.favIc {
			fmt.Println(p1.first, "favourite Ice cream flavours is", v)

		}

		for _, v := range p2.favIc {
			fmt.Println(p2.first, "favourite Ice cream flavour is", v)
		}
		//Take the code from the previous exercise, then store the VALUES of type person in a map with the KEY of last name. Access each value in the map. Print out the values, ranging over the slice.
		m := map[string]person{
			p1.last: p1,
			p2.last: p2,
		}

		for _, v := range p1.favIc {
			fmt.Println(v)
		}

		for _, v := range p2.favIc {
			fmt.Println(v)
		}
		for _, v := range m {
			fmt.Println(v)
			for _, v2 := range v.favIc {
				fmt.Println(v.first, v.last, v2)
			}

		}



			//Use a for range loop to loop over these words, then count how many times each word occurs.

			x := []string{"in", "my", "younger", "and", "more", "vulnerable", "years", "my", "father", "gave", "me", "some", "advice", "that", "i’ve", "been", "turning", "over", "in", "my", "mind", "ever", "since.", "whenever"}

			m := make(map[string]int) // create empty map with correct Key and Value types

			for _, v := range x { //loop over spilce and increase value for word count
				m[v]++
			}

			for k, v := range m { // print each index and value to display word count
				fmt.Println(k, v)
			}


					// Creating empty int map, then creating and assigning value to each. Multiple lines of code
							m := map[string]int{}
							m[`bond_james`] = 2
							m[`moneypenny_jenny`] = 3
							m[`no_dr`] = 4
							fmt.Println(m)

						m := map[string][]string{}
						m[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
						m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
						m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}

					// creating a map with key and values assigned
					m := map[string]int{
						`bond_james`:       2,
						`moneypenny_jenny`: 3,
						`no_dr`:            4}

				m := map[string][]string{
					`bond_james`:       {`shaken, not stirred`, `martinis`, `fast cars`},
					`moneypenny_jenny`: {`intelligence`, `literature`, `computer science`},
					`no_dr`:            {`cats`, `ice cream`, `sunsets`},
				}
				fmt.Println(m)


					// Map of String
					n := make(map[string]string)
					n[`bond_james`] = `2`
					n[`moneypenny_jenny`] = `3`
					n[`no_dr`] = `4`

					//
					m := make(map[string][]string)
					m[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
					m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
					m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}




					for i, v := range m {
						fmt.Printf("The Index is %v and the Value is %v\n", i, v)
					}



						//Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice

						x := []string{"James", "Bond", "Shaken, not stirred"}
						y := []string{"Miss", "Moneypenny", "I'm 008."}

						fmt.Println(x)
						fmt.Println(y)

						z := [][]string{x, y}
						fmt.Println(z)

						//Range over the records, then range over the data in each record

						for i, v := range z {
							fmt.Println(i, v)
							for a, b:= range v {
								fmt.Println(a, b)
							}
						}


							//Append Slice
							x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

							x = append(x, 52)
							fmt.Println(x)

							x = append(x, 53, 54, 55)
							fmt.Println(x)

							y := []int{56, 57, 58, 59, 60}

							x = append(x, y...)
							fmt.Println(x)

							// Delete Slice

							z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
							fmt.Println(z)

							z = append(z[:3], z[6:]...)
							fmt.Println(z)

							//Make Slice

							// Create a slice to store the names of all of the states in the United States of America
							// Use make and append to do this
							// Goal: do not have the array that underlies the slice created more than once.

							d := make([]string, 0, 50)
							fmt.Println(d)
							fmt.Println(len(d))
							fmt.Println(cap(d))

							d = append(d, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

							fmt.Println(len(d))
							fmt.Println(cap(d))

							for i := 0; i < len(d); i++ {
								fmt.Println(d[i], i)
							}



								// create a SLICE of TYPE int
								y := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
								fmt.Println(y)

								//assign these 10 VALUES
								// Doesnt work in an empty array as theparameters for the array would not be set
								//y[0] = 42
								//y[1] = 43
								//y[2] = 44
								//y[3] = 45
								//y[4] = 46
								//y[5] = 47
								//y[6] = 48
								//y[7] = 49
								//y[8] = 50
								//y[9] = 51
								//fmt.Println(y)

								// Range over the slice and print the values out.
								for _, v := range y {
									fmt.Printf("The type is %T and the Value is %v\n", v, v)
								}

								// create the following new slices which are then printed:
								//[42 43 44 45 46]
								y = y[0:5]
								fmt.Println(y)

								// [47 48 49 50 51]
								y = y[5:]
								fmt.Println(y)

								// [44 45 46 47 48]
								y = y[2:7]
								fmt.Println(y)

								// [43 44 45 46 47]
								y = y[1:6]
								fmt.Println(y)


									// create an ARRAY which holds 5 VALUES of TYPE int
									x := [5]int{}
									fmt.Println(x)

									// assign VALUES to each index position.

									x[0] = 10
									x[1] = 20
									x[2] = 30
									x[3] = 40
									x[4] = 50
									fmt.Println(x)

									// Range over the array and print the values out.
									for _, v := range x {
										fmt.Printf("The type is %T and the value is %v\n", v, v)
									}

									//x = append(x, 1, 2, 3, 4)

									a := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}

									fmt.Println(a)
									fmt.Println(len(a))
									for k, v := range a {
										fmt.Printf("The key unit is %v  and the Value unit is %v\n", k, v)
									}

									fmt.Println(true && true)
									fmt.Println(true && false)
									fmt.Println(true || true)
									fmt.Println(true || false)
									fmt.Println(!true)

									for i := 0; i < 100; i++ {
										fmt.Printf("Loop number %v\n", i)

										if x := rand.Intn(5); x == 3 {
											fmt.Printf("x is 3\n")
										}
									}

									m := map[string]int{
											"James":      42,
											"Moneypenny": 32,
										}

										for k, v := range m {
											fmt.Printf("The Key is %v and the value is %v\n", k, v)
										}

										age := m["James"]
										fmt.Println(age)

										letter := m["Q"]
										fmt.Println(letter)

										if v, ok := m["Q"]; !ok {
											fmt.Printf("Not in Map %v", v)
									    }



										i := 10
										for {

											if i < 1 {
												break
											}
											i--
											fmt.Println(i)
										}

										for i := 0; i <= 100; i++ {
											fmt.Printf("variable i is %v\t", i)
										}


											x := rand.Intn(10)
											y := rand.Intn(10)

											fmt.Println(x)
											fmt.Println(y)

												if x < 4 && y < 4 {
													fmt.Println("x and y are both less than 4")
												} else if x > 6 && y > 6 {
													fmt.Println(" x and y are both greater than 6")
												} else if x >= 4 && x <= 6 {
													fmt.Println(" x is greater than or equal to 4 and less than or equal to 6")
												} else if y != 5 {
													fmt.Println(" y is not 5")
												} else {
													fmt.Println(" none of the previous cases were met")
												}

											switch {
											case x < 4 && y < 4:
												fmt.Println("x and y are both less than 4")
											case x > 6 && y > 6:
												fmt.Println(" x and y are both greater than 6")
											case x >= 4 && x <= 6:
												fmt.Println(" x is greater than or equal to 4 and less than or equal to 6")
											case y != 5:
												fmt.Println(" y is not 5")
											default:
												fmt.Println(" none of the previous cases were met")

											}

											fmt.Println("Hello ")
											x := rand.Intn(250)
											y := rand.Intn(250)

											fmt.Println(x)
											fmt.Println(y)

											if x >= 0 && x <= 100{
												fmt.Println("x between 0 and 100")
											} else if x >= 101 && x <=200{
												fmt.Println("x between 101 and 200")
											} else if x >= 201 && x <= 250{
												fmt.Println("x between 201 and 250")
											} else {
												fmt.Println("x Error")
											}

											if y >= 0 && y <= 100{
												fmt.Println("y between 0 and 100")
											} else if y >= 101 && y <=200{
												fmt.Println("y between 101 and 200")
											} else if y >= 201 && y <= 250{
												fmt.Println("y between 201 and 250")
											} else {
												fmt.Println("y Error")
											}
											z := 45
											fmt.Println(x)
											fmt.Println(y)
											fmt.Println(z)

											if  a := rand.Intn(20); a <= 10{
												fmt.Println("less than rando")
											} else {
												fmt.Println("More than rando")
											}
}

func test(x int, y int) {
	fmt.Println(x + y)
}
*/
