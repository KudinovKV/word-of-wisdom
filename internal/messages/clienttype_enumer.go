// Code generated by "enumer -type=ClientType -json"; DO NOT EDIT.

package messages

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _ClientTypeName = "ClientTypeRequestClientTypeResponse"

var _ClientTypeIndex = [...]uint8{0, 17, 35}

const _ClientTypeLowerName = "clienttyperequestclienttyperesponse"

func (i ClientType) String() string {
	if i < 0 || i >= ClientType(len(_ClientTypeIndex)-1) {
		return fmt.Sprintf("ClientType(%d)", i)
	}
	return _ClientTypeName[_ClientTypeIndex[i]:_ClientTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ClientTypeNoOp() {
	var x [1]struct{}
	_ = x[ClientTypeRequest-(0)]
	_ = x[ClientTypeResponse-(1)]
}

var _ClientTypeValues = []ClientType{ClientTypeRequest, ClientTypeResponse}

var _ClientTypeNameToValueMap = map[string]ClientType{
	_ClientTypeName[0:17]:       ClientTypeRequest,
	_ClientTypeLowerName[0:17]:  ClientTypeRequest,
	_ClientTypeName[17:35]:      ClientTypeResponse,
	_ClientTypeLowerName[17:35]: ClientTypeResponse,
}

var _ClientTypeNames = []string{
	_ClientTypeName[0:17],
	_ClientTypeName[17:35],
}

// ClientTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ClientTypeString(s string) (ClientType, error) {
	if val, ok := _ClientTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ClientTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ClientType values", s)
}

// ClientTypeValues returns all values of the enum
func ClientTypeValues() []ClientType {
	return _ClientTypeValues
}

// ClientTypeStrings returns a slice of all String values of the enum
func ClientTypeStrings() []string {
	strs := make([]string, len(_ClientTypeNames))
	copy(strs, _ClientTypeNames)
	return strs
}

// IsAClientType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ClientType) IsAClientType() bool {
	for _, v := range _ClientTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for ClientType
func (i ClientType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClientType
func (i *ClientType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ClientType should be a string, got %s", data)
	}

	var err error
	*i, err = ClientTypeString(s)
	return err
}
