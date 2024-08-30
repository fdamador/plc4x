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

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// TransportSize is an enum
type TransportSize uint8

type ITransportSize interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Supported_S7_300() bool
	Supported_LOGO() bool
	Code() uint8
	SizeInBytes() uint8
	Supported_S7_400() bool
	Supported_S7_1200() bool
	ShortName() uint8
	Supported_S7_1500() bool
	DataTransportSize() DataTransportSize
	DataProtocolId() string
	BaseType() TransportSize
}

const (
	TransportSize_BOOL           TransportSize = 0x01
	TransportSize_BYTE           TransportSize = 0x02
	TransportSize_WORD           TransportSize = 0x03
	TransportSize_DWORD          TransportSize = 0x04
	TransportSize_LWORD          TransportSize = 0x05
	TransportSize_INT            TransportSize = 0x06
	TransportSize_UINT           TransportSize = 0x07
	TransportSize_SINT           TransportSize = 0x08
	TransportSize_USINT          TransportSize = 0x09
	TransportSize_DINT           TransportSize = 0x0A
	TransportSize_UDINT          TransportSize = 0x0B
	TransportSize_LINT           TransportSize = 0x0C
	TransportSize_ULINT          TransportSize = 0x0D
	TransportSize_REAL           TransportSize = 0x0E
	TransportSize_LREAL          TransportSize = 0x0F
	TransportSize_CHAR           TransportSize = 0x10
	TransportSize_WCHAR          TransportSize = 0x11
	TransportSize_STRING         TransportSize = 0x12
	TransportSize_WSTRING        TransportSize = 0x13
	TransportSize_S5TIME         TransportSize = 0x14
	TransportSize_TIME           TransportSize = 0x15
	TransportSize_LTIME          TransportSize = 0x16
	TransportSize_DATE           TransportSize = 0x17
	TransportSize_TIME_OF_DAY    TransportSize = 0x18
	TransportSize_TOD            TransportSize = 0x19
	TransportSize_LTIME_OF_DAY   TransportSize = 0x1A
	TransportSize_LTOD           TransportSize = 0x1B
	TransportSize_DATE_AND_TIME  TransportSize = 0x1C
	TransportSize_DT             TransportSize = 0x1D
	TransportSize_DATE_AND_LTIME TransportSize = 0x1E
	TransportSize_LDT            TransportSize = 0x1F
	TransportSize_DTL            TransportSize = 0x21
)

var TransportSizeValues []TransportSize

func init() {
	_ = errors.New
	TransportSizeValues = []TransportSize{
		TransportSize_BOOL,
		TransportSize_BYTE,
		TransportSize_WORD,
		TransportSize_DWORD,
		TransportSize_LWORD,
		TransportSize_INT,
		TransportSize_UINT,
		TransportSize_SINT,
		TransportSize_USINT,
		TransportSize_DINT,
		TransportSize_UDINT,
		TransportSize_LINT,
		TransportSize_ULINT,
		TransportSize_REAL,
		TransportSize_LREAL,
		TransportSize_CHAR,
		TransportSize_WCHAR,
		TransportSize_STRING,
		TransportSize_WSTRING,
		TransportSize_S5TIME,
		TransportSize_TIME,
		TransportSize_LTIME,
		TransportSize_DATE,
		TransportSize_TIME_OF_DAY,
		TransportSize_TOD,
		TransportSize_LTIME_OF_DAY,
		TransportSize_LTOD,
		TransportSize_DATE_AND_TIME,
		TransportSize_DT,
		TransportSize_DATE_AND_LTIME,
		TransportSize_LDT,
		TransportSize_DTL,
	}
}

func (e TransportSize) Supported_S7_300() bool {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return true
		}
	case 0x02:
		{ /* '0x02' */
			return true
		}
	case 0x03:
		{ /* '0x03' */
			return true
		}
	case 0x04:
		{ /* '0x04' */
			return true
		}
	case 0x05:
		{ /* '0x05' */
			return false
		}
	case 0x06:
		{ /* '0x06' */
			return true
		}
	case 0x07:
		{ /* '0x07' */
			return false
		}
	case 0x08:
		{ /* '0x08' */
			return false
		}
	case 0x09:
		{ /* '0x09' */
			return false
		}
	case 0x0A:
		{ /* '0x0A' */
			return true
		}
	case 0x0B:
		{ /* '0x0B' */
			return false
		}
	case 0x0C:
		{ /* '0x0C' */
			return false
		}
	case 0x0D:
		{ /* '0x0D' */
			return false
		}
	case 0x0E:
		{ /* '0x0E' */
			return true
		}
	case 0x0F:
		{ /* '0x0F' */
			return false
		}
	case 0x10:
		{ /* '0x10' */
			return true
		}
	case 0x11:
		{ /* '0x11' */
			return false
		}
	case 0x12:
		{ /* '0x12' */
			return true
		}
	case 0x13:
		{ /* '0x13' */
			return false
		}
	case 0x14:
		{ /* '0x14' */
			return true
		}
	case 0x15:
		{ /* '0x15' */
			return true
		}
	case 0x16:
		{ /* '0x16' */
			return false
		}
	case 0x17:
		{ /* '0x17' */
			return true
		}
	case 0x18:
		{ /* '0x18' */
			return true
		}
	case 0x19:
		{ /* '0x19' */
			return true
		}
	case 0x1A:
		{ /* '0x1A' */
			return false
		}
	case 0x1B:
		{ /* '0x1B' */
			return false
		}
	case 0x1C:
		{ /* '0x1C' */
			return true
		}
	case 0x1D:
		{ /* '0x1D' */
			return true
		}
	case 0x1E:
		{ /* '0x1E' */
			return false
		}
	case 0x1F:
		{ /* '0x1F' */
			return false
		}
	case 0x21:
		{ /* '0x21' */
			return false
		}
	default:
		{
			return false
		}
	}
}

