package worker

import (
	"context"
	"log"
	"sync"
)

// internal/worker/pool.go
type JobFunc func(ctx context.Context) error

type Pool struct {
	jobs chan JobFunc
	wg   sync.WaitGroup
}

func NewPool(workers int) *Pool {
	p := &Pool{jobs: make(chan JobFunc, 100)} // 100 = buffer, tune as needed
	for i := 0; i < workers; i++ {
		p.wg.Add(1)
		go p.run()
	}
	return p
}

func (p *Pool) run() {
	defer p.wg.Done()
	for job := range p.jobs {
		if err := job(context.Background()); err != nil {
			log.Printf("job error: %v", err)
		}
	}
}

func (p *Pool) Submit(job JobFunc) {
	p.jobs <- job
}

func (p *Pool) Shutdown() {
	close(p.jobs)
	p.wg.Wait()
}
