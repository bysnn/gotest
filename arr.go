package gotest

import(
"errors"
)


func ArrFirst(arr []int) (int, error) {
    if len(arr)>0 {
        return arr[0], nil
    } else {
        return 0, errors.New("arr is empty")
    }
}

func ArrEnd(arr []int) (int, error) {
    count=len(arr)
    if count>0 {
        return arr[count-1], nil
    } else {
        return 0, errors.New("arr is empty")
    }
    
}