func TransportSizeFirstEnumForFieldSupported_S7_300(value bool) (TransportSize, bool) {
	for _, sizeValue := range TransportSizeValues {
		if sizeValue.Supported_S7_300() == value {
			return sizeValue, true
		}
	}
	return 0, false
}

func (e TransportSize) Supported_LOGO() bool {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return true
		}
	case 0x02:
		{ /* '0x02' */
			return true
		}
	case 0x03:
		{ /* '0x03' */
			return true
		}
	case 0x04:
		{ /* '0x04' */
			return true
		}
	case 0x05:
		{ /* '0x05' */
			return false
		}
	case 0x06:
		{ /* '0x06' */
			return true
		}
	case 0x07:
		{ /* '0x07' */
			return true
		}
	case 0x08:
		{ /* '0x08' */
			return true
		}
	case 0x09:
		{ /* '0x09' */
			return true
		}
	case 0x0A:
		{ /* '0x0A' */
			return true
		}
	case 0x0B:
		{ /* '0x0B' */
			return true
		}
	case 0x0C:
		{ /* '0x0C' */
			return false
		}
	case 0x0D:
		{ /* '0x0D' */
			return false
		}
	case 0x0E:
		{ /* '0x0E' */
			return true
		}
	case 0x0F:
		{ /* '0x0F' */
			return false
		}
	case 0x10:
		{ /* '0x10' */
			return true
		}
	case 0x11:
		{ /* '0x11' */
			return true
		}
	case 0x12:
		{ /* '0x12' */
			return true
		}
	case 0x13:
		{ /* '0x13' */
			return true
		}
	case 0x14:
		{ /* '0x14' */
			return false
		}
	case 0x15:
		{ /* '0x15' */
			return true
		}
	case 0x16:
		{ /* '0x16' */
			return false
		}
	case 0x17:
		{ /* '0x17' */
			return true
		}
	case 0x18:
		{ /* '0x18' */
			return true
		}
	case 0x19:
		{ /* '0x19' */
			return true
		}
	case 0x1A:
		{ /* '0x1A' */
			return true
		}
	case 0x1B:
		{ /* '0x1B' */
			return true
		}
	case 0x1C:
		{ /* '0x1C' */
			return false
		}
	case 0x1D:
		{ /* '0x1D' */
			return false
		}
	case 0x1E:
		{ /* '0x1E' */
			return false
		}
	case 0x1F:
		{ /* '0x1F' */
			return false
		}
	case 0x21:
		{ /* '0x21' */
			return false
		}
	default:
		{
			return false
		}
	}
}

func TransportSizeFirstEnumForFieldSupported_LOGO(value bool) (TransportSize, bool) {
	for _, sizeValue := range TransportSizeValues {
		if sizeValue.Supported_LOGO() == value {
			return sizeValue, true
		}
	}
	return 0, false
}

func (e TransportSize) Code() uint8 {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return 0x01
		}
	case 0x02:
		{ /* '0x02' */
			return 0x02
		}
	case 0x03:
		{ /* '0x03' */
			return 0x04
		}
	case 0x04:
		{ /* '0x04' */
			return 0x06
		}
	case 0x05:
		{ /* '0x05' */
			return 0x00
		}
	case 0x06:
		{ /* '0x06' */
			return 0x05
		}
	case 0x07:
		{ /* '0x07' */
			return 0x05
		}
	case 0x08:
		{ /* '0x08' */
			return 0x02
		}
	case 0x09:
		{ /* '0x09' */
			return 0x02
		}
	case 0x0A:
		{ /* '0x0A' */
			return 0x07
		}
	case 0x0B:
		{ /* '0x0B' */
			return 0x07
		}
	case 0x0C:
		{ /* '0x0C' */
			return 0x00
		}
	case 0x0D:
		{ /* '0x0D' */
			return 0x00
		}
	case 0x0E:
		{ /* '0x0E' */
			return 0x08
		}
	case 0x0F:
		{ /* '0x0F' */
			return 0x00
		}
	case 0x10:
		{ /* '0x10' */
			return 0x03
		}
	case 0x11:
		{ /* '0x11' */
			return 0x13
		}
	case 0x12:
		{ /* '0x12' */
			return 0x03
		}
	case 0x13:
		{ /* '0x13' */
			return 0x00
		}
	case 0x14:
		{ /* '0x14' */
			return 0x00
		}
	case 0x15:
		{ /* '0x15' */
			return 0x00
		}
	case 0x16:
		{ /* '0x16' */
			return 0x00
		}
	case 0x17:
		{ /* '0x17' */
			return 0x00
		}
	case 0x18:
		{ /* '0x18' */
			return 0x00
		}
	case 0x19:
		{ /* '0x19' */
			return 0x0A
		}
	case 0x1A:
		{ /* '0x1A' */
			return 0x00
		}
	case 0x1B:
		{ /* '0x1B' */
			return 0x00
		}
	case 0x1C:
		{ /* '0x1C' */
			return 0x00
		}
	case 0x1D:
		{ /* '0x1D' */
			return 0x00
		}
	case 0x1E:
		{ /* '0x1E' */
			return 0x00
		}
	case 0x1F:
		{ /* '0x1F' */
			return 0x00
		}
	case 0x21:
		{ /* '0x21' */
			return 0x00
		}
	default:
		{
			return 0
		}
	}
}

