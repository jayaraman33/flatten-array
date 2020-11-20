package flatten

func Flatten(nested interface{}) []interface{} {

	flatten := []interface{}{}

	switch n := nested.(type) {
	case nil:
	case []interface{}:
		for _, v := range n {
			flatten = append(flatten, Flatten(v)...)
		}
	default:
		flatten = append(flatten, n)
	}

	return flatten
}


