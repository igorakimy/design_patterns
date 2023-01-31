package main

func main() {

	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowsMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}

// Client inserts Lightning connector into computer.
// Lightning connector is plugged into mac machine.
// Client inserts Lightning connector into computer.
// Adapter converts Lightning signal to USB.
// USB connector is plugged into windows machine.
