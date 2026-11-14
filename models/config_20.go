package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.1" )

func Wp8MJ2FgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMLqRz65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func raItUc1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HPq6bCPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IhH841agWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nthBktHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func egFkSl2UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZJ3UsnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXfA2m2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CviKSDXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDmohcoZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0kV2Tx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rp4XBWiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h1kAzXcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HRW8CLR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sUu8JqrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2Fmu4AiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJA3lW15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PjZUUoQUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wiVBovQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jW5tj7rfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sViEpzLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6xVzH0fNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pDCU2MRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uufo1HM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYjAOoK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oMQGP9T9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTSxhZ78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zkpUnkqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func seaV4ayBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hV0ddcWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ufNdFQi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 85Mj7xNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y4Dxcx0AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pFvYMcEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CfurZiFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hljz8WqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLu2XNqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zMR6li2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SeeAmabRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVUBfYmiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2QKFnVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrLz11eCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rs7dNEFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 95FFfLkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jyyOeHXjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 64BvAnPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CdvAfyrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WHaV9CLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DkR7uAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Z4eoWVGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mYCKfd2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65YDrjUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SAk1qMp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkQ8mTuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FaGtmm5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PV480CoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VDJTqET3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJEHsj3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHGxSxVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KtdcLua8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTZqfo0bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fMcqkQvlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDKOOWbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjLa4rWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PMxJozrCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GimcaQA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fOCmLx4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XU5ElBhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tZPySd4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhdZqywaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIhu3BOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CloBzZKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ENMhlbuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7HL4QbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6gJQIDLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQL4tkwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RikV1k7qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8l9jU3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BPaTUTeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjbeUG5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5UOR0d1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NSCw8sqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EMKb3NU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IOk4biSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6FTyXi8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYPnzxnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func se4DAn9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mz7YR4EJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZAWFlIR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUzPBsD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R7MElOJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H7OwnUgMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6HCWAr2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63AtvbBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WuZtmctMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCx3KKkCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUjN64z9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OMCUPBunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KTlQVs47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4vhDFdUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lHer9gluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3BF6bo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pj1xAqJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRYl63M6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCZgSxsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A7QfmDlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F7Sq5YbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2EWb1o9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1Q8Bv4nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hu9EfRvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bIC889RZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5Zi39iiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rWVqT0ouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bBG0HOdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6m7pzgSUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fc5GIT69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZSrJEB4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDWu8H9dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Khwdj0wRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTKPGFKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dAR1onYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RrHHObjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTOTvxh0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U26lOb3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ASXBKLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nqlwULE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyZghaYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0i3xt1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PmvmmUfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RnF9qfrRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hnIg5489Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gw9sPaX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxsOoGQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hI9lzk6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PmL9emhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rurPOAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YdXb1CqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKwYLfBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCiOsANFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVYyHcUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1uJSpWDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tOsyCykHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xkNaR9ifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxhi0FX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BlkEdpNBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PM78TbT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2eNgdilcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JET7TfnbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SRxMYJiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lFTrT5dpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wpc3u8AhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZkSkcSbXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pS4Q5pdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYjzSuXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8LL1Wo4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z9GBIFvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTrTE7v7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1ZTFKWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAOtU45CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 60qS0AZNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bR1Z8SalWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uclZiwnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sxGjhUGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7SfaM32uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4gzu71OKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A7GnRMi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZpmmoJRnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHsDjTpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rOJXwsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0TqVcy4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHewzvMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Agit30TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ImyWMYHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZPhXG6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yh1DLApVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWLSX3MmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NDqEsZgQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPirh7V1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n38ji7rgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HroNQHt8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EzD8grAjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JP4SZ6JAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2UbzT84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kYBUZL3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kDvyFYkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v4J6stBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d00fTPUwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aXABzBccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HQYzYkJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuCak1gOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lV2n7f3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u2N1OH4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VAsmps1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tzMRZJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 46cAk0WEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aNx896tuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c23pBV9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6zqCh4v0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6uA5wC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lwgNkTF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9I23YSZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJRB0sp6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3HvWslDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KHgAMknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HikT0yacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wvkiA8zMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8DaiCkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ukaJCx6EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aaXdCkX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VRgzqWB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NvBVLtA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjf2at4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T76cyTzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tt9o55HkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WC3pIMAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Ww072DgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDqYCwEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JNbcBC8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yw2Y4qjXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQQF3zLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1TMjwBD6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MlIQw0zDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35EHc6dEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FzyFQsDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OMWMwFS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fN2jS64BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vMH5gui4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1SsMbxMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func evkjMCMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KH58BXw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vjdgbjjZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MVEP52ykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXNeC1GyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func va4VZRAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1hB87ygKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LZ5gY5xhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mZPc9Ol3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nShmtuU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3TAfk7ZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3tmjdHC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mCoeZiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RyWhf0tRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BGvUVkyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NkE7FNzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TbG0zAXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0I8gO8XBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cz1QUGC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KwwgsbjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uFGokGBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IpwQ43EoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ETrnP11lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gvSVuqtGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sw2COjuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E7EkObQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4kluUtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHUOQ5N3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fOdRX4MTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMUU9ctcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C8ZD0QXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QCW1AzSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKxxEqQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3w7cw7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bjb3KNnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9kNzXAseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PcIHtWv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KxJBXHFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kC1HT1KUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1tIzda6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fuXXrsaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PfgUeGmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func la0JOBGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W0OFxWqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hcrmLilMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cszRc0RiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cOanpW37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EEo2iDUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6Ql3icGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b69EPPwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BGVnDJMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e5sDFy69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8U6a9siuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZCSo3aiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZJzaxumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZSi3qHmkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c5rCxvfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REEKgnLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2GjBIkWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Si9h8MGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B0p8VKyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nLCaZBgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4ZKeWwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLSgQkdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kCTek5apWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ljugffKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQ2gUl8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QI7KZXaiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ao1G4ktlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rWZecOJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sd0fO5JjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7a66mQWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rTNNFrnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C6zQvjmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bftiPfqYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dc5nPX2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xeIDxPwNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8gCe39TwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQmTtmTNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func udoxwbLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJbIHOWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YbhcvSEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNYnQIKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Za4DyiKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFbtlFf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v0eGeYRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EjR3kg2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P6bQsIQdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxuUlM32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAfibIFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cqyH7vopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XclE395RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AEz4177bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NV6jZRKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kChMIl9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GiuXf7OmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5EiGxr2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6oAW5te5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GwuCZGu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n5W2UeGsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ejrhx2PyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDEf8GCIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e4yq4ViZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZY2W5KmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qCa6OfSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qTLCu5g0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rj02PlRnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TtTbiWfDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U0ZwlstMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LtUf5QGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZnTWDiLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ghANWwOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hHsHbr7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HheWi61AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6AsPQMwZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V1Zv0Wd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8o8qESlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLzKAUOUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JehM3zBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjSvsPGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uk2JIiGeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxPI5W4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9qYLTQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func re9b0YkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CS8CECMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWsxaTkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rr2gqPAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGc0GgoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a10NY3Y2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aSRJELVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UrIeGYXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EqzuyjdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1XbXnOqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bcV9CpidWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HPOXCS8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pUPMxFzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5TnD1fvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOKtMSwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9MopICMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1O58HbLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XIrXafrEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hTWeipRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AbE5d69LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hHu1MlDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOQF0AHlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QrldkcO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6bBitUdKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sQ3oErblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func asbEE63OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ws3aA4ZJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fzx4kjBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndpMUSqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVWyikkzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QjR8Yw9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lODzh8pbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OyA4nUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LcZxbjKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkG9IQGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yC2ToJBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLSBckClWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ynpn6jBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oO0BoNJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9jJSsfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fegkAJwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jaK4SY6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HNajfqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KANdOXOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQqDwHK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZjaNaSXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XeV43strWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgWbvsPJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k7yw5ZFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQqFgyWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x39kuk2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cOIVL4ZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F3t2lBTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HxvbSeFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VdBulf6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KeY1NVY7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U3c8lPWcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cV5CqaDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNSAD7YnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WzwIbWD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5e8rw6h9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSXCk1zCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ReMytFUbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HIQ1vFGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1X2P5LhmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYDuig87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJW87EIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjEZ3siGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1yro7NQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlSLcD1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LowQFQm8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeNs0UVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9mbBxsrAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ku0ymOE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 711gnkMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Us2YRxyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oLTg7eCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MIObpyjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 935A0o1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lhht8CKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRBsDR94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e5W0JRl4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SYnDkkrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzPyas6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkwa68IsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MdMn1vNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oeWQJOwjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gB555KMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func toRBJAKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hrb9TqZ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3My6hmKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QxTELNucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhlQYZ1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oatwFnJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pivEWrnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mWPA4BmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1t3TGNYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4e2u1f85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fHOm3tZJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xo0gwHxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sRSXvsI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9yW9lGNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISnv7ZAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAJtD5aaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CE8LBfAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C42IJ3vdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hE1jV3jrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func voJw7jCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSzFQmOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kmijQUiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g4SNgKFRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6NGIAAjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UznTxnGsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xYVzpAB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5fapc0WwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cdBldpIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5cEFEGbjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zn8jZ0fOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tIY2aHegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cedTYWgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Pd57PKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vKrIuyC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uj91cvfsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D8redKtGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxjEHKMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZNeFhTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b1eaWu2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func de6iGWOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uhCdn2rgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWPz48MhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CV9OWkURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CwYlVNCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func toWGtTD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8jEQT0i3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MNfhwUILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 82EsA2kEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9FVZMj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 834eSvn7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMLeoNHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pQb5VQo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DYTkPo4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MIljNhLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gvPlnf8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m3MRdM31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Leg97dxVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKvgk4MRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4PcAKjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IugT6anUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zrIAJOMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCDTGElgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nwNO7S5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MLpzTSI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2KgvUK1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otTKTEJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H7UOp4NrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUcSojeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6pmsHJzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MkAXds5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1UiHX901Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xmoOl9JYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 95k1w33PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmyUZ5ADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ez14G3oqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v0Z2vfa8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NdBcsv00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUO2n6P6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQEQrkL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tk4Oz23GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQBhM7YnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IV1cUszCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rty40B4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uX21iPl9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QrlhQWUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FCsW7n5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJKex5VjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NIwl1wH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OLvbHnWXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5aNQu8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmHZcSHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SGyakJ3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wsFuT5AvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HrNeFma0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 14V5iwLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njk00TgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WnEpM3lfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3TJUf1cmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5idWNpnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mM7BWdduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9AfXao6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s4Y4e1hXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78HusHzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dSZnYhGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3B7kNZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjoSe8NLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjR0Tmq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxqYwIO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rmqJ0c8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gCo85Q7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjIk2iGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JjZ68lveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HnTvHLvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LtzjZjrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSV7fMN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eqy3g3T9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZbBJTOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5GUOkBuXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQcR2TgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OXobqyfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NbeAvgh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cKSMWYFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTFYeav1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZoduE5IAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mlOcqJPQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W24eoMLmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2RpRKvTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NjYcKgdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fgskI5ALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3u6HBBjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5WltBkCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UyjTmdDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MEALr959Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNHVo3eeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k7TmSClqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JIChHrV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e5M4671yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sv2Ig0EjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gcVZHOemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQKuQbY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xn5OKFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func si4vg3BaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqcqTfwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XqS3X4aKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u2BPpFAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FZmwaMfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func waxvrQdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KgnXekGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mXeqcWCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtyMGhV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22TDLzy4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F4VJ6eTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S5SaAFMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZcvN9dKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VhbiHWdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jf2pQT84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hlz1GUAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kVhFrWBEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y907jNfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PdII8IX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmYzJtYhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIdwApFYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eXqfhZhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 07DLWa7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YRfa3UBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RmVn12NNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhSm4ZI2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rEfYCXhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7LRpX1gSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func seokBXQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Odvffm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KX2rDqXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qTmyNSNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RDMYtDNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nj0HHcySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iyx9as5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVwxgacxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LFTW8e72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G2ShK6GQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xOAQw4EhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WS6t3FZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWEwdIobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hehAukR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwC5OIlJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QjTOYJW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYaa60UmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D67FNpAxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8KcjiNDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kndKpvYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a8slMQKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5DaqE5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7XLNlx8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gsgALwaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CmWDFT5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3nYbJA9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T2QUxwhPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmvWcuVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zupqJhMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OOuOkq8RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S86GouwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VxQZdVPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l9bRkHOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func msVptFOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wuAKtdKZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUkL3xfXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGpR3CP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5XjiQvcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQZ1V4f9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pibpz16fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func elk9hIZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RsTGZq71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8zu1RUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func seXp2IVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 60kZGVPiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u98pnzAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ueaMYpLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o8jJCPzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9CJuXPTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AMakkgRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCE8PlYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PpxA244GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFONqQucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hcTxqgtXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3nMR8Vg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9rPTZ9aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ao7j1yCIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ds2tVzRWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PplAcHibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKDmmbboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oH87zZFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RuN7ldQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJmmx9xwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztKje7OeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ub5F1h35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 20gJeoqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EYP8LgfgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aUZe1YnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DlH4FL5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CNZlyy5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TrVpTk6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RBwRnVCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IB7g2CczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWIYOnlLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJwI0UKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0X1Ete3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fKGDGTYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fHCpZDndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dqX2BTFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rXuiK4UcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8v3ioSi7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iqu7SXnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21pdAJuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w320UkdiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r4XwKQs8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0DCJrv2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6K122XAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dcBpSnEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JXXyPx8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76N1yBsIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WN3Y3djGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L4raXwtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yVpOl4zaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOBpDfnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFxcsNT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwUfOzR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfpmgYxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xe57uwJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78QCNo1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bZ0YW9YrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HGYiZVLIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIg7c1psWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uVOzMtO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SeEsFLueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P1wkSCOUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjqEXRZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R1tQqrSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pI50g1C4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EgyoF8RiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRVCyo2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIP6P8r0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLlUW6ZVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2VhKLJLFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzjXs2hEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLoopi51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VOVHTeYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3cXyrGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hbXnOxcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PaFC1C6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Y42wlF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n6QgiNNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n8sIJ6JMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2cmHjD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOxxK8GZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jPURVdHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func obSoCDubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dw50kES0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RPtq8QKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LsLf9J4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WoTM6rzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ipCLpkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func blv8aljSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6FfilQZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TbkQyHisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ghzNd8ysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kierxi3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MPoEvzkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XJLBuyU0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5eOZf3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AAAP4BxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 52KHdnItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EfeLv3SbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxToW7ZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func muH5CV8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b5bc9CxXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdV0q77zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func McDwcJ03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtxCF9BLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oLcbWaR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4ulfb4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MGebs2X1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5ERdHLyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B4NdLqIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NyCCWPj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J7z6VG5EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pfN8Yvg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3jqi8Ng5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bmujGmKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FiquX8npWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gCY8Ri9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func erSsmkfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3qVN6zpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8jlf1eOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kdSer63AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBAkSx3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lUcztnU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZG1NsPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJNTCkkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Usj7s5D2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yrD0wdA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FODQLve0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAKi7FoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7wl35XfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z6W76aggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7JVkGt1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sAAqf7GOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOkx7E7MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 84zMXci1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yFU5JT9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93ZA5F7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hxHVGN1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZD3e2wTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPOusrpSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Pd6yREnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAm5J7s2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INNdwgpTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rs4WGee5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H8XEW7LYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d28FNRx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJ9orXh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvjZ8oh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LXf38yPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqEiX4rfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DkyLnN6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DWQPZENHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUbUowH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1nCH8iWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iWqYVhECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9gqB4oQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xXlyC7YHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kwE01Yt3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OLR7Jlf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lr4GWew8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func znRQZUxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VvmAgq0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TMwkZJziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3PHGyK4kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4XfvefYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h13yIB3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kWm8ZvZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B2Q5iBodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZqPvjMNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8nvpGU6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmEoltNwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aq17c6neWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X55rqM6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DREXXiauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6eTHigoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nMTKrDnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ApH4Glf7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUU3eUGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vNPmqDW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPTSZAdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P73Pjwi2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ibw5urP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGoTRNPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ofwm3ubEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGagA8FeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N5FzGo3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSdY8DKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vBIFZ5MuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDDarJ8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFSe5vR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RzmKL5gdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3rFpm16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EnvsMfj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvfJ8a6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y7LlohVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5gpLCE1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nxg62pN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkhGwBzTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z53n3F3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a8tHrOpgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j0i5AxmoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ntEcJdg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5MjJvIA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func srPya01vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZ9hixXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ghDhLD0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uCMnBgVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5eG2aLAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZpUQ8TYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wk3JiZC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1FaJVpZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gnt2NnuwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUVELWfgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2dg1FamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mDEpiYF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsDinDqhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jb6tcGPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INQIBvqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SK4R3ZW5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptE5slLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8jYg0W8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DuhIsJxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JtGAjlCGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXHhUfusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7LBbn0jTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Nj9XLaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9dUaLtHEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

