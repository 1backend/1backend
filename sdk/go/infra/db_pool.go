/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package infra

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	defaultPostgresMaxOpenConns    = 16
	defaultPostgresMaxIdleConns    = 4
	defaultPostgresConnMaxLifetime = 30 * time.Minute
	defaultPostgresConnMaxIdleTime = 5 * time.Minute
	minPostgresDatastoreOpenConns  = 2
	postgresApplicationNameMaxLen  = 63
)

type DbPoolConfig struct {
	// MaxOpenConns bounds the total number of open connections per sql.DB pool.
	// Leave zero to use the SDK's PostgreSQL default.
	MaxOpenConns int

	// MaxIdleConns bounds idle connections retained per sql.DB pool.
	// Leave zero to use the SDK's PostgreSQL default.
	MaxIdleConns int

	// ConnMaxLifetime closes connections after this age.
	// Leave zero to use the SDK's PostgreSQL default.
	ConnMaxLifetime time.Duration

	// ConnMaxIdleTime closes connections after this idle duration.
	// Leave zero to use the SDK's PostgreSQL default.
	ConnMaxIdleTime time.Duration
}

func (cfg *DataStoreConfig) loadDbConnectionRuntimeOptionsFromEnv() error {
	if cfg.DbApplicationName == "" {
		cfg.DbApplicationName = os.Getenv("OB_DB_APPLICATION_NAME")
	}

	var err error
	if cfg.DbPool.MaxOpenConns == 0 {
		cfg.DbPool.MaxOpenConns, err = optionalIntEnv("OB_DB_MAX_OPEN_CONNS")
		if err != nil {
			return err
		}
	}
	if cfg.DbPool.MaxIdleConns == 0 {
		cfg.DbPool.MaxIdleConns, err = optionalIntEnv("OB_DB_MAX_IDLE_CONNS")
		if err != nil {
			return err
		}
	}
	if cfg.DbPool.ConnMaxLifetime == 0 {
		cfg.DbPool.ConnMaxLifetime, err = optionalDurationEnv("OB_DB_CONN_MAX_LIFETIME")
		if err != nil {
			return err
		}
	}
	if cfg.DbPool.ConnMaxIdleTime == 0 {
		cfg.DbPool.ConnMaxIdleTime, err = optionalDurationEnv("OB_DB_CONN_MAX_IDLE_TIME")
		if err != nil {
			return err
		}
	}

	if err := cfg.validateDbPoolConfig(); err != nil {
		return err
	}
	cfg.applyDbPoolDefaults()
	return nil
}

func optionalIntEnv(name string) (int, error) {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return 0, nil
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer: %w", name, err)
	}
	if parsed < 0 {
		return 0, fmt.Errorf("%s must be greater than or equal to 0", name)
	}
	return parsed, nil
}

func optionalDurationEnv(name string) (time.Duration, error) {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return 0, nil
	}
	parsed, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a Go duration: %w", name, err)
	}
	if parsed < 0 {
		return 0, fmt.Errorf("%s must be greater than or equal to 0", name)
	}
	return parsed, nil
}

func (cfg DataStoreConfig) validateDbPoolConfig() error {
	if cfg.DbPool.MaxOpenConns < 0 {
		return fmt.Errorf("DbPool.MaxOpenConns must be greater than or equal to 0")
	}
	if cfg.DbPool.MaxIdleConns < 0 {
		return fmt.Errorf("DbPool.MaxIdleConns must be greater than or equal to 0")
	}
	if cfg.DbPool.ConnMaxLifetime < 0 {
		return fmt.Errorf("DbPool.ConnMaxLifetime must be greater than or equal to 0")
	}
	if cfg.DbPool.ConnMaxIdleTime < 0 {
		return fmt.Errorf("DbPool.ConnMaxIdleTime must be greater than or equal to 0")
	}
	return nil
}

func (cfg DataStoreConfig) validateDatastoreDbPoolConfig() error {
	if !isPostgresDriver(cfg.Db) {
		return nil
	}
	if cfg.DbPool.MaxOpenConns > 0 && cfg.DbPool.MaxOpenConns < minPostgresDatastoreOpenConns {
		return fmt.Errorf(
			"postgres datastore MaxOpenConns must be at least %d because the advisory lock reserves one writer connection",
			minPostgresDatastoreOpenConns,
		)
	}
	return nil
}

func (cfg *DataStoreConfig) applyDbPoolDefaults() {
	if !isPostgresDriver(cfg.Db) {
		return
	}

	if cfg.DbPool.MaxOpenConns == 0 {
		cfg.DbPool.MaxOpenConns = defaultPostgresMaxOpenConns
	}
	if cfg.DbPool.MaxIdleConns == 0 {
		cfg.DbPool.MaxIdleConns = defaultPostgresMaxIdleConns
	}
	if cfg.DbPool.ConnMaxLifetime == 0 {
		cfg.DbPool.ConnMaxLifetime = defaultPostgresConnMaxLifetime
	}
	if cfg.DbPool.ConnMaxIdleTime == 0 {
		cfg.DbPool.ConnMaxIdleTime = defaultPostgresConnMaxIdleTime
	}
	if cfg.DbPool.MaxOpenConns > 0 && cfg.DbPool.MaxIdleConns > cfg.DbPool.MaxOpenConns {
		cfg.DbPool.MaxIdleConns = cfg.DbPool.MaxOpenConns
	}
}

