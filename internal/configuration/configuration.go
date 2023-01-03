package configuration

import "golang.org/x/exp/slog"

// Application is the wide application configuration.
type Application struct {
	// Application wide configuration
	// More fields can be added here
	Logger *slog.Logger
}
