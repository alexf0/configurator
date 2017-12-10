package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"encoding/json"
	"github.com/kylelemons/go-gypsy/yaml"
)

type configRepository struct {
	*sql.DB
}

var insertStmt *sql.Stmt

func (r *configRepository) Store(c *Config) error {
	sqlStmt, err := prepareInsertStmt(r)

	if err != nil {
		return err
	}

	defer sqlStmt.Close()

	params, err := json.Marshal(c.Params)

	if err != nil {
		return err
	}

	_, err = sqlStmt.Exec(c.Type, c.Data, params)

	return err
}

func (r *configRepository) GetConfig(tp, data string) (*Config, error) {
	rows, err := r.Query(`SELECT cs.params FROM configs cs where cs.type=$1 and cs.data=$2 limit 1`, tp, data)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		s      string
		params map[string]interface{}
		found  bool
	)

	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			return nil, err
		}
		found = true
	}

	if !found  {
		return nil, ErrUnknown
	}

	err = json.Unmarshal([]byte(s), &params)

	return &Config{tp, data, params}, err
}

func NewConfigRepository(conf *yaml.File) (Repository, error) {
	connConf, err := postgresConf(conf)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", connConf)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}


	return &configRepository{db}, nil
}

func postgresConf(conf *yaml.File) (connConf string, err error) {
	keys := []string{"host", "port", "user", "password", "dbname", "sslmode"}

	for _, k := range keys {
		val, err := conf.Get("database." + k)
		if err != nil {
			return "", err
		}
		connConf += k + "=" + val + " "
	}

	return connConf, nil
}

func prepareInsertStmt(c *configRepository) (*sql.Stmt, error) {
	if insertStmt == nil {
		return c.Prepare(`
			INSERT INTO configs (type, data, params)
			VALUES ($1, $2, $3) ON CONFLICT (type, data) DO NOTHING`)
	}

	return insertStmt, nil
}
