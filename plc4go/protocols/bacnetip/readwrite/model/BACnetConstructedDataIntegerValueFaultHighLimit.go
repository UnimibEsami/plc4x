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

// BACnetConstructedDataIntegerValueFaultHighLimit is the corresponding interface of BACnetConstructedDataIntegerValueFaultHighLimit
type BACnetConstructedDataIntegerValueFaultHighLimit interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultHighLimit returns FaultHighLimit (property field)
	GetFaultHighLimit() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
}

// BACnetConstructedDataIntegerValueFaultHighLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIntegerValueFaultHighLimit.
// This is useful for switch cases.
type BACnetConstructedDataIntegerValueFaultHighLimitExactly interface {
	BACnetConstructedDataIntegerValueFaultHighLimit
	isBACnetConstructedDataIntegerValueFaultHighLimit() bool
}

// _BACnetConstructedDataIntegerValueFaultHighLimit is the data-structure of this message
type _BACnetConstructedDataIntegerValueFaultHighLimit struct {
	*_BACnetConstructedData
	FaultHighLimit BACnetApplicationTagSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueFaultHighLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_INTEGER_VALUE
}

func (m *_BACnetConstructedDataIntegerValueFaultHighLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_HIGH_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegerValueFaultHighLimit) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIntegerValueFaultHighLimit) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueFaultHighLimit) GetFaultHighLimit() BACnetApplicationTagSignedInteger {
	return m.FaultHighLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueFaultHighLimit) GetActualValue() BACnetApplicationTagSignedInteger {
	return CastBACnetApplicationTagSignedInteger(m.GetFaultHighLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIntegerValueFaultHighLimit factory function for _BACnetConstructedDataIntegerValueFaultHighLimit
func NewBACnetConstructedDataIntegerValueFaultHighLimit(faultHighLimit BACnetApplicationTagSignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntegerValueFaultHighLimit {
	_result := &_BACnetConstructedDataIntegerValueFaultHighLimit{
		FaultHighLimit:         faultHighLimit,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegerValueFaultHighLimit(structType interface{}) BACnetConstructedDataIntegerValueFaultHighLimit {
	if casted, ok := structType.(BACnetConstructedDataIntegerValueFaultHighLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegerValueFaultHighLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegerValueFaultHighLimit) GetTypeName() string {
	return "BACnetConstructedDataIntegerValueFaultHighLimit"
}

func (m *_BACnetConstructedDataIntegerValueFaultHighLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataIntegerValueFaultHighLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (faultHighLimit)
	lengthInBits += m.FaultHighLimit.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntegerValueFaultHighLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIntegerValueFaultHighLimitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIntegerValueFaultHighLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegerValueFaultHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegerValueFaultHighLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (faultHighLimit)
	if pullErr := readBuffer.PullContext("faultHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for faultHighLimit")
	}
	_faultHighLimit, _faultHighLimitErr := BACnetApplicationTagParse(readBuffer)
	if _faultHighLimitErr != nil {
		return nil, errors.Wrap(_faultHighLimitErr, "Error parsing 'faultHighLimit' field")
	}
	faultHighLimit := _faultHighLimit.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("faultHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for faultHighLimit")
	}

	// Virtual field
	_actualValue := faultHighLimit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegerValueFaultHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegerValueFaultHighLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIntegerValueFaultHighLimit{
		FaultHighLimit: faultHighLimit,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIntegerValueFaultHighLimit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegerValueFaultHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegerValueFaultHighLimit")
		}

		// Simple Field (faultHighLimit)
		if pushErr := writeBuffer.PushContext("faultHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for faultHighLimit")
		}
		_faultHighLimitErr := writeBuffer.WriteSerializable(m.GetFaultHighLimit())
		if popErr := writeBuffer.PopContext("faultHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for faultHighLimit")
		}
		if _faultHighLimitErr != nil {
			return errors.Wrap(_faultHighLimitErr, "Error serializing 'faultHighLimit' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegerValueFaultHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegerValueFaultHighLimit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntegerValueFaultHighLimit) isBACnetConstructedDataIntegerValueFaultHighLimit() bool {
	return true
}

func (m *_BACnetConstructedDataIntegerValueFaultHighLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}