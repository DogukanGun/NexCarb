package connectorHandlers

import (
	"SensorManager/database"
	"SensorManager/utils"
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/time/rate"
	"time"
)

type VerifyPayment struct {
	TxHash              string `json:"tx_hash"`
	SenderWallet        string `json:"sender_wallet"`
	IsPaymentWithSolana bool   `json:"chain"`
}

func VerifyPaymentHandler(c *fiber.Ctx) error {
	p := new(VerifyPayment)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	//receive database instance
	db, ok := c.Locals("db").(*database.Database)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).SendString("Database connection error")
	}
	if p.IsPaymentWithSolana {
		cluster := rpc.MainNetBeta
		rpcClient := rpc.NewWithCustomRPCClient(rpc.NewWithLimiter(
			cluster.RPC,
			rate.Every(time.Second), // time frame
			5,                       // limit of requests per time frame
		))
		version := uint64(0)
		tx, err := rpcClient.GetTransaction(
			context.TODO(),
			solana.MustSignatureFromBase58(p.TxHash),
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
				transfer.GetRecipientAccount().PublicKey.String() == p.SenderWallet {
				utils.RunWithHandlingError(db.Write("user_"+p.SenderWallet+"_status", "approved"))
				return c.Status(fiber.StatusAccepted).SendString("payment is verified")
			}
		}
		return c.Status(fiber.StatusForbidden).SendString("payment is not verified")
	} else {
		//TODO handle with ethena
		return c.SendStatus(500)
	}
}
