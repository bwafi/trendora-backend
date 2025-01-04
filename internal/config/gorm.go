package config

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	name := viper.GetString("database.name")
	host := viper.GetString("database.host")
	port := viper.GetInt32("database.port")
	ssl := viper.GetString("database.ssl")
	idleConnection := viper.GetInt("database.pool.idle")
	maxConnection := viper.GetInt("database.pool.max")
	lifetimeConnection := viper.GetInt("database.pool.lifetime")

	// logger := logger.New(, )

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Jakarta", host, username, password, name, port, ssl)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed error connect database: %v \n", err)
	}
	connection, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to connect database: %v \n", err)
	}

	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxConnection)
	connection.SetConnMaxLifetime(time.Second * time.Duration(lifetimeConnection))

	return db
}

type LogrusWriter struct {
	*logger.Config
	Logger *logrus.Logger
}

func (l *LogrusWriter) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *LogrusWriter) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.Logger.WithContext(ctx).Infof(msg, data...)
	}
}

func (l *LogrusWriter) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.Logger.WithContext(ctx).Warnf(msg, data...)
	}
}

func (l *LogrusWriter) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.Logger.WithContext(ctx).Errorf(msg, data...)
	}
}

func (l *LogrusWriter) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)

	sql, rows := fc()
	entry := l.Logger.WithFields(logrus.Fields{
		"elapsed":      elapsed,
		"rowsAffected": rows,
	})

	if err != nil {
		entry.WithError(err).Error("SQL Execution Failed")
	} else {
		entry.Info("SQL Executed")
	}
	entry.WithField("sql", sql).Debug("SQL Trace")
}
