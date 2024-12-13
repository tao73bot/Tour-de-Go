package main

import (
	"fmt"
	"time"
	"sync"
)

type Message struct {
	chats []string
	friends []string
}

func main() {
	fmt.Println("Go Routings")
	now := time.Now()

	id := getUserByName("Taohid")
	fmt.Println(id)

	wg := &sync.WaitGroup{}
	ch := make(chan *Message, 2)
	wg.Add(2)

	go getUserChats(id,ch,wg)

	go getUserFriends(id,ch,wg)
	wg.Wait()
	close(ch)
	for msg := range ch {
		fmt.Println(msg)
	}
	fmt.Println(time.Since(now))
}

func getUserFriends(id string, ch chan <- *Message, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	ch <- &Message {
		friends: []string{"Friend1", "Friend2", "Friend3"},
	}
	wg.Done()
}

func getUserChats(id string, ch chan <- *Message, wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	ch <- &Message {
		chats: []string{"Chat1", "Chat2", "Chat3"},
	}
	wg.Done()
}

func getUserByName(name string) string {
	time.Sleep(1 * time.Second)
	return "User: " + name
}