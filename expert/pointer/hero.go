package main

// just a struct with 3 string fields
type Hero struct {
	Name        string
	Description string
	Class       string
}

//go:noinline
func NewHero() Hero {
	return Hero{
		Name:        "Hero",
		Description: "Hero description",
		Class:       "Hero class",
	}
}

//go:noinline
func NewHeroPtr() *Hero {
	return &Hero{
		Name:        "Hero",
		Description: "Hero description",
		Class:       "Hero class",
	}
}

//go:noinline
func ReturnHero() Hero {
	h := NewHero()
	return h
}


//go:noinline
func ReturnHeroPtrWithNewHero() *Hero {
	h := NewHero()
	return &h
}


//go:noinline
func ReturnHeroPtr() *Hero {
	h := NewHeroPtr()
	return h
}

//go:noinline
func returnHero(h Hero) Hero {
	return h
}

//go:noinline
func returnHeroPtr(h *Hero) *Hero {
	return h
}
