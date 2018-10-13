package main

import (
    "fmt"
    "time"
	"unsafe"
)

func call() {
    for true {
        g, g1, g2, g3 := getg()
        g4 := getg1()
		p := (*g000)(g1)
        fmt.Printf("test %v, %v, %v==%v, %v, guid %d\n", g, g1, g2, g4, g3, p.goid)
        time.Sleep(time.Second * 10)
    }
}

func main() {
    for i := 0; i < 100; i++ {
    go call();
    }
    for true {
        g, g1, g2, g3 := getg()
        fmt.Printf("main %v, %v, %v, %v\n", g, g1, g2, g3)
        time.Sleep(time.Second*5)
    }
}

type stack struct {
	lo uintptr
	hi uintptr
}

type _panic struct {
	argp      unsafe.Pointer // pointer to arguments of deferred call run during panic; cannot move - known to liblink
	arg       interface{}    // argument to panic
	link      *_panic        // link to earlier panic
	recovered bool           // whether this panic is over
	aborted   bool           // the panic was aborted
}

type g000 struct {
	// Stack parameters.
	// stack describes the actual stack memory: [stack.lo, stack.hi).
	// stackguard0 is the stack pointer compared in the Go stack growth prologue.
	// It is stack.lo+StackGuard normally, but can be StackPreempt to trigger a preemption.
	// stackguard1 is the stack pointer compared in the C stack growth prologue.
	// It is stack.lo+StackGuard on g0 and gsignal stacks.
	// It is ~0 on other goroutine stacks, to trigger a call to morestackc (and crash).
	stack       stack   // offset known to runtime/cgo
	stackguard0 uintptr // offset known to liblink
	stackguard1 uintptr // offset known to liblink

	_panic         uintptr // innermost panic - offset known to liblink
	_defer         uintptr // innermost defer
	m              uintptr      // current m; offset known to arm liblink
	sched          gobuf
	syscallsp      uintptr        // if status==Gsyscall, syscallsp = sched.sp to use during gc
	syscallpc      uintptr        // if status==Gsyscall, syscallpc = sched.pc to use during gc
	stktopsp       uintptr        // expected sp at top of stack, to check in traceback
	param          unsafe.Pointer // passed parameter on wakeup
	atomicstatus   uint32
	stackLock      uint32 // sigprof/scang lock; TODO: fold in to atomicstatus
	goid           int64
	waitsince      int64  // approx time when the g become blocked
	waitreason     string // if status==Gwaiting
}

type gobuf struct {
	// The offsets of sp, pc, and g are known to (hard-coded in) libmach.
	//
	// ctxt is unusual with respect to GC: it may be a
	// heap-allocated funcval so write require a write barrier,
	// but gobuf needs to be cleared from assembly. We take
	// advantage of the fact that the only path that uses a
	// non-nil ctxt is morestack. As a result, gogo is the only
	// place where it may not already be nil, so gogo uses an
	// explicit write barrier. Everywhere else that resets the
	// gobuf asserts that ctxt is already nil.
	sp   uintptr
	pc   uintptr
	g    guintptr
	ctxt unsafe.Pointer // this has to be a pointer so that gc scans it
	ret  Uintreg
	lr   uintptr
	bp   uintptr // for GOEXPERIMENT=framepointer
}

type guintptr uintptr
type Uintreg uint64
