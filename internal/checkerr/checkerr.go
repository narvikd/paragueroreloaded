package checkerr

import "log"

// Println checks for an error and if it's found it "log.Println" it
func Println(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Fatalln checks for an error and if it's found it "log.Fatalln" it
func Fatalln(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