func TransportSizeFirstEnumForFieldCode(value uint8) (TransportSize, bool) {
	for _, sizeValue := range TransportSizeValues {
		if sizeValue.Code() == value {
			return sizeValue, true
		}
	}
	return 0, false
}

func (e TransportSize) SizeInBytes() uint8 {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return 1
		}
	case 0x02:
		{ /* '0x02' */
			return 1
		}
	case 0x03:
		{ /* '0x03' */
			return 2
		}
	case 0x04:
		{ /* '0x04' */
			return 4
		}
	case 0x05:
		{ /* '0x05' */
			return 8
		}
	case 0x06:
		{ /* '0x06' */
			return 2
		}
	case 0x07:
		{ /* '0x07' */
			return 2
		}
	case 0x08:
		{ /* '0x08' */
			return 1
		}
	case 0x09:
		{ /* '0x09' */
			return 1
		}
	case 0x0A:
		{ /* '0x0A' */
			return 4
		}
	case 0x0B:
		{ /* '0x0B' */
			return 4
		}
	case 0x0C:
		{ /* '0x0C' */
			return 8
		}
	case 0x0D:
		{ /* '0x0D' */
			return 8
		}
	case 0x0E:
		{ /* '0x0E' */
			return 4
		}
	case 0x0F:
		{ /* '0x0F' */
			return 8
		}
	case 0x10:
		{ /* '0x10' */
			return 1
		}
	case 0x11:
		{ /* '0x11' */
			return 2
		}
	case 0x12:
		{ /* '0x12' */
			return 1
		}
	case 0x13:
		{ /* '0x13' */
			return 2
		}
	case 0x14:
		{ /* '0x14' */
			return 2
		}
	case 0x15:
		{ /* '0x15' */
			return 4
		}
	case 0x16:
		{ /* '0x16' */
			return 8
		}
	case 0x17:
		{ /* '0x17' */
			return 2
		}
	case 0x18:
		{ /* '0x18' */
			return 4
		}
	case 0x19:
		{ /* '0x19' */
			return 4
		}
	case 0x1A:
		{ /* '0x1A' */
			return 8
		}
	case 0x1B:
		{ /* '0x1B' */
			return 8
		}
	case 0x1C:
		{ /* '0x1C' */
			return 8
		}
	case 0x1D:
		{ /* '0x1D' */
			return 8
		}
	case 0x1E:
		{ /* '0x1E' */
			return 8
		}
	case 0x1F:
		{ /* '0x1F' */
			return 8
		}
	case 0x21:
		{ /* '0x21' */
			return 12
		}
	default:
		{
			return 0
		}
	}
}

func TransportSizeFirstEnumForFieldSizeInBytes(value uint8) (TransportSize, bool) {
	for _, sizeValue := range TransportSizeValues {
		if sizeValue.SizeInBytes() == value {
			return sizeValue, true
		}
	}
	return 0, false
}

