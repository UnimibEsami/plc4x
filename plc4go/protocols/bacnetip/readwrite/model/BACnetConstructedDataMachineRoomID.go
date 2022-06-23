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

// BACnetConstructedDataMachineRoomID is the corresponding interface of BACnetConstructedDataMachineRoomID
type BACnetConstructedDataMachineRoomID interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMachineRoomId returns MachineRoomId (property field)
	GetMachineRoomId() BACnetApplicationTagObjectIdentifier
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagObjectIdentifier
}

// BACnetConstructedDataMachineRoomIDExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMachineRoomID.
// This is useful for switch cases.
type BACnetConstructedDataMachineRoomIDExactly interface {
	BACnetConstructedDataMachineRoomID
	isBACnetConstructedDataMachineRoomID() bool
}

// _BACnetConstructedDataMachineRoomID is the data-structure of this message
type _BACnetConstructedDataMachineRoomID struct {
	*_BACnetConstructedData
	MachineRoomId BACnetApplicationTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMachineRoomID) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMachineRoomID) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MACHINE_ROOM_ID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMachineRoomID) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMachineRoomID) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMachineRoomID) GetMachineRoomId() BACnetApplicationTagObjectIdentifier {
	return m.MachineRoomId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMachineRoomID) GetActualValue() BACnetApplicationTagObjectIdentifier {
	return CastBACnetApplicationTagObjectIdentifier(m.GetMachineRoomId())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMachineRoomID factory function for _BACnetConstructedDataMachineRoomID
func NewBACnetConstructedDataMachineRoomID(machineRoomId BACnetApplicationTagObjectIdentifier, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMachineRoomID {
	_result := &_BACnetConstructedDataMachineRoomID{
		MachineRoomId:          machineRoomId,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMachineRoomID(structType interface{}) BACnetConstructedDataMachineRoomID {
	if casted, ok := structType.(BACnetConstructedDataMachineRoomID); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMachineRoomID); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMachineRoomID) GetTypeName() string {
	return "BACnetConstructedDataMachineRoomID"
}

func (m *_BACnetConstructedDataMachineRoomID) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataMachineRoomID) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (machineRoomId)
	lengthInBits += m.MachineRoomId.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMachineRoomID) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMachineRoomIDParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMachineRoomID, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMachineRoomID"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMachineRoomID")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (machineRoomId)
	if pullErr := readBuffer.PullContext("machineRoomId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for machineRoomId")
	}
	_machineRoomId, _machineRoomIdErr := BACnetApplicationTagParse(readBuffer)
	if _machineRoomIdErr != nil {
		return nil, errors.Wrap(_machineRoomIdErr, "Error parsing 'machineRoomId' field")
	}
	machineRoomId := _machineRoomId.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("machineRoomId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for machineRoomId")
	}

	// Virtual field
	_actualValue := machineRoomId
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMachineRoomID"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMachineRoomID")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMachineRoomID{
		MachineRoomId: machineRoomId,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMachineRoomID) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMachineRoomID"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMachineRoomID")
		}

		// Simple Field (machineRoomId)
		if pushErr := writeBuffer.PushContext("machineRoomId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for machineRoomId")
		}
		_machineRoomIdErr := writeBuffer.WriteSerializable(m.GetMachineRoomId())
		if popErr := writeBuffer.PopContext("machineRoomId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for machineRoomId")
		}
		if _machineRoomIdErr != nil {
			return errors.Wrap(_machineRoomIdErr, "Error serializing 'machineRoomId' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMachineRoomID"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMachineRoomID")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMachineRoomID) isBACnetConstructedDataMachineRoomID() bool {
	return true
}

func (m *_BACnetConstructedDataMachineRoomID) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}