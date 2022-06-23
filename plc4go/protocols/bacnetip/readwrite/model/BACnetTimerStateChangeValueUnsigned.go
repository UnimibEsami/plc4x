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

// BACnetTimerStateChangeValueUnsigned is the corresponding interface of BACnetTimerStateChangeValueUnsigned
type BACnetTimerStateChangeValueUnsigned interface {
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetApplicationTagUnsignedInteger
}

// BACnetTimerStateChangeValueUnsignedExactly can be used when we want exactly this type and not a type which fulfills BACnetTimerStateChangeValueUnsigned.
// This is useful for switch cases.
type BACnetTimerStateChangeValueUnsignedExactly interface {
	BACnetTimerStateChangeValueUnsigned
	isBACnetTimerStateChangeValueUnsigned() bool
}

// _BACnetTimerStateChangeValueUnsigned is the data-structure of this message
type _BACnetTimerStateChangeValueUnsigned struct {
	*_BACnetTimerStateChangeValue
	UnsignedValue BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueUnsigned) InitializeParent(parent BACnetTimerStateChangeValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimerStateChangeValueUnsigned) GetParent() BACnetTimerStateChangeValue {
	return m._BACnetTimerStateChangeValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueUnsigned) GetUnsignedValue() BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateChangeValueUnsigned factory function for _BACnetTimerStateChangeValueUnsigned
func NewBACnetTimerStateChangeValueUnsigned(unsignedValue BACnetApplicationTagUnsignedInteger, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueUnsigned {
	_result := &_BACnetTimerStateChangeValueUnsigned{
		UnsignedValue:                unsignedValue,
		_BACnetTimerStateChangeValue: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueUnsigned(structType interface{}) BACnetTimerStateChangeValueUnsigned {
	if casted, ok := structType.(BACnetTimerStateChangeValueUnsigned); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueUnsigned); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueUnsigned) GetTypeName() string {
	return "BACnetTimerStateChangeValueUnsigned"
}

func (m *_BACnetTimerStateChangeValueUnsigned) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetTimerStateChangeValueUnsigned) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueUnsigned) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimerStateChangeValueUnsignedParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueUnsigned, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueUnsigned"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueUnsigned")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unsignedValue)
	if pullErr := readBuffer.PullContext("unsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unsignedValue")
	}
	_unsignedValue, _unsignedValueErr := BACnetApplicationTagParse(readBuffer)
	if _unsignedValueErr != nil {
		return nil, errors.Wrap(_unsignedValueErr, "Error parsing 'unsignedValue' field")
	}
	unsignedValue := _unsignedValue.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("unsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unsignedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueUnsigned"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueUnsigned")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimerStateChangeValueUnsigned{
		UnsignedValue: unsignedValue,
		_BACnetTimerStateChangeValue: &_BACnetTimerStateChangeValue{
			ObjectTypeArgument: objectTypeArgument,
		},
	}
	_child._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimerStateChangeValueUnsigned) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueUnsigned"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueUnsigned")
		}

		// Simple Field (unsignedValue)
		if pushErr := writeBuffer.PushContext("unsignedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for unsignedValue")
		}
		_unsignedValueErr := writeBuffer.WriteSerializable(m.GetUnsignedValue())
		if popErr := writeBuffer.PopContext("unsignedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for unsignedValue")
		}
		if _unsignedValueErr != nil {
			return errors.Wrap(_unsignedValueErr, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueUnsigned"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueUnsigned")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueUnsigned) isBACnetTimerStateChangeValueUnsigned() bool {
	return true
}

func (m *_BACnetTimerStateChangeValueUnsigned) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}