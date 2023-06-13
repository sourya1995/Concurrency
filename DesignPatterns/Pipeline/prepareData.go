package main

type Operation struct {
	id int64
	multiply int64
	addition int64
}

func prepareData(ic <-chan int64) <-chan Operation { 
	oc := make(chan Operation)
	go func() {
		for id := range ic {
			input := Operation{id: id, multiply: id *2, addition: id + 5}
			oc <- input
		}
		close(oc)
	}()
	return oc
}