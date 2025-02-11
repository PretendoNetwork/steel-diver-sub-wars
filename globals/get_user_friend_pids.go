package globals

import (
	"context"

	pb_friends "github.com/PretendoNetwork/grpc-go/friends"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	"google.golang.org/grpc/metadata"
)

func GetUserFriendPIDs(pid uint32) []uint32 {
	ctx := metadata.NewOutgoingContext(context.Background(), GRPCFriendsCommonMetadata)

	response, err := GRPCFriendsClient.GetUserFriendPIDs(ctx, &pb_friends.GetUserFriendPIDsRequest{Pid: pid})
	if err != nil {
		globals.Logger.Error(err.Error())
		return make([]uint32, 0)
	}

	return response.Pids
}
