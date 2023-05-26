package tss_test

import (
	"encoding/hex"

	"github.com/bandprotocol/chain/v2/pkg/tss"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

func (suite *TSSTestSuite) TestComputeLagrangeCoefficient() {
	expValue := tss.ParseScalar(new(secp256k1.ModNScalar).SetInt(1140))
	value := tss.ComputeLagrangeCoefficient(
		tss.MemberID(3),
		[]tss.MemberID{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
	)

	suite.Require().Equal(expValue, value)
}

func (suite *TSSTestSuite) TestComputeOwnLo() {
	lo := tss.ComputeOwnLo(suite.member1.mid, suite.data, suite.bytes)
	suite.Require().Equal(suite.lo, lo)
}

func (suite *TSSTestSuite) TestComputeOwnPublicNonce() {
	pubNonce, err := tss.ComputeOwnPublicNonce(suite.member1.d.PublicKey, suite.member1.e.PublicKey, suite.lo)
	suite.Require().NoError(err)
	suite.Require().Equal(suite.member1.ownNonce.PublicKey, pubNonce)

	pubNonce, err = tss.ComputeOwnPublicNonce(suite.member2.d.PublicKey, suite.member2.e.PublicKey, suite.lo)
	suite.Require().NoError(err)
	suite.Require().Equal(suite.member2.ownNonce.PublicKey, pubNonce)
}

func (suite *TSSTestSuite) TestComputeOwnPrivateNonce() {
	privNonce, err := tss.ComputeOwnPrivateNonce(suite.member1.d.PrivateKey, suite.member1.e.PrivateKey, suite.lo)
	suite.Require().NoError(err)
	suite.Require().Equal(suite.member1.ownNonce.PrivateKey, privNonce)

	privNonce, err = tss.ComputeOwnPrivateNonce(suite.member2.d.PrivateKey, suite.member2.e.PrivateKey, suite.lo)
	suite.Require().NoError(err)
	suite.Require().Equal(suite.member2.ownNonce.PrivateKey, privNonce)
}

func (suite *TSSTestSuite) TestComputeGroupPublicNonce() {
	groupPubNonce, err := tss.ComputeGroupPublicNonce(
		suite.member1.ownNonce.PublicKey, suite.member2.ownNonce.PublicKey,
	)
	suite.Require().NoError(err)
	suite.Require().Equal(suite.groupPubNonce, groupPubNonce)
}

func (suite *TSSTestSuite) TestSignAndVerifySigning() {
	// Sign
	sig, err := tss.SignSigning(
		suite.groupPubNonce,
		suite.groupPubKey,
		suite.data,
		suite.member1.lagrange,
		suite.member1.ownNonce.PrivateKey,
		suite.member1.ownKey.PrivateKey,
	)
	suite.Require().NoError(err)

	// Success case
	err = tss.VerifySigningSig(
		suite.groupPubNonce,
		suite.groupPubKey,
		suite.data,
		suite.member1.lagrange,
		sig,
		suite.member1.ownKey.PublicKey,
	)
	suite.Require().NoError(err)
}

func (suite *TSSTestSuite) TestSignAndVerifyGroupSigning() {
	// Sign by member1
	sig1, err := tss.SignSigning(
		suite.groupPubNonce,
		suite.groupPubKey,
		suite.data,
		suite.member1.lagrange,
		suite.member1.ownNonce.PrivateKey,
		suite.member1.ownKey.PrivateKey,
	)
	suite.Require().NoError(err)

	// Sign by member2
	sig2, err := tss.SignSigning(
		suite.groupPubNonce,
		suite.groupPubKey,
		suite.data,
		suite.member2.lagrange,
		suite.member2.ownNonce.PrivateKey,
		suite.member2.ownKey.PrivateKey,
	)
	suite.Require().NoError(err)

	sig, err := tss.CombineSignatures(sig1, sig2)
	suite.Require().Equal(suite.groupPubNonce, tss.PublicKey(sig.R()))

	// Success case
	err = tss.VerifyGroupSigningSig(suite.groupPubKey, suite.data, sig)
	suite.Require().NoError(err)
}

func (suite *TSSTestSuite) TestGenerateMessageGroupSigning() {
	msg := tss.GenerateMessageGroupSigning(suite.groupPubKey, suite.data)
	suite.Require().
		Equal("7369676e696e6703534dfb533fedd09a97cbedeab70ae895399ed48be0ad7f789a705ec023dcf04464617461", hex.EncodeToString(msg))
}
