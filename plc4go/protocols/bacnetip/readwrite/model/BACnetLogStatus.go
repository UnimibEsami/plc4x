/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetLogStatus is an enum
type BACnetLogStatus uint8

type IBACnetLogStatus interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetLogStatus_LOG_DISABLED    BACnetLogStatus = 0
	BACnetLogStatus_BUFFER_PURGED   BACnetLogStatus = 1
	BACnetLogStatus_LOG_INTERRUPTED BACnetLogStatus = 2
)

var BACnetLogStatusValues []BACnetLogStatus

func init() {
	_ = errors.New
	BACnetLogStatusValues = []BACnetLogStatus{
		BACnetLogStatus_LOG_DISABLED,
		BACnetLogStatus_BUFFER_PURGED,
		BACnetLogStatus_LOG_INTERRUPTED,
	}
}

func BACnetLogStatusByValue(value uint8) BACnetLogStatus {
	switch value {
	case 0:
		return BACnetLogStatus_LOG_DISABLED
	case 1:
		return BACnetLogStatus_BUFFER_PURGED
	case 2:
		return BACnetLogStatus_LOG_INTERRUPTED
	}
	return 0
}

func BACnetLogStatusByName(value string) BACnetLogStatus {
	switch value {
	case "LOG_DISABLED":
		return BACnetLogStatus_LOG_DISABLED
	case "BUFFER_PURGED":
		return BACnetLogStatus_BUFFER_PURGED
	case "LOG_INTERRUPTED":
		return BACnetLogStatus_LOG_INTERRUPTED
	}
	return 0
}

func BACnetLogStatusKnows(value uint8) bool {
	for _, typeValue := range BACnetLogStatusValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLogStatus(structType interface{}) BACnetLogStatus {
	castFunc := func(typ interface{}) BACnetLogStatus {
		if sBACnetLogStatus, ok := typ.(BACnetLogStatus); ok {
			return sBACnetLogStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLogStatus) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetLogStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogStatusParse(readBuffer utils.ReadBuffer) (BACnetLogStatus, error) {
	val, err := readBuffer.ReadUint8("BACnetLogStatus", 8)
	if err != nil {
		return 0, nil
	}
	return BACnetLogStatusByValue(val), nil
}

func (e BACnetLogStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetLogStatus", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e BACnetLogStatus) name() string {
	switch e {
	case BACnetLogStatus_LOG_DISABLED:
		return "LOG_DISABLED"
	case BACnetLogStatus_BUFFER_PURGED:
		return "BUFFER_PURGED"
	case BACnetLogStatus_LOG_INTERRUPTED:
		return "LOG_INTERRUPTED"
	}
	return ""
}

func (e BACnetLogStatus) String() string {
	return e.name()
}