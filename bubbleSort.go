package main

import (
    "fmt"
)

func main() {
    // ソートする値
    arr := []int{29, 10, 1, 37}
    // ソート終了判定するための配列
    arrTF := []bool{false, false, false, true}

    for {
        for i := 0; i < len(arr)-1; i++ {
            // 「左＞右」の場合、左右入れ替える
            if arr[i] > arr[i+1] {
                v1 := arr[i]
                v2 := arr[i+1]
                arr[i] = v2
                arr[i+1] = v1
            } else {
                // 「左＜右」になっていることを示すため、FをTにする
                arrTF[i] = true
            }
        }
        // ソート終了か判定する
        if finishSort(arrTF) {
            break
        }
    }

    fmt.Printf("Finish: %v %v\n", arr, arrTF)
}

func finishSort(arr []bool) bool {
    for _, v := range arr {
        if !v {
            return false
        }
    }
    return true
}