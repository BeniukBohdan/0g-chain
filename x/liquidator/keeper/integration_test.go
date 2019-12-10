package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/kava-labs/kava/app"
	"github.com/kava-labs/kava/x/cdp"
	"github.com/kava-labs/kava/x/liquidator"
	"github.com/kava-labs/kava/x/liquidator/types"
	"github.com/kava-labs/kava/x/pricefeed"
)

// Avoid cluttering test cases with long function name
func i(in int64) sdk.Int                    { return sdk.NewInt(in) }
func d(str string) sdk.Dec                  { return sdk.MustNewDecFromStr(str) }
func c(denom string, amount int64) sdk.Coin { return sdk.NewInt64Coin(denom, amount) }
func cs(coins ...sdk.Coin) sdk.Coins        { return sdk.NewCoins(coins...) }

// Genesis states to initialize test apps

func NewPricefeedGenState(asset string, price sdk.Dec) app.GenesisState {
	pfGenesis := pricefeed.GenesisState{
		Params: pricefeed.Params{
			Markets: []pricefeed.Market{
				pricefeed.Market{MarketID: asset, BaseAsset: asset, QuoteAsset: "usd", Oracles: pricefeed.Oracles{}, Active: true}},
		},
		PostedPrices: []pricefeed.PostedPrice{
			pricefeed.PostedPrice{
				MarketID:      asset,
				OracleAddress: sdk.AccAddress{},
				Price:         price,
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{pricefeed.ModuleName: pricefeed.ModuleCdc.MustMarshalJSON(pfGenesis)}
}

func NewCDPGenState() app.GenesisState {
	cdpGenesis := cdp.GenesisState{
		Params: cdp.CdpParams{
			GlobalDebtLimit: sdk.NewInt(1000000),
			CollateralParams: []cdp.CollateralParams{
				{
					Denom:            "btc",
					LiquidationRatio: sdk.MustNewDecFromStr("1.5"),
					DebtLimit:        sdk.NewInt(500000),
				},
			},
		},
		GlobalDebt: sdk.ZeroInt(),
		CDPs:       cdp.CDPs{},
	}
	return app.GenesisState{cdp.ModuleName: cdp.ModuleCdc.MustMarshalJSON(cdpGenesis)}
}

func NewLiquidatorGenState() app.GenesisState {
	liquidatorGenesis := types.GenesisState{
		Params: types.LiquidatorParams{
			DebtAuctionSize: sdk.NewInt(1000),
			CollateralParams: []types.CollateralParams{
				{
					Denom:       "btc",
					AuctionSize: sdk.NewInt(1),
				},
			},
		},
	}
	return app.GenesisState{liquidator.ModuleName: liquidator.ModuleCdc.MustMarshalJSON(liquidatorGenesis)}
}
