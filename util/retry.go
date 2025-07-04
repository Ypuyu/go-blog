package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/avast/retry-go"
)

// 重试次数，attempts次数，delay时间间隔，nil是随机间隔
func Retry(f func() error, attempts uint, opts ...retry.Option) error {
	newOpts := []retry.Option{retry.Attempts(attempts), retry.LastErrorOnly(true)}
	newOpts = append(newOpts, opts...)
	return retry.Do(
		f,
		newOpts...,
	)
}

func UntilDeadline(deadline time.Time, f func() error) error {
	err := errors.New(fmt.Sprintf("deadline has been exceeded, now: %v, deadline: %v", time.Now(), deadline))
	sleepTime := time.Millisecond * 100
	for ; deadline.After(time.Now()) && err != nil; time.Sleep(sleepTime) {
		sleepTime *= 3
		err = f()
	}
	return err
}
