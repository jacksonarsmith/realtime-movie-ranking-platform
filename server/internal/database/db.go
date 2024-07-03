// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.checkMovieExistsStmt, err = db.PrepareContext(ctx, checkMovieExists); err != nil {
		return nil, fmt.Errorf("error preparing query CheckMovieExists: %w", err)
	}
	if q.createMovieStmt, err = db.PrepareContext(ctx, createMovie); err != nil {
		return nil, fmt.Errorf("error preparing query CreateMovie: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.getMovieByFieldsStmt, err = db.PrepareContext(ctx, getMovieByFields); err != nil {
		return nil, fmt.Errorf("error preparing query GetMovieByFields: %w", err)
	}
	if q.getMovieByIdStmt, err = db.PrepareContext(ctx, getMovieById); err != nil {
		return nil, fmt.Errorf("error preparing query GetMovieById: %w", err)
	}
	if q.getMoviesStmt, err = db.PrepareContext(ctx, getMovies); err != nil {
		return nil, fmt.Errorf("error preparing query GetMovies: %w", err)
	}
	if q.getMoviesUpdatedMoreThanAnHourAgoStmt, err = db.PrepareContext(ctx, getMoviesUpdatedMoreThanAnHourAgo); err != nil {
		return nil, fmt.Errorf("error preparing query GetMoviesUpdatedMoreThanAnHourAgo: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.updateMovieStmt, err = db.PrepareContext(ctx, updateMovie); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMovie: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.checkMovieExistsStmt != nil {
		if cerr := q.checkMovieExistsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing checkMovieExistsStmt: %w", cerr)
		}
	}
	if q.createMovieStmt != nil {
		if cerr := q.createMovieStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createMovieStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.getMovieByFieldsStmt != nil {
		if cerr := q.getMovieByFieldsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMovieByFieldsStmt: %w", cerr)
		}
	}
	if q.getMovieByIdStmt != nil {
		if cerr := q.getMovieByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMovieByIdStmt: %w", cerr)
		}
	}
	if q.getMoviesStmt != nil {
		if cerr := q.getMoviesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMoviesStmt: %w", cerr)
		}
	}
	if q.getMoviesUpdatedMoreThanAnHourAgoStmt != nil {
		if cerr := q.getMoviesUpdatedMoreThanAnHourAgoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMoviesUpdatedMoreThanAnHourAgoStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.updateMovieStmt != nil {
		if cerr := q.updateMovieStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMovieStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                    DBTX
	tx                                    *sql.Tx
	checkMovieExistsStmt                  *sql.Stmt
	createMovieStmt                       *sql.Stmt
	createUserStmt                        *sql.Stmt
	getMovieByFieldsStmt                  *sql.Stmt
	getMovieByIdStmt                      *sql.Stmt
	getMoviesStmt                         *sql.Stmt
	getMoviesUpdatedMoreThanAnHourAgoStmt *sql.Stmt
	getUserStmt                           *sql.Stmt
	updateMovieStmt                       *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                    tx,
		tx:                                    tx,
		checkMovieExistsStmt:                  q.checkMovieExistsStmt,
		createMovieStmt:                       q.createMovieStmt,
		createUserStmt:                        q.createUserStmt,
		getMovieByFieldsStmt:                  q.getMovieByFieldsStmt,
		getMovieByIdStmt:                      q.getMovieByIdStmt,
		getMoviesStmt:                         q.getMoviesStmt,
		getMoviesUpdatedMoreThanAnHourAgoStmt: q.getMoviesUpdatedMoreThanAnHourAgoStmt,
		getUserStmt:                           q.getUserStmt,
		updateMovieStmt:                       q.updateMovieStmt,
	}
}
