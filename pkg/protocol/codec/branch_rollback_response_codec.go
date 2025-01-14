/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package codec

import (
	"github.com/seata/seata-go/pkg/common/bytes"
	"github.com/seata/seata-go/pkg/protocol/branch"
	"github.com/seata/seata-go/pkg/protocol/message"
)

func init() {
	GetCodecManager().RegisterCodec(CodecTypeSeata, &BranchRollbackResponseCodec{})
}

type BranchRollbackResponseCodec struct {
}

func (g *BranchRollbackResponseCodec) Decode(in []byte) interface{} {
	data := message.BranchRollbackResponse{}
	buf := bytes.NewByteBuffer(in)

	data.Xid = bytes.ReadString16Length(buf)
	data.BranchId = int64(bytes.ReadUInt64(buf))
	data.BranchStatus = branch.BranchStatus(bytes.ReadByte(buf))

	return data
}

func (g *BranchRollbackResponseCodec) Encode(in interface{}) []byte {
	data, _ := in.(message.BranchRollbackResponse)
	buf := bytes.NewByteBuffer([]byte{})

	bytes.WriteString16Length(data.Xid, buf)
	buf.WriteInt64(data.BranchId)
	buf.WriteByte(byte(data.BranchStatus))

	return buf.Bytes()
}

func (g *BranchRollbackResponseCodec) GetMessageType() message.MessageType {
	return message.MessageType_BranchRollbackResult
}