func (e TransportSize) Supported_S7_400() bool {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return true
		}
	case 0x02:
		{ /* '0x02' */
			return true
		}
	case 0x03:
		{ /* '0x03' */
			return true
		}
	case 0x04:
		{ /* '0x04' */
			return true
		}
	case 0x05:
		{ /* '0x05' */
			return false
		}
	case 0x06:
		{ /* '0x06' */
			return true
		}
	case 0x07:
		{ /* '0x07' */
			return false
		}
	case 0x08:
		{ /* '0x08' */
			return false
		}
	case 0x09:
		{ /* '0x09' */
			return false
		}
	case 0x0A:
		{ /* '0x0A' */
			return true
		}
	case 0x0B:
		{ /* '0x0B' */
			return false
		}
	case 0x0C:
		{ /* '0x0C' */
			return false
		}
	case 0x0D:
		{ /* '0x0D' */
			return false
		}
	case 0x0E:
		{ /* '0x0E' */
			return true
		}
	case 0x0F:
		{ /* '0x0F' */
			return false
		}
	case 0x10:
		{ /* '0x10' */
			return true
		}
	case 0x11:
		{ /* '0x11' */
			return false
		}
	case 0x12:
		{ /* '0x12' */
			return true
		}
	case 0x13:
		{ /* '0x13' */
			return false
		}
	case 0x14:
		{ /* '0x14' */
			return true
		}
	case 0x15:
		{ /* '0x15' */
			return true
		}
	case 0x16:
		{ /* '0x16' */
			return false
		}
	case 0x17:
		{ /* '0x17' */
			return true
		}
	case 0x18:
		{ /* '0x18' */
			return true
		}
	case 0x19:
		{ /* '0x19' */
			return true
		}
	case 0x1A:
		{ /* '0x1A' */
			return false
		}
	case 0x1B:
		{ /* '0x1B' */
			return false
		}
	case 0x1C:
		{ /* '0x1C' */
			return true
		}
	case 0x1D:
		{ /* '0x1D' */
			return true
		}
	case 0x1E:
		{ /* '0x1E' */
			return false
		}
	case 0x1F:
		{ /* '0x1F' */
			return false
		}
	case 0x21:
		{ /* '0x21' */
			return false
		}
	default:
		{
			return false
		}
	}
}

func TransportSizeFirstEnumForFieldSupported_S7_400(value bool) (TransportSize, bool) {
	for _, sizeValue := range TransportSizeValues {
		if sizeValue.Supported_S7_400() == value {
			return sizeValue, true
		}
	}
	return 0, false
}

func (e TransportSize) Supported_S7_1200() bool {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return true
		}
	case 0x02:
		{ /* '0x02' */
			return true
		}
	case 0x03:
		{ /* '0x03' */
			return true
		}
	case 0x04:
		{ /* '0x04' */
			return true
		}
	case 0x05:
		{ /* '0x05' */
			return false
		}
	case 0x06:
		{ /* '0x06' */
			return true
		}
	case 0x07:
		{ /* '0x07' */
			return true
		}
	case 0x08:
		{ /* '0x08' */
			return true
		}
	case 0x09:
		{ /* '0x09' */
			return true
		}
	case 0x0A:
		{ /* '0x0A' */
			return true
		}
	case 0x0B:
		{ /* '0x0B' */
			return true
		}
	case 0x0C:
		{ /* '0x0C' */
			return false
		}
	case 0x0D:
		{ /* '0x0D' */
			return false
		}
	case 0x0E:
		{ /* '0x0E' */
			return true
		}
	case 0x0F:
		{ /* '0x0F' */
			return true
		}
	case 0x10:
		{ /* '0x10' */
			return true
		}
	case 0x11:
		{ /* '0x11' */
			return true
		}
	case 0x12:
		{ /* '0x12' */
			return true
		}
	case 0x13:
		{ /* '0x13' */
			return true
		}
	case 0x14:
		{ /* '0x14' */
			return false
		}
	case 0x15:
		{ /* '0x15' */
			return true
		}
	case 0x16:
		{ /* '0x16' */
			return false
		}
	case 0x17:
		{ /* '0x17' */
			return true
		}
	case 0x18:
		{ /* '0x18' */
			return true
		}
	case 0x19:
		{ /* '0x19' */
			return true
		}
	case 0x1A:
		{ /* '0x1A' */
			return false
		}
	case 0x1B:
		{ /* '0x1B' */
			return false
		}
	case 0x1C:
		{ /* '0x1C' */
			return false
		}
	case 0x1D:
		{ /* '0x1D' */
			return false
		}
	case 0x1E:
		{ /* '0x1E' */
			return false
		}
	case 0x1F:
		{ /* '0x1F' */
			return false
		}
	case 0x21:
		{ /* '0x21' */
			return true
		}
	default:
		{
			return false
		}
	}
}

func TransportSizeFirstEnumForFieldSupported_S7_1200(value bool) (TransportSize, bool) {
	for _, sizeValue := range TransportSizeValues {
		if sizeValue.Supported_S7_1200() == value {
			return sizeValue, true
		}
	}
	return 0, false
}

