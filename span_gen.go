package main

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Payload) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0003 uint32
	zb0003, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0003) {
		(*z) = (*z)[:zb0003]
	} else {
		(*z) = make(Payload, zb0003)
	}
	for zb0001 := range *z {
		var zb0004 uint32
		zb0004, err = dc.ReadArrayHeader()
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
		if cap((*z)[zb0001]) >= int(zb0004) {
			(*z)[zb0001] = ((*z)[zb0001])[:zb0004]
		} else {
			(*z)[zb0001] = make([]Span, zb0004)
		}
		for zb0002 := range (*z)[zb0001] {
			err = (*z)[zb0001][zb0002].DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, zb0001, zb0002)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Payload) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0005 := range z {
		err = en.WriteArrayHeader(uint32(len(z[zb0005])))
		if err != nil {
			err = msgp.WrapError(err, zb0005)
			return
		}
		for zb0006 := range z[zb0005] {
			err = z[zb0005][zb0006].EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, zb0005, zb0006)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Payload) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0005 := range z {
		o = msgp.AppendArrayHeader(o, uint32(len(z[zb0005])))
		for zb0006 := range z[zb0005] {
			o, err = z[zb0005][zb0006].MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, zb0005, zb0006)
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Payload) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0003) {
		(*z) = (*z)[:zb0003]
	} else {
		(*z) = make(Payload, zb0003)
	}
	for zb0001 := range *z {
		var zb0004 uint32
		zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
		if cap((*z)[zb0001]) >= int(zb0004) {
			(*z)[zb0001] = ((*z)[zb0001])[:zb0004]
		} else {
			(*z)[zb0001] = make([]Span, zb0004)
		}
		for zb0002 := range (*z)[zb0001] {
			bts, err = (*z)[zb0001][zb0002].UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, zb0001, zb0002)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Payload) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0005 := range z {
		s += msgp.ArrayHeaderSize
		for zb0006 := range z[zb0005] {
			s += z[zb0005][zb0006].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Span) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "service":
			z.Service, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Service")
				return
			}
		case "resource":
			z.Resource, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Resource")
				return
			}
		case "type":
			z.Type, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Type")
				return
			}
		case "start":
			z.Start, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Start")
				return
			}
		case "duration":
			z.Duration, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Duration")
				return
			}
		case "meta":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Meta")
				return
			}
			if z.Meta == nil {
				z.Meta = make(map[string]string, zb0002)
			} else if len(z.Meta) > 0 {
				for key := range z.Meta {
					delete(z.Meta, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 string
				za0001, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Meta")
					return
				}
				za0002, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Meta", za0001)
					return
				}
				z.Meta[za0001] = za0002
			}
		case "metrics":
			var zb0003 uint32
			zb0003, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Metrics")
				return
			}
			if z.Metrics == nil {
				z.Metrics = make(map[string]float64, zb0003)
			} else if len(z.Metrics) > 0 {
				for key := range z.Metrics {
					delete(z.Metrics, key)
				}
			}
			for zb0003 > 0 {
				zb0003--
				var za0003 string
				var za0004 float64
				za0003, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Metrics")
					return
				}
				za0004, err = dc.ReadFloat64()
				if err != nil {
					err = msgp.WrapError(err, "Metrics", za0003)
					return
				}
				z.Metrics[za0003] = za0004
			}
		case "span_id":
			z.SpanID, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "SpanID")
				return
			}
		case "trace_id":
			z.TraceID, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "TraceID")
				return
			}
		case "parent_id":
			z.ParentID, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "ParentID")
				return
			}
		case "error":
			z.Error, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "Error")
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
func (z *Span) EncodeMsg(en *msgp.Writer) (err error) {
	// omitempty: check for empty values
	zb0001Len := uint32(12)
	var zb0001Mask uint16 /* 12 bits */
	if z.Meta == nil {
		zb0001Len--
		zb0001Mask |= 0x40
	}
	if z.Metrics == nil {
		zb0001Len--
		zb0001Mask |= 0x80
	}
	// variable map header, size zb0001Len
	err = en.Append(0x80 | uint8(zb0001Len))
	if err != nil {
		return
	}
	if zb0001Len == 0 {
		return
	}
	// write "name"
	err = en.Append(0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "service"
	err = en.Append(0xa7, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Service)
	if err != nil {
		err = msgp.WrapError(err, "Service")
		return
	}
	// write "resource"
	err = en.Append(0xa8, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Resource)
	if err != nil {
		err = msgp.WrapError(err, "Resource")
		return
	}
	// write "type"
	err = en.Append(0xa4, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Type)
	if err != nil {
		err = msgp.WrapError(err, "Type")
		return
	}
	// write "start"
	err = en.Append(0xa5, 0x73, 0x74, 0x61, 0x72, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Start)
	if err != nil {
		err = msgp.WrapError(err, "Start")
		return
	}
	// write "duration"
	err = en.Append(0xa8, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Duration)
	if err != nil {
		err = msgp.WrapError(err, "Duration")
		return
	}
	if (zb0001Mask & 0x40) == 0 { // if not empty
		// write "meta"
		err = en.Append(0xa4, 0x6d, 0x65, 0x74, 0x61)
		if err != nil {
			return
		}
		err = en.WriteMapHeader(uint32(len(z.Meta)))
		if err != nil {
			err = msgp.WrapError(err, "Meta")
			return
		}
		for za0001, za0002 := range z.Meta {
			err = en.WriteString(za0001)
			if err != nil {
				err = msgp.WrapError(err, "Meta")
				return
			}
			err = en.WriteString(za0002)
			if err != nil {
				err = msgp.WrapError(err, "Meta", za0001)
				return
			}
		}
	}
	if (zb0001Mask & 0x80) == 0 { // if not empty
		// write "metrics"
		err = en.Append(0xa7, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73)
		if err != nil {
			return
		}
		err = en.WriteMapHeader(uint32(len(z.Metrics)))
		if err != nil {
			err = msgp.WrapError(err, "Metrics")
			return
		}
		for za0003, za0004 := range z.Metrics {
			err = en.WriteString(za0003)
			if err != nil {
				err = msgp.WrapError(err, "Metrics")
				return
			}
			err = en.WriteFloat64(za0004)
			if err != nil {
				err = msgp.WrapError(err, "Metrics", za0003)
				return
			}
		}
	}
	// write "span_id"
	err = en.Append(0xa7, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.SpanID)
	if err != nil {
		err = msgp.WrapError(err, "SpanID")
		return
	}
	// write "trace_id"
	err = en.Append(0xa8, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.TraceID)
	if err != nil {
		err = msgp.WrapError(err, "TraceID")
		return
	}
	// write "parent_id"
	err = en.Append(0xa9, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.ParentID)
	if err != nil {
		err = msgp.WrapError(err, "ParentID")
		return
	}
	// write "error"
	err = en.Append(0xa5, 0x65, 0x72, 0x72, 0x6f, 0x72)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.Error)
	if err != nil {
		err = msgp.WrapError(err, "Error")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Span) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(12)
	var zb0001Mask uint16 /* 12 bits */
	if z.Meta == nil {
		zb0001Len--
		zb0001Mask |= 0x40
	}
	if z.Metrics == nil {
		zb0001Len--
		zb0001Mask |= 0x80
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len == 0 {
		return
	}
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "service"
	o = append(o, 0xa7, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65)
	o = msgp.AppendString(o, z.Service)
	// string "resource"
	o = append(o, 0xa8, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65)
	o = msgp.AppendString(o, z.Resource)
	// string "type"
	o = append(o, 0xa4, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.Type)
	// string "start"
	o = append(o, 0xa5, 0x73, 0x74, 0x61, 0x72, 0x74)
	o = msgp.AppendInt64(o, z.Start)
	// string "duration"
	o = append(o, 0xa8, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendInt64(o, z.Duration)
	if (zb0001Mask & 0x40) == 0 { // if not empty
		// string "meta"
		o = append(o, 0xa4, 0x6d, 0x65, 0x74, 0x61)
		o = msgp.AppendMapHeader(o, uint32(len(z.Meta)))
		for za0001, za0002 := range z.Meta {
			o = msgp.AppendString(o, za0001)
			o = msgp.AppendString(o, za0002)
		}
	}
	if (zb0001Mask & 0x80) == 0 { // if not empty
		// string "metrics"
		o = append(o, 0xa7, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73)
		o = msgp.AppendMapHeader(o, uint32(len(z.Metrics)))
		for za0003, za0004 := range z.Metrics {
			o = msgp.AppendString(o, za0003)
			o = msgp.AppendFloat64(o, za0004)
		}
	}
	// string "span_id"
	o = append(o, 0xa7, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64)
	o = msgp.AppendUint64(o, z.SpanID)
	// string "trace_id"
	o = append(o, 0xa8, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64)
	o = msgp.AppendUint64(o, z.TraceID)
	// string "parent_id"
	o = append(o, 0xa9, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64)
	o = msgp.AppendUint64(o, z.ParentID)
	// string "error"
	o = append(o, 0xa5, 0x65, 0x72, 0x72, 0x6f, 0x72)
	o = msgp.AppendInt32(o, z.Error)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Span) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "service":
			z.Service, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Service")
				return
			}
		case "resource":
			z.Resource, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Resource")
				return
			}
		case "type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Type")
				return
			}
		case "start":
			z.Start, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Start")
				return
			}
		case "duration":
			z.Duration, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Duration")
				return
			}
		case "meta":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Meta")
				return
			}
			if z.Meta == nil {
				z.Meta = make(map[string]string, zb0002)
			} else if len(z.Meta) > 0 {
				for key := range z.Meta {
					delete(z.Meta, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 string
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Meta")
					return
				}
				za0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Meta", za0001)
					return
				}
				z.Meta[za0001] = za0002
			}
		case "metrics":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Metrics")
				return
			}
			if z.Metrics == nil {
				z.Metrics = make(map[string]float64, zb0003)
			} else if len(z.Metrics) > 0 {
				for key := range z.Metrics {
					delete(z.Metrics, key)
				}
			}
			for zb0003 > 0 {
				var za0003 string
				var za0004 float64
				zb0003--
				za0003, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Metrics")
					return
				}
				za0004, bts, err = msgp.ReadFloat64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Metrics", za0003)
					return
				}
				z.Metrics[za0003] = za0004
			}
		case "span_id":
			z.SpanID, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SpanID")
				return
			}
		case "trace_id":
			z.TraceID, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TraceID")
				return
			}
		case "parent_id":
			z.ParentID, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ParentID")
				return
			}
		case "error":
			z.Error, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Error")
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
func (z *Span) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 8 + msgp.StringPrefixSize + len(z.Service) + 9 + msgp.StringPrefixSize + len(z.Resource) + 5 + msgp.StringPrefixSize + len(z.Type) + 6 + msgp.Int64Size + 9 + msgp.Int64Size + 5 + msgp.MapHeaderSize
	if z.Meta != nil {
		for za0001, za0002 := range z.Meta {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.StringPrefixSize + len(za0002)
		}
	}
	s += 8 + msgp.MapHeaderSize
	if z.Metrics != nil {
		for za0003, za0004 := range z.Metrics {
			_ = za0004
			s += msgp.StringPrefixSize + len(za0003) + msgp.Float64Size
		}
	}
	s += 8 + msgp.Uint64Size + 9 + msgp.Uint64Size + 10 + msgp.Uint64Size + 6 + msgp.Int32Size
	return
}
