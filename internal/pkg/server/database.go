package server

import (
	"errors"
	"fmt"
	"syscall"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	defaultMaxOpenConnections    = 5
	defaultMaxIdelConnections    = 3
	defaultMaxConnectionLifeTime = time.Duration(1 * time.Hour)
)

func NewDataBase(env Env) (*gorm.DB, error) {
	db := (*gorm.DB)(nil)
	err := (error)(nil)
	for attempts := 0; attempts < 3; attempts++ {
		db, err = gorm.Open(mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				env.DBUser, env.DBPassword, env.DBHost, env.DBPort, env.DBName)), &gorm.Config{})
		if errors.Is(err, syscall.ECONNRESET) {
			time.Sleep(time.Second)
			continue
		}
		break
	}

	if err != nil {
		return nil, fmt.Errorf("with database: %w", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("with database: %w", err)
	}

	sqlDb.SetMaxIdleConns(defaultMaxIdelConnections)
	sqlDb.SetMaxOpenConns(defaultMaxOpenConnections)
	sqlDb.SetConnMaxLifetime(defaultMaxConnectionLifeTime)

	return db, nil
}
