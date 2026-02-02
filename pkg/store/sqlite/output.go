package store

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// ClearOutputs clear all outputs
func (store *SQLiteStore) ClearOutputs(ctx context.Context) error {
	_, err := store.db.ExecContext(ctx, "DELETE FROM outputs")
	return err
}

// GetOutput returns a stored Output.
func (store *SQLiteStore) GetOutput(ctx context.Context, ID string) (*model.OutputDef, error) {
	var output model.OutputDef
	var propsJSON, filtersJSON sql.NullString

	err := store.db.QueryRowContext(ctx, `
		SELECT id, alias, name, description, condition, props, filters, enabled, nb_success, nb_error
		FROM outputs WHERE id = ?
	`, ID).Scan(
		&output.ID,
		&output.Alias,
		&output.Spec.Name,
		&output.Spec.Desc,
		&output.Condition,
		&propsJSON,
		&filtersJSON,
		&output.Enabled,
		&output.NbSuccess,
		&output.NbError,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrOutputNotFound
		}
		return nil, err
	}

	// Unmarshal props
	if propsJSON.Valid && propsJSON.String != "" {
		if err := json.Unmarshal([]byte(propsJSON.String), &output.Props); err != nil {
			return nil, err
		}
	}

	// Unmarshal filters
	if filtersJSON.Valid && filtersJSON.String != "" {
		if err := json.Unmarshal([]byte(filtersJSON.String), &output.Filters); err != nil {
			return nil, err
		}
	}

	return &output, nil
}

// DeleteOutput removes a output.
func (store *SQLiteStore) DeleteOutput(ctx context.Context, ID string) (*model.OutputDef, error) {
	output, err := store.GetOutput(ctx, ID)
	if err != nil {
		return nil, err
	}

	_, err = store.db.ExecContext(ctx, "DELETE FROM outputs WHERE id = ?", ID)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (store *SQLiteStore) assertOutputQuota(ctx context.Context, output *model.OutputDef) error {
	if store.quota.MaxNbOutputs > 0 {
		var exists bool
		err := store.db.QueryRowContext(ctx, "SELECT COUNT(*) > 0 FROM outputs WHERE id = ?", output.ID).Scan(&exists)
		if err != nil {
			return err
		}

		if !exists {
			var total int
			err := store.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM outputs").Scan(&total)
			if err != nil {
				return err
			}
			if total >= store.quota.MaxNbOutputs {
				return common.ErrOutputQuotaExceeded
			}
		}
	}
	return nil
}

// SaveOutput stores a output.
func (store *SQLiteStore) SaveOutput(ctx context.Context, output model.OutputDef) (*model.OutputDef, error) {
	if err := store.assertOutputQuota(ctx, &output); err != nil {
		return nil, err
	}

	// Marshal props to JSON
	propsJSON, err := json.Marshal(output.Props)
	if err != nil {
		return nil, err
	}

	// Marshal filters to JSON
	filtersJSON, err := json.Marshal(output.Filters)
	if err != nil {
		return nil, err
	}

	_, err = store.db.ExecContext(ctx, `
		INSERT INTO outputs (id, alias, name, description, condition, props, filters, enabled, nb_success, nb_error)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			alias = excluded.alias,
			name = excluded.name,
			description = excluded.description,
			condition = excluded.condition,
			props = excluded.props,
			filters = excluded.filters,
			enabled = excluded.enabled,
			nb_success = excluded.nb_success,
			nb_error = excluded.nb_error
	`,
		output.ID,
		output.Alias,
		output.Spec.Name,
		output.Spec.Desc,
		output.Condition,
		string(propsJSON),
		string(filtersJSON),
		output.Enabled,
		output.NbSuccess,
		output.NbError,
	)

	if err != nil {
		return nil, err
	}

	return &output, nil
}

// ListOutputs returns a paginated list of outputs.
func (store *SQLiteStore) ListOutputs(ctx context.Context, page, limit int) (*model.OutputDefCollection, error) {
	offset := (page - 1) * limit

	rows, err := store.db.QueryContext(ctx, `
		SELECT id, alias, name, description, condition, props, filters, enabled, nb_success, nb_error
		FROM outputs
		LIMIT ? OFFSET ?
	`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	outputs := model.OutputDefCollection{}

	for rows.Next() {
		var output model.OutputDef
		var propsJSON, filtersJSON sql.NullString

		err := rows.Scan(
			&output.ID,
			&output.Alias,
			&output.Spec.Name,
			&output.Spec.Desc,
			&output.Condition,
			&propsJSON,
			&filtersJSON,
			&output.Enabled,
			&output.NbSuccess,
			&output.NbError,
		)
		if err != nil {
			return nil, err
		}

		// Unmarshal props
		if propsJSON.Valid && propsJSON.String != "" {
			if err := json.Unmarshal([]byte(propsJSON.String), &output.Props); err != nil {
				return nil, err
			}
		}

		// Unmarshal filters
		if filtersJSON.Valid && filtersJSON.String != "" {
			if err := json.Unmarshal([]byte(filtersJSON.String), &output.Filters); err != nil {
				return nil, err
			}
		}

		outputs = append(outputs, &output)
	}

	return &outputs, rows.Err()
}

// ForEachOutput iterates over all outputs
func (store *SQLiteStore) ForEachOutput(ctx context.Context, cb func(*model.OutputDef) error) error {
	rows, err := store.db.QueryContext(ctx, `
		SELECT id, alias, name, description, condition, props, filters, enabled, nb_success, nb_error
		FROM outputs
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var output model.OutputDef
		var propsJSON, filtersJSON sql.NullString

		err := rows.Scan(
			&output.ID,
			&output.Alias,
			&output.Spec.Name,
			&output.Spec.Desc,
			&output.Condition,
			&propsJSON,
			&filtersJSON,
			&output.Enabled,
			&output.NbSuccess,
			&output.NbError,
		)
		if err != nil {
			return err
		}

		// Unmarshal props
		if propsJSON.Valid && propsJSON.String != "" {
			if err := json.Unmarshal([]byte(propsJSON.String), &output.Props); err != nil {
				return err
			}
		}

		// Unmarshal filters
		if filtersJSON.Valid && filtersJSON.String != "" {
			if err := json.Unmarshal([]byte(filtersJSON.String), &output.Filters); err != nil {
				return err
			}
		}

		if err := cb(&output); err != nil {
			return err
		}
	}

	return rows.Err()
}
