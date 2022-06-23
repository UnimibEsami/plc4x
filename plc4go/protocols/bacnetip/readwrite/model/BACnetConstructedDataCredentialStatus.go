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

// BACnetConstructedDataCredentialStatus is the corresponding interface of BACnetConstructedDataCredentialStatus
type BACnetConstructedDataCredentialStatus interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBinaryPv returns BinaryPv (property field)
	GetBinaryPv() BACnetBinaryPVTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetBinaryPVTagged
}

// BACnetConstructedDataCredentialStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCredentialStatus.
// This is useful for switch cases.
type BACnetConstructedDataCredentialStatusExactly interface {
	BACnetConstructedDataCredentialStatus
	isBACnetConstructedDataCredentialStatus() bool
}

// _BACnetConstructedDataCredentialStatus is the data-structure of this message
type _BACnetConstructedDataCredentialStatus struct {
	*_BACnetConstructedData
	BinaryPv BACnetBinaryPVTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCredentialStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCredentialStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CREDENTIAL_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCredentialStatus) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCredentialStatus) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCredentialStatus) GetBinaryPv() BACnetBinaryPVTagged {
	return m.BinaryPv
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCredentialStatus) GetActualValue() BACnetBinaryPVTagged {
	return CastBACnetBinaryPVTagged(m.GetBinaryPv())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCredentialStatus factory function for _BACnetConstructedDataCredentialStatus
func NewBACnetConstructedDataCredentialStatus(binaryPv BACnetBinaryPVTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCredentialStatus {
	_result := &_BACnetConstructedDataCredentialStatus{
		BinaryPv:               binaryPv,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCredentialStatus(structType interface{}) BACnetConstructedDataCredentialStatus {
	if casted, ok := structType.(BACnetConstructedDataCredentialStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCredentialStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCredentialStatus) GetTypeName() string {
	return "BACnetConstructedDataCredentialStatus"
}

func (m *_BACnetConstructedDataCredentialStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataCredentialStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (binaryPv)
	lengthInBits += m.BinaryPv.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCredentialStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCredentialStatusParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCredentialStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCredentialStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCredentialStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (binaryPv)
	if pullErr := readBuffer.PullContext("binaryPv"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for binaryPv")
	}
	_binaryPv, _binaryPvErr := BACnetBinaryPVTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _binaryPvErr != nil {
		return nil, errors.Wrap(_binaryPvErr, "Error parsing 'binaryPv' field")
	}
	binaryPv := _binaryPv.(BACnetBinaryPVTagged)
	if closeErr := readBuffer.CloseContext("binaryPv"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for binaryPv")
	}

	// Virtual field
	_actualValue := binaryPv
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCredentialStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCredentialStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCredentialStatus{
		BinaryPv: binaryPv,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCredentialStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCredentialStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCredentialStatus")
		}

		// Simple Field (binaryPv)
		if pushErr := writeBuffer.PushContext("binaryPv"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for binaryPv")
		}
		_binaryPvErr := writeBuffer.WriteSerializable(m.GetBinaryPv())
		if popErr := writeBuffer.PopContext("binaryPv"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for binaryPv")
		}
		if _binaryPvErr != nil {
			return errors.Wrap(_binaryPvErr, "Error serializing 'binaryPv' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCredentialStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCredentialStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCredentialStatus) isBACnetConstructedDataCredentialStatus() bool {
	return true
}

func (m *_BACnetConstructedDataCredentialStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}