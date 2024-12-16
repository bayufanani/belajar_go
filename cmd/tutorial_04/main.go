package main

import "fmt"

func main() {
	var intArr [3]int32
	intArr[1] = 111
	fmt.Println(intArr[0])
	fmt.Println(intArr[0:2])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// initialize array
	intArray2 := [...]int32{1, 22, 3}
	fmt.Println(intArray2)

	// slices
	intSlices := []int32{4, 5, 6}
	fmt.Println(intSlices)
	fmt.Printf("panjang slice adalah %v, kapasitas %v", len(intSlices), cap(intSlices))
	intSlices = append(intSlices, 7)
	fmt.Println(intSlices)
	fmt.Printf("panjang slice adalah %v, kapasitas %v\n\n", len(intSlices), cap(intSlices))

	var intSlices2 []int32 = []int32{8, 9}
	intSlices = append(intSlices, intSlices2...)
	fmt.Println(intSlices)

	var intSLice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSLice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"bayu": 30, "erin": 29}
	fmt.Println(myMap2["bayu"])
	fmt.Println(myMap2["john"])
	var umur, ok = myMap2["john"]
	if ok {
		fmt.Printf("umur %v\n", umur)
	} else {
		fmt.Println("tidak ada")
	}

	for name, age := range myMap2 {
		fmt.Println(name, age)
	}

	for i, v := range intArr {
		fmt.Println(i, v)
	}

	// no while in go, this is alternative
	var i = 0
	for i < 5 {
		fmt.Print(i)
		i++
	}

	fmt.Println()

	var j = 0
	for {
		if j > 5 {
			break
		}
		fmt.Print(j)
		j++
	}

	fmt.Println()

	for k := 0; k < 5; k++ {
		fmt.Print(k)
	}
	fmt.Println()

	startBenchmark()
}
