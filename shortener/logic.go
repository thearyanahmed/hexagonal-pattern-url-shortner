package shortener

import (
	"errors"
	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
	"time"
)

var (
	RedirectNotFound = errors.New("Redirect not found.")
	InvalidRedirect = errors.New("Invalid redirect.")
)

// redirect service needs to have Find & Store method
// Our repository contract / interface also has the same methods
// this is a customer redirect service that implements the redirect repository
// which also needs to satisfy the Find & Store methods
type redirectService struct {
	redirectRepo RedirectRepository
}

// generate a new redirect service insntance with the repository
// the repository is outside the inner hexagon
func NewRedirectService(redirectRepository RedirectRepository) RedirectService {
	return &redirectService{
		redirectRepository,
	}
}

// satisfy the interface
func (r *redirectService) Find(code string) (*Redirect, error) {
	return r.redirectRepo.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		return errs.Wrap(InvalidRedirect,"service.Redirect.store")
	}

	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()

	return r.redirectRepo.Store(redirect)
}