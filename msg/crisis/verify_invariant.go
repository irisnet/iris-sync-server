package crisis

import (
	"encoding/json"
	"github.com/irisnet/irishub-sync/store/document"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/irishub-sync/types"
	. "github.com/irisnet/irishub-sync/util/constant"
	"github.com/irisnet/irishub-sync/store"
)

type DocMsgVerifyInvariant struct {
	Sender              string `bson:"sender"`
	InvariantModuleName string `bson:"invariant_module_name" yaml:"invariant_module_name"`
	InvariantRoute      string `bson:"invariant_route" yaml:"invariant_route"`
}

func (m *DocMsgVerifyInvariant) Type() string {
	return TxTypeVerifyInvariant
}

func (m *DocMsgVerifyInvariant) BuildMsg(v interface{}) {
	var msg types.MsgVerifyInvariant
	data, _ := json.Marshal(v)
	json.Unmarshal(data, &msg)

	m.Sender = msg.Sender.String()
	m.InvariantModuleName = msg.InvariantModuleName
	m.InvariantRoute = msg.InvariantRoute

}

func (m *DocMsgVerifyInvariant) HandleTxMsg(msgData sdk.Msg, tx *document.CommonTx) *document.CommonTx {

	m.BuildMsg(msgData)
	tx.Msgs = append(tx.Msgs, document.DocTxMsg{
		Type: m.Type(),
		Msg:  m,
	})
	tx.Addrs = append(tx.Addrs, m.Sender)
	tx.Types = append(tx.Types, m.Type())
	if len(tx.Msgs) > 1 {
		return tx
	}
	tx.Type = m.Type()
	tx.From = m.Sender
	tx.To = ""
	tx.Amount = []store.Coin{}
	return tx
}
