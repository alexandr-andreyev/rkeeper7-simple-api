package app

import (
	"fmt"
	"rkeeper7-simpleapi-service/internal/config"

	"golang.org/x/sys/windows/svc/debug"
)

// if setup returns an error, the service doesn't start
func setup(wl debug.Log, svcName, sha1ver string) (config.Server, error) {
	var s config.Server

	// did we get a full SHA1?
	if len(sha1ver) == 40 {
		sha1ver = sha1ver[0:7]
	}

	if sha1ver == "" {
		sha1ver = "dev"
	}

	s.Winlog = wl

	// Note: any logging here goes to Windows App Log
	// I suggest you setup local logging
	s.Winlog.Info(1, fmt.Sprintf("%s: setup (%s)", svcName, sha1ver))

	// read configuration
	cfg, err := config.New()
	if err != nil {
		s.Winlog.Error(1, err.Error())
	}
	s.Config = cfg
	// configure more logging

	return s, nil
}
