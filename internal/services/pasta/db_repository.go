package pasta

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type dBRepository struct {
	conn *sql.DB
}

func NewDBRepository(conn *sql.DB) *dBRepository {
	return &dBRepository{
		conn: conn,
	}
}

func (r *dBRepository) GetAll() ([]*Pasta, error) {
	result := make([]*Pasta, 0, 1000)

	rows, err := r.conn.Query("select * from pastes")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var h string
		var p string
		var c int64

		err = rows.Scan(&h, &p, &c)
		if err != nil {
			fmt.Println(err)
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

func (r *dBRepository) GetByHash(hash string) (*Pasta, error) {
	stmt, err := r.conn.Prepare("select * from pastes where hash = ?")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var h string
	var p string
	var c int64

	err = stmt.QueryRow(hash).Scan(&h, &p, &c)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer stmt.Close()

	return &Pasta{
		Hash:      h,
		Pasta:     p,
		CreatedAt: c,
	}, nil
}

func (r *dBRepository) Store(pasta *Pasta) (*Pasta, error) {
	stmt, err := r.conn.Prepare("insert into pastes (hash, pasta, created_at) values (?,?,?);")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = stmt.Exec(pasta.Hash, pasta.Pasta, pasta.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return pasta, nil
}
