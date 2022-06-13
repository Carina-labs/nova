package keeper

import (
	"github.com/Carina-labs/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetWithdrawRecord(ctx sdk.Context) {

}

func (k Keeper) SetWithdrawRecord(ctx sdk.Context, record types.WithdrawRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyWithdrawRecordInfo)
	bz := k.cdc.MustMarshal(&record)
	store.Set([]byte(record.ZoneId+record.Withdrawer), bz)
}

func (k Keeper) DeleteWithdrawRecord(ctx sdk.Context) {

}

func (k Keeper) GetWithdrawReceipt(ctx sdk.Context) {

}

func (k Keeper) SetWithdrawReceipt(ctx sdk.Context, receipt types.MsgWithdrawReceipt) {

}

func (k Keeper) DeleteWithdrawReceipt(ctx sdk.Context) {

}
