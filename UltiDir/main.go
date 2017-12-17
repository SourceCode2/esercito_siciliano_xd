package main

import (
	"os"
)

func main() {
	for i := 0; true; i++ {
		var path = "/home/stud/Desktop/Ciao" + string(i)
		os.Mkdir(path, 0777)
	}
}
