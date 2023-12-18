package service

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/TechBowl-japan/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	stmt, err := s.db.PrepareContext(ctx, insert)
	if err != nil {
		return nil, err
	}

	var id int64
	sqlRes, err := stmt.ExecContext(ctx, subject, description)
	if err != nil {
		return nil, err
	}

	id, err = sqlRes.LastInsertId()

	var res model.TODO
	res.ID = id
	err = s.db.QueryRowContext(ctx, confirm, id).Scan(&res.Subject, &res.Description, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)

	var rows *sql.Rows

	fmt.Println(prevID,size)

	if size == 0 {
		// bit を反転させて、最大値を取得する
		size = int64(^uint64(0) >> 1)
	}

	if prevID == 0 {
		r , err := s.db.QueryContext(ctx, read, size)
		if err != nil {
			return nil, err
		}
		rows = r
	} else {
		r , err := s.db.QueryContext(ctx, readWithID, prevID, size)
		if err != nil {
			return nil, err
		}
		rows = r
	}

	defer rows.Close()

	res := make([]*model.TODO, 0)

	for rows.Next() {
		todo := &model.TODO{}
		err := rows.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, todo)
		fmt.Println(todo)
	}

	err := rows.Err()

	if err != nil {
		return nil, err
	}

	fmt.Println(res)

	return res, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject string, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	if id == 0 {
		return nil, &model.ErrNotFound{Message: "id is not found"}
	}


	stmt, err := s.db.PrepareContext(ctx, update)
	if err != nil {
		return nil, err
	}

	sqlRes, err := stmt.ExecContext(ctx, subject, description, id)
	if err != nil {
		return nil, err
	}

	row, err := sqlRes.RowsAffected()
	if err != nil {
		return nil, err
	}

	if row == 0 {
		return nil, fmt.Errorf("Not changed")
	}

	var res model.TODO
	res.ID = id
	err = s.db.QueryRowContext(ctx, confirm, id).Scan(&res.Subject, &res.Description, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	if len(ids) == 0 {
		return nil
	}

	queryStr := strings.Repeat(",?", len(ids))
	delete := fmt.Sprintf(deleteFmt, queryStr)

	fmt.Println(len(ids),delete)

	stmt, err := s.db.PrepareContext(ctx, delete)
	if err != nil {
		return err
	}

	args := make([]interface{}, 0)

	for _, id := range ids {
		args = append(args, id)
	}

	fmt.Println(args)

	sqlRes, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}

	row, err := sqlRes.RowsAffected()

	if row == 0 {
		return &model.ErrNotFound{Message: "Not changed"}
	}

	return nil
}
