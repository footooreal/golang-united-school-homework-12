package batch

import (
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	chan_user := make(chan user, n)
	chan_int := make(chan int, n)
	arr_user := []user{}

	routine := func(i chan int, u chan user) {
		for v := range i {
			u <- getOne(int64(v))
		}
	}

	for i := 0; i < int(pool); i++ {
		go routine(chan_int, chan_user)
	}

	for i := 0; i < int(n); i++ {
		chan_int <- i
	}

	for i := 0; i < int(n); i++ {
		arr_user = append(arr_user, <-chan_user)
	}

	return arr_user
}
