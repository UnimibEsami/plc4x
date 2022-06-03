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

// BACnetConstructedDataMaskedAlarmValues is the data-structure of this message
type BACnetConstructedDataMaskedAlarmValues struct {
	*BACnetConstructedData
	MaskedAlarmValues []*BACnetDoorAlarmStateTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataMaskedAlarmValues is the corresponding interface of BACnetConstructedDataMaskedAlarmValues
type IBACnetConstructedDataMaskedAlarmValues interface {
	IBACnetConstructedData
	// GetMaskedAlarmValues returns MaskedAlarmValues (property field)
	GetMaskedAlarmValues() []*BACnetDoorAlarmStateTagged
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

func (m *BACnetConstructedDataMaskedAlarmValues) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMaskedAlarmValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MASKED_ALARM_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMaskedAlarmValues) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMaskedAlarmValues) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMaskedAlarmValues) GetMaskedAlarmValues() []*BACnetDoorAlarmStateTagged {
	return m.MaskedAlarmValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaskedAlarmValues factory function for BACnetConstructedDataMaskedAlarmValues
func NewBACnetConstructedDataMaskedAlarmValues(maskedAlarmValues []*BACnetDoorAlarmStateTagged, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataMaskedAlarmValues {
	_result := &BACnetConstructedDataMaskedAlarmValues{
		MaskedAlarmValues:     maskedAlarmValues,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMaskedAlarmValues(structType interface{}) *BACnetConstructedDataMaskedAlarmValues {
	if casted, ok := structType.(BACnetConstructedDataMaskedAlarmValues); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaskedAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaskedAlarmValues(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMaskedAlarmValues(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMaskedAlarmValues) GetTypeName() string {
	return "BACnetConstructedDataMaskedAlarmValues"
}

func (m *BACnetConstructedDataMaskedAlarmValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMaskedAlarmValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.MaskedAlarmValues) > 0 {
		for _, element := range m.MaskedAlarmValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataMaskedAlarmValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMaskedAlarmValuesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataMaskedAlarmValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaskedAlarmValues"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (maskedAlarmValues)
	if pullErr := readBuffer.PullContext("maskedAlarmValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	maskedAlarmValues := make([]*BACnetDoorAlarmStateTagged, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDoorAlarmStateTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'maskedAlarmValues' field")
			}
			maskedAlarmValues = append(maskedAlarmValues, CastBACnetDoorAlarmStateTagged(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("maskedAlarmValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaskedAlarmValues"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMaskedAlarmValues{
		MaskedAlarmValues:     maskedAlarmValues,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMaskedAlarmValues) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaskedAlarmValues"); pushErr != nil {
			return pushErr
		}

		// Array Field (maskedAlarmValues)
		if m.MaskedAlarmValues != nil {
			if pushErr := writeBuffer.PushContext("maskedAlarmValues", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.MaskedAlarmValues {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'maskedAlarmValues' field")
				}
			}
			if popErr := writeBuffer.PopContext("maskedAlarmValues", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaskedAlarmValues"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMaskedAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}