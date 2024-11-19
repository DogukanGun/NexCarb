package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func CheckDatabaseSetup(conn *pgx.Conn) error {
	// Check and create `search_node_registry` table if not exists
	err := checkAndCreateTable(conn, "search_node_registery", `
		CREATE TABLE search_node_registery (
			wallet_address TEXT,
			country TEXT,
			smart_contract TEXT
		);`)
	if err != nil {
		return err
	}

	// Check and create `node_assignments` table if not exists
	err = checkAndCreateTable(conn, "node_assignments", `
		CREATE TABLE node_assignments (
			node_wallet_address TEXT,
			search_text TEXT,
			user_contract TEXT,
			is_active BIT
		);`)

	return err
}

func checkAndCreateTable(conn *pgx.Conn, tableName, createQuery string) error {
	checkTableQuery := `
		SELECT EXISTS (
			SELECT FROM pg_tables
			WHERE schemaname = 'public' AND tablename = $1
		);`
	var exists bool
	err := conn.QueryRow(context.Background(), checkTableQuery, tableName).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if table %s exists: %v", tableName, err)
	}

	if !exists {
		_, err = conn.Exec(context.Background(), createQuery)
		if err != nil {
			return fmt.Errorf("failed to create table %s: %v", tableName, err)
		}
	}
	return nil
}
