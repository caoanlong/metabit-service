package config

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	G_DB     *gorm.DB
	G_CONFIG Conf
	G_LOG    *zap.Logger
)
