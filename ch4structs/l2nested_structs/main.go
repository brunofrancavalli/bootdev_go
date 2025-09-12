package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	//if( mToSend == null ) {
	//	return false
	//}
	//if( (mToSend.sender == null) || (mToSend.recipient == null) ) {
	//	return false
	//}
	if( ( mToSend.sender.name == "" ) ||  mToSend.recipient.name == "" ) {
		return false
	}
	if( (mToSend.sender.number <= 0) || (mToSend.recipient.number <= 0) ) {
		return false
	}
	if( mToSend.message == "" ) {
		return false
	}
	return true
}
