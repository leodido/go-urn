package scimschema

type Type int

const (
	Unsupported Type = iota
	Schemas
	API
)

func (t Type) String() string {
	switch t {
	case Schemas:
		return "schemas"
	case API:
		return "api"
	}

	return ""
}

func TypeFromString(input string) Type {
	switch input {
	case "schemas":
		return Schemas
	case "api":
		return API
	}

	return Unsupported
}
