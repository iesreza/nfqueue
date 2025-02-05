/**
 * @license
 * Copyright 2018 Telefónica Investigación y Desarrollo, S.A.U
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package nfqueue

import "C"
import (
	"unsafe"
)

//export handle
func handle(id uint32, buffer *C.uchar, len C.int, queueID int) int {
	q := queueRegistry.Get(uint16(queueID))
	if q == nil {
		return 0
	}
	packet := &Packet{
		id:   id,
		Data: C.GoBytes(unsafe.Pointer(buffer), len),
		Q:    q,
	}
	q.handler.Handle(packet)
	return 0
}
