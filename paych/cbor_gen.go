// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package paych

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

func (t *VoucherInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	// t.Voucher (paych.SignedVoucher) (struct)
	if len("Voucher") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Voucher\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Voucher")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Voucher")); err != nil {
		return err
	}

	if err := t.Voucher.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Proof ([]uint8) (slice)
	if len("Proof") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Proof\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Proof")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Proof")); err != nil {
		return err
	}

	if len(t.Proof) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Proof was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Proof)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Proof); err != nil {
		return err
	}
	return nil
}

func (t *VoucherInfo) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("VoucherInfo: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(br)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Voucher (paych.SignedVoucher) (struct)
		case "Voucher":

			{

				pb, err := br.PeekByte()
				if err != nil {
					return err
				}
				if pb == cbg.CborNull[0] {
					var nbuf [1]byte
					if _, err := br.Read(nbuf[:]); err != nil {
						return err
					}
				} else {
					t.Voucher = new(SignedVoucher)
					if err := t.Voucher.UnmarshalCBOR(br); err != nil {
						return err
					}
				}

			}
			// t.Proof ([]uint8) (slice)
		case "Proof":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Proof: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}
			t.Proof = make([]byte, extra)
			if _, err := io.ReadFull(br, t.Proof); err != nil {
				return err
			}

		default:
			return fmt.Errorf("unknown struct field %d: '%s'", i, name)
		}
	}

	return nil
}
func (t *ChannelInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{166}); err != nil {
		return err
	}

	// t.Channel (address.Address) (struct)
	if len("Channel") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Channel\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Channel")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Channel")); err != nil {
		return err
	}

	if err := t.Channel.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Control (address.Address) (struct)
	if len("Control") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Control\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Control")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Control")); err != nil {
		return err
	}

	if err := t.Control.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Target (address.Address) (struct)
	if len("Target") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Target\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Target")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Target")); err != nil {
		return err
	}

	if err := t.Target.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Direction (uint64) (uint64)
	if len("Direction") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Direction\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Direction")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Direction")); err != nil {
		return err
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Direction))); err != nil {
		return err
	}

	// t.Vouchers ([]*paych.VoucherInfo) (slice)
	if len("Vouchers") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Vouchers\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Vouchers")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Vouchers")); err != nil {
		return err
	}

	if len(t.Vouchers) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Vouchers was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Vouchers)))); err != nil {
		return err
	}
	for _, v := range t.Vouchers {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.NextLane (uint64) (uint64)
	if len("NextLane") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"NextLane\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("NextLane")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("NextLane")); err != nil {
		return err
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.NextLane))); err != nil {
		return err
	}
	return nil
}

func (t *ChannelInfo) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("ChannelInfo: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(br)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Channel (address.Address) (struct)
		case "Channel":

			{

				if err := t.Channel.UnmarshalCBOR(br); err != nil {
					return err
				}

			}
			// t.Control (address.Address) (struct)
		case "Control":

			{

				if err := t.Control.UnmarshalCBOR(br); err != nil {
					return err
				}

			}
			// t.Target (address.Address) (struct)
		case "Target":

			{

				if err := t.Target.UnmarshalCBOR(br); err != nil {
					return err
				}

			}
			// t.Direction (uint64) (uint64)
		case "Direction":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}
			if maj != cbg.MajUnsignedInt {
				return fmt.Errorf("wrong type for uint64 field")
			}
			t.Direction = uint64(extra)
			// t.Vouchers ([]*paych.VoucherInfo) (slice)
		case "Vouchers":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.Vouchers: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}
			if extra > 0 {
				t.Vouchers = make([]*VoucherInfo, extra)
			}
			for i := 0; i < int(extra); i++ {

				var v VoucherInfo
				if err := v.UnmarshalCBOR(br); err != nil {
					return err
				}

				t.Vouchers[i] = &v
			}

			// t.NextLane (uint64) (uint64)
		case "NextLane":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}
			if maj != cbg.MajUnsignedInt {
				return fmt.Errorf("wrong type for uint64 field")
			}
			t.NextLane = uint64(extra)

		default:
			return fmt.Errorf("unknown struct field %d: '%s'", i, name)
		}
	}

	return nil
}
