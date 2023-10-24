package globals

import (
	pb_account "github.com/PretendoNetwork/grpc-go/account"
	pb_friends "github.com/PretendoNetwork/grpc-go/friends"
	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/plogger-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var Logger = plogger.NewLogger()
var KerberosPassword = "password" // * Default password
var AuthenticationServer *nex.Server
var SecureServer *nex.Server
var GRPCAccountClientConnection *grpc.ClientConn
var GRPCAccountClient pb_account.AccountClient
var GRPCAccountCommonMetadata metadata.MD
var GRPCFriendsClientConnection *grpc.ClientConn
var GRPCFriendsClient pb_friends.FriendsClient
var GRPCFriendsCommonMetadata metadata.MD
