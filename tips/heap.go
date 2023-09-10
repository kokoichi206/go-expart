package main

//go:noinline
func sumValue(a, b int) int {
	z := a + b
	return z
}

//go:noinline
func sumPointer(a, b int) *int {
	z := a + b
	return &z
}
