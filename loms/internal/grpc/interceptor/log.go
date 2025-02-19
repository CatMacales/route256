package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
)

func Logger(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	rawReq, _ := protojson.Marshal((req).(proto.Message))
	log.Printf("request: method: %s, req: %s\n", info.FullMethod, string(rawReq))

	resp, err = handler(ctx, req)

	return
}
