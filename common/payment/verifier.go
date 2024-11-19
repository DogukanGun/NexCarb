package payment

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"golang.org/x/time/rate"
	"time"
)

func VerifyTheTransactions(txHash string, wallet string) bool {
	cluster := rpc.MainNetBeta
	rpcClient := rpc.NewWithCustomRPCClient(rpc.NewWithLimiter(
		cluster.RPC,
		rate.Every(time.Second), // time frame
		5,                       // limit of requests per time frame
	))
	version := uint64(0)
	tx, err := rpcClient.GetTransaction(
		context.TODO(),
		solana.MustSignatureFromBase58(txHash),
		&rpc.GetTransactionOpts{
			MaxSupportedTransactionVersion: &version,
			Encoding:                       solana.EncodingBase64,
		},
	)
	if err != nil {
		panic(err)
	}
	parsed, err := tx.Transaction.GetTransaction()
	if err != nil {
		panic(err)
	}
	spew.Dump(tx)
	for _, instr := range parsed.Message.Instructions {

		// Find the accounts of this instruction:
		accounts, err := instr.ResolveInstructionAccounts(&parsed.Message)
		if err != nil {
			panic(err)
		}

		inst, err := system.DecodeInstruction(accounts, instr.Data)
		if err != nil {
			//panic(err)
			continue
		}

		spew.Dump(inst)
		if transfer, ok := inst.Impl.(*system.Transfer); ok &&
			transfer.GetRecipientAccount().PublicKey.String() == wallet {
			return true
		}
	}
	return false
}
