package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	registerAndLogin "registerAndLogin/proto/registerAndLogin"
)

type RegisterAndLogin struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *RegisterAndLogin) Call(ctx context.Context, req *registerAndLogin.Request, rsp *registerAndLogin.Response) error {
	log.Info("Received RegisterAndLogin.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *RegisterAndLogin) Stream(ctx context.Context, req *registerAndLogin.StreamingRequest, stream registerAndLogin.RegisterAndLogin_StreamStream) error {
	log.Infof("Received RegisterAndLogin.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&registerAndLogin.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *RegisterAndLogin) PingPong(ctx context.Context, stream registerAndLogin.RegisterAndLogin_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&registerAndLogin.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
