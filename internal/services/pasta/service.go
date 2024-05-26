package pasta

import (
	"crypto/md5"
	"fmt"
	"time"
)

type Pasta struct {
	Hash      string `json:"hash"`
	Pasta     string `json:"pasta"`
	CreatedAt int64  `json:"created_at"`
}

type repositoryInterface interface {
	GetAll() []Pasta
	GetByHash(hash string) (*Pasta, error)
	Store(pasta *Pasta) (*Pasta, error)
}

type PastaService struct {
	repository repositoryInterface
}

func NewPastaService(repository repositoryInterface) *PastaService {
	return &PastaService{
		repository: repository,
	}
}

func (p *PastaService) GetByHash(hash string) (*Pasta, error) {
	pasta, err := p.repository.GetByHash(hash)
	if err != nil {
		return nil, err
	}

	return pasta, nil
}

func (p *PastaService) Store(pastaText string) (*Pasta, error) {
	pasta, err := p.repository.Store(&Pasta{
		Hash:      p.createHashFromPastaText(pastaText),
		Pasta:     pastaText,
		CreatedAt: time.Now().Unix(),
	})
	if err != nil {
		return nil, err
	}

	return pasta, nil
}

func (p *PastaService) createHashFromPastaText(pastaText string) string {
	l := len(pastaText)
	if l > 20 {
		l = 20
	}

	hash := md5.Sum([]byte(pastaText[:l]))

	return fmt.Sprintf("%x", hash)
}
