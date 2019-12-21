package payment

const salePath = "/1/sales"

// ICieloRepository is a interface for Cielo Repository
type ICieloRepository interface {
	Sale(CieloRequestBody) (*CieloResponseBody, error)
}

// CieloRepository repository to comunicate with cielo
type CieloRepository struct {
	r IHTTPRequester
}

// NewCieloRepository constructor
func NewCieloRepository(r IHTTPRequester) ICieloRepository {
	return &CieloRepository{
		r: r,
	}
}

// Sale .
func (c *CieloRepository) Sale(crb CieloRequestBody) (*CieloResponseBody, error) {
	out := &CieloResponseBody{}
	if err := c.r.Post(salePath, crb, out); err != nil {
		return out, err
	}
	return out, nil
}
