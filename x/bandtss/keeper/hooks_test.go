package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	bandtesting "github.com/bandprotocol/chain/v2/testing"
	"github.com/bandprotocol/chain/v2/x/tss/types"
)

func (s *KeeperTestSuite) TestAfterSigningFailed() {
	ctx, k := s.ctx, s.app.BandtssKeeper
	hook := k.Hooks()

	testCases := []struct {
		name     string
		signing  types.Signing
		expCoins sdk.Coins
	}{
		{
			"10uband with 2 members",
			types.Signing{
				ID:      1,
				GroupID: 1,
				AssignedMembers: []types.AssignedMember{
					{MemberID: 1},
					{MemberID: 2},
				},
			},
			sdk.NewCoins(sdk.NewInt64Coin("uband", 20)),
		},
		{
			"10uband,15token with 2 members",
			types.Signing{
				ID:      1,
				GroupID: 1,
				AssignedMembers: []types.AssignedMember{
					{MemberID: 1},
					{MemberID: 2},
				},
			},
			sdk.NewCoins(sdk.NewInt64Coin("uband", 20), sdk.NewInt64Coin("token", 30)),
		},
		{
			"0uband with 2 members",
			types.Signing{
				ID:      2,
				GroupID: 1,
				AssignedMembers: []types.AssignedMember{
					{MemberID: 1},
					{MemberID: 2},
				},
			},
			sdk.NewCoins(),
		},
		{
			"10uband with 0 member",
			types.Signing{
				ID:              3,
				GroupID:         1,
				AssignedMembers: []types.AssignedMember{},
			},
			sdk.NewCoins(),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			balancesBefore := s.app.BankKeeper.GetAllBalances(ctx, bandtesting.FeePayer.Address)
			balancesModuleBefore := s.app.BankKeeper.GetAllBalances(ctx, k.GetBandtssAccount(ctx).GetAddress())

			err := hook.AfterSigningFailed(ctx, tc.signing)
			s.Require().NoError(err)

			balancesAfter := s.app.BankKeeper.GetAllBalances(ctx, bandtesting.FeePayer.Address)
			balancesModuleAfter := s.app.BankKeeper.GetAllBalances(ctx, k.GetBandtssAccount(ctx).GetAddress())

			gain := balancesAfter.Sub(balancesBefore...)
			s.Require().Equal(tc.expCoins, gain)

			lose := balancesModuleBefore.Sub(balancesModuleAfter...)
			s.Require().Equal(tc.expCoins, lose)
		})
	}
}
