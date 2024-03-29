// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package state

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

func easyjson48287f42DecodeGithubComScribbleRsScribbleRsInternalState(in *jlexer.Lexer, out *PageStats) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "activeLobbyCount":
			out.ActiveLobbyCount = int(in.Int())
		case "playersCount":
			out.PlayersCount = uint64(in.Uint64())
		case "occupiedPlayerSlotCount":
			out.OccupiedPlayerSlotCount = uint64(in.Uint64())
		case "connectedPlayersCount":
			out.ConnectedPlayersCount = uint64(in.Uint64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson48287f42EncodeGithubComScribbleRsScribbleRsInternalState(out *jwriter.Writer, in PageStats) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"activeLobbyCount\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ActiveLobbyCount))
	}
	{
		const prefix string = ",\"playersCount\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.PlayersCount))
	}
	{
		const prefix string = ",\"occupiedPlayerSlotCount\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.OccupiedPlayerSlotCount))
	}
	{
		const prefix string = ",\"connectedPlayersCount\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ConnectedPlayersCount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PageStats) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson48287f42EncodeGithubComScribbleRsScribbleRsInternalState(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PageStats) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson48287f42EncodeGithubComScribbleRsScribbleRsInternalState(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PageStats) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson48287f42DecodeGithubComScribbleRsScribbleRsInternalState(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PageStats) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson48287f42DecodeGithubComScribbleRsScribbleRsInternalState(l, v)
}
