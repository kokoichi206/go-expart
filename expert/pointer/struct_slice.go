package main

type pien struct {
	paon   string
	hoge   string
	taihen string
}

//go:noinline
func sf(v pien) pien {
	return v
}

//go:noinline
func sg(v *pien) *pien {
	return v
}

//go:noinline
func f(v []pien) []pien {
	return v
}

//go:noinline
func g(v []*pien) []*pien {
	paon := []*pien{}
	return paon
}
