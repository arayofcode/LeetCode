func nextGreatestLetter(letters []byte, target byte) byte {
    for _, letter := range letters {
        if letter > target {
            return letter
        }
    }
    return letters[0]
}

func nextGreatestLetterBinarySearchImplementation(letters []byte, target byte) byte {
    low, high := 0, len(letters) - 1
    mid := low + (high - low) / 2
    for low <= high {
        if letters[mid] == target {
            return nextLetter(letters, target, mid)
        } else if letters[mid] > target {
            high = mid - 1
        } else {
            low = mid + 1
        }
        mid = low + (high - low) / 2
    }
    if low < len(letters) {
        return letters[low]
    } else {
        return letters[0]
    }
}

func nextLetter(letters []byte, target byte, pos int) byte {
    for i := pos + 1; i < len(letters); i++ {
        if letters[i] != target {
            return letters[i]
        }
    }
    return letters[0]
}