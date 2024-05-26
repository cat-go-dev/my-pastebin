package pasta

type dBRepository struct {
	// todo: add connection
}

// GetAll() []Pasta
// GetById(id string) *Pasta
// Store(pasta Pasta)

func newDBRepository() *dBRepository {
	return &dBRepository{}
}

func (r *dBRepository) GetAll() []Pasta {
	result := make([]Pasta, 5, 5)

	return result
}

func (r *dBRepository) GetById(id string) *Pasta {
	return &Pasta{id: id}
}

func (r *dBRepository) Store(pasta Pasta) {

}
