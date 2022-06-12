package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	user "user/proto/user"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) Info(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Info("Received User.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

