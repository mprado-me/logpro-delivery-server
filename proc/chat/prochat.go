package prochat

import (
	"mem/short"
	"msg"
)

func Process(inChatMsg msg.InChat) {
	outMsg := msg.OutChat{
		SenderEmail: inChatMsg.Sender.Email,
		Text:        inChatMsg.Text,
	}
	// TODO: Enviar para todos os usuario ativos
	smem.MsgTo("test@test.com", outMsg)
}
