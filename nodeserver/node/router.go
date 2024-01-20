package node

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) RegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionReq: // 服务注册
		s.ServiceConnectionReq(serviceMsg)
	case cmd.GetServerOuterAddrReq: // 心跳
		s.GetServerOuterAddrReq(serviceMsg)
	case cmd.PlayerLoginReq: // 玩家登录通知
		s.PlayerLoginReq(serviceMsg)
	}
}