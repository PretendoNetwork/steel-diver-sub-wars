package nex

import (
	"os"
	"strconv"

	nex "github.com/PretendoNetwork/nex-go"
	ticket_granting "github.com/PretendoNetwork/nex-protocols-common-go/ticket-granting"
	"github.com/PretendoNetwork/steel-diver-sub-wars/globals"
)

func registerCommonAuthenticationServerProtocols() {
	ticketGrantingProtocol := ticket_granting.NewCommonTicketGrantingProtocol(globals.AuthenticationServer)

	port, _ := strconv.Atoi(os.Getenv("PN_SDSB_SECURE_SERVER_PORT"))

	secureStationURL := nex.NewStationURL("")
	secureStationURL.SetScheme("prudps")
	secureStationURL.SetAddress(os.Getenv("PN_SDSB_SECURE_SERVER_HOST"))
	secureStationURL.SetPort(uint32(port))
	secureStationURL.SetCID(1)
	secureStationURL.SetPID(2)
	secureStationURL.SetSID(1)
	secureStationURL.SetStream(10)
	secureStationURL.SetType(2)

	ticketGrantingProtocol.SetSecureStationURL(secureStationURL)
	ticketGrantingProtocol.SetBuildName(serverBuildString)

	globals.AuthenticationServer.SetPasswordFromPIDFunction(globals.PasswordFromPID)
}
