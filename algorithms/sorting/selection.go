// ----------------------------------------------------------------------------
// Algorithms - Selection Sort
// ----------------------------------------------------------------------------
//
// This is an example of Selection Sort. Selection sort does an in place sort
// on a set (slice). It's easiest to explain how it works with an example:
//
// If we have a set: [3, 1, 6, 0, 5]
//
// We can use selection sort to sort this set. First, we start at index 0 and
// see that we have 3. We then loop through indices 1 to len(set) finding the
// next smallest number. 1 is small, 6 is larger than 1 so we don't care, 0 is
// smaller than 1 so 0 is now the smallest, 5 is larger than 0 and we are at
// the end. This means at index 3 we have 0 as the smallest integer. We then
// swap the 3 with the 0 and move forward to the next index (1).
//
// We repeat this process of finding the smallest integer in the rest of the
// list and replacing it with our current position until we are fully sorted.
//
// While this process is certainly tedious, it does give us a way to
// programatticaly go through a set and sort it.
//
// ----------------------------------------------------------------------------
// Time Complexity
// ----------------------------------------------------------------------------
//
// Worst case: Θ(n^2)
// Best case: Θ(n^2)
// Average case: Θ(n^2)
//
// ----------------------------------------------------------------------------
package sort

// SelectionSort takes a slice of integers and sorts it in place.
func SelectionSort(data []int) {
	for i, _ := range data {
		min := MinIndex(data, i)
		Swap(data, i, min)
	}
}

// MinIndex finds the index of the next smallest number in a slice.
func MinIndex(data []int, initial int) int {
	minValue := data[initial]
	minIndex := initial

	for i := initial; i < len(data); i++ {
		if data[i] < minValue {
			minIndex = i
			minValue = data[i]
		}
	}

	return minIndex
}
