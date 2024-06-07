package pasta

import (
	"context"
	"crypto/md5"
	"fmt"
	"log/slog"
	"time"
)

type Pasta struct {
	Hash      string `json:"hash"`
	Pasta     string `json:"pasta"`
	CreatedAt int64  `json:"created_at"`
}

type repositoryInterface interface {
	GetAll(ctx context.Context) ([]*Pasta, error)
	GetByHash(ctx context.Context, hash string) (*Pasta, error)
	Store(ctx context.Context, pasta *Pasta) (*Pasta, error)
}

type PastaService struct {
	repository repositoryInterface
	logger     *slog.Logger
}

func NewPastaService(repository repositoryInterface, logger *slog.Logger) *PastaService {
	return &PastaService{
		repository: repository,
		logger:     logger,
	}
}

func (p *PastaService) GetAll(ctx context.Context) ([]*Pasta, error) {
	res, err := p.repository.GetAll(ctx)
	if err != nil {
		p.logger.Error(err.Error())
		return nil, err
	}

	return res, nil
}

func (p *PastaService) GetByHash(ctx context.Context, hash string) (*Pasta, error) {
	pasta, err := p.repository.GetByHash(ctx, hash)
	if err != nil {
		p.logger.Error(err.Error())
		return nil, err
	}

	return pasta, nil
}

func (p *PastaService) Store(ctx context.Context, pastaText string) (*Pasta, error) {
	pasta, err := p.repository.Store(ctx, &Pasta{
		Hash:      p.createHashFromPastaText(pastaText),
		Pasta:     pastaText,
		CreatedAt: time.Now().Unix(),
	})
	if err != nil {
		p.logger.Error(err.Error())
		return nil, err
	}

	return pasta, nil
}

func (p *PastaService) createHashFromPastaText(pastaText string) string {
	l := len(pastaText)
	if l > 20 {
		l = 20
	}

	hash := md5.Sum([]byte(pastaText[:len(pastaText)-l]))

	return fmt.Sprintf("%x", hash)
}
