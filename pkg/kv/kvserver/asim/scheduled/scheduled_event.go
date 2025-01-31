// Copyright 2023 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package scheduled

import (
	"time"

	"github.com/cockroachdb/cockroach/pkg/kv/kvserver/asim/event"
)

type ScheduledEventList []ScheduledEvent

// Len implements sort.Interface.
func (sel ScheduledEventList) Len() int { return len(sel) }

// Less implements sort.Interface.
func (sel ScheduledEventList) Less(i, j int) bool {
	if sel[i].At == sel[j].At {
		return i < j
	}
	return sel[i].At.Before(sel[j].At)
}

// Swap implements sort.Interface.
func (sel ScheduledEventList) Swap(i, j int) {
	sel[i], sel[j] = sel[j], sel[i]
}

// ScheduledEvent contains the target event to be executed at the specified At
// time.
type ScheduledEvent struct {
	At          time.Time
	TargetEvent event.Event
}

// IsMutationEvent returns whether the scheduled event is a mutation event or an
// assertion event.
func (s ScheduledEvent) IsMutationEvent() bool {
	return s.TargetEvent.Func().GetType() == event.MutationType
}

// ScheduledMutationWithAssertionEvent contains the MutationWithAssertionEvent
// event to be executed at the specified At time.
type ScheduledMutationWithAssertionEvent struct {
	At                         time.Time
	MutationWithAssertionEvent event.MutationWithAssertionEvent
}
