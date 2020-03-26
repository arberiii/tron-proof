package tron_test

import (
	"crypto/sha256"
	"testing"

	"github.com/arberiii/tron-proof"

	"github.com/gogo/protobuf/proto"
	"github.com/sasaxie/go-client-api/service"
)

// 18318292
func TestCreateMerkleTree(t *testing.T) {
	for i := int64(18318273); i < 18318274; i++ {
		client := service.NewGrpcClient("grpc.trongrid.io:50051")
		client.Start()

		block := client.GetBlockByNum(i)
		client.Conn.Close()
		var hashes [][]byte
		for _, tx := range block.Transactions {
			rawData, err := proto.Marshal(tx)
			if err != nil {
				t.Fatal(err)
			}
			h256h := sha256.New()
			h256h.Write(rawData)
			hash := h256h.Sum(nil)

			hashes = append(hashes, hash)
		}

		m := tron.CreateMerkleTree(hashes)
		if !cmp(m.Root(), block.GetBlockHeader().GetRawData().GetTxTrieRoot()) {
			t.Errorf("merkle tree of block %d is not calculated correctly", i)
		}
	}
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
