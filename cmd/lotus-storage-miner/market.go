package main

import (
	"fmt"

	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
	"gopkg.in/urfave/cli.v2"
)

var marketCmd = &cli.Command{
	Name:  "market",
	Usage: "Interact with the storage market as a miner",
	Subcommands: []*cli.Command{
		setPriceCmd,
	},
}

var setPriceCmd = &cli.Command{
	Name:  "set-price",
	Usage: "Set price that miner will accept storage deals at",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := lcli.DaemonContext(cctx)

		if !cctx.Args().Present() {
			return fmt.Errorf("must specify price to set")
		}

		fp, err := types.ParseFIL(cctx.Args().First())
		if err != nil {
			return err
		}

		return api.MarketSetPrice(ctx, types.BigInt(fp))
	},
}
