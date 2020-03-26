package tron

type node struct {
	Hash []byte
	Dir  bool
}

func (m *MerkleTree) GenerateProof(i int) []*node {
	var ret []*node
	level := m.Size
	pointer := 0
	index := i
	for pointer < len(m.Leaves) {
		if index%2 == 1 {
			ret = append(ret, &node{
				Hash: m.Leaves[i-1].Hash,
				Dir:  false,
			})
		} else {
			if index+1 < level {
				ret = append(ret, &node{
					Hash: m.Leaves[i+1].Hash,
					Dir:  true,
				})
			}
		}
		index = index / 2
		pointer += level
		i = pointer + index
		if level%2 == 0 {
			level = level / 2
		} else {
			level = (level / 2) + 1
		}
	}

	return ret
}

func VerifyProof(elem, root []byte, proof []*node) bool {
	for _, n := range proof {
		if n.Dir == false {
			elem = computeHash(n.Hash, elem)
		} else {
			elem = computeHash(elem, n.Hash)
		}
	}
	return cmp(elem, root)
}

func cmp(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
