package mcp2221a

import (
	"errors"
	"fmt"
)

// Chybové konštanty
var (
	ErrInvalidPin  = errors.New("invalid pin for IOC (iba G1 podporovaný)")
	ErrInvalidEdge = errors.New("invalid edge type for IOC")
)

// IOC poskytuje API pre Interrupt-On-Change
type IOC struct {
	mcp *MCP2221A
}

// Enable zapne IOC pre GP1 (G1) s daným edge (RisingEdge/FallingEdge)
func (ioc *IOC) Enable(pin byte, edge IntEdge) error {
	if pin != 1 {
		return ErrInvalidPin
	}
	if edge > RisingFallingEdge {
		return ErrInvalidEdge
	}
	cmd := []byte{0xB2, byte(edge)}
	return ioc.sendFeatureReport(cmd)
}

// Read načíta IOC flag (vracia aktuálny stav G1 pri IOC evente)
func (ioc *IOC) Read() (IntEdge, error) {
	buf := make([]byte, 1)
	_, err := ioc.mcp.Device.GetFeatureReport(buf)
	if err != nil {
		return NoInterrupt, fmt.Errorf("IOC.Read GetFeatureReport: %w", err)
	}
	return IntEdge(buf[0]), nil
}

// Clear resetuje IOC flag
func (ioc *IOC) Clear() error {
	cmd := []byte{0xB1, 0x00}
	return ioc.sendFeatureReport(cmd)
}

// Pomocná funkcia na odoslanie FeatureReport
func (ioc *IOC) sendFeatureReport(data []byte) error {
	return ioc.mcp.Device.SendFeatureReport(data)
}

