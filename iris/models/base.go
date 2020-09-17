package controllers

import (
	"strings"

	"github.com/sirupsen/logrus"
)

var OrmLogger *logrus.Logger

type Base struct {
}

func (m *Base) ErrReport(err error) {
	if !strings.HasSuffix(err.Error(), "record not found") {
		OrmLogger.Error(err.Error())
	}
}
