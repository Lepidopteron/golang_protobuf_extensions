// Copyright 2013 Matt T. Proud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ext

import (
	"code.google.com/p/goprotobuf/proto"
	"io"
)

// WriteDelimited encodes and dumps a message to the provided writer prefixed
// with a 64-bit varint indicating the length of the encoded message, producing
// a length-delimited record stream, which can be used to chain together
// encoded messages of the same type together in a file.  It returns the total
// number of bytes written and any applicable error.
func WriteDelimited(w io.Writer, m proto.Message) (n int, err error) {
	buffer, err := proto.Marshal(m)
	if err != nil {
		return 0, err
	}

	length := proto.EncodeVarint(uint64(len(buffer)))

	sync, err := w.Write(length)
	if err != nil {
		return sync, err
	}

	n, err = w.Write(buffer)
	return n + sync, err
}