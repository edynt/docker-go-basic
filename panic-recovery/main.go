package main

import (
	"log"
)

// func main() {
// 	log.Println("X...1")

// 	// Panic and Recovery
// 	defer func() {
// 		log.Println("Close Connect MySQL")
// 	}()
// 	// panic("Something went wrong")

// 	log.Println("X...2")
// 	log.Println("X...3")
// }

func order() {
	log.Println("Start API Order")

	sendEmail()
	sendSMS()

	log.Println("Success API Order")
}

func sendEmail() {
	log.Println("Start API Send Email")
	defer func() {
		if r := recover(); r != nil {
			log.Println("Notes: hot fix")
		}
	}()
	panic("Send Email Error")
	log.Println("Success API Send Email")
}

func sendSMS() {
	log.Println("Start API Send SMS")

	log.Println("Success API Send SMS")
}

func main() {
	log.Println("Process API Order ...")
	order()
	log.Println("End API Order ...")
}
