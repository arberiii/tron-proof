package tron_test

import (
	"crypto/sha256"
	"testing"

	"github.com/arberiii/tron-proof"
	"github.com/gogo/protobuf/proto"
	"github.com/sasaxie/go-client-api/service"
)

func TestMerkleProof(t *testing.T) {
	client := service.NewGrpcClient("grpc.trongrid.io:50051")
	client.Start()

	block := client.GetBlockByNum(17318273)
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
	for i := range hashes {
		proof := m.GenerateProof(i)
		if !tron.VerifyProof(hashes[i], m.Root(), proof) {
			t.Errorf("merkle proof of transaction %d is not correct", i)
		}
	}

}