func (e TransportSize) ShortName() uint8 {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return 'X'
		}
	case 0x02:
		{ /* '0x02' */
			return 'B'
		}
	case 0x03:
		{ /* '0x03' */
			return 'W'
		}
	case 0x04:
		{ /* '0x04' */
			return 'D'
		}
	case 0x05:
		{ /* '0x05' */
			return 'X'
		}
	case 0x06:
		{ /* '0x06' */
			return 'W'
		}
	case 0x07:
		{ /* '0x07' */
			return 'W'
		}
	case 0x08:
		{ /* '0x08' */
			return 'B'
		}
	case 0x09:
		{ /* '0x09' */
			return 'B'
		}
	case 0x0A:
		{ /* '0x0A' */
			return 'D'
		}
	case 0x0B:
		{ /* '0x0B' */
			return 'D'
		}
	case 0x0C:
		{ /* '0x0C' */
			return 'X'
		}
	case 0x0D:
		{ /* '0x0D' */
			return 'X'
		}
	case 0x0E:
		{ /* '0x0E' */
			return 'D'
		}
	case 0x0F:
		{ /* '0x0F' */
			return 'X'
		}
	case 0x10:
		{ /* '0x10' */
			return 'B'
		}
	case 0x11:
		{ /* '0x11' */
			return 'X'
		}
	case 0x12:
		{ /* '0x12' */
			return 'X'
		}
	case 0x13:
		{ /* '0x13' */
			return 'X'
		}
	case 0x14:
		{ /* '0x14' */
			return 'X'
		}
	case 0x15:
		{ /* '0x15' */
			return 'X'
		}
	case 0x16:
		{ /* '0x16' */
			return 'X'
		}
	case 0x17:
		{ /* '0x17' */
			return 'X'
		}
	case 0x18:
		{ /* '0x18' */
			return 'X'
		}
	case 0x19:
		{ /* '0x19' */
			return 'X'
		}
	case 0x1A:
		{ /* '0x1A' */
			return 'X'
		}
	case 0x1B:
		{ /* '0x1B' */
			return 'X'
		}
	case 0x1C:
		{ /* '0x1C' */
			return 'X'
		}
	case 0x1D:
		{ /* '0x1D' */
			return 'X'
		}
	case 0x1E:
		{ /* '0x1E' */
			return 'X'
		}
	case 0x1F:
		{ /* '0x1F' */
			return 'X'
		}
	case 0x21:
		{ /* '0x21' */
			return 'X'
		}
	default:
		{
			return 0
		}
	}
}

func TransportSizeFirstEnumForFieldShortName(value uint8) (TransportSize, bool) {
	for _, sizeValue := range TransportSizeValues {
		if sizeValue.ShortName() == value {
			return sizeValue, true
		}
	}
	return 0, false
}

func (e TransportSize) Supported_S7_1500() bool {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return true
		}
	case 0x02:
		{ /* '0x02' */
			return true
		}
	case 0x03:
		{ /* '0x03' */
			return true
		}
	case 0x04:
		{ /* '0x04' */
			return true
		}
	case 0x05:
		{ /* '0x05' */
			return true
		}
	case 0x06:
		{ /* '0x06' */
			return true
		}
	case 0x07:
		{ /* '0x07' */
			return true
		}
	case 0x08:
		{ /* '0x08' */
			return true
		}
	case 0x09:
		{ /* '0x09' */
			return true
		}
	case 0x0A:
		{ /* '0x0A' */
			return true
		}
	case 0x0B:
		{ /* '0x0B' */
			return true
		}
	case 0x0C:
		{ /* '0x0C' */
			return true
		}
	case 0x0D:
		{ /* '0x0D' */
			return true
		}
	case 0x0E:
		{ /* '0x0E' */
			return true
		}
	case 0x0F:
		{ /* '0x0F' */
			return true
		}
	case 0x10:
		{ /* '0x10' */
			return true
		}
	case 0x11:
		{ /* '0x11' */
			return true
		}
	case 0x12:
		{ /* '0x12' */
			return true
		}
	case 0x13:
		{ /* '0x13' */
			return true
		}
	case 0x14:
		{ /* '0x14' */
			return true
		}
	case 0x15:
		{ /* '0x15' */
			return true
		}
	case 0x16:
		{ /* '0x16' */
			return true
		}
	case 0x17:
		{ /* '0x17' */
			return true
		}
	case 0x18:
		{ /* '0x18' */
			return true
		}
	case 0x19:
		{ /* '0x19' */
			return true
		}
	case 0x1A:
		{ /* '0x1A' */
			return true
		}
	case 0x1B:
		{ /* '0x1B' */
			return true
		}
	case 0x1C:
		{ /* '0x1C' */
			return true
		}
	case 0x1D:
		{ /* '0x1D' */
			return true
		}
	case 0x1E:
		{ /* '0x1E' */
			return true
		}
	case 0x1F:
		{ /* '0x1F' */
			return true
		}
	case 0x21:
		{ /* '0x21' */
			return true
		}
	default:
		{
			return false
		}
	}
}

func TransportSizeFirstEnumForFieldSupported_S7_1500(value bool) (TransportSize, bool) {
	for _, sizeValue := range TransportSizeValues {
		if sizeValue.Supported_S7_1500() == value {
			return sizeValue, true
		}
	}
	return 0, false
}

