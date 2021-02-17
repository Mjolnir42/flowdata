/*-
 * Copyright (c) 2021, Jörg Pernfuß
 *
 * Use of this source code is governed by a 2-clause BSD license
 * that can be found in the LICENSE file.
 */

package flowdata // import "github.com/mjolnir42/flowdata"

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

type Bitmap [2]byte

func ParseBitmap(s string) Bitmap {
	i, _ := strconv.ParseInt(s, 0, 16)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	var bm Bitmap
	copy(bm[:], b[len(b)-2:])
	return bm
}

func (bm *Bitmap) Bytes() []byte {
	b := make([]byte, 2)
	copy(b, bm[:])
	return b
}

func (bm *Bitmap) String() string {
	return fmt.Sprintf("%#x", bm)
}

func (bm *Bitmap) Copy() Bitmap {
	b := Bitmap{}
	copy(b[:], bm[:])
	return b
}

func (bm *Bitmap) Check(s string) bool {
	i, _ := strconv.ParseInt(s, 0, 16)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	var chk, res Bitmap
	copy(chk[:], b[len(b)-2:])
	res[0] = bm[0] & chk[0]
	res[1] = bm[1] & chk[1]
	return binary.BigEndian.Uint16(res.Bytes()) > 0
}

func (bm *Bitmap) IsSet(i uint16) bool {
	// this counts from the back
	var p int
	switch {
	case i < 8:
		p = 1
	case i >= 8:
		p = 0
	}
	remainder := i % 8
	if remainder == 1 {
		return bm[p] > bm[p]^1
	}
	return bm[p] > bm[p]^(1<<uint(remainder-1))
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix
