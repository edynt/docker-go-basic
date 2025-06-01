package main

import (
	"fmt"
	"time"
)

type Message struct {
	OrderId string
	Title   string
	Price   int64
}

func publisher(channel chan<- Message, orders []Message) {
	for _, order := range orders {
		fmt.Printf("Pub:::%s\n", order.OrderId)
		channel <- order
		time.Sleep(time.Second * 1)
	}
	close(channel)
}

func subscriber(channel <-chan Message, userName string) {
	for msg := range channel {
		fmt.Printf("userName %s, Order %s, Title %s, Price %d \n", userName, msg.OrderId, msg.Title, msg.Price)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// 1 - channel order
	orderChannel := make(chan Message)

	// 2 - simulate orders
	orders := []Message{
		{OrderId: "Order-01", Title: "Tips GO", Price: 30},
		{OrderId: "Order-02", Title: "Tips NodeJS", Price: 40},
		{OrderId: "Order-03", Title: "Tips Java", Price: 50},
	}

	// send order to pub
	go publisher(orderChannel, orders)
	go subscriber(orderChannel, "anoystick user")

	time.Sleep(time.Second * 3)
	fmt.Println(("End pub sub"))
}

// type Course struct {
// 	Title string
// 	Price int
// }

// func main() {
// 	// 1. add channel
// 	ch := make(chan Course)

// 	// 2. create goroutine
// 	go func() {
// 		course := Course{Title: "Tips GO", Price: 30}
// 		ch <- course // send data to channel
// 	}()

// 	c := <-ch // receive data
// 	fmt.Printf("Receive data: title %s, price %d \n", c.Title, c.Price)
// }

// func main() { // goroutine
// 	// r := gin.Default()
// 	// r.GET("/ping", func(c *gin.Context) {
// 	// 	c.JSON(http.StatusOK, gin.H{
// 	// 		"message": "pong",
// 	// 	})
// 	// })

// 	// r.Run()

// 	fmt.Println("Starting...")

// 	var wg sync.WaitGroup

// 	ids := []int{1, 2, 3, 4, 5, 6}

// 	start := time.Now()
// 	for _, id := range ids {
// 		wg.Add(1)
// 		go getProductByIdAPI(id, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("Total time: ", time.Since(start))
// }

// func getProductByIdAPI(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	url := fmt.Sprintf("https://fakestoreapi.com/products/%d", id)
// 	resp, err := http.Get(url)

// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		return
// 	}

// 	fmt.Printf(">>> Data Product ID: %d - %s \n ", id, string(body))
// }
