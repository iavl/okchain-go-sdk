package staking

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/module/staking/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/crypto/keys"
	"github.com/okex/okchain-go-sdk/types/params"
	"github.com/okex/okchain-go-sdk/utils"
)

// Delegate delegates okt for voting
func (sc stakingClient) Delegate(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	coin, err := sdk.ParseDecCoin(coinsStr)
	if err != nil {
		return resp, fmt.Errorf("failed : parse Coins [%s] error: %s", coinsStr, err)
	}

	msg := types.NewMsgDelegate(fromInfo.GetAddress(), coin)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// Unbond unbonds the delegation on okchain
func (sc stakingClient) Unbond(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	coin, err := sdk.ParseDecCoin(coinsStr)
	if err != nil {
		return resp, fmt.Errorf("failed : parse Coins [%s] error: %s", coinsStr, err)
	}

	msg := types.NewMsgUndelegate(fromInfo.GetAddress(), coin)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// Vote votes to the some specific validators
func (sc stakingClient) Vote(fromInfo keys.Info, passWd string, valAddrsStr []string, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckVoteParams(fromInfo, passWd, valAddrsStr); err != nil {
		return
	}

	valAddrs, err := utils.ParseValAddresses(valAddrsStr)
	if err != nil {
		return resp, fmt.Errorf("failed. validator address parsed error: %s", err.Error())
	}

	msg := types.NewMsgVote(fromInfo.GetAddress(), valAddrs)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// DestroyValidator deregisters the validator and unbond the min-self-delegation
func (sc stakingClient) DestroyValidator(fromInfo keys.Info, passWd string, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	msg := types.NewMsgDestroyValidator(fromInfo.GetAddress())

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// CreateValidator creates a new validator
func (sc stakingClient) CreateValidator(fromInfo keys.Info, passWd, pubkeyStr, moniker, identity, website, details,
	memo string, accNum, seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	pubkey, err := sdk.GetConsPubKeyBech32(pubkeyStr)
	if err != nil {
		return
	}

	description := types.NewDescription(moniker, identity, website, details)
	msg := types.NewMsgCreateValidator(sdk.ValAddress(fromInfo.GetAddress()), pubkey, description)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// EditValidator edits the description on a validator by the owner
func (sc stakingClient) EditValidator(fromInfo keys.Info, passWd, moniker, identity, website, details, memo string, accNum,
	seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	description := types.NewDescription(moniker, identity, website, details)
	msg := types.NewMsgEditValidator(sdk.ValAddress(fromInfo.GetAddress()), description)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// RegisterProxy registers the identity of proxy
func (sc stakingClient) RegisterProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	msg := types.NewMsgRegProxy(fromInfo.GetAddress(), true)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// UnregisterProxy registers the identity of proxy
func (sc stakingClient) UnregisterProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	msg := types.NewMsgRegProxy(fromInfo.GetAddress(), false)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// BindProxy binds the staking tokens to a proxy
func (sc stakingClient) BindProxy(fromInfo keys.Info, passWd, proxyAddrStr, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckSendParams(fromInfo, passWd, proxyAddrStr); err != nil {
		return
	}

	proxyAddr, err := sdk.AccAddressFromBech32(proxyAddrStr)
	if err != nil {
		return resp, fmt.Errorf("failed. parse Address [%s] error: %s", proxyAddrStr, err)
	}

	msg := types.NewMsgBindProxy(fromInfo.GetAddress(), proxyAddr)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// UnbindProxy unbinds the staking tokens from a proxy
func (sc stakingClient) UnbindProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	msg := types.NewMsgUnbindProxy(fromInfo.GetAddress())

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}
