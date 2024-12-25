package concurrency

type result struct {
	string
	bool
}

func CheckWebsites(wc func(url string) bool, websites []string) map[string]bool {
	// var m sync.Mutex

	ch := make(chan result)

	results := map[string]bool{}

	for _, w := range websites {
		go func() {
			// Option 1 (Mutex)
			// m.Lock()
			// results[w] = wc(w)
			// m.Unlock()

			// Option 2 (Channels)
			ch <- result{w, wc(w)}
		}()
	}

	for i := 0; i < len(websites); i++ {
		e := <-ch
		results[e.string] = e.bool
	}

	return results
}
