package urn

import (
    "fmt"
    "strings"
)

var (
    errPrefix         = "expecting the prefix to be the \"urn\" string (whatever case) [col %d]"
    errIdentifier     = "expecting the identifier to be string (1..31 alnum chars, also containing dashes but not at its start) [col %d]"
    errSpecificString = "expecting the specific string to be a string containing alnum, hex, or others ([()+,-.:=@;$_!*']) chars [col %d]"
    errNoUrnWithinID  = "expecting the identifier to not contain the \"urn\" reserved string [col %d]"
    errHex            = "expecting the specific string hex chars to be well-formed (%%alnum{2}) [col %d]"
    errParse          = "parsing error [col %d]"
)

%%{
machine urn;

# unsigned alphabet
alphtype uint8;

action mark {
    m.pb = m.p
}

action set_pre {
    m.output.prefix = string(m.text())
}

action set_nid {
    m.output.ID = string(m.text())
}

action set_hex {
    m.output.norm += strings.ToLower(string(m.text()))
    m.output.SS += string(m.text())
}

action set_sss {
    m.output.norm += string(m.text())
    m.output.SS += string(m.text())
}

action err_pre {
    m.err = fmt.Errorf(errPrefix, m.p)
    fhold;
    fgoto fail;
}

action err_nid {
    m.err = fmt.Errorf(errIdentifier, m.p)
    fhold;
    fgoto fail;
}

action err_nss {
    m.err = fmt.Errorf(errSpecificString, m.p)
    fhold;
    fgoto fail;
}

action err_urn {
    m.err = fmt.Errorf(errNoUrnWithinID, m.p)
    fhold;
    fgoto fail;
}

action err_hex {
    m.err = fmt.Errorf(errHex, m.p)
    fhold;
    fgoto fail;
}

action err_parse {
    m.err = fmt.Errorf(errParse, m.p)
    fhold;
    fgoto fail;
}

pre = ([uU][rR][nN] @err(err_pre)) >mark %set_pre;

nid = (alnum >mark (alnum | '-'){0,31}) %set_nid;

hex = '%' >mark alnum{2} %set_hex $err(err_hex);

sss = (alnum | [()+,\-.:=@;$_!*']) >mark %set_sss;

nss = (sss | hex)+ $err(err_nss);

fail := (any - [\n\r])* @err{ fgoto main; };

main := (pre ':' (nid - pre %err(err_urn)) $err(err_nid) ':' nss) $err(err_parse);

}%%

%% write data;

// Machine is the interface representing the FSM
type Machine interface {
    Err() error
    Parse(input []byte) (*URN, error)
}

type machine struct {
    data         []byte
    cs           int
    p, pe, eof   int
    pb           int
    err          error
    output       *URN
}

// NewMachine creates a new FSM able to parse RFC 2141 strings.
func NewMachine() Machine {
    m := &machine{}

    %% access m.;
    %% variable p m.p;
    %% variable pe m.pe;
    %% variable eof m.eof;
    %% variable data m.data;

    return m
}

// Err returns the error that occurred on the last call to Parse.
//
// If the result is nil, then the line was parsed successfully.
func (m *machine) Err() error {
    return m.err
}

func (m *machine) text() []byte {
    return m.data[m.pb:m.p]
}

// Parse parses the input byte array as a RFC 2141 string.
func (m *machine) Parse(input []byte) (*URN, error) {
    m.data = input
    m.p = 0
    m.pb = 0
    m.pe = len(input)
    m.eof = len(input)
    m.err = nil
    m.output = &URN{}

    %% write init;
    %% write exec;

    if m.cs < urn_first_final || m.cs == urn_en_fail {
        return nil, m.err
    }

    return m.output, nil
}
