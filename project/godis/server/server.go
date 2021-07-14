package server

import (
	"context"
	"example/project/godis/config"
	"example/project/godis/interface/db"
	"example/project/godis/sync/atomic"
	"fmt"
	"net"
	"sync"
)

var (
	unknownErrReplyBytes = []byte("-ERR unknow\r\n")
)

type Handler struct {
	activeConn sync.Map
	db         db.DB
	closing    atomic.Boolean
}

func (h *Handler) Handle(ctx context.Context, conn net.Conn) {

}

func (h *Handler) Close() error {
	return nil
}


func MakeHandler() *Handler {
	// TODO 还没写
	var db *db.DB
	if config.Properties.Self != "" && len(config.Properties.Peers) > 0 {
		db = nil
	}else {
		db = nil
	}
	fmt.Println(db)
	return nil
}
