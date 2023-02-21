package timeseries

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *FieldDataType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 byte
		zb0001, err = dc.ReadByte()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = FieldDataType(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z FieldDataType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteByte(byte(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z FieldDataType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendByte(o, byte(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FieldDataType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 byte
		zb0001, bts, err = msgp.ReadByteBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = FieldDataType(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z FieldDataType) Msgsize() (s int) {
	s = msgp.ByteSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FieldDefinition) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "type":
			{
				var zb0002 byte
				zb0002, err = dc.ReadByte()
				if err != nil {
					err = msgp.WrapError(err, "DataType")
					return
				}
				z.DataType = FieldDataType(zb0002)
			}
		case "pos":
			z.OutputPosition, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "OutputPosition")
				return
			}
		case "stype":
			z.SDataType, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "SDataType")
				return
			}
		case "provider1":
			z.ProviderData1, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "ProviderData1")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *FieldDefinition) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "name"
	err = en.Append(0x85, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "type"
	err = en.Append(0xa4, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteByte(byte(z.DataType))
	if err != nil {
		err = msgp.WrapError(err, "DataType")
		return
	}
	// write "pos"
	err = en.Append(0xa3, 0x70, 0x6f, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt(z.OutputPosition)
	if err != nil {
		err = msgp.WrapError(err, "OutputPosition")
		return
	}
	// write "stype"
	err = en.Append(0xa5, 0x73, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.SDataType)
	if err != nil {
		err = msgp.WrapError(err, "SDataType")
		return
	}
	// write "provider1"
	err = en.Append(0xa9, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x31)
	if err != nil {
		return
	}
	err = en.WriteInt(z.ProviderData1)
	if err != nil {
		err = msgp.WrapError(err, "ProviderData1")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FieldDefinition) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "name"
	o = append(o, 0x85, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "type"
	o = append(o, 0xa4, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendByte(o, byte(z.DataType))
	// string "pos"
	o = append(o, 0xa3, 0x70, 0x6f, 0x73)
	o = msgp.AppendInt(o, z.OutputPosition)
	// string "stype"
	o = append(o, 0xa5, 0x73, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.SDataType)
	// string "provider1"
	o = append(o, 0xa9, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x31)
	o = msgp.AppendInt(o, z.ProviderData1)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FieldDefinition) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "type":
			{
				var zb0002 byte
				zb0002, bts, err = msgp.ReadByteBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "DataType")
					return
				}
				z.DataType = FieldDataType(zb0002)
			}
		case "pos":
			z.OutputPosition, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "OutputPosition")
				return
			}
		case "stype":
			z.SDataType, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SDataType")
				return
			}
		case "provider1":
			z.ProviderData1, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ProviderData1")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldDefinition) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.ByteSize + 4 + msgp.IntSize + 6 + msgp.StringPrefixSize + len(z.SDataType) + 10 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FieldDefinitions) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(FieldDefinitions, zb0002)
	}
	for zb0001 := range *z {
		err = (*z)[zb0001].DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z FieldDefinitions) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0003 := range z {
		err = z[zb0003].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, zb0003)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z FieldDefinitions) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		o, err = z[zb0003].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, zb0003)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FieldDefinitions) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(FieldDefinitions, zb0002)
	}
	for zb0001 := range *z {
		bts, err = (*z)[zb0001].UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z FieldDefinitions) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		s += z[zb0003].Msgsize()
	}
	return
}
