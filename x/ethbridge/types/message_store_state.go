package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStoreState = "store_state"

var _ sdk.Msg = &MsgStoreState{}

func NewMsgStoreState(creator string, address string, slot string) *MsgStoreState {
	return &MsgStoreState{
		Creator: creator,
		Address: address,
		Slot:    slot,
	}
}

func (msg *MsgStoreState) Route() string {
	return RouterKey
}

func (msg *MsgStoreState) Type() string {
	return TypeMsgStoreState
}

func (msg *MsgStoreState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStoreState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStoreState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
