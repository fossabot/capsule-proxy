// Copyright 2022 Clastix Labs
// SPDX-License-Identifier: Apache-2.0

package watchdog

import "context"

type watchItem struct {
	triggers []chan event.GenericEvent
	cancelFn context.CancelFunc
}

type watchMap map[string]watchItem

type Manager struct {
	client   client.Client
	watchMap watchMap
	// managerErrChan is the channel that is going to be used
	// when the watchdog manager cannot start due to any kind of problem.
	managerErrChan chan event.GenericEvent

	MigrateCABundle         []byte
	MigrateServiceName      string
	MigrateServiceNamespace string
	AdminClient             client.Client
}
