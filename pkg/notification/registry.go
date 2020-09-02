package notification

import "github.com/oeoen/push-notifications/driver/config"

type Registry interface {
	NotificationManager() Manager
	Configuration() Configuration
	// Validator() *Validator
}

type Configuration interface {
	config.Provider
}