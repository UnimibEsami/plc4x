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

// BACnetPropertyStateActionUnknown is the data-structure of this message
type BACnetPropertyStateActionUnknown struct {
	*BACnetPropertyStates
	UnknownValue *BACnetContextTagUnknown
}

// IBACnetPropertyStateActionUnknown is the corresponding interface of BACnetPropertyStateActionUnknown
type IBACnetPropertyStateActionUnknown interface {
	IBACnetPropertyStates
	// GetUnknownValue returns UnknownValue (property field)
	GetUnknownValue() *BACnetContextTagUnknown
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetPropertyStateActionUnknown) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStateActionUnknown) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStateActionUnknown) GetUnknownValue() *BACnetContextTagUnknown {
	return m.UnknownValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStateActionUnknown factory function for BACnetPropertyStateActionUnknown
func NewBACnetPropertyStateActionUnknown(unknownValue *BACnetContextTagUnknown, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStateActionUnknown {
	_result := &BACnetPropertyStateActionUnknown{
		UnknownValue:         unknownValue,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStateActionUnknown(structType interface{}) *BACnetPropertyStateActionUnknown {
	if casted, ok := structType.(BACnetPropertyStateActionUnknown); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStateActionUnknown); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStateActionUnknown(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStateActionUnknown(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStateActionUnknown) GetTypeName() string {
	return "BACnetPropertyStateActionUnknown"
}

func (m *BACnetPropertyStateActionUnknown) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStateActionUnknown) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (unknownValue)
	lengthInBits += m.UnknownValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStateActionUnknown) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStateActionUnknownParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStateActionUnknown, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStateActionUnknown"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unknownValue)
	if pullErr := readBuffer.PullContext("unknownValue"); pullErr != nil {
		return nil, pullErr
	}
	_unknownValue, _unknownValueErr := BACnetContextTagParse(readBuffer, uint8(peekedTagNumber), BACnetDataType(BACnetDataType_UNKNOWN))
	if _unknownValueErr != nil {
		return nil, errors.Wrap(_unknownValueErr, "Error parsing 'unknownValue' field")
	}
	unknownValue := CastBACnetContextTagUnknown(_unknownValue)
	if closeErr := readBuffer.CloseContext("unknownValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStateActionUnknown"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStateActionUnknown{
		UnknownValue:         CastBACnetContextTagUnknown(unknownValue),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStateActionUnknown) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStateActionUnknown"); pushErr != nil {
			return pushErr
		}

		// Simple Field (unknownValue)
		if pushErr := writeBuffer.PushContext("unknownValue"); pushErr != nil {
			return pushErr
		}
		_unknownValueErr := m.UnknownValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("unknownValue"); popErr != nil {
			return popErr
		}
		if _unknownValueErr != nil {
			return errors.Wrap(_unknownValueErr, "Error serializing 'unknownValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStateActionUnknown"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStateActionUnknown) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}