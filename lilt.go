package lilt

import "testing"

type TInterface interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Parallel()
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
}

type sinkT struct{}

func (s sinkT) Error(args ...interface{})                 {}
func (s sinkT) Errorf(format string, args ...interface{}) {}
func (s sinkT) Fail()                                     {}
func (s sinkT) FailNow()                                  {}
func (s sinkT) Fatal(args ...interface{})                 {}
func (s sinkT) Fatalf(format string, args ...interface{}) {}
func (s sinkT) Log(args ...interface{})                   {}
func (s sinkT) Logf(format string, args ...interface{})   {}
func (s sinkT) Parallel()                                 {}
func (s sinkT) Skip(args ...interface{})                  {}
func (s sinkT) SkipNow()                                  {}
func (s sinkT) Skipf(format string, args ...interface{})  {}

var sink sinkT = sinkT{}

func NewIf(t *testing.T) func(bool) TInterface {
	return func(condition bool) TInterface {
		if condition {
			return t
		} else {
			return sink
		}
	}
}

func NewIfNot(t *testing.T) func(bool) TInterface {
	return func(condition bool) TInterface {
		if !condition {
			return t
		} else {
			return sink
		}
	}
}

func New(t *testing.T) (func(bool) TInterface, func(bool) TInterface) {
	return NewIf(t), NewIfNot(t)
}
