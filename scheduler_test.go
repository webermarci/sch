package sch

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSchedulerRunningOnce(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	s := NewScheduler(func(n int) error {
		wg.Done()
		return nil
	}, time.Millisecond, time.Second)

	s.Schedule(42)
	wg.Wait()
}

func TestSchedulerRunningUntilTimeout(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(10)

	s := NewScheduler(func(n int) error {
		wg.Done()
		return fmt.Errorf("not good")
	}, 100*time.Millisecond, time.Second)

	s.Schedule(42)
	wg.Wait()
}
