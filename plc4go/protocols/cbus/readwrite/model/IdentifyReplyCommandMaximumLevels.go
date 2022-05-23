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
)

// Code generated by code-generation. DO NOT EDIT.

// IdentifyReplyCommandMaximumLevels is the data-structure of this message
type IdentifyReplyCommandMaximumLevels struct {
	*IdentifyReplyCommand
}

// IIdentifyReplyCommandMaximumLevels is the corresponding interface of IdentifyReplyCommandMaximumLevels
type IIdentifyReplyCommandMaximumLevels interface {
	IIdentifyReplyCommand
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

func (m *IdentifyReplyCommandMaximumLevels) GetAttribute() Attribute {
	return Attribute_MaximumLevels
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *IdentifyReplyCommandMaximumLevels) InitializeParent(parent *IdentifyReplyCommand) {}

func (m *IdentifyReplyCommandMaximumLevels) GetParent() *IdentifyReplyCommand {
	return m.IdentifyReplyCommand
}

// NewIdentifyReplyCommandMaximumLevels factory function for IdentifyReplyCommandMaximumLevels
func NewIdentifyReplyCommandMaximumLevels() *IdentifyReplyCommandMaximumLevels {
	_result := &IdentifyReplyCommandMaximumLevels{
		IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	_result.Child = _result
	return _result
}

func CastIdentifyReplyCommandMaximumLevels(structType interface{}) *IdentifyReplyCommandMaximumLevels {
	if casted, ok := structType.(IdentifyReplyCommandMaximumLevels); ok {
		return &casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandMaximumLevels); ok {
		return casted
	}
	if casted, ok := structType.(IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandMaximumLevels(casted.Child)
	}
	if casted, ok := structType.(*IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandMaximumLevels(casted.Child)
	}
	return nil
}

func (m *IdentifyReplyCommandMaximumLevels) GetTypeName() string {
	return "IdentifyReplyCommandMaximumLevels"
}

func (m *IdentifyReplyCommandMaximumLevels) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *IdentifyReplyCommandMaximumLevels) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *IdentifyReplyCommandMaximumLevels) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandMaximumLevelsParse(readBuffer utils.ReadBuffer, attribute Attribute) (*IdentifyReplyCommandMaximumLevels, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandMaximumLevels"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandMaximumLevels"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &IdentifyReplyCommandMaximumLevels{
		IdentifyReplyCommand: &IdentifyReplyCommand{},
	}
	_child.IdentifyReplyCommand.Child = _child
	return _child, nil
}

func (m *IdentifyReplyCommandMaximumLevels) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandMaximumLevels"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandMaximumLevels"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *IdentifyReplyCommandMaximumLevels) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}