func (cfg DataStoreConfig) openSQLDB(role string, connectionString string) (*sql.DB, error) {
	db, err := sql.Open(cfg.Db, cfg.connectionStringForRole(role, connectionString))
	if err != nil {
		return nil, err
	}
	cfg.configureSQLDB(db)
	return db, nil
}

func (cfg DataStoreConfig) configureSQLDB(db *sql.DB) {
	if cfg.DbPool.MaxOpenConns > 0 {
		db.SetMaxOpenConns(cfg.DbPool.MaxOpenConns)
	}
	if cfg.DbPool.MaxIdleConns > 0 {
		db.SetMaxIdleConns(cfg.DbPool.MaxIdleConns)
	}
	if cfg.DbPool.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(cfg.DbPool.ConnMaxLifetime)
	}
	if cfg.DbPool.ConnMaxIdleTime > 0 {
		db.SetConnMaxIdleTime(cfg.DbPool.ConnMaxIdleTime)
	}
}

func discardSQLConn(conn *sql.Conn) {
	_ = conn.Raw(func(driverConn any) error {
		return driver.ErrBadConn
	})
}

func (cfg DataStoreConfig) connectionStringForRole(role string, connectionString string) string {
	if !isPostgresDriver(cfg.Db) {
		return connectionString
	}

	applicationName := cfg.applicationNameForRole(role)
	if applicationName == "" {
		return connectionString
	}

	return connectionStringWithPostgresApplicationName(connectionString, applicationName)
}

func (cfg DataStoreConfig) applicationNameForRole(role string) string {
	name := strings.TrimSpace(cfg.DbApplicationName)
	if name == "" {
		name = defaultDbApplicationName()
	}
	name = sanitizePostgresApplicationName(name)

	role = sanitizePostgresApplicationName(role)
	if role == "" {
		return name
	}

	suffix := ":" + role
	if len(suffix) >= postgresApplicationNameMaxLen {
		return strings.TrimLeft(suffix[:postgresApplicationNameMaxLen], ":")
	}

	if len(name)+len(suffix) > postgresApplicationNameMaxLen {
		name = strings.TrimRight(name[:postgresApplicationNameMaxLen-len(suffix)], "-:")
	}
	if name == "" {
		return role
	}
	return name + suffix
}

func defaultDbApplicationName() string {
	serviceName := firstNonEmptyEnv("OTEL_SERVICE_NAME", "OB_SERVICE_NAME")
	nodeName := firstNonEmptyEnv("OB_NODE_ID", "HOSTNAME")
	if serviceName != "" && nodeName != "" && serviceName != nodeName {
		return serviceName + "/" + nodeName
	}
	if serviceName != "" {
		return serviceName
	}
	if nodeName != "" {
		return nodeName
	}
	if len(os.Args) > 0 {
		if exe := filepath.Base(os.Args[0]); exe != "" && exe != "." {
			return exe
		}
	}
	return "1backend"
}

func firstNonEmptyEnv(names ...string) string {
	for _, name := range names {
		value := strings.TrimSpace(os.Getenv(name))
		if value != "" {
			return value
		}
	}
	return ""
}

func sanitizePostgresApplicationName(name string) string {
	name = strings.TrimSpace(name)
	var b strings.Builder
	b.Grow(len(name))
	for _, r := range name {
		switch {
		case r >= 'a' && r <= 'z':
			b.WriteRune(r)
		case r >= 'A' && r <= 'Z':
			b.WriteRune(r)
		case r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == '.' || r == '_' || r == '-' || r == ':':
			b.WriteRune(r)
		default:
			b.WriteByte('-')
		}
		if b.Len() >= postgresApplicationNameMaxLen {
			break
		}
	}
	return strings.Trim(b.String(), "-")
}

func connectionStringWithPostgresApplicationName(connectionString string, applicationName string) string {
	parsed, err := url.Parse(connectionString)
	if err == nil && isPostgresScheme(parsed.Scheme) {
		query := parsed.Query()
		query.Set("application_name", applicationName)
		parsed.RawQuery = query.Encode()
		return parsed.String()
	}

	if strings.Contains(connectionString, "=") && !strings.Contains(connectionString, "://") {
		return upsertPostgresKeywordValue(connectionString, "application_name", applicationName)
	}

	return connectionString
}

func upsertPostgresKeywordValue(connectionString string, key string, value string) string {
	if !strings.Contains(connectionString, key+"=") {
		return strings.TrimSpace(connectionString) + " " + key + "=" + quotePostgresKeywordValue(value)
	}

	parts := strings.Fields(connectionString)
	for i, part := range parts {
		if strings.HasPrefix(part, key+"=") {
			parts[i] = key + "=" + quotePostgresKeywordValue(value)
			break
		}
	}
	return strings.Join(parts, " ")
}

func quotePostgresKeywordValue(value string) string {
	value = strings.ReplaceAll(value, `\`, `\\`)
	value = strings.ReplaceAll(value, `'`, `\'`)
	return "'" + value + "'"
}

func isPostgresDriver(driver string) bool {
	switch strings.ToLower(strings.TrimSpace(driver)) {
	case "postgres", "postgresql", "pgx":
		return true
	default:
		return false
	}
}

func isPostgresScheme(scheme string) bool {
	switch strings.ToLower(strings.TrimSpace(scheme)) {
	case "postgres", "postgresql":
		return true
	default:
		return false
	}
}
