package pasta

type Pasta struct {
	id string
}

type repositoryInterface interface {
	GetAll() []Pasta
	GetById(id string) *Pasta
	Store(pasta Pasta)
}

type PastaService struct {
	repository repositoryInterface
}

func NewPastaService() *PastaService {
	return &PastaService{
		repository: newDBRepository(),
	}
}
