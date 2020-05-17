package boilerplate

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/and07/boilerplate-go/api/gen-boilerplate-go/api"
	"github.com/golang/protobuf/jsonpb"
	_struct "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BlockchainServer ...
type BlockchainServer struct {
	eventBus <-chan interface{}
}

// NewBlockchainServer ...
func NewBlockchainServer(eventBus <-chan interface{}) *BlockchainServer {
	return &BlockchainServer{eventBus: eventBus}
}

// Address ...
func (b *BlockchainServer) Address(_ context.Context, req *api.AddressRequest) (*api.AddressResponse, error) {
	if req.Address != "Mxb9a117e772a965a3fddddf83398fd8d71bf57ff6" {
		return &api.AddressResponse{}, status.Error(codes.FailedPrecondition, "wallet not found")
	}
	return &api.AddressResponse{
		Balance: map[string]string{
			"BIP": "12345678987654321",
		},
		TransactionsCount: "120",
	}, nil
}

// Subscribe ...
func (b *BlockchainServer) Subscribe(req *api.SubscribeRequest, stream api.BlockchainService_SubscribeServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case event := <-b.eventBus:
			byteData, err := json.Marshal(event)
			if err != nil {
				return err
			}
			var bb bytes.Buffer
			bb.Write(byteData)
			data := &_struct.Struct{Fields: make(map[string]*_struct.Value)}
			if err := (&jsonpb.Unmarshaler{}).Unmarshal(&bb, data); err != nil {
				return err
			}

			if err := stream.Send(&api.SubscribeResponse{
				Query: req.Query,
				Data:  data,
				Events: []*api.SubscribeResponse_Event{
					{
						Key:    "tx.hash",
						Events: []string{"01EFD8EEF507A5BFC4A7D57ECA6F61B96B7CDFF559698639A6733D25E2553539"},
					},
				},
			}); err != nil {
				return err
			}
		case <-time.After(5 * time.Second):
			return nil
		}
	}
}
