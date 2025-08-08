//go:build !windows
// +build !windows

/* SPDX-License-Identifier: MIT */
/*
 * Author: Jianhui Zhao <zhaojh329@gmail.com>
 */

package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func signalHandle() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGUSR1)

	for s := range c {
		switch s {
		case syscall.SIGUSR1:
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
			log.Debug().Msg("Debug mode enabled")
		}
	}
}
