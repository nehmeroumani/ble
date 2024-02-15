package linux

import (
	"github.com/nehmeroumani/ble"
	"github.com/nehmeroumani/ble/linux/gatt"
	"github.com/nehmeroumani/ble/linux/hci"
	"github.com/nehmeroumani/ble/linux/rfcomm"
)

var testDevice ble.Device = &Device{}
var testConn ble.Conn = &hci.Conn{}
var testClient ble.ClientBLE = &gatt.Client{}
var testRFCOMMClient ble.ClientRFCOMM = &rfcomm.Client{}
var testAdv ble.Advertisement = &hci.Advertisement{}
