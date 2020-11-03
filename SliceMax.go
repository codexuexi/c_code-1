package c_code

func SliceMax(intArr []int) (val, index int) {
	//var intArr  = [...]int{3,-4,93,8,12,29}
	maxVal := intArr[0]
	maxValIndex := 0
	for i := 0; i < len(intArr); i++ {
		//从第二个元素开始循环比较，如果发现有更大的数，则交换
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	return maxVal, maxValIndex
}
