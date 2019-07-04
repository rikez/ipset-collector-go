package ipset

import "testing"

func TestCIDRParserShouldReturnARange(t *testing.T) {
	iprange, err := CIDRParser("189.68.26.0/24")

	if err != nil {
		t.Errorf("Error must be nil, got %s", err.Error())
	}
	if len(iprange) != 254 {
		t.Errorf("Unexpected length for the range of IPs /24, got %d, expected 254", len(iprange))
	}
}

func TestCIDRParserShouldReturnEmptyArray(t *testing.T) {
	iprange, err := CIDRParser("189.68.26.36")

	if err == nil {
		t.Errorf("Error must not be nil, got nil, expected invalid CIDR address: 189.68.26.36")
	}
	if iprange != nil {
		t.Errorf("The iprange must be nil, got %v", iprange)
	}
}
