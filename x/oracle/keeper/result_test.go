package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bandprotocol/bandchain/chain/pkg/obi"
	"github.com/bandprotocol/bandchain/chain/x/oracle/types"
)

func TestResultBasicFunctions(t *testing.T) {
	_, ctx, k := createTestInput()

	req := types.NewOracleRequestPacketData("alice", 1, BasicCalldata, 1, 1)
	res := types.NewOracleResponsePacketData("alice", 1, 1, 1589535020, 1589535022, 1, BasicCalldata)
	encodedResult := obi.MustEncode(req, res)
	k.SetResult(ctx, types.RequestID(1), encodedResult)

	// Test GetResult func
	result, err := k.GetResult(ctx, types.RequestID(1))
	require.NoError(t, err)
	require.Equal(t, result, types.Result{RequestPacketData: req, ResponsePacketData: res})

	// Test MustGetResult func
	result = k.MustGetResult(ctx, types.RequestID(1))
	require.NoError(t, err)
	require.Equal(t, result, types.Result{RequestPacketData: req, ResponsePacketData: res})

	_, err = k.GetResult(ctx, types.RequestID(2))
	require.Error(t, err)

	require.Panics(t, func() { k.MustGetResult(ctx, types.RequestID(2)) })

	// Test HasResult func
	require.True(t, k.HasResult(ctx, types.RequestID(1)))
	require.False(t, k.HasResult(ctx, types.RequestID(2)))
}

func TestGetAllResults(t *testing.T) {
	_, ctx, k := createTestInput()

	req1 := types.NewOracleRequestPacketData("alice", 1, BasicCalldata, 1, 1)
	res1 := types.NewOracleResponsePacketData("alice", 1, 1, 1589535020, 1589535022, 1, BasicCalldata)
	encodedResult1 := obi.MustEncode(req1, res1)
	k.SetResult(ctx, types.RequestID(1), encodedResult1)
	req4 := types.NewOracleRequestPacketData("bob", 1, BasicCalldata, 1, 1)
	res4 := types.NewOracleResponsePacketData("bob", 4, 1, 1589535020, 1589535022, 1, BasicCalldata)
	encodedResult4 := obi.MustEncode(req4, res4)
	k.SetResult(ctx, types.RequestID(4), encodedResult4)

	results := k.GetAllResults(ctx)

	require.Equal(t, 4, len(results))
	require.Equal(t, encodedResult1, results[0])

	// result of reqID 2 and 3 should be nil
	require.Empty(t, results[1])
	require.Empty(t, results[2])

	require.Equal(t, encodedResult4, results[3])
}
