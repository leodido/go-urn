package urn

type SCIMType int

const (
	Unsupported SCIMType = iota
	Schemas
	API
	Param
)

func (t SCIMType) String() string {
	switch t {
	case Schemas:
		return "schemas"
	case API:
		return "api"
	case Param:
		return "param"
	}

	return ""
}

func SCIMTypeFromString(input string) SCIMType {
	switch input {
	case "schemas":
		return Schemas
	case "api":
		return API
	case "param":
		return Param
	}

	return Unsupported
}

type SCIM struct {
	Type  SCIMType
	Name  string
	Other string
	pos   int
}
