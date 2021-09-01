package cube

func upper3(x float64) float64 {
	return x * x * x
}

func Volume(lado float64) float64 {
	return upper3(lado)
}
