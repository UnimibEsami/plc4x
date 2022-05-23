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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple is the data-structure of this message
type BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple struct {
	*BACnetConfirmedServiceRequest
	SubscriberProcessIdentifier *BACnetContextTagUnsignedInteger
	InitiatingDeviceIdentifier  *BACnetContextTagObjectIdentifier
	TimeRemaining               *BACnetContextTagUnsignedInteger
	Timestamp                   *BACnetTimeStampEnclosed
	ListOfCovNotifications      *ListOfCovNotificationsList

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple is the corresponding interface of BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
type IBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple interface {
	IBACnetConfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() *BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier (property field)
	GetInitiatingDeviceIdentifier() *BACnetContextTagObjectIdentifier
	// GetTimeRemaining returns TimeRemaining (property field)
	GetTimeRemaining() *BACnetContextTagUnsignedInteger
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() *BACnetTimeStampEnclosed
	// GetListOfCovNotifications returns ListOfCovNotifications (property field)
	GetListOfCovNotifications() *ListOfCovNotificationsList
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

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetParent() *BACnetConfirmedServiceRequest {
	return m.BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetSubscriberProcessIdentifier() *BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetInitiatingDeviceIdentifier() *BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetTimeRemaining() *BACnetContextTagUnsignedInteger {
	return m.TimeRemaining
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetTimestamp() *BACnetTimeStampEnclosed {
	return m.Timestamp
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetListOfCovNotifications() *ListOfCovNotificationsList {
	return m.ListOfCovNotifications
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple factory function for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
func NewBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple(subscriberProcessIdentifier *BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier *BACnetContextTagObjectIdentifier, timeRemaining *BACnetContextTagUnsignedInteger, timestamp *BACnetTimeStampEnclosed, listOfCovNotifications *ListOfCovNotificationsList, serviceRequestLength uint16) *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple {
	_result := &BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple{
		SubscriberProcessIdentifier:   subscriberProcessIdentifier,
		InitiatingDeviceIdentifier:    initiatingDeviceIdentifier,
		TimeRemaining:                 timeRemaining,
		Timestamp:                     timestamp,
		ListOfCovNotifications:        listOfCovNotifications,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple(structType interface{}) *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits()

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits()

	// Simple field (timeRemaining)
	lengthInBits += m.TimeRemaining.GetLengthInBits()

	// Optional Field (timestamp)
	if m.Timestamp != nil {
		lengthInBits += (*m.Timestamp).GetLengthInBits()
	}

	// Simple field (listOfCovNotifications)
	lengthInBits += m.ListOfCovNotifications.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (subscriberProcessIdentifier)
	if pullErr := readBuffer.PullContext("subscriberProcessIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_subscriberProcessIdentifier, _subscriberProcessIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _subscriberProcessIdentifierErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierErr, "Error parsing 'subscriberProcessIdentifier' field")
	}
	subscriberProcessIdentifier := CastBACnetContextTagUnsignedInteger(_subscriberProcessIdentifier)
	if closeErr := readBuffer.CloseContext("subscriberProcessIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (initiatingDeviceIdentifier)
	if pullErr := readBuffer.PullContext("initiatingDeviceIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_initiatingDeviceIdentifier, _initiatingDeviceIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _initiatingDeviceIdentifierErr != nil {
		return nil, errors.Wrap(_initiatingDeviceIdentifierErr, "Error parsing 'initiatingDeviceIdentifier' field")
	}
	initiatingDeviceIdentifier := CastBACnetContextTagObjectIdentifier(_initiatingDeviceIdentifier)
	if closeErr := readBuffer.CloseContext("initiatingDeviceIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (timeRemaining)
	if pullErr := readBuffer.PullContext("timeRemaining"); pullErr != nil {
		return nil, pullErr
	}
	_timeRemaining, _timeRemainingErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _timeRemainingErr != nil {
		return nil, errors.Wrap(_timeRemainingErr, "Error parsing 'timeRemaining' field")
	}
	timeRemaining := CastBACnetContextTagUnsignedInteger(_timeRemaining)
	if closeErr := readBuffer.CloseContext("timeRemaining"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (timestamp) (Can be skipped, if a given expression evaluates to false)
	var timestamp *BACnetTimeStampEnclosed = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("timestamp"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetTimeStampEnclosedParse(readBuffer, uint8(3))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'timestamp' field")
		default:
			timestamp = CastBACnetTimeStampEnclosed(_val)
			if closeErr := readBuffer.CloseContext("timestamp"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Simple Field (listOfCovNotifications)
	if pullErr := readBuffer.PullContext("listOfCovNotifications"); pullErr != nil {
		return nil, pullErr
	}
	_listOfCovNotifications, _listOfCovNotificationsErr := ListOfCovNotificationsListParse(readBuffer, uint8(uint8(4)))
	if _listOfCovNotificationsErr != nil {
		return nil, errors.Wrap(_listOfCovNotificationsErr, "Error parsing 'listOfCovNotifications' field")
	}
	listOfCovNotifications := CastListOfCovNotificationsList(_listOfCovNotifications)
	if closeErr := readBuffer.CloseContext("listOfCovNotifications"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple{
		SubscriberProcessIdentifier:   CastBACnetContextTagUnsignedInteger(subscriberProcessIdentifier),
		InitiatingDeviceIdentifier:    CastBACnetContextTagObjectIdentifier(initiatingDeviceIdentifier),
		TimeRemaining:                 CastBACnetContextTagUnsignedInteger(timeRemaining),
		Timestamp:                     CastBACnetTimeStampEnclosed(timestamp),
		ListOfCovNotifications:        CastListOfCovNotificationsList(listOfCovNotifications),
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); pushErr != nil {
			return pushErr
		}

		// Simple Field (subscriberProcessIdentifier)
		if pushErr := writeBuffer.PushContext("subscriberProcessIdentifier"); pushErr != nil {
			return pushErr
		}
		_subscriberProcessIdentifierErr := m.SubscriberProcessIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("subscriberProcessIdentifier"); popErr != nil {
			return popErr
		}
		if _subscriberProcessIdentifierErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierErr, "Error serializing 'subscriberProcessIdentifier' field")
		}

		// Simple Field (initiatingDeviceIdentifier)
		if pushErr := writeBuffer.PushContext("initiatingDeviceIdentifier"); pushErr != nil {
			return pushErr
		}
		_initiatingDeviceIdentifierErr := m.InitiatingDeviceIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("initiatingDeviceIdentifier"); popErr != nil {
			return popErr
		}
		if _initiatingDeviceIdentifierErr != nil {
			return errors.Wrap(_initiatingDeviceIdentifierErr, "Error serializing 'initiatingDeviceIdentifier' field")
		}

		// Simple Field (timeRemaining)
		if pushErr := writeBuffer.PushContext("timeRemaining"); pushErr != nil {
			return pushErr
		}
		_timeRemainingErr := m.TimeRemaining.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timeRemaining"); popErr != nil {
			return popErr
		}
		if _timeRemainingErr != nil {
			return errors.Wrap(_timeRemainingErr, "Error serializing 'timeRemaining' field")
		}

		// Optional Field (timestamp) (Can be skipped, if the value is null)
		var timestamp *BACnetTimeStampEnclosed = nil
		if m.Timestamp != nil {
			if pushErr := writeBuffer.PushContext("timestamp"); pushErr != nil {
				return pushErr
			}
			timestamp = m.Timestamp
			_timestampErr := timestamp.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("timestamp"); popErr != nil {
				return popErr
			}
			if _timestampErr != nil {
				return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
			}
		}

		// Simple Field (listOfCovNotifications)
		if pushErr := writeBuffer.PushContext("listOfCovNotifications"); pushErr != nil {
			return pushErr
		}
		_listOfCovNotificationsErr := m.ListOfCovNotifications.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("listOfCovNotifications"); popErr != nil {
			return popErr
		}
		if _listOfCovNotificationsErr != nil {
			return errors.Wrap(_listOfCovNotificationsErr, "Error serializing 'listOfCovNotifications' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}