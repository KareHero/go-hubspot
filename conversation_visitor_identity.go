package hubspot

const (
	visitorIdentificationBasePath = "/conversations/v3/visitor-identification"
)

type IdentificationTokenResponse struct {
	Token string `json:"token"`
}

type IdentificationTokenRequest struct {
	firstName string
	lastName  string
	email     string
}

// Getter methods
func (r *IdentificationTokenRequest) FirstName() string {
	return r.firstName
}

func (r *IdentificationTokenRequest) LastName() string {
	return r.lastName
}

func (r *IdentificationTokenRequest) Email() string {
	return r.email
}

// Setter methods
func (r *IdentificationTokenRequest) SetFirstName(firstName string) {
	r.firstName = firstName
}

func (r *IdentificationTokenRequest) SetLastName(lastName string) {
	r.lastName = lastName
}

func (r *IdentificationTokenRequest) SetEmail(email string) {
	r.email = email
}

type VisitorIdentificationService interface {
	GenerateIdentificationToken(request *IdentificationTokenRequest) (*IdentificationTokenResponse, error)
}

type VisitorIdentificationServiceOp struct {
	client *Client
}

var _ VisitorIdentificationService = (*VisitorIdentificationServiceOp)(nil)

func (s *VisitorIdentificationServiceOp) GenerateIdentificationToken(request *IdentificationTokenRequest) (*IdentificationTokenResponse, error) {
	response := &IdentificationTokenResponse{}
	path := visitorIdentificationBasePath + "/tokens/create"
	if err := s.client.Post(path, request, response); err != nil {
		return nil, err
	}
	return response, nil
}
