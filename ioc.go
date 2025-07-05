package mcp2221a

// IOCEdge typ reprezentuje detekovaný edge
type IOCEdge byte

const (
	DisableEdge        IOCEdge = 0x00
	RisingEdge         IOCEdge = 0x01
	FallingEdge        IOCEdge = 0x02
	RisingFallingEdge  IOCEdge = 0x03
)

// IOC poskytuje API pre Interrupt-On-Change funkciu
type IOC struct {
	mcp *MCP2221A
}

// Enable povolí Interrupt-On-Change pre GP1 (G1)
func (i *IOC) Enable(pin byte, edge IOCEdge) error {
	// GP1 = pin 1, podľa MCP2221A datasheetu
	if pin != 1 {
		return ErrInvalidPin
	}
	if edge > RisingFallingEdge {
		return ErrInvalidEdge
	}
	// Príkaz 0xB2 (Enable Interrupt-On-Change)
	cmd := []byte{0xB2, edge}
	return i.mcp.sendFeatureReport(cmd)
}

// Read načíta IOC flag pre GP1 (G1)
func (i *IOC) Read() (IOCEdge, error) {
	// Príkaz 0xB0 (Get Interrupt Flag)
	buf, err := i.mcp.getFeatureReport(0xB0, 1)
	if err != nil {
		return DisableEdge, err
	}
	return IOCEdge(buf[0]), nil
}

// Clear resetuje IOC flag
func (i *IOC) Clear() error {
	// Príkaz 0xB1 (Clear Interrupt Flag)
	cmd := []byte{0xB1, 0x00}
	return i.mcp.sendFeatureReport(cmd)
}

