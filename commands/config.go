package commands

import "code.cloudfoundry.org/cli/utils/config"

//go:generate counterfeiter . Config

// Config a way of getting basic CF configuration
type Config interface {
	APIVersion() string
	BinaryName() string
	ColorEnabled() config.ColorSetting
	CurrentUser() (config.User, error)
	Locale() string
	Plugins() map[string]config.Plugin
	SetTargetInformation(api string, apiVersion string, auth string, loggregator string, doppler string, uaa string, skipSSLValidation bool)
	Target() string
	TargetedOrganization() config.Organization
	TargetedSpace() config.Space
	SkipSSLValidation() bool
}
