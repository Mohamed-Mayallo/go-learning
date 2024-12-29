package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("using unbuffered channels and wait groups", func(t *testing.T) {
		counter := Counter{}
		wantedCount := 1000

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		ch := make(chan struct{})

		// should be in another goroutine to make sure that the synchronization is done
		go func() {
			ch <- struct{}{} // send the first signal to start the first goroutine
		}()

		for i := 0; i < wantedCount; i++ {
			go func() {
				<-ch // wait for the next signal

				counter.Inc()

				// should be in another goroutine to make sure that the synchronization is done
				go func() {
					ch <- struct{}{} // signal the next goroutine to start
				}()

				wg.Done()
			}()
		}

		wg.Wait()

		if counter.Value() != wantedCount {
			t.Errorf("got %d, want %d", counter.Value(), wantedCount)
		}
	})

	t.Run("using buffered channels and wait groups", func(t *testing.T) {
		counter := Counter{}
		wantedCount := 1000

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		// only one goroutine can access the channel at a time because the channel has a buffer size of 1.
		ch := make(chan struct{}, 1)

		ch <- struct{}{} // send the first signal to start the first goroutine

		for i := 0; i < wantedCount; i++ {
			go func() {
				<-ch // wait for the next signal
				counter.Inc()
				ch <- struct{}{} // signal the next goroutine to start
				wg.Done()
			}()
		}

		wg.Wait()

		if counter.Value() != wantedCount {
			t.Errorf("got %d, want %d", counter.Value(), wantedCount)
		}
	})

	t.Run("using buffered channels without wait groups (infinite loop)", func(t *testing.T) {
		counter := Counter{}
		wantedCount := 1000

		ch := make(chan struct{}, 1)

		ch <- struct{}{}

		for i := 0; i < wantedCount; i++ {
			go func() {
				<-ch
				counter.Inc()
				ch <- struct{}{}
			}()
		}

		for {
			if counter.Value() == wantedCount {
				break
			}
		}

		if counter.Value() != wantedCount {
			t.Errorf("got %d, want %d", counter.Value(), wantedCount)
		}
	})

	t.Run("using buffered channels without wait groups (another channel)", func(t *testing.T) {
		counter := Counter{}
		wantedCount := 1000

		ch := make(chan struct{}, 1)
		wc := make(chan struct{}, 1)

		ch <- struct{}{}

		for i := 0; i < wantedCount; i++ {
			go func() {
				<-ch
				counter.Inc()
				ch <- struct{}{}

				if counter.Value() == wantedCount {
					close(wc)
				}
			}()
		}

		<-wc

		if counter.Value() != wantedCount {
			t.Errorf("got %d, want %d", counter.Value(), wantedCount)
		}
	})

	t.Run("using mutexes and wait groups", func(t *testing.T) {
		counter := Counter{}
		wantedCount := 1000

		var wg sync.WaitGroup
		var mut sync.Mutex

		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				mut.Lock()
				counter.Inc()
				mut.Unlock()
				wg.Done()
			}()
		}

		wg.Wait()

		if counter.Value() != wantedCount {
			t.Errorf("got %d, want %d", counter.Value(), wantedCount)
		}
	})
}
