// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package cupti

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson8ffcabdbDecodeGithubComRaiProjectGoCupti(in *jlexer.Lexer, out *events) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(events, 0, 1)
			} else {
				*out = events{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Event
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8ffcabdbEncodeGithubComRaiProjectGoCupti(out *jwriter.Writer, in events) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v events) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8ffcabdbEncodeGithubComRaiProjectGoCupti(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v events) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8ffcabdbEncodeGithubComRaiProjectGoCupti(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *events) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8ffcabdbDecodeGithubComRaiProjectGoCupti(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *events) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8ffcabdbDecodeGithubComRaiProjectGoCupti(l, v)
}
func easyjson8ffcabdbDecodeGithubComRaiProjectGoCupti1(in *jlexer.Lexer, out *Event) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "event":
			out.Event = int(in.Int())
		case "id":
			out.ID = int(in.Int())
		case "domain_id":
			out.DomainID = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "short_description":
			out.ShortDescription = string(in.String())
		case "long_description":
			out.LongDescription = string(in.String())
		case "category":
			out.Category = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8ffcabdbEncodeGithubComRaiProjectGoCupti1(out *jwriter.Writer, in Event) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Event != 0 {
		const prefix string = ",\"event\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Event))
	}
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ID))
	}
	if in.DomainID != 0 {
		const prefix string = ",\"domain_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.DomainID))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.ShortDescription != "" {
		const prefix string = ",\"short_description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ShortDescription))
	}
	if in.LongDescription != "" {
		const prefix string = ",\"long_description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LongDescription))
	}
	if in.Category != "" {
		const prefix string = ",\"category\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Category))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Event) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8ffcabdbEncodeGithubComRaiProjectGoCupti1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8ffcabdbEncodeGithubComRaiProjectGoCupti1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Event) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8ffcabdbDecodeGithubComRaiProjectGoCupti1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8ffcabdbDecodeGithubComRaiProjectGoCupti1(l, v)
}