package nex

import (
	"fmt"
	"os"
	"strconv"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/steel-diver-sub-wars/globals"
)

var serverBuildString string

func StartAuthenticationServer() {
	globals.AuthenticationServer = nex.NewPRUDPServer()

	globals.AuthenticationEndpoint = nex.NewPRUDPEndPoint(1)
	globals.AuthenticationEndpoint.ServerAccount = globals.AuthenticationServerAccount
	globals.AuthenticationEndpoint.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.AuthenticationEndpoint.AccountDetailsByUsername = globals.AccountDetailsByUsername
	globals.AuthenticationServer.BindPRUDPEndPoint(globals.AuthenticationEndpoint)

	globals.AuthenticationServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(3, 7, 0))
	globals.AuthenticationServer.ByteStreamSettings.UseStructureHeader = true
	globals.AuthenticationServer.AccessKey = "fb9537fe"

	globals.AuthenticationEndpoint.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		fmt.Println("== Steel Diver: Sub Wars - Auth ==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID)
		fmt.Printf("Method ID: %#v\n", request.MethodID)
		fmt.Println("==================================")
	})

	registerCommonAuthenticationServerProtocols()

	port, _ := strconv.Atoi(os.Getenv("PN_SDSB_AUTHENTICATION_SERVER_PORT"))

	globals.AuthenticationServer.Listen(port)
}
