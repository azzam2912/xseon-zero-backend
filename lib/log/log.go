package log

import (
	"github.com/skortech/st-kit/logging"
	"github.com/skortech/st-kit/service"
)

func WriteLog(severity, code, message, userID string, reference map[string]interface{}) {
	switch severity {
	case "ERROR":
		logging.Error(
			code,
			message,
			logging.WithService(service.Application),
			logging.WithIdentity(userID),
			logging.WithReference(reference),
		)
	case "INFO":
		logging.Info(
			code,
			message,
			logging.WithService(service.Application),
			logging.WithIdentity(userID),
			logging.WithReference(reference),
		)

	}
}
