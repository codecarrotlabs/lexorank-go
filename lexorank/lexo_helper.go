package lexorank

func ArrayCopy(source []interface{}, sourceIndex int, destination []interface{}, destinationIndex int, length int) {
    for i := 0; i < length; i++ {
        if sourceIndex+i < len(source) && destinationIndex+i < len(destination) {
            destination[destinationIndex+i] = source[sourceIndex+i]
        }
    }
}