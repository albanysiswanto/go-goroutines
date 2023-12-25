package belajar_go_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Albany Siswanto"
		fmt.Println("Data berhasil di simpan")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// 2. Channel as Parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Albany Siswanto"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// 3. Channel IN and OUT
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Albany Siswanto"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// 4. Channel Buffer

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Albany"
		channel <- "Siswanto"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	//fmt.Println(len(channel)) // Melihat jumlah data di buffer
	//fmt.Println(cap(channel)) // Melihat panjang buffer

	fmt.Println("Selesai")
}

// 5. Range Channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke-" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("DONE")
}

// 6. Select Channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	conunter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			conunter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			conunter++
		}

		if conunter == 2 {
			fmt.Println("DONE")
			break
		}
	}
}

// 7. Default Select Channel
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	conunter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			conunter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			conunter++
		default:
			fmt.Println("Menunggu data")
			// ini dilakukan saat ingin mengerjakan proses lain sembari menunggu channel mengirim data
		}

		if conunter == 2 {
			fmt.Println("DONE")
			break
		}
	}
}
