package repository

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

// Gormを使ったRepositoryインターフェースの実装
type GormRepository struct {
	db *gorm.DB
}

// Repository用Config
type Config struct {
	MariaDBHostname string
	MariaDBPort     int
	MariaDBUsername string
	MariaDBPassword string
	MariaDBDatabase string
}

// 新しいGormで実装したRepositoryを作成
func NewGormRepository(c *Config, logger *zap.Logger) (*GormRepository, error) {
	db, err := newDBConnection(c, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db :%w", err)
	}

	return &GormRepository{db: db}, nil
}

// DBとのコネクションを作成
func newDBConnection(c *Config, logger *zap.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.MariaDBUsername, c.MariaDBPassword, c.MariaDBHostname, c.MariaDBPort, c.MariaDBDatabase) + "?parseTime=true&loc=Local&charset=utf8mb4"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: zapgorm2.New(logger)})
	if err != nil {
		return nil, fmt.Errorf("failed to connect DB : %w", err)
	}
	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")

	return db, nil
}

// sqlパッケージのDBインスタンスを取得
func GetSqlDB(repo *GormRepository) (*sql.DB, error) {
	db, err := repo.db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql db : %w", err)
	}

	return db, nil
}
