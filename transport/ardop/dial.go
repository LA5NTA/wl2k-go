// Copyright 2015 Martin Hebnes Pedersen (LA5NTA). All rights reserved.
// Use of this source code is governed by the MIT-license that can be
// found in the LICENSE file.

package ardop

import (
	"fmt"
	"net"

	"github.com/la5nta/wl2k-go/transport"
)

// DialURL dials ardop:// URLs
func (tnc *TNC) DialURL(url *transport.URL) (net.Conn, error) {
	if url.Scheme != "ardop" {
		return nil, transport.ErrUnsupportedScheme
	}

	// Set bandwidth from the URL
	bw := url.Params.Get("bw")
	if bw != "" {
		bandwidth, err := StrToBandwidth(bw)
		if err != nil {
			return nil, err
		}
		if err = tnc.SetARQBandwidth(bandwidth); err != nil {
			return nil, err
		}
	}

	return tnc.Dial(url.Target)
}

func (tnc *TNC) Dial(targetcall string) (net.Conn, error) {
	if tnc.closed {
		return nil, ErrTNCClosed
	}

	if err := tnc.arqCall(targetcall, 10); err != nil {
		return nil, err
	}

	mycall, err := tnc.MyCall()
	if err != nil {
		return nil, fmt.Errorf("Error when getting mycall: %s", err)
	}

	tnc.data = &tncConn{
		remoteAddr: Addr{targetcall},
		localAddr:  Addr{mycall},
		ctrlOut:    tnc.out,
		dataOut:    tnc.dataOut,
		ctrlIn:     tnc.in,
		dataIn:     tnc.dataIn,
		eofChan:    make(chan struct{}),
		isTCP:      tnc.isTCP,
	}

	return tnc.data, nil
}
