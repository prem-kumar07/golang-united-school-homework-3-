package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	reversed:= make( []int64,len(input))
	copy(reversed,input)
	for i := len(input)/2 - 1; i >= 0; i-- {
        rev := len(input) - 1 - i
        reversed[i], reversed[rev] = reversed[rev], reversed[i]
    }
	return reversed
}
