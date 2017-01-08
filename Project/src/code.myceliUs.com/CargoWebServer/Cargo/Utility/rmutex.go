package Utility

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Token used to identify the caller
type Token int32

var tokenGender int32 = 0

func getToken() int32 {
	return atomic.AddInt32(&tokenGender, 1)
}

// Get a newly generated token
func NewToken() Token {
	return Token(getToken())
}

type Rmutex struct {
	lock      sync.Mutex
	owner     int32 // current lock owner
	recursion int32 //current recursion level
}

func NewRmutex() *Rmutex {
	return new(Rmutex)
}

func (r *Rmutex) Lock(me Token) {
	// fast path
	if atomic.LoadInt32(&r.owner) == int32(me) {
		r.recursion++
		return
	}

	r.lock.Lock()
	// we are now inside the lock
	atomic.StoreInt32(&r.owner, int32(me))
	r.recursion = 1
}

func (r *Rmutex) Unlock(me Token) {
	if atomic.LoadInt32(&r.owner) != int32(me) {
		panic(fmt.Sprintf("you are not the owner(%d): %d!", r.owner, me))
	}

	r.recursion--
	// fast path
	if r.recursion != 0 {
		return
	}

	// we are going to release the lock
	atomic.StoreInt32(&r.owner, 0)
	r.lock.Unlock()
}

