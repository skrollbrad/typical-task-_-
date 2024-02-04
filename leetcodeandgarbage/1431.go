package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	testArray := []int{}
	for _, v := range candies {
		testArray = append(testArray, v)
	}

	for i := 0; i < len(candies)-1; i++ {
		if testArray[i] > testArray[i+1] {
			testArray[i+1], testArray[i] = testArray[i], testArray[i+1]
		}

	}
	boolArray := []bool{}
	maxValue := testArray[len(candies)-1]
	for _, v := range candies {
		if v+extraCandies >= maxValue {
			boolArray = append(boolArray, true)
		} else {
			boolArray = append(boolArray, false)
		}
	}

	return boolArray
}
