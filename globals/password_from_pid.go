package globals

import (
	"context"

	pb_account "github.com/PretendoNetwork/grpc-go/account"
	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	"google.golang.org/grpc/metadata"
)

func PasswordFromPID(pid uint32) (string, uint32) {
	ctx := metadata.NewOutgoingContext(context.Background(), GRPCAccountCommonMetadata)

	response, err := GRPCAccountClient.GetNEXData(ctx, &pb_account.GetNEXDataRequest{Pid: pid})
	if err != nil {
		globals.Logger.Error(err.Error())
		return "", nex.Errors.RendezVous.InvalidUsername
	}

	// * We only allow tester accounts for now
	if response.AccessLevel < 1 {
		globals.Logger.Errorf("PID %d is not a tester!", response.Pid)
		return "", nex.Errors.RendezVous.AccountDisabled
	}

	return response.Password, 0
}
