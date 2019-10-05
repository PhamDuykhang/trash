package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	tunel := make(chan string)
	tunel2 := make(chan string)
	listData := []string{"as", "ds", "as", "ds", "as", "ds", "fail", "as", "ds", "as", "ds", "fail", "fail"}
	go save(tunel2)
	go retry(tunel2, tunel)
	for _, data := range listData {
		fmt.Println("sending : ", data)
		go sendMessage(tunel, data)

	}

	time.Sleep(1 * time.Minute)
	return
}

func sendMessage(out chan<- string, in string) {
	err := send(in)
	if err != nil {
		out <- in
		return
	}
	return
}

func send(data string) error {
	if data == "fail" {
		err := errors.New("have errors")
		fmt.Println("got ", err)
		return err
	}
	return nil
}

func retry(inDataBase chan<- string, inC <-chan string) {

	select {
	case a, ok := <-inC:
		if ok {
			fmt.Println("stopped")
			close(inDataBase)
			return
		}
		for i := 1; i <= 3; i++ {
			fmt.Printf("resending data time %d \n", i)
			err := resend(a)
			if err == nil {
				return
			}
		}
		inDataBase <- a
		return
	}
}

func save(data <-chan string) {
	select {
	case a, ok := <-data:
		if ok {
			fmt.Println("done")
			return
		}
		fmt.Printf("data : %s is saved", a)
		return
	}
}

func resend(data string) error {
	if data == "fail" {
		return errors.New("have errors")
	}
	return nil
}
