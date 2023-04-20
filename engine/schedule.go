package engine

import "github.com/dreamerjackson/newbiecrawler/collect"

type Scheduler interface {
	Schedule()
	Push(...*collect.Request)
	Pull() *collect.Request
}

type Schedule struct {
	requestCh chan *collect.Request
	workerCh  chan *collect.Request
	reqQueue  []*collect.Request
}

func NewSchedule() *Schedule {
	s := &Schedule{}
	requestCh := make(chan *collect.Request)
	workerCh := make(chan *collect.Request)
	s.requestCh = requestCh
	s.workerCh = workerCh

	return s
}

func (s *Schedule) Push(reqs ...*collect.Request) {
	for _, req := range reqs {
		s.requestCh <- req
	}
}

func (s *Schedule) Pull() *collect.Request {
	r := <-s.workerCh

	return r
}

func (s *Schedule) Schedule() {
	var ch chan *collect.Request

	var req *collect.Request

	for {
		if req == nil && len(s.reqQueue) > 0 {
			req = s.reqQueue[0]
			s.reqQueue = s.reqQueue[1:]
			ch = s.workerCh
		}

		select {
		case r := <-s.requestCh:
			s.reqQueue = append(s.reqQueue, r)
		case ch <- req:
			req = nil
			ch = nil
		}
	}
}
