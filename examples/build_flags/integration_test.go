// +build integration

package buildflags

import "testing"

func Test_Integration(t *testing.T) {
	t.Log("integration test only runs with integration flag")
}
