package urn

import (
	"bytes"
	"fmt"
	"strings"
)

var (
	subNamespaceForSCIM      = []byte("urn:ietf:params:scim")
	subNamespaceForSCIMIndex = len(subNamespaceForSCIM) + 1
	subNamespaceForSCIMTypes = []string{"schemas", "api"}
)

// SCIM represents an URN Sub-namespace for SCIM described in https://www.ietf.org/rfc/rfc7643.txt
type SCIM struct {
	urn   *URN   // Uniform Resource Name
	Type  string // The entity type for SCIM
	Name  string // Defined a major namespace of a schema within SCIM
	Other string // Defined a sub-namespace as needed to uniquely identify a schema
}

// URN returns general URN
func (s *SCIM) URN() *URN {
	return s.urn
}

// String reassembles the URN into a valid URN string.
func (s *SCIM) String() string {
	return s.urn.String()
}

// MarshalJSON marshals the URN to JSON string form (e.g. `"urn:oid:1.2.3.4"`).
func (s *SCIM) MarshalJSON() ([]byte, error) {
	return s.urn.MarshalJSON()
}

// UnmarshalJSON unmarshals a URN from JSON string form (e.g. `"urn:oid:1.2.3.4"`).
func (s *SCIM) UnmarshalJSON(bytes []byte) error {
	return s.urn.UnmarshalJSON(bytes)
}

func hasSCIMTypes(s string) bool {
	for i := range subNamespaceForSCIMTypes {
		if subNamespaceForSCIMTypes[i] == s {
			return true
		}
	}
	return false
}

// ParseAsSCIM is responsible to create an SCIM instance from a byte array.
func ParseAsSCIM(b []byte) (*SCIM, error) {
	urn, ok := Parse(b)
	if !ok {
		return nil, fmt.Errorf("failed to parse: invalid URN")
	}
	if !bytes.HasPrefix(b, subNamespaceForSCIM) {
		return nil, fmt.Errorf("faled to parse: invalid SCIM Sub-namespace")
	}
	r := strings.SplitN(string(b[subNamespaceForSCIMIndex:]), ":", 3)
	typ, name, other := r[0], r[1], r[2]
	if !hasSCIMTypes(typ) {
		return nil, fmt.Errorf(
			`got invalid type in SCIM Sub-namespace, `+
				`either "schemas" or "api": %s`, typ)
	}
	return &SCIM{
		urn:   urn,
		Type:  typ,
		Name:  name,
		Other: other,
	}, nil
}
