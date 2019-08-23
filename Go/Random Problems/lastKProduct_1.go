package main

import (
    "fmt"
)

func main() {
    kp := New(2)
    kp.Add(0)
    kp.Add(2)
    kp.Add(3)
    kp.Add(5)
}

// Actual implementation
type KProduct struct {
    k, p, zeros int
    nums []int
}

func New (k int) KProduct {
    return KProduct{
        k: k,
        p: 1,
        zeros: 0,
        nums: []int{},
    }
}

func (kp *KProduct) Add (n int) {
    if n == 0 {
        kp.zeros++
    }

    if len(kp.nums) < kp.k {
        kp.nums = append(kp.nums, n)
    } else {
        if kp.nums[0] == 0 {
            kp.zeros--
        } else {
            kp.p /= kp.nums[0]
        }
        kp.nums = append(kp.nums[1:kp.k], n)
    }

    if n != 0 {
        kp.p *= n
    }

    fmt.Printf("%+v\nActual Product: %d\n", *kp, kp.Get())
}

func (kp *KProduct) Get () int {
    if kp.zeros > 0 {
        return 0
    }
    return kp.p
}

