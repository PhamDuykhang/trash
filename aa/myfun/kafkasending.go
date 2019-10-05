package myfun

import (
	"errors"
	"fmt"
	"time"
)

const (
	MaxRetry = 3
)

func Send(data string) {
	r := make(chan string)
	go func() {
		_ = RetryAndSave(r)
		return
	}()
	go func() {
		_ = SendToKafka(r, data)
	}()
	fmt.Println("sending complete")
	return
}

func SendToKafka(out chan<- string, d string) error {
	if d == "f" {
		err := errors.New("just error")
		fmt.Println("put data into retry channel", d)
		out <- d
		return err
	}
	fmt.Printf("send %s is successfuly ", d)
	close(out)
	return nil
}

func RetryAndSave(in chan string) error {
	data, ok := <-in
	if !ok {
		fmt.Println("done ", data)
		return nil
	}
	for i := 1; i <= MaxRetry; i++ {
		fmt.Println("retry ")
		time.Sleep(2 * time.Second)
		err := kafkaSend(data)
		if err == nil {
			return nil
		}
	}
	err := sendToMongo(data)
	if err != nil {
		return err
	}
	fmt.Println("saved")
	return nil
}

func kafkaSend(d string) error {
	if d == "f" {
		fmt.Println("got error")
		return errors.New("got err")
	}
	return nil
}

func sendToMongo(a string) error {
	if a == "fff" {
		fmt.Println("save fail")
		return errors.New("save to database fail")
	}
	fmt.Println("save suss")
	return nil
}
