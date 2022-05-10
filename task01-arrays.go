package homework

func average(input []float32) (result float32) {
	//Place your code here
	 var sum float32;
	for _,nums := range input{
		sum+=nums
	}
	average := sum/float32(len(input))
	return average
}
