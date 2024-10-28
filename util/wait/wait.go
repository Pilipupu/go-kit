package wait

import "time"

type ConditionFunc func() (done bool, err error)

type Backoff struct {
	Duration time.Duration
	Factor   float64
	Jitter   float64
	Steps    int
	Cap      time.Duration
}

func ExponentialBackoff(backoff Backoff, condition ConditionFunc) Backoff {
	for backoff.Steps > 0 {
		runCon
		time.Sleep(backoff.Duration)
	}
}

func (backoff *Backoff) Step() time.Duration {

}

func (backoff *Backoff) delay() {

}
