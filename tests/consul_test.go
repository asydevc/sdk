// author: asydevc <asydev@163.com>
// date: 2021-03-05

package tests

import (
	"testing"

	"github.com/asydevc/log/v2"
	"github.com/asydevc/sdk"
)

func TestConsul(t *testing.T) {
	ctx := log.NewContext()
	for i := 0; i < 2; i++ {
		host, err := sdk.Consul.Get(ctx, "activity.module")
		if err != nil {
			t.Errorf("get service error: %v.", err)
		} else {
			t.Logf("get service found: %s.", host)
		}
	}
}
