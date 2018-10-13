package main

import (
        "unsafe"
)

func getg() (unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
func getg1() unsafe.Pointer

// G returns current g (the goroutine struct) to user space.
func G() (unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) {
    return getg()
}

func G1() unsafe.Pointer {
    return getg1()
}
