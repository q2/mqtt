/*
 * Copyright (c) 2013 IBM Corp.
 *
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v1.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v10.html
 *
 * Contributors:
 *    Seth Hoenig
 *    Allan Stockdill-Mander
 *    Mike Robertson
 */

package mqtt

import "testing"

func Test_NewClient_simple(t *testing.T) {
	ops := NewClientOptions().SetClientId("foo").SetBroker("tcp://10.10.0.1:1883")
	c := NewClient(ops)

	if c == nil {
		t.Fatalf("ops is nil")
	}

	if c.options.clientId != "foo" {
		t.Fatalf("bad client id")
	}

	if c.options.server.Scheme != "tcp" {
		t.Fatalf("bad server scheme")
	}

	if c.options.server.Host != "10.10.0.1:1883" {
		t.Fatalf("bad server host")
	}
}
