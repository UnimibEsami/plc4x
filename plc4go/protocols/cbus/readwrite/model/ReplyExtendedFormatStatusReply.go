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

// ReplyExtendedFormatStatusReply is the corresponding interface of ReplyExtendedFormatStatusReply
type ReplyExtendedFormatStatusReply interface {
	utils.LengthAware
	utils.Serializable
	Reply
	// GetReply returns Reply (property field)
	GetReply() ExtendedFormatStatusReply
	// GetPayloadLength returns PayloadLength (virtual field)
	GetPayloadLength() uint16
}

// ReplyExtendedFormatStatusReplyExactly can be used when we want exactly this type and not a type which fulfills ReplyExtendedFormatStatusReply.
// This is useful for switch cases.
type ReplyExtendedFormatStatusReplyExactly interface {
	ReplyExtendedFormatStatusReply
	isReplyExtendedFormatStatusReply() bool
}

// _ReplyExtendedFormatStatusReply is the data-structure of this message
type _ReplyExtendedFormatStatusReply struct {
	*_Reply
	Reply ExtendedFormatStatusReply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReplyExtendedFormatStatusReply) InitializeParent(parent Reply, peekedByte byte) {
	m.PeekedByte = peekedByte
}

func (m *_ReplyExtendedFormatStatusReply) GetParent() Reply {
	return m._Reply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReplyExtendedFormatStatusReply) GetReply() ExtendedFormatStatusReply {
	return m.Reply
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_ReplyExtendedFormatStatusReply) GetPayloadLength() uint16 {
	return uint16(m.ReplyLength)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReplyExtendedFormatStatusReply factory function for _ReplyExtendedFormatStatusReply
func NewReplyExtendedFormatStatusReply(reply ExtendedFormatStatusReply, peekedByte byte, cBusOptions CBusOptions, replyLength uint16, requestContext RequestContext) *_ReplyExtendedFormatStatusReply {
	_result := &_ReplyExtendedFormatStatusReply{
		Reply:  reply,
		_Reply: NewReply(peekedByte, cBusOptions, replyLength, requestContext),
	}
	_result._Reply._ReplyChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastReplyExtendedFormatStatusReply(structType interface{}) ReplyExtendedFormatStatusReply {
	if casted, ok := structType.(ReplyExtendedFormatStatusReply); ok {
		return casted
	}
	if casted, ok := structType.(*ReplyExtendedFormatStatusReply); ok {
		return *casted
	}
	return nil
}

func (m *_ReplyExtendedFormatStatusReply) GetTypeName() string {
	return "ReplyExtendedFormatStatusReply"
}

func (m *_ReplyExtendedFormatStatusReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ReplyExtendedFormatStatusReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Manual Field (reply)
	lengthInBits += uint16(int32(m.GetLengthInBytes()) * int32(int32(2)))

	return lengthInBits
}

func (m *_ReplyExtendedFormatStatusReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ReplyExtendedFormatStatusReplyParse(readBuffer utils.ReadBuffer, cBusOptions CBusOptions, replyLength uint16, requestContext RequestContext) (ReplyExtendedFormatStatusReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReplyExtendedFormatStatusReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReplyExtendedFormatStatusReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_payloadLength := replyLength
	payloadLength := uint16(_payloadLength)
	_ = payloadLength

	// Manual Field (reply)
	_reply, _replyErr := ReadExtendedFormatStatusReply(readBuffer, payloadLength)
	if _replyErr != nil {
		return nil, errors.Wrap(_replyErr, "Error parsing 'reply' field of ReplyExtendedFormatStatusReply")
	}
	reply := _reply.(ExtendedFormatStatusReply)

	if closeErr := readBuffer.CloseContext("ReplyExtendedFormatStatusReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReplyExtendedFormatStatusReply")
	}

	// Create a partially initialized instance
	_child := &_ReplyExtendedFormatStatusReply{
		Reply: reply,
		_Reply: &_Reply{
			CBusOptions:    cBusOptions,
			ReplyLength:    replyLength,
			RequestContext: requestContext,
		},
	}
	_child._Reply._ReplyChildRequirements = _child
	return _child, nil
}

func (m *_ReplyExtendedFormatStatusReply) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReplyExtendedFormatStatusReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReplyExtendedFormatStatusReply")
		}
		// Virtual field
		if _payloadLengthErr := writeBuffer.WriteVirtual("payloadLength", m.GetPayloadLength()); _payloadLengthErr != nil {
			return errors.Wrap(_payloadLengthErr, "Error serializing 'payloadLength' field")
		}

		// Manual Field (reply)
		_replyErr := WriteExtendedFormatStatusReply(writeBuffer, m.GetReply())
		if _replyErr != nil {
			return errors.Wrap(_replyErr, "Error serializing 'reply' field")
		}

		if popErr := writeBuffer.PopContext("ReplyExtendedFormatStatusReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReplyExtendedFormatStatusReply")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ReplyExtendedFormatStatusReply) isReplyExtendedFormatStatusReply() bool {
	return true
}

func (m *_ReplyExtendedFormatStatusReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}