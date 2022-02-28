package gorm2

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

// Repository Gorm v2リポジトリ
type Repository struct {
	db     *gorm.DB
	logger *zap.Logger
}

// Config Gorm v2リポジトリ用Config
type Config struct {
	Hostname string
	Port     int
	Username string
	Password string
	Database string
}

// NewGorm2Repository 新しいGorm v2リポジトリを生成
func NewGorm2Repository(c *Config, logger *zap.Logger) (*Repository, error) {
	db, err := newDBConnection(c, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db :%w", err)
	}

	return &Repository{
		db:     db,
		logger: logger,
	}, nil
}

//  newDBConnection 新しいDBとのコネクションを作成
func newDBConnection(c *Config, logger *zap.Logger) (*gorm.DB, error) {
	// DSNの生成
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Username, c.Password, c.Hostname, c.Port, c.Database) + "?parseTime=true&loc=Local&charset=utf8mb4"

	// DBとのコネクションを作成
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: zapgorm2.New(logger)})
	if err != nil {
		return nil, fmt.Errorf("failed to connect DB : %w", err)
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")

	logger.Info("Successfully connected to DB")

	return db, nil
}

// GetSQLDb sqlパッケージのDBインスタンスを取得
func GetSQLDb(repo *Repository) (*sql.DB, error) {
	db, err := repo.db.DB()
	if err != nil {
		repo.logger.Error("failed to get sql db", zap.Error(err))
		return nil, fmt.Errorf("failed to get sql db : %w", err)
	}

	return db, nil
}
