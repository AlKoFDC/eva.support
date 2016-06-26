package handler

// MessageHandler is an interface, that handles messages.
type MessageHandler interface {
	Handle()
	Close()
}

// AsynchMessageHandler is an interface, that is designed to handle messages
// asynchronously.
type AsynchMessageHandler interface {
	// This starts the handling by starting the correct go routines
	// and doing setup and cleanup on closing the channel or receiving
	// anything on it.
	Start(chan struct{})
	// These functions synchronize the receiving and sending of messages.
	receiveMessage()
	sendMessage()
	// This function provides asynchronous message handling.
	handle()
	// Close frees underlying structures.
	close()
}
