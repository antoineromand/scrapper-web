package repository

import (
	"database/sql"
	"fmt"
)
type BaseRepository struct {
	DB *sql.DB
}

func (r *BaseRepository) GetEntityByID(id int, tableName string, entity interface{}) error {
	
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", tableName)
	err := r.DB.QueryRow(query, id).Scan(entity)
	if (err != nil) {
		return err
	}
	return nil
}
