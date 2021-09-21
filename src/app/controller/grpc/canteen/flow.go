package canteen

import (
	"context"
	"funnel/app/service/canteen_service"
	rpc "funnel/rpc"
)

type CanteenRpc struct{}

func (CanteenRpc) CanteenFolw(ctx context.Context, in *rpc.CanteenFolwRequest) (*rpc.CanteenFolwReply, error) {
	data,
		err := canteen_service.FetchFlow()
	if err != nil {
		return nil, err
	}
	var rpcData []*rpc.CanteenFolwReply_Canteen

	for _, item := range data.Data.Data {
		rpcData = append(rpcData, &rpc.CanteenFolwReply_Canteen{RestaurantName: item.RestaurantName, DealCount: item.DealCount})
	}

	return &rpc.CanteenFolwReply{Data: rpcData, BeginTime: data.Data.BeginTime, EndTime: data.Data.EndTime}, nil
}
