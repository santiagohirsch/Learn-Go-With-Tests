package arrayswithgenerics

func Reduce[A, B any](collection []A, action func(B, A) B, initialValue B) B {
	var toReturn = initialValue

	for _, x := range collection {
		toReturn = action(toReturn, x)
	}

	return toReturn
}

func Find[A any](colleciton []A, predicate func(A) bool) (elem A, found bool) {
	for _, elem := range colleciton {
		if predicate(elem) {
			return elem, true
		}
	}
	return
}
