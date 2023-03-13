package sch

import (
	"time"
)

type task[T any] struct {
	data  T
	until int64
}

func (t *task[T]) run(scheduler *Scheduler[T]) {
	if time.Now().UnixNano() > t.until {
		return
	}

	if err := scheduler.operation(t.data); err == nil {
		return
	}

	time.Sleep(scheduler.periodicity)
	t.run(scheduler)
}

type Scheduler[T any] struct {
	operation   func(T) error
	periodicity time.Duration
	limit       time.Duration
}

func NewScheduler[T any](operation func(T) error, periodicity, limit time.Duration) *Scheduler[T] {
	return &Scheduler[T]{
		operation:   operation,
		periodicity: periodicity,
		limit:       limit,
	}
}

func (s *Scheduler[T]) Schedule(data T) {
	t := task[T]{
		data:  data,
		until: time.Now().Add(s.limit).UnixNano(),
	}
	go t.run(s)
}
