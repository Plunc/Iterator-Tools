package lib

//Generators
func Apply[T, R any](Iterable []T, Closure func(T) R) []R {
	result := make([]R, len(Iterable))
	for index, value := range Iterable {
		result[index] = Closure(value)
	}

	return result
}

func Filter[T any](Iterable []T, Closure func(T) bool) []T {
	var result []T
	for _, value := range Iterable {
		if Closure(value) {
			result = append(result, value)
		}
	}

	return result
}

//Booleans
func Any(Slice []bool) bool {
	for _, value := range Slice {
		if value {
			return true
		}
	}

	return false
}

func All(Slice []bool) bool {
	for _, value := range Slice {
		if !value {
			return false
		}
	}

	return true
}

//Searchers
type Number interface {
	~uint | ~int | ~float32 | ~float64
}

func Min[T Number](Slice []T) T {
	smallest := Slice[0]
	for _, value := range Slice {
		if value <= smallest {
			smallest = value
		}
	}

	return smallest
}

func Max[T Number](Slice []T) T {
	largest := Slice[0]
	for _, value := range Slice {
		if value >= largest {
			largest = value
		}
	}

	return largest
}

func minByKey[T any, R Number](Slice []T, Closure func(T) R) T {
	smallest := Slice[0]
	for _, value := range Slice {
		if Closure(value) <= Closure(smallest) {
			smallest = value
		}
	}

	return smallest
}

func maxByKey[T any, R Number](Slice []T, Closure func(T) R) T {
	largest := Slice[0]
	for _, value := range Slice {
		if Closure(value) >= Closure(largest) {
			largest = value
		}
	}

	return largest
}

//Sorters
func Sort[T Number](Slice []T) []T{
	copy := Slice
    for pass := range copy[:len(copy)-1] { // Outer loop using range
        for currentIndex := range copy[:len(copy)-1-pass] { // Inner loop using range
            nextIndex := currentIndex + 1
            if copy[currentIndex] > copy[nextIndex] {
                // Swap elements if they are out of order
                copy[currentIndex], copy[nextIndex] = copy[nextIndex], copy[currentIndex]
            }
        }
    }
	return copy
}

func sortByKey[T any, R Number](Slice []T, Closure func(T) R) []T{
	copy := Slice
	for pass := range copy[:len(copy)-1] { // Outer loop using range
        for currentIndex := range copy[:len(copy)-1-pass] { // Inner loop using range
            nextIndex := currentIndex + 1
            if Closure(copy[currentIndex]) > Closure(copy[nextIndex]) {
                // Swap elements if they are out of order
                copy[currentIndex], copy[nextIndex] = copy[nextIndex], copy[currentIndex]
            }
        }
    }
	return copy
}
