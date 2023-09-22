package asyncQueue

type Job func() interface{}

func NewAsyncQueue(capacity int) (chan<- Job, <-chan interface{}) {
	jobChan := make(chan Job, capacity)
	resChan := make(chan interface{}, capacity)

	go func() {
		for job := range jobChan {
			resChan <- job()
		}
		close(resChan)
	}()

	return jobChan, resChan
}
