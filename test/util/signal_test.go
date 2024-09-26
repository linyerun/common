package util

import (
	"fmt"
	"testing"

	"github.com/linyerun/common/util"
)

func TestMonitorStopSignal(t *testing.T) {
	signal := util.MonitorStopSignal()

	<-signal
	fmt.Println("received stop signal")
}
