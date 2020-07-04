package main

func main() {
	// loopTilCondition()
	// loopTilConditionWithPostClause()
	// infiniteLoop()
	collectionLoop()
}

func loopTilCondition() {
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 4 {
			break
		}
		if i == 2 {
			continue
		}
		println("continuing...")
	}
	println(i)
}

func loopTilConditionWithPostClause() {
	for i := 0; i < 5; i++ {
		println(i)
	}
	// compiler error below
	// println(i)
}

func infiniteLoop() {
	var i int
	for {
		if i == 5 {
			break
		}
		println(i)
		i++
	}
}

func collectionLoop() {
	// with traditional loop
	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	// collection loop with slice
	for i, v := range slice {
		println(i, v)
	}

	// collection loop with map
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for key, value := range wellKnownPorts {
		println(key, value)
	}

	// collection loop with map: key only
	for key := range wellKnownPorts {
		println(key)
	}

	// collection loop with map: value only
	for _, value := range wellKnownPorts {
		println(value)
	}
}
