package utils

//Check if array contains a value
func IsContains[T comparable](arr []T, x T) bool {
    for _, v := range arr {
        if v == x {
            return true
        }
    }
    return false
}