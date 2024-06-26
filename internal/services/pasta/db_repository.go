package pasta

import (
	"context"
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

type dBRepository struct {
	conn   *sql.DB
	logger *slog.Logger
}

func NewDBRepository(conn *sql.DB, logger *slog.Logger) *dBRepository {
	return &dBRepository{
		conn:   conn,
		logger: logger,
	}
}

func (r *dBRepository) GetAll(ctx context.Context) ([]*Pasta, error) {
	result := make([]*Pasta, 0, 1000)

	rows, err := r.conn.Query("select * from pastes")
	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var h string
		var p string
		var c int64

		err = rows.Scan(&h, &p, &c)
		if err != nil {
			r.logger.Error(err.Error())
			return nil, err
		}

		result = append(result, &Pasta{
			Hash:      h,
			Pasta:     p,
			CreatedAt: c,
		})
	}

	return result, nil
}

func (r *dBRepository) GetByHash(ctx context.Context, hash string) (*Pasta, error) {
	stmt, err := r.conn.Prepare("select * from pastes where hash = ?")
	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	var h string
	var p string
	var c int64

	err = stmt.QueryRow(hash).Scan(&h, &p, &c)
	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}
	defer stmt.Close()

	// todo: not found error

	return &Pasta{
		Hash:      h,
		Pasta:     p,
		CreatedAt: c,
	}, nil
}

func (r *dBRepository) Store(ctx context.Context, pasta *Pasta) (*Pasta, error) {
	stmt, err := r.conn.Prepare("insert into pastes (hash, pasta, created_at) values (?,?,?);")
	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	// todo: unique error
	// todo: random hash

	_, err = stmt.Exec(pasta.Hash, pasta.Pasta, pasta.CreatedAt)
	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return pasta, nil
}
