package tron

import (
	"crypto/sha256"
)

type MerkleTree struct {
	Leaves []*leaf
}

type leaf struct {
	Hash  []byte
	left  *leaf
	right *leaf
}

func CreateMerkleTree(hashes [][]byte) *MerkleTree {
	m := &MerkleTree{}
	leaves := m.createLeaves(hashes)

	for len(leaves) > 1 {
		leaves = m.createParentLeaves(leaves)
	}

	return m
}

func (m *MerkleTree) createParentLeaves(list []*leaf) []*leaf {
	var ret []*leaf
	step := 2
	length := len(list)
	for i := 0; i < length; i += step {
		if i+1 < length {
			ret = append(ret, m.createLeaf2(list[i], list[i+1]))
		} else {
			ret = append(ret, m.createLeaf1(list[i].Hash))
		}
	}
	return ret
}

func (m *MerkleTree) createLeaves(list [][]byte) []*leaf {
	var ret []*leaf
	step := 2
	length := len(list)
	for i := 0; i < length; i += step {
		if i+1 < length {
			ret = append(ret, m.createLeaf2(m.createLeaf1(list[i]), m.createLeaf1(list[i+1])))
		} else {
			ret = append(ret, m.createLeaf1(list[i]))
		}
	}
	return ret
}

func (m *MerkleTree) createLeaf1(hash []byte) *leaf {
	l := &leaf{Hash: hash}
	m.Leaves = append(m.Leaves, l)
	return l
}

func (m *MerkleTree) createLeaf2(left, right *leaf) *leaf {
	l := &leaf{}
	if len(right.Hash) == 0 {
		l.Hash = left.Hash
	} else {
		l.Hash = computeHash(left.Hash, right.Hash)
	}
	l.left = left
	l.right = right
	m.Leaves = append(m.Leaves, l)
	return l
}

func computeHash(left, right []byte) []byte {
	h := sha256.Sum256([]byte(string(left[:]) + string(right[:])))
	return h[:]
}

func (m *MerkleTree) Root() []byte {
	return m.Leaves[len(m.Leaves)-1].Hash
}
