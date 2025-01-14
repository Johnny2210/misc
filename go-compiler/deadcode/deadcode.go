package main

const COND = 1

func main() {
	a := 2
	if a < COND {
		println("COMPILER SHOULD DELETE THIS")
	}

	println("Only this should get compiled")
}
