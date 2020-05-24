package BubbleSort

import "reflect"

func PerformBubbleSortForInt(inputArr []int) []int {
	performIteration := func(inArr []int) []int {
		resultArr := make([]int, len(inArr))
		copy(resultArr, inArr)
		for i := range inArr {
			//if it's the last item in the array, carry on
			if i == len(inArr) -1 {
				break
			}

			if inArr[i] > inArr[i + 1] {
				resultArr[i] = inArr[i + 1]
				resultArr[(i + 1)] = inArr[i]
			}
		}
		return resultArr
	}

	currentIteration := make([]int, len(inputArr))
	copy(currentIteration, inputArr)
	for {
		nextIteration := performIteration(currentIteration)
		if  reflect.DeepEqual(currentIteration, nextIteration) {
			return nextIteration
		}else{
			copy(currentIteration, nextIteration)
		}
	}
}