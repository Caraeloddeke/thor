// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package log

import (
	"math/rand"
	"testing"
)

var sink []byte

func BenchmarkPrettyInt64Logfmt(b *testing.B) {
	buf := make([]byte, 100)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = appendInt64(buf, rand.Int63())
	}
}

func BenchmarkPrettyUint64Logfmt(b *testing.B) {
	buf := make([]byte, 100)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = appendUint64(buf, rand.Uint64(), false)
	}
}