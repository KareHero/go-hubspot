package hubspot

type Conversation struct {
	VisitorIdentification VisitorIdentificationService
}

func newConversation(c *Client) *Conversation {
	return &Conversation{
		VisitorIdentification: &VisitorIdentificationServiceOp{
			client: c,
		},
	}
}
