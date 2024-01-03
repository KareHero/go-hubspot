package hubspot

const (
	visitorIdentificationBasePath = "/conversations/v3/visitor-identification"
)

type IdentificationTokenResponse struct {
	Token string `json:"token"`
}

type IdentificationTokenRequest struct {
	FirstName string
	LastName  string
	Email     string
}

type VisitorIdentificationService interface {
	GenerateIdentificationToken(request *IdentificationTokenRequest) (*IdentificationTokenResponse, error)
}

type VisitorIdentificationServiceOp struct {
	client *Client
}

func (s *VisitorIdentificationServiceOp) GenerateIdentificationToken(request *IdentificationTokenRequest) (*IdentificationTokenResponse, error) {
	response := &IdentificationTokenResponse{}
	path := visitorIdentificationBasePath + "/tokens/create"
	if err := s.client.Post(path, request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Embed the VisitorIdentificationService interface in VisitorIdentificationServiceOp
var _ VisitorIdentificationService = &VisitorIdentificationServiceOp{}
