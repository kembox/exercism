package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if len(l1) == 0 {
		if len(l2) == 0 {
			return RelationEqual
		} else {
			return RelationSublist
		}
	}
	if len(l2) == 0 {
		return RelationSuperlist
	}

	if len(l1) == len(l2) {
		if equalList(l1, l2) {
			return RelationEqual
		}
	} else if len(l1) < len(l2) {
		if isSubList(l1, l2) {
			return RelationSublist
		} else {
			return RelationUnequal
		}
	} else if len(l1) > len(l2) {
		if isSubList(l2, l1) {
			return RelationSuperlist
		} else {
			return RelationUnequal
		}
	}

	return RelationUnequal
	//should never happen, I'm just too lazy to indent if blocks
}

func equalList(l1 []int, l2 []int) bool {
	count := 0
	for i := 0; i < len(l1); i++ {
		if l1[i] == l2[i] {
			count += 1
		}
	}
	return count == len(l1)
}

func isSubList(small []int, big []int) bool {
	isSubList := 0
	for i := 0; i < (len(big) - len(small) + 1); i++ {
		if equalList(big[i:i+len(small)], small) {
			isSubList = 1
			break
		}
	}
	return isSubList == 1
}
