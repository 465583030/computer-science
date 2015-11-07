// ----------------------------------------------------------------------------
// Algorithms - Insertion Sort
// ----------------------------------------------------------------------------
//
// This is an example of Insertion Sort. It takes a set of data starting at
// the second index. As it moves through the array, it checks the values
// before the current index for a value less than itself and if it finds it,
// it inserts itself in that index. For example:
//
// If we have a set: [3, 1, 6, 0, 5]
//
// We can use insertion sort to sort this. We start at index 1 with value 1.
// We move backwards to index 0 and see that 3 is greater than 1. Since this
// is the beginning of the set we place our current value (1) at index 0 and
// move the value at index 0 to index 1. Our new set is:
//
// [1, 3, 6, 0, 5]
//
// One the next iteration, we move to index 2 where we see the value is 6. We
// check index 1 and 0 and see that both values are smaller than 6 so we don't
// do anything. Our set remains the same:
//
// [1, 3, 6, 0, 5]
//
// On the third iteration, we start at index 4 and see the value 0. We check
// the values at index 3, 2, and 1 and see that they are all larger than our
// value so we move the set to the right and place our value at the front of
// the set giving us a new set:
//
// [0, 1, 3, 6, 5]
//
// On our last iteration, we start at index 5 and see the value 5. We check
// index 4 and see the value greater than 5 but at index 3 we see the value is
// less than our value (3). We stop there, move the 6 over to the right and
// place 5 in the index 4. Our set is now sorted!
//
// ----------------------------------------------------------------------------
// Time Complexity
// ----------------------------------------------------------------------------
//
// Worst case: Θ(n^2)
// Best case: Θ(n)
// Average case: Θ(n^2)
//
// ----------------------------------------------------------------------------
package sort

// InsertionSort takes a slice of integers and sorts it in place.
func InsertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		Insert(data, i-1, data[i])
	}
}

// Insert loops backwards through and array until it finds a value less than
// itself and places itself directly after it.
func Insert(data []int, rightBound, val int) {
	var i int
	for i = rightBound; i >= 0 && data[i] > val; i-- {
		data[i+1] = data[i]
	}
	data[i+1] = val
}
