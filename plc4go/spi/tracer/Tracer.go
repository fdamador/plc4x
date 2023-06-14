/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package tracer

import (
	"fmt"
	"sync"
	"time"

	"github.com/apache/plc4x/plc4go/spi/options"
	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/rs/zerolog"
)

type TraceEntry struct {
	Timestamp     time.Time
	ConnectionId  string
	TransactionId string
	Operation     string
	Message       string
}

type Provider interface {
	EnableTracer()
	GetTracer() *Tracer
}

type Tracer interface {
	GetConnectionId() string
	SetConnectionId(connectionId string)
	ResetTraces()
	GetTraces() []TraceEntry
	AddTrace(operation string, message string)
	AddTransactionalStartTrace(operation string, message string) string
	AddTransactionalTrace(transactionId string, operation string, message string)
	FilterTraces(traces []TraceEntry, connectionIdFilter string, transactionIdFilter string, operationFilter string, messageFilter string) []TraceEntry
}

func NewTracer(connectionId string, _options ...options.WithOption) Tracer {
	return &tracer{
		connectionId: connectionId,
		traceEntries: []TraceEntry{},
		log:          options.ExtractCustomLogger(_options...),
	}
}

type tracer struct {
	connectionId string
	traceEntries []TraceEntry
	m            sync.Mutex

	log zerolog.Logger
}

func (t *tracer) GetConnectionId() string {
	return t.connectionId
}

func (t *tracer) SetConnectionId(connectionId string) {
	t.connectionId = connectionId
}

func (t *tracer) ResetTraces() {
	t.m.Lock()
	t.traceEntries = []TraceEntry{}
	t.m.Unlock()
}

func (t *tracer) GetTraces() []TraceEntry {
	t.m.Lock()
	entries := t.traceEntries
	t.m.Unlock()
	return entries
}

func (t *tracer) AddTrace(operation string, message string) {
	t.m.Lock()
	t.traceEntries = append(t.traceEntries, TraceEntry{
		Timestamp:     time.Now(),
		ConnectionId:  t.connectionId,
		TransactionId: "",
		Operation:     operation,
		Message:       message,
	})
	t.m.Unlock()
}

func (t *tracer) AddTransactionalStartTrace(operation string, message string) string {
	t.m.Lock()
	transactionId := utils.GenerateId(t.log, 4)
	t.traceEntries = append(t.traceEntries, TraceEntry{
		Timestamp:     time.Now(),
		ConnectionId:  t.connectionId,
		TransactionId: transactionId,
		Operation:     operation,
		Message:       message,
	})
	t.m.Unlock()
	return transactionId
}

func (t *tracer) AddTransactionalTrace(transactionId string, operation string, message string) {
	t.m.Lock()
	t.traceEntries = append(t.traceEntries, TraceEntry{
		Timestamp:     time.Now(),
		ConnectionId:  t.connectionId,
		TransactionId: transactionId,
		Operation:     operation,
		Message:       message,
	})
	t.m.Unlock()
}

func (t *tracer) FilterTraces(traces []TraceEntry, connectionIdFilter string, transactionIdFilter string, operationFilter string, messageFilter string) []TraceEntry {
	var result []TraceEntry
traceFiltering:
	for _, trace := range traces {
		if connectionIdFilter != "" && trace.ConnectionId != connectionIdFilter {
			continue traceFiltering
		}
		if transactionIdFilter != "" && trace.TransactionId != transactionIdFilter {
			continue traceFiltering
		}
		if operationFilter != "" && trace.Operation != operationFilter {
			continue traceFiltering
		}
		if messageFilter != "" && trace.Message != messageFilter {
			continue traceFiltering
		}
		result = append(result, trace)
	}
	return result
}

func (t *tracer) String() string {
	return fmt.Sprintf("Tracer for %s", t.connectionId)
}
