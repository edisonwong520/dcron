package driver

import (
	"context"

	redis "github.com/redis/go-redis/v9"
)

// There is only one driver for one dcron.
// Tips for write a user-defined Driver by yourself.
//  1. Confirm that `Stop` and `Start` can be called for more times.
//  2. Must make `GetNodes` will return error when timeout.
type DriverV2 interface {
	// init driver
	Init(serviceName string, opts ...Option)
	// get nodeID
	NodeID() string
	// get nodes
	GetNodes(ctx context.Context) (nodes []string, err error)

	// register node to remote server (like redis),
	// will create a goroutine to keep the connection.
	// And then continue for other work.
	Start(ctx context.Context) (err error)

	// stop the goroutine of keep connection.
	Stop(ctx context.Context) (err error)

	withOption(opt Option) (err error)
}

func NewRedisDriver(redisClient *redis.Client) DriverV2 {
	return newRedisDriver(redisClient)
}

func NewRedisZSetDriver(redisClient *redis.Client) DriverV2 {
	return newRedisZSetDriver(redisClient)
}
