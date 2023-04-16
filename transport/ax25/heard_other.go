// Copyright 2015 Martin Hebnes Pedersen (LA5NTA). All rights reserved.
// Use of this source code is governed by the MIT-license that can be
// found in the LICENSE file.

//go:build !(libax25 && cgo)
// +build !libax25 !cgo

package ax25

import "time"

func Heard(axPort string) (map[string]time.Time, error) { return nil, ErrNoLibax25 }
