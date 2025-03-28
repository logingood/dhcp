package dhcpv6

import (
	"fmt"

	"github.com/logingood/dhcp/rfc1035label"
	"github.com/u-root/uio/uio"
)

// OptFQDN implements OptionFQDN option.
//
// https://tools.ietf.org/html/rfc4704
type OptFQDN struct {
	Flags      uint8
	DomainName *rfc1035label.Labels
}

// Code returns the option code.
func (op *OptFQDN) Code() OptionCode {
	return OptionFQDN
}

// ToBytes serializes the option and returns it as a sequence of bytes
func (op *OptFQDN) ToBytes() []byte {
	buf := uio.NewBigEndianBuffer(nil)
	buf.Write8(op.Flags)
	buf.WriteBytes(op.DomainName.ToBytes())
	return buf.Data()
}

func (op *OptFQDN) String() string {
	return fmt.Sprintf("%s: {Flags=%d DomainName=%s}", op.Code(), op.Flags, op.DomainName)
}

// FromBytes deserializes from bytes to build a OptFQDN structure.
func (op *OptFQDN) FromBytes(data []byte) error {
	var err error
	buf := uio.NewBigEndianBuffer(data)
	op.Flags = buf.Read8()
	op.DomainName, err = rfc1035label.FromBytes(buf.ReadAll())
	if err != nil {
		return err
	}
	return buf.FinError()
}
