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

// BACnetProgramRequestTagged is the data-structure of this message
type BACnetProgramRequestTagged struct {
	Header *BACnetTagHeader
	Value  BACnetProgramRequest

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

// IBACnetProgramRequestTagged is the corresponding interface of BACnetProgramRequestTagged
type IBACnetProgramRequestTagged interface {
	// GetHeader returns Header (property field)
	GetHeader() *BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetProgramRequest
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetProgramRequestTagged) GetHeader() *BACnetTagHeader {
	return m.Header
}

func (m *BACnetProgramRequestTagged) GetValue() BACnetProgramRequest {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetProgramRequestTagged factory function for BACnetProgramRequestTagged
func NewBACnetProgramRequestTagged(header *BACnetTagHeader, value BACnetProgramRequest, tagNumber uint8, tagClass TagClass) *BACnetProgramRequestTagged {
	return &BACnetProgramRequestTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

func CastBACnetProgramRequestTagged(structType interface{}) *BACnetProgramRequestTagged {
	if casted, ok := structType.(BACnetProgramRequestTagged); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetProgramRequestTagged); ok {
		return casted
	}
	return nil
}

func (m *BACnetProgramRequestTagged) GetTypeName() string {
	return "BACnetProgramRequestTagged"
}

func (m *BACnetProgramRequestTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetProgramRequestTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *BACnetProgramRequestTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetProgramRequestTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (*BACnetProgramRequestTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetProgramRequestTagged"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, pullErr
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field")
	}
	header := CastBACnetTagHeader(_header)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, closeErr
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, utils.ParseValidationError{"tag class doesn't match"}
	}

	// Validation
	if !(bool(bool(bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool(bool(bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, utils.ParseAssertError{"tagnumber doesn't match"}
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetProgramRequest_READY)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value.(BACnetProgramRequest)

	if closeErr := readBuffer.CloseContext("BACnetProgramRequestTagged"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetProgramRequestTagged(header, value, tagNumber, tagClass), nil
}

func (m *BACnetProgramRequestTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetProgramRequestTagged"); pushErr != nil {
		return pushErr
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return pushErr
	}
	_headerErr := m.Header.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return popErr
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetProgramRequestTagged"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetProgramRequestTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}