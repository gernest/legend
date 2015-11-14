package core

type BaseTalk struct {
	req Payload
	res chan Payload
}

func NewTalk(req Payload) Talk {
	return &BaseTalk{
		req: req,
		res: make(chan Payload),
	}
}

func (b *BaseTalk) Request() Payload {
	return b.req
}

func (b *BaseTalk) Respond() chan Payload {
	return b.res
}