func (e TransportSize) DataTransportSize() DataTransportSize {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return DataTransportSize_BIT
		}
	case 0x02:
		{ /* '0x02' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x03:
		{ /* '0x03' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x04:
		{ /* '0x04' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x05:
		{ /* '0x05' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x06:
		{ /* '0x06' */
			return DataTransportSize_INTEGER
		}
	case 0x07:
		{ /* '0x07' */
			return DataTransportSize_INTEGER
		}
	case 0x08:
		{ /* '0x08' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x09:
		{ /* '0x09' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x0A:
		{ /* '0x0A' */
			return DataTransportSize_INTEGER
		}
	case 0x0B:
		{ /* '0x0B' */
			return DataTransportSize_INTEGER
		}
	case 0x0C:
		{ /* '0x0C' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x0D:
		{ /* '0x0D' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x0E:
		{ /* '0x0E' */
			return DataTransportSize_REAL
		}
	case 0x0F:
		{ /* '0x0F' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x10:
		{ /* '0x10' */
			return DataTransportSize_OCTET_STRING
		}
	case 0x11:
		{ /* '0x11' */
			return DataTransportSize_OCTET_STRING
		}
	case 0x12:
		{ /* '0x12' */
			return DataTransportSize_OCTET_STRING
		}
	case 0x13:
		{ /* '0x13' */
			return DataTransportSize_OCTET_STRING
		}
	case 0x14:
		{ /* '0x14' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x15:
		{ /* '0x15' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x16:
		{ /* '0x16' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x17:
		{ /* '0x17' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x18:
		{ /* '0x18' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x19:
		{ /* '0x19' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x1A:
		{ /* '0x1A' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x1B:
		{ /* '0x1B' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x1C:
		{ /* '0x1C' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x1D:
		{ /* '0x1D' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x1E:
		{ /* '0x1E' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x1F:
		{ /* '0x1F' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	case 0x21:
		{ /* '0x21' */
			return DataTransportSize_BYTE_WORD_DWORD
		}
	default:
		{
			return 0
		}
	}
}

func TransportSizeFirstEnumForFieldDataTransportSize(value DataTransportSize) (TransportSize, bool) {
	for _, sizeValue := range TransportSizeValues {
		if sizeValue.DataTransportSize() == value {
			return sizeValue, true
		}
	}
	return 0, false
}

func (e TransportSize) DataProtocolId() string {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return "IEC61131_BOOL"
		}
	case 0x02:
		{ /* '0x02' */
			return "IEC61131_BYTE"
		}
	case 0x03:
		{ /* '0x03' */
			return "IEC61131_WORD"
		}
	case 0x04:
		{ /* '0x04' */
			return "IEC61131_DWORD"
		}
	case 0x05:
		{ /* '0x05' */
			return "IEC61131_LWORD"
		}
	case 0x06:
		{ /* '0x06' */
			return "IEC61131_INT"
		}
	case 0x07:
		{ /* '0x07' */
			return "IEC61131_UINT"
		}
	case 0x08:
		{ /* '0x08' */
			return "IEC61131_SINT"
		}
	case 0x09:
		{ /* '0x09' */
			return "IEC61131_USINT"
		}
	case 0x0A:
		{ /* '0x0A' */
			return "IEC61131_DINT"
		}
	case 0x0B:
		{ /* '0x0B' */
			return "IEC61131_UDINT"
		}
	case 0x0C:
		{ /* '0x0C' */
			return "IEC61131_LINT"
		}
	case 0x0D:
		{ /* '0x0D' */
			return "IEC61131_ULINT"
		}
	case 0x0E:
		{ /* '0x0E' */
			return "IEC61131_REAL"
		}
	case 0x0F:
		{ /* '0x0F' */
			return "IEC61131_LREAL"
		}
	case 0x10:
		{ /* '0x10' */
			return "IEC61131_CHAR"
		}
	case 0x11:
		{ /* '0x11' */
			return "IEC61131_WCHAR"
		}
	case 0x12:
		{ /* '0x12' */
			return "IEC61131_STRING"
		}
	case 0x13:
		{ /* '0x13' */
			return "IEC61131_WSTRING"
		}
	case 0x14:
		{ /* '0x14' */
			return "S7_S5TIME"
		}
	case 0x15:
		{ /* '0x15' */
			return "IEC61131_TIME"
		}
	case 0x16:
		{ /* '0x16' */
			return "IEC61131_LTIME"
		}
	case 0x17:
		{ /* '0x17' */
			return "IEC61131_DATE"
		}
	case 0x18:
		{ /* '0x18' */
			return "IEC61131_TIME_OF_DAY"
		}
	case 0x19:
		{ /* '0x19' */
			return "IEC61131_TIME_OF_DAY"
		}
	case 0x1A:
		{ /* '0x1A' */
			return "IEC61131_LTIME_OF_DAY"
		}
	case 0x1B:
		{ /* '0x1B' */
			return "IEC61131_LTIME_OF_DAY"
		}
	case 0x1C:
		{ /* '0x1C' */
			return "IEC61131_DATE_AND_TIME"
		}
	case 0x1D:
		{ /* '0x1D' */
			return "IEC61131_DATE_AND_TIME"
		}
	case 0x1E:
		{ /* '0x1E' */
			return "IEC61131_DATE_AND_LTIME"
		}
	case 0x1F:
		{ /* '0x1F' */
			return "IEC61131_DATE_AND_LTIME"
		}
	case 0x21:
		{ /* '0x21' */
			return "IEC61131_DTL"
		}
	default:
		{
			return ""
		}
	}
}

