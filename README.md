# tron-proof
This is a merkle-proof package that supports tron transaction proofs.

### Installation
```
go get github.com/arberiii/tron-proof
```

### Usage
```go
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
log.Println(m.Root())

// Generate proof for transaction number 0
proof := m.GenerateProof(0)
if !tron.VerifyProof(hashes[0], m.Root(), proof) {
    t.Errorf("merkle proof of transaction %d is not correct", 0)
}

```
