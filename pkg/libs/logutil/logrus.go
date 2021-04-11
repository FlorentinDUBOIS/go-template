package logutil

import (
	"github.com/FlorentinDUBOIS/go-template/pkg/libs/math/i"
	"github.com/sirupsen/logrus"
)

func GetLogrusLevel(level int) logrus.Level {
	return logrus.AllLevels[i.Min(len(logrus.AllLevels), level)]
}