func TransportSizeFirstEnumForFieldDataProtocolId(value string) (TransportSize, bool) {
	for _, sizeValue := range TransportSizeValues {
		if sizeValue.DataProtocolId() == value {
			return sizeValue, true
		}
	}
	return 0, false
}

func (e TransportSize) BaseType() TransportSize {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return 0
		}
	case 0x02:
		{ /* '0x02' */
			return 0
		}
	case 0x03:
		{ /* '0x03' */
			return 0
		}
	case 0x04:
		{ /* '0x04' */
			return TransportSize_WORD
		}
	case 0x05:
		{ /* '0x05' */
			return 0
		}
	case 0x06:
		{ /* '0x06' */
			return 0
		}
	case 0x07:
		{ /* '0x07' */
			return TransportSize_INT
		}
	case 0x08:
		{ /* '0x08' */
			return TransportSize_INT
		}
	case 0x09:
		{ /* '0x09' */
			return TransportSize_INT
		}
	case 0x0A:
		{ /* '0x0A' */
			return TransportSize_INT
		}
	case 0x0B:
		{ /* '0x0B' */
			return TransportSize_INT
		}
	case 0x0C:
		{ /* '0x0C' */
			return TransportSize_INT
		}
	case 0x0D:
		{ /* '0x0D' */
			return TransportSize_INT
		}
	case 0x0E:
		{ /* '0x0E' */
			return 0
		}
	case 0x0F:
		{ /* '0x0F' */
			return TransportSize_REAL
		}
	case 0x10:
		{ /* '0x10' */
			return 0
		}
	case 0x11:
		{ /* '0x11' */
			return 0
		}
	case 0x12:
		{ /* '0x12' */
			return 0
		}
	case 0x13:
		{ /* '0x13' */
			return 0
		}
	case 0x14:
		{ /* '0x14' */
			return 0
		}
	case 0x15:
		{ /* '0x15' */
			return 0
		}
	case 0x16:
		{ /* '0x16' */
			return TransportSize_TIME
		}
	case 0x17:
		{ /* '0x17' */
			return 0
		}
	case 0x18:
		{ /* '0x18' */
			return 0
		}
	case 0x19:
		{ /* '0x19' */
			return 0
		}
	case 0x1A:
		{ /* '0x1A' */
			return 0
		}
	case 0x1B:
		{ /* '0x1B' */
			return 0
		}
	case 0x1C:
		{ /* '0x1C' */
			return 0
		}
	case 0x1D:
		{ /* '0x1D' */
			return 0
		}
	case 0x1E:
		{ /* '0x1E' */
			return 0
		}
	case 0x1F:
		{ /* '0x1F' */
			return 0
		}
	case 0x21:
		{ /* '0x21' */
			return 0
		}
	default:
		{
			return 0
		}
	}
}

func TransportSizeFirstEnumForFieldBaseType(value TransportSize) (TransportSize, bool) {
	for _, sizeValue := range TransportSizeValues {
		if sizeValue.BaseType() == value {
			return sizeValue, true
		}
	}
	return 0, false
}
func TransportSizeByValue(value uint8) (enum TransportSize, ok bool) {
	switch value {
	case 0x01:
		return TransportSize_BOOL, true
	case 0x02:
		return TransportSize_BYTE, true
	case 0x03:
		return TransportSize_WORD, true
	case 0x04:
		return TransportSize_DWORD, true
	case 0x05:
		return TransportSize_LWORD, true
	case 0x06:
		return TransportSize_INT, true
	case 0x07:
		return TransportSize_UINT, true
	case 0x08:
		return TransportSize_SINT, true
	case 0x09:
		return TransportSize_USINT, true
	case 0x0A:
		return TransportSize_DINT, true
	case 0x0B:
		return TransportSize_UDINT, true
	case 0x0C:
		return TransportSize_LINT, true
	case 0x0D:
		return TransportSize_ULINT, true
	case 0x0E:
		return TransportSize_REAL, true
	case 0x0F:
		return TransportSize_LREAL, true
	case 0x10:
		return TransportSize_CHAR, true
	case 0x11:
		return TransportSize_WCHAR, true
	case 0x12:
		return TransportSize_STRING, true
	case 0x13:
		return TransportSize_WSTRING, true
	case 0x14:
		return TransportSize_S5TIME, true
	case 0x15:
		return TransportSize_TIME, true
	case 0x16:
		return TransportSize_LTIME, true
	case 0x17:
		return TransportSize_DATE, true
	case 0x18:
		return TransportSize_TIME_OF_DAY, true
	case 0x19:
		return TransportSize_TOD, true
	case 0x1A:
		return TransportSize_LTIME_OF_DAY, true
	case 0x1B:
		return TransportSize_LTOD, true
	case 0x1C:
		return TransportSize_DATE_AND_TIME, true
	case 0x1D:
		return TransportSize_DT, true
	case 0x1E:
		return TransportSize_DATE_AND_LTIME, true
	case 0x1F:
		return TransportSize_LDT, true
	case 0x21:
		return TransportSize_DTL, true
	}
	return 0, false
}

