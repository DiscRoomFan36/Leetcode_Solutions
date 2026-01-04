
package main


func sumOfThree(num int64) []int64 {
    if num % 3 != 0 { return []int64{} }

    div_3 := num/3
    return []int64{div_3-1, div_3, div_3+1}
}
