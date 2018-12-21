package types

import (
	"fmt"
	"github.com/irisnet/irishub-sync/conf/server"
	"github.com/irisnet/irishub-sync/logger"
	"github.com/irisnet/irishub-sync/store"
	"github.com/irisnet/irishub/client/utils"
	"github.com/irisnet/irishub/codec"
	"github.com/irisnet/irishub/modules/auth"
	"github.com/irisnet/irishub/modules/bank"
	"github.com/irisnet/irishub/modules/distribution"
	dtags "github.com/irisnet/irishub/modules/distribution/tags"
	"github.com/irisnet/irishub/modules/gov"
	"github.com/irisnet/irishub/modules/gov/tags"
	"github.com/irisnet/irishub/modules/slashing"
	"github.com/irisnet/irishub/modules/stake"
	staketypes "github.com/irisnet/irishub/modules/stake/types"
	"github.com/irisnet/irishub/modules/upgrade"
	"github.com/irisnet/irishub/types"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tm "github.com/tendermint/tendermint/types"
	"regexp"
	"strconv"
	"strings"
)

type (
	MsgTransfer = bank.MsgSend

	MsgStakeCreate                 = stake.MsgCreateValidator
	MsgStakeEdit                   = stake.MsgEditValidator
	MsgStakeDelegate               = stake.MsgDelegate
	MsgStakeBeginUnbonding         = stake.MsgBeginUnbonding
	MsgBeginRedelegate             = stake.MsgBeginRedelegate
	MsgUnjail                      = slashing.MsgUnjail
	MsgSetWithdrawAddress          = distribution.MsgSetWithdrawAddress
	MsgWithdrawDelegatorReward     = distribution.MsgWithdrawDelegatorReward
	MsgWithdrawDelegatorRewardsAll = distribution.MsgWithdrawDelegatorRewardsAll
	MsgWithdrawValidatorRewardsAll = distribution.MsgWithdrawValidatorRewardsAll
	StakeValidator                 = stake.Validator
	Delegation                     = stake.Delegation
	UnbondingDelegation            = stake.UnbondingDelegation

	MsgDeposit        = gov.MsgDeposit
	MsgSubmitProposal = gov.MsgSubmitProposal
	MsgVote           = gov.MsgVote
	Proposal          = gov.Proposal
	SdkVote           = gov.Vote

	ResponseDeliverTx = abci.ResponseDeliverTx

	StdTx      = auth.StdTx
	SdkCoins   = types.Coins
	KVPair     = types.KVPair
	AccAddress = types.AccAddress
	ValAddress = types.ValAddress
	Dec        = types.Dec
	Validator  = tm.Validator
	Tx         = tm.Tx
	Block      = tm.Block
	BlockMeta  = tm.BlockMeta
	HexBytes   = cmn.HexBytes

	ABCIQueryOptions = rpcclient.ABCIQueryOptions
	Client           = rpcclient.Client
	HTTP             = rpcclient.HTTP
	ResultStatus     = ctypes.ResultStatus
)

var (
	ValidatorsKey        = stake.ValidatorsKey
	GetValidatorKey      = stake.GetValidatorKey
	GetDelegationKey     = stake.GetDelegationKey
	GetDelegationsKey    = stake.GetDelegationsKey
	GetUBDKey            = stake.GetUBDKey
	GetUBDsKey           = stake.GetUBDsKey
	TagProposalID        = tags.ProposalID
	TagReward            = dtags.Reward
	ValAddressFromBech32 = types.ValAddressFromBech32

	UnmarshalValidator      = staketypes.UnmarshalValidator
	MustUnmarshalValidator  = staketypes.MustUnmarshalValidator
	UnmarshalDelegation     = staketypes.UnmarshalDelegation
	MustUnmarshalDelegation = staketypes.MustUnmarshalDelegation
	MustUnmarshalUBD        = staketypes.MustUnmarshalUBD

	Bech32ifyValPub      = types.Bech32ifyValPub
	RegisterCodec        = types.RegisterCodec
	AccAddressFromBech32 = types.AccAddressFromBech32
	BondStatusToString   = types.BondStatusToString

	NewDecFromStr = types.NewDecFromStr

	AddressStoreKey   = auth.AddressStoreKey
	GetAccountDecoder = utils.GetAccountDecoder

	KeyProposal      = gov.KeyProposal
	KeyVotesSubspace = gov.KeyVotesSubspace

	NewHTTP = rpcclient.NewHTTP

	cdc *codec.Codec
)

// 初始化账户地址前缀
func init() {
	config := types.GetConfig()
	config.SetBech32PrefixForAccount(server.Bech32.PrefixAccAddr, server.Bech32.PrefixAccPub)
	config.SetBech32PrefixForValidator(server.Bech32.PrefixValAddr, server.Bech32.PrefixValPub)
	config.SetBech32PrefixForConsensusNode(server.Bech32.PrefixAccAddr, server.Bech32.PrefixConsPub)
	config.Seal()

	cdc = codec.New()

	bank.RegisterCodec(cdc)
	stake.RegisterCodec(cdc)
	slashing.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	gov.RegisterCodec(cdc)
	upgrade.RegisterCodec(cdc)
	distribution.RegisterCodec(cdc)

	types.RegisterCodec(cdc)

	codec.RegisterCrypto(cdc)
}

func GetCodec() *codec.Codec {
	return cdc
}

//
func ParseCoins(coinsStr string) (coins store.Coins) {
	coinsStr = strings.TrimSpace(coinsStr)
	if len(coinsStr) == 0 {
		return
	}

	coinStrs := strings.Split(coinsStr, ",")
	for _, coinStr := range coinStrs {
		coin := ParseCoin(coinStr)
		coins = append(coins, coin)
	}
	return coins
}

func ParseCoin(coinStr string) (coin store.Coin) {
	var (
		reDnm  = `[A-Za-z\-]{2,15}`
		reAmt  = `[0-9]+[.]?[0-9]*`
		reSpc  = `[[:space:]]*`
		reCoin = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reAmt, reSpc, reDnm))
	)

	coinStr = strings.TrimSpace(coinStr)

	matches := reCoin.FindStringSubmatch(coinStr)
	if matches == nil {
		logger.Error("invalid coin expression", logger.Any("coin", coinStr))
		return coin
	}
	denom, amount := matches[2], matches[1]

	amt, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		logger.Error("Convert str to int failed", logger.Any("amount", amount))
		return coin
	}

	return store.Coin{
		Denom:  denom,
		Amount: amt,
	}
}

func BuildFee(fee auth.StdFee) store.Fee {
	return store.Fee{
		Amount: ParseCoins(fee.Amount.String()),
		Gas:    int64(fee.Gas),
	}
}
