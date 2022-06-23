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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetLandingDoorStatusLandingDoorsListEntry is the corresponding interface of BACnetLandingDoorStatusLandingDoorsListEntry
type BACnetLandingDoorStatusLandingDoorsListEntry interface {
	utils.LengthAware
	utils.Serializable
	// GetFloorNumber returns FloorNumber (property field)
	GetFloorNumber() BACnetContextTagUnsignedInteger
	// GetDoorStatus returns DoorStatus (property field)
	GetDoorStatus() BACnetDoorStatusTagged
}

// BACnetLandingDoorStatusLandingDoorsListEntryExactly can be used when we want exactly this type and not a type which fulfills BACnetLandingDoorStatusLandingDoorsListEntry.
// This is useful for switch cases.
type BACnetLandingDoorStatusLandingDoorsListEntryExactly interface {
	BACnetLandingDoorStatusLandingDoorsListEntry
	isBACnetLandingDoorStatusLandingDoorsListEntry() bool
}

// _BACnetLandingDoorStatusLandingDoorsListEntry is the data-structure of this message
type _BACnetLandingDoorStatusLandingDoorsListEntry struct {
	FloorNumber BACnetContextTagUnsignedInteger
	DoorStatus  BACnetDoorStatusTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLandingDoorStatusLandingDoorsListEntry) GetFloorNumber() BACnetContextTagUnsignedInteger {
	return m.FloorNumber
}

func (m *_BACnetLandingDoorStatusLandingDoorsListEntry) GetDoorStatus() BACnetDoorStatusTagged {
	return m.DoorStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLandingDoorStatusLandingDoorsListEntry factory function for _BACnetLandingDoorStatusLandingDoorsListEntry
func NewBACnetLandingDoorStatusLandingDoorsListEntry(floorNumber BACnetContextTagUnsignedInteger, doorStatus BACnetDoorStatusTagged) *_BACnetLandingDoorStatusLandingDoorsListEntry {
	return &_BACnetLandingDoorStatusLandingDoorsListEntry{FloorNumber: floorNumber, DoorStatus: doorStatus}
}

// Deprecated: use the interface for direct cast
func CastBACnetLandingDoorStatusLandingDoorsListEntry(structType interface{}) BACnetLandingDoorStatusLandingDoorsListEntry {
	if casted, ok := structType.(BACnetLandingDoorStatusLandingDoorsListEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLandingDoorStatusLandingDoorsListEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLandingDoorStatusLandingDoorsListEntry) GetTypeName() string {
	return "BACnetLandingDoorStatusLandingDoorsListEntry"
}

func (m *_BACnetLandingDoorStatusLandingDoorsListEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLandingDoorStatusLandingDoorsListEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (floorNumber)
	lengthInBits += m.FloorNumber.GetLengthInBits()

	// Simple field (doorStatus)
	lengthInBits += m.DoorStatus.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetLandingDoorStatusLandingDoorsListEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLandingDoorStatusLandingDoorsListEntryParse(readBuffer utils.ReadBuffer) (BACnetLandingDoorStatusLandingDoorsListEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLandingDoorStatusLandingDoorsListEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLandingDoorStatusLandingDoorsListEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (floorNumber)
	if pullErr := readBuffer.PullContext("floorNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for floorNumber")
	}
	_floorNumber, _floorNumberErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _floorNumberErr != nil {
		return nil, errors.Wrap(_floorNumberErr, "Error parsing 'floorNumber' field")
	}
	floorNumber := _floorNumber.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("floorNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for floorNumber")
	}

	// Simple Field (doorStatus)
	if pullErr := readBuffer.PullContext("doorStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doorStatus")
	}
	_doorStatus, _doorStatusErr := BACnetDoorStatusTaggedParse(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _doorStatusErr != nil {
		return nil, errors.Wrap(_doorStatusErr, "Error parsing 'doorStatus' field")
	}
	doorStatus := _doorStatus.(BACnetDoorStatusTagged)
	if closeErr := readBuffer.CloseContext("doorStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doorStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetLandingDoorStatusLandingDoorsListEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLandingDoorStatusLandingDoorsListEntry")
	}

	// Create the instance
	return NewBACnetLandingDoorStatusLandingDoorsListEntry(floorNumber, doorStatus), nil
}

func (m *_BACnetLandingDoorStatusLandingDoorsListEntry) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLandingDoorStatusLandingDoorsListEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLandingDoorStatusLandingDoorsListEntry")
	}

	// Simple Field (floorNumber)
	if pushErr := writeBuffer.PushContext("floorNumber"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for floorNumber")
	}
	_floorNumberErr := writeBuffer.WriteSerializable(m.GetFloorNumber())
	if popErr := writeBuffer.PopContext("floorNumber"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for floorNumber")
	}
	if _floorNumberErr != nil {
		return errors.Wrap(_floorNumberErr, "Error serializing 'floorNumber' field")
	}

	// Simple Field (doorStatus)
	if pushErr := writeBuffer.PushContext("doorStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for doorStatus")
	}
	_doorStatusErr := writeBuffer.WriteSerializable(m.GetDoorStatus())
	if popErr := writeBuffer.PopContext("doorStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for doorStatus")
	}
	if _doorStatusErr != nil {
		return errors.Wrap(_doorStatusErr, "Error serializing 'doorStatus' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLandingDoorStatusLandingDoorsListEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLandingDoorStatusLandingDoorsListEntry")
	}
	return nil
}

func (m *_BACnetLandingDoorStatusLandingDoorsListEntry) isBACnetLandingDoorStatusLandingDoorsListEntry() bool {
	return true
}

func (m *_BACnetLandingDoorStatusLandingDoorsListEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}