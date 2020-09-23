// env_test.go.

// +build test

////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019..2020 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package env

import (
	"os"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_GetEnv(t *testing.T) {

	var aTest *tester.Test
	var err error
	var result string

	aTest = tester.New(t)

	// Test #1. Normal Data.
	err = os.Setenv("TEST_ENV_A", "XYZ")
	aTest.MustBeNoError(err)
	result, err = GetEnv("TEST_ENV_A")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, "XYZ")

	// Test #2. Empty Variable.
	// Ensure that it is really empty before this Test.
	aTest.MustBeEqual(len(os.Getenv("TEST_ENV_B")), 0)
	result, err = GetEnv("TEST_ENV_B")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, "")
}
