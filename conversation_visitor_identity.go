package hubspot

const (
	visitorIdentificationBasePath = "/conversations/v3/visitor-identification"
)

type VisitorIdentificationServiceOp struct {
	client *Client
}

type IdentificationTokenResponse struct {
	Token string `json:"token"`
}

type IdentificationTokenRequest struct {
	FirstName string
	LastName  string
	Email     string
}

type VisitorIdentificationService interface {
	GenerateIdentificationToken(request IdentificationTokenRequest) (*IdentificationTokenResponse, error)
}

func (s *VisitorIdentificationServiceOp) GenerateIdentificationToken(request IdentificationTokenRequest) (*IdentificationTokenResponse, error) {
	response := &IdentificationTokenResponse{}
	path := visitorIdentificationBasePath + "/tokens/create"
	if err := s.client.Post(path, request, response); err != nil {
		return nil, err
	}
	return response, nil
}
