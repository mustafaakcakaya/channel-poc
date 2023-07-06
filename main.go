package main

import (
	"fmt"
)

// Customer havuzda depolanacak müşteri veri yapısıdır
type Customer struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	// Customer kanalı oluşturulur
	customerChan := make(chan *Customer)

	// Havuzda müşteri verilerini depolamak için bir goroutine oluşturulur
	go func() {
		// Havuzda bir Customer nesnesi oluşturulur
		customer := &Customer{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Address:   "123 Main St",
		}

		// Customer kanalına müşteri bilgileri gönderilir
		customerChan <- customer
	}()

	// Havuzdaki müşteri verilerini kullanmak için bir goroutine daha oluşturulur
	go func() {
		// Customer kanalından müşteri bilgileri alınır
		customerFromChan := <-customerChan
		fmt.Println("Havuzdan Alınan Müşteri:", customerFromChan)

		// Havuzdaki müşteri bilgileri kullanılır
		// ...

		// Kanal kapatılır
		close(customerChan)
	}()

	// Kanalın kapatılmasını bekler
	<-customerChan
}
