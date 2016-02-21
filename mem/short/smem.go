// TODO: smem contem um map dos motoboys ativos
package smem

type OutMsgs []interface{}

var outMsgsByReceiverEmail map[string]OutMsgs

func init() {
	outMsgsByReceiverEmail = make(map[string]OutMsgs)
}

func MsgTo(email string, msg interface{}) {
	_, ok := outMsgsByReceiverEmail[email]
	if ok {
		outMsgsByReceiverEmail[email] = append(outMsgsByReceiverEmail[email], msg)
	} else {
		outMsgsByReceiverEmail[email] = OutMsgs{msg}
	}
}

func GetMsgsTo(email string) (OutMsgs, bool) {
	msgs, ok := outMsgsByReceiverEmail[email]
	return msgs, ok
}

func ClearMsgsTo(email string) {
	delete(outMsgsByReceiverEmail, email)
}
