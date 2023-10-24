package nex

import (
	"fmt"
	"os"

	"github.com/PretendoNetwork/steel-diver-sub-wars/globals"

	nex "github.com/PretendoNetwork/nex-go"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewServer()
	globals.SecureServer.SetPRUDPVersion(1)
	globals.SecureServer.SetPRUDPProtocolMinorVersion(4)
	globals.SecureServer.SetDefaultNEXVersion(nex.NewNEXVersion(3, 7, 0))
	globals.SecureServer.SetKerberosPassword(globals.KerberosPassword)
	globals.SecureServer.SetAccessKey("fb9537fe")

	globals.SecureServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("== Steel Diver: Sub Wars - Secure ==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("====================================")
	})

	registerCommonSecureServerProtocols()

	globals.SecureServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_SDSB_SECURE_SERVER_PORT")))
}
