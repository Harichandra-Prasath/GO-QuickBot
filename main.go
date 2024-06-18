package main

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	Listen()
}
