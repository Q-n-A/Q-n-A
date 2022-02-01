package gorm2

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

// Gorm v2を使ったRepositoryインターフェースの実装
type Gorm2Repository struct {
	db     *gorm.DB
	logger *zap.Logger
}

// Gorm2Repository用Config
type Config struct {
	MariaDBHostname string
	MariaDBPort     int
	MariaDBUsername string
	MariaDBPassword string
	MariaDBDatabase string
}

// Gorm2Repositoryを生成
func NewGorm2Repository(c *Config, logger *zap.Logger) (*Gorm2Repository, error) {
	db, err := newDBConnection(c, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db :%w", err)
	}

	return &Gorm2Repository{
		db:     db,
		logger: logger,
	}, nil
}

// DBとのコネクションを作成
func newDBConnection(c *Config, logger *zap.Logger) (*gorm.DB, error) {
	// DSNの生成
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.MariaDBUsername, c.MariaDBPassword, c.MariaDBHostname, c.MariaDBPort, c.MariaDBDatabase) + "?parseTime=true&loc=Local&charset=utf8mb4"

	// DBとのコネクションを作成
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: zapgorm2.New(logger)})
	if err != nil {
		return nil, fmt.Errorf("failed to connect DB : %w", err)
	}
	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")
	logger.Info("Successfully connected to DB")

	return db, nil
}

// sqlパッケージのDBインスタンスを取得
func GetSqlDB(repo *Gorm2Repository) (*sql.DB, error) {
	db, err := repo.db.DB()
	if err != nil {
		repo.logger.Error("failed to get sql db", zap.Error(err))
		return nil, fmt.Errorf("failed to get sql db : %w", err)
	}

	return db, nil
}
