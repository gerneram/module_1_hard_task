package fan_in

import (
	"sync"
)

// MergeChannels - принимает несколько каналов на вход и объединяет их в один
// Fan-in и merge channels синонимы
func MergeChannels(channels ...<-chan int) <-chan int {
	// TODO: ваш код
	var wg sync.WaitGroup
	out := make(chan int)
	for _, channel := range channels {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for c := range ch {
				out <- c
			}

		}(channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