func TransportSizeByName(value string) (enum TransportSize, ok bool) {
	switch value {
	case "BOOL":
		return TransportSize_BOOL, true
	case "BYTE":
		return TransportSize_BYTE, true
	case "WORD":
		return TransportSize_WORD, true
	case "DWORD":
		return TransportSize_DWORD, true
	case "LWORD":
		return TransportSize_LWORD, true
	case "INT":
		return TransportSize_INT, true
	case "UINT":
		return TransportSize_UINT, true
	case "SINT":
		return TransportSize_SINT, true
	case "USINT":
		return TransportSize_USINT, true
	case "DINT":
		return TransportSize_DINT, true
	case "UDINT":
		return TransportSize_UDINT, true
	case "LINT":
		return TransportSize_LINT, true
	case "ULINT":
		return TransportSize_ULINT, true
	case "REAL":
		return TransportSize_REAL, true
	case "LREAL":
		return TransportSize_LREAL, true
	case "CHAR":
		return TransportSize_CHAR, true
	case "WCHAR":
		return TransportSize_WCHAR, true
	case "STRING":
		return TransportSize_STRING, true
	case "WSTRING":
		return TransportSize_WSTRING, true
	case "S5TIME":
		return TransportSize_S5TIME, true
	case "TIME":
		return TransportSize_TIME, true
	case "LTIME":
		return TransportSize_LTIME, true
	case "DATE":
		return TransportSize_DATE, true
	case "TIME_OF_DAY":
		return TransportSize_TIME_OF_DAY, true
	case "TOD":
		return TransportSize_TOD, true
	case "LTIME_OF_DAY":
		return TransportSize_LTIME_OF_DAY, true
	case "LTOD":
		return TransportSize_LTOD, true
	case "DATE_AND_TIME":
		return TransportSize_DATE_AND_TIME, true
	case "DT":
		return TransportSize_DT, true
	case "DATE_AND_LTIME":
		return TransportSize_DATE_AND_LTIME, true
	case "LDT":
		return TransportSize_LDT, true
	case "DTL":
		return TransportSize_DTL, true
	}
	return 0, false
}

func TransportSizeKnows(value uint8) bool {
	for _, typeValue := range TransportSizeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastTransportSize(structType any) TransportSize {
	castFunc := func(typ any) TransportSize {
		if sTransportSize, ok := typ.(TransportSize); ok {
			return sTransportSize
		}
		return 0
	}
	return castFunc(structType)
}

func (m TransportSize) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m TransportSize) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TransportSizeParse(ctx context.Context, theBytes []byte) (TransportSize, error) {
	return TransportSizeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TransportSizeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TransportSize, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("TransportSize", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TransportSize")
	}
	if enum, ok := TransportSizeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for TransportSize")
		return TransportSize(val), nil
	} else {
		return enum, nil
	}
}

func (e TransportSize) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e TransportSize) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("TransportSize", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TransportSize) PLC4XEnumName() string {
	switch e {
	case TransportSize_BOOL:
		return "BOOL"
	case TransportSize_BYTE:
		return "BYTE"
	case TransportSize_WORD:
		return "WORD"
	case TransportSize_DWORD:
		return "DWORD"
	case TransportSize_LWORD:
		return "LWORD"
	case TransportSize_INT:
		return "INT"
	case TransportSize_UINT:
		return "UINT"
	case TransportSize_SINT:
		return "SINT"
	case TransportSize_USINT:
		return "USINT"
	case TransportSize_DINT:
		return "DINT"
	case TransportSize_UDINT:
		return "UDINT"
	case TransportSize_LINT:
		return "LINT"
	case TransportSize_ULINT:
		return "ULINT"
	case TransportSize_REAL:
		return "REAL"
	case TransportSize_LREAL:
		return "LREAL"
	case TransportSize_CHAR:
		return "CHAR"
	case TransportSize_WCHAR:
		return "WCHAR"
	case TransportSize_STRING:
		return "STRING"
	case TransportSize_WSTRING:
		return "WSTRING"
	case TransportSize_S5TIME:
		return "S5TIME"
	case TransportSize_TIME:
		return "TIME"
	case TransportSize_LTIME:
		return "LTIME"
	case TransportSize_DATE:
		return "DATE"
	case TransportSize_TIME_OF_DAY:
		return "TIME_OF_DAY"
	case TransportSize_TOD:
		return "TOD"
	case TransportSize_LTIME_OF_DAY:
		return "LTIME_OF_DAY"
	case TransportSize_LTOD:
		return "LTOD"
	case TransportSize_DATE_AND_TIME:
		return "DATE_AND_TIME"
	case TransportSize_DT:
		return "DT"
	case TransportSize_DATE_AND_LTIME:
		return "DATE_AND_LTIME"
	case TransportSize_LDT:
		return "LDT"
	case TransportSize_DTL:
		return "DTL"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e TransportSize) String() string {
	return e.PLC4XEnumName()
}
