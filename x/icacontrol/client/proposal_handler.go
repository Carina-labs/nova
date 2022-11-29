package client

import (
	"github.com/Carina-labs/nova/v2/x/icacontrol/client/cli"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	"net/http"
)

var (
	ZoneProposalHandler = govclient.NewProposalHandler(cli.NewProposalZoneTxCmd, emptyRestHandler)
)

func emptyRestHandler(client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "unsupported-icacontrol-client",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Legacy REST Routes are not supported for icacontrol proposals")
		},
	}
}
