// Package repository handles interactions with the persistence layer
package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/mcrors/secret-santa-picker-server/domain"
)

const groupColumns = "uuid, name, created_at"
const groupInsertReturning = "INSERT INTO groups (uuid, name) VALUES ($1, $2) RETURNING " + groupColumns
const groupSelect = "SELECT " + groupColumns + " FROM groups"
const groupSelectByUUID = groupSelect + " WHERE uuid = $1"

type Groups struct {
	db DBTX
}

func NewGroupRepository(db DBTX) *Groups {
	return &Groups{
		db: db,
	}
}

func (g *Groups) ListGroups(ctx context.Context) ([]domain.Group, error) {
	rows, err := g.db.QueryContext(ctx, groupSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]domain.Group, 0)
	for rows.Next() {
		record, err := scanGroup(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, record)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (g *Groups) GetGroup(ctx context.Context, id domain.ID) (domain.Group, error) {
	row := g.db.QueryRowContext(ctx, groupSelectByUUID, id)
	result, err := scanGroup(row)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return domain.Group{}, domain.ErrGroupNotFound
	case err != nil:
		return domain.Group{}, err
	default:
		return result, nil
	}
}

func (g *Groups) CreateGroup(ctx context.Context, group domain.Group) (domain.Group, error) {
	row := g.db.QueryRowContext(ctx, groupInsertReturning, group.ID, group.Name)
	result, err := scanGroup(row)
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23505" {
		return domain.Group{}, domain.ErrGroupConflict
	}
	if err != nil {
		return domain.Group{}, err
	}
	return result, nil
}

func (g *Groups) RenameGroup(ctx context.Context, id domain.ID, newName string) error {
	query := "UPDATE groups SET name = $1, updated_at = NOW() WHERE uuid = $2"
	result, err := g.db.ExecContext(ctx, query, newName, id)
	if err != nil {
		return err
	}
	return checkRowsAffected(result.RowsAffected())
}

func (g *Groups) DeleteGroup(ctx context.Context, id domain.ID) error {
	query := "DELETE FROM groups WHERE uuid = $1"
	result, err := g.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return checkRowsAffected(result.RowsAffected())
}

func checkRowsAffected(rowCount int64, err error) error {
	if err != nil {
		return err
	}
	if rowCount == 0 {
		return domain.ErrGroupNotFound
	}
	return nil
}

func scanGroup(s scanner) (domain.Group, error) {
	var result domain.Group
	err := s.Scan(
		&result.ID,
		&result.Name,
		&result.CreatedAt,
	)
	return result, err
}
