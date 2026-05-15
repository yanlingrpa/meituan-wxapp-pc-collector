package meituan

import (
	"fmt"

	"github.com/yanlingrpa/wxapp-pc-toolkits/wxapputils"
	"yanlingrpa.com/yanling/protocol/script"
)

func onAppReady(event script.Event) {
	data := event.Data.(wxapputils.AppReadyData)
	fmt.Printf("app is ready, payload: %+v\n", data)
}
