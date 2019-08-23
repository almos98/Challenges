// Last K Product
package main

import (
    "fmt"
)

func main() {
    kp := New()
    kp.Add(2)
    kp.Add(0)
    kp.Add(3)
    kp.Add(4)
    kp.Add(5)
    fmt.Printf("Get(5): %d\nGet(4): %d\nGet(3): %d\nGet(2): %d\nGet(1): %d\n",
        kp.Get(5),
        kp.Get(4),
        kp.Get(3),
        kp.Get(2),
        kp.Get(1),
    )
}

// Actual implementation
type KProduct struct {
    p int
    zc, div []int
}

func New () KProduct {
    return KProduct{
        p: 1,
        div: []int{1},
        zc: []int{0},
    }
}

func (kp *KProduct) Add (n int) {
    if n == 0 {
        kp.zc = append(kp.zc, kp.zc[len(kp.zc)-1] + 1)
        kp.div = append(kp.div, kp.p)
    } else {
        kp.zc = append(kp.zc, kp.zc[len(kp.zc)-1])
        kp.p *= n
        kp.div = append(kp.div, kp.p)
    }

    fmt.Printf("%+v\n", *kp)
}

func (kp *KProduct) Get (k int) int {
    if kp.zc[len(kp.zc)-1] - kp.zc[len(kp.zc)-1-k] > 0 {
        return 0
    }
    return kp.p / kp.div[len(kp.div)-k-1]
}
