package main

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {

	sender := &fedex.FedexSender{}
	SendBook("foo", sender)
}
