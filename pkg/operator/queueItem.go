package operator

type queueItem struct {
	operation string
	item      string
	initiator string
	message   string
}

func (qi queueItem) getOperation() string {
	return qi.operation
}
