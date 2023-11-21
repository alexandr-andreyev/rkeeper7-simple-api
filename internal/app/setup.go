package app

import (
	"rkeeper7-simpleapi-service/internal/config"

	"golang.org/x/sys/windows/svc/debug"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// if setup returns an error, the service doesn't start
func setup(wl debug.Log, svcName, sha1ver string) (config.Server, error) {
	var s config.Server

	// did we get a full SHA1?
	if len(sha1ver) == 40 {
		sha1ver = sha1ver[0:7]
	}

	if sha1ver == "" {
		sha1ver = envDev
	}

	// setup logger
	s.Logger = setupLogger(envLocal)

	// read configuration
	cfg, err := config.New()
	if err != nil {
		s.Logger.Error(err.Error())
		return s, err
	}
	s.Config = &cfg
	// configure more logging

	return s, nil
}
