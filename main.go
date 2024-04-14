package main

import (
	"fmt"
	"time"
)

func main() {
	resultch := make(chan string)

	//this code is unblockable because here resultch knows that there is person-A (go anonymous func) to accept the cookie when person-B
	//(resultch <- "fooo") gives resultch
	//and it passs this cookie to the person-A because resultch is unbuffered means it has no box to store cookie but to handover to someone else
	go func() {
		result := <-resultch
		fmt.Println(result)
	}()

	resultch <- "fooo"

}

func fetchRes(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf(`waited for..%d`, n)
}
