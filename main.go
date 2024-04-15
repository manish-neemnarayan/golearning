package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	ID       int
	Comments []string
	Likes    int
	Friends  []int
}

func main() {
	start := time.Now()
	userProfile, err := handleGetUserProfile(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", userProfile)
	fmt.Printf("time taken in fetching user profile: %v\n", time.Since(start))
}

type Response struct {
	data any
	err  error
}

func handleGetUserProfile(id int) (*UserProfile, error) {
	var (
		respch = make(chan Response, 3)
		wg     = &sync.WaitGroup{}
	)

	//we are doing 3 request inside their own goroutine
	go getComments(id, respch, wg)
	go getLikes(id, respch, wg)
	go getFriends(id, respch, wg)
	//adding 3 to the waitgroup
	wg.Add(3)
	wg.Wait() // block until the wg counter == 0 we unblock
	close(respch)

	///keep ranging but when to stop??
	// by close(respch) it knows that now it should break
	userProfile := &UserProfile{}
	for resp := range respch {
		if resp.err != nil {
			return nil, resp.err
		}
		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		case []string:
			userProfile.Comments = msg

		}
	}

	return userProfile, nil
}

func getComments(id int, respch chan Response, wg *sync.WaitGroup) {
	_ = id

	time.Sleep(time.Millisecond * 200)
	comments := []string{
		"hey buddy..",
		"where are you??",
		"comming for cricket match today???",
	}
	respch <- Response{
		data: comments,
		err:  nil,
	}

	//work is done
	wg.Done()
}

func getFriends(id int, respch chan Response, wg *sync.WaitGroup) {
	_ = id
	time.Sleep(time.Millisecond * 100)
	friendIds := []int{23, 434, 342, 23}

	respch <- Response{
		data: friendIds,
		err:  nil,
	}

	wg.Done()
}

func getLikes(id int, respch chan Response, wg *sync.WaitGroup) {
	_ = id
	time.Sleep(time.Millisecond * 200)
	respch <- Response{
		data: 23,
		err:  nil,
	}

	//work is done
	wg.Done()
}
