package main

import (
	"sync"
	"thread-safe-queue/queue"
)

func main() {
	println("Testing ConcurrentQueue")

	cq := queue.NewConcurrentQueue()

	// waitgroup
	var wgE sync.WaitGroup //waitgroup enqueue
	var wgD sync.WaitGroup //waitgroup dequeue

	var goRoutines = 1_00_00_000

	for i := 0; i < goRoutines; i++ {
		wgE.Add(1)
		go func() {
			cq.Enqueue(int32(i))
			// println("Added to queue", i)
			wgE.Done()
		}()
	}

	wgE.Wait()

	println("Queue size after enqueue", cq.Size())

	for i := 0; i < goRoutines; i++ {
		wgD.Add(1)
		go func() {
			_ = cq.Dequeue()
			// println("Removed from queue", x)
			wgD.Done()
		}()
	}
	wgD.Wait()
	println("Queue size after dequeue", cq.Size())

}
