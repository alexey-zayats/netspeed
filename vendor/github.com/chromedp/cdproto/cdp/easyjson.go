// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package cdp

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

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp(in *jlexer.Lexer, out *RGBA) {
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
		case "r":
			out.R = int64(in.Int64())
		case "g":
			out.G = int64(in.Int64())
		case "b":
			out.B = int64(in.Int64())
		case "a":
			out.A = float64(in.Float64())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp(out *jwriter.Writer, in RGBA) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"r\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.R))
	}
	{
		const prefix string = ",\"g\":"
		out.RawString(prefix)
		out.Int64(int64(in.G))
	}
	{
		const prefix string = ",\"b\":"
		out.RawString(prefix)
		out.Int64(int64(in.B))
	}
	{
		const prefix string = ",\"a\":"
		out.RawString(prefix)
		out.Float64(float64(in.A))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RGBA) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RGBA) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RGBA) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RGBA) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp1(in *jlexer.Lexer, out *OriginTrialTokenWithStatus) {
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
		case "rawTokenText":
			out.RawTokenText = string(in.String())
		case "parsedToken":
			if in.IsNull() {
				in.Skip()
				out.ParsedToken = nil
			} else {
				if out.ParsedToken == nil {
					out.ParsedToken = new(OriginTrialToken)
				}
				(*out.ParsedToken).UnmarshalEasyJSON(in)
			}
		case "status":
			(out.Status).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp1(out *jwriter.Writer, in OriginTrialTokenWithStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"rawTokenText\":"
		out.RawString(prefix[1:])
		out.String(string(in.RawTokenText))
	}
	if in.ParsedToken != nil {
		const prefix string = ",\"parsedToken\":"
		out.RawString(prefix)
		(*in.ParsedToken).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		(in.Status).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OriginTrialTokenWithStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OriginTrialTokenWithStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OriginTrialTokenWithStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OriginTrialTokenWithStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp2(in *jlexer.Lexer, out *OriginTrialToken) {
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
		case "origin":
			out.Origin = string(in.String())
		case "matchSubDomains":
			out.MatchSubDomains = bool(in.Bool())
		case "trialName":
			out.TrialName = string(in.String())
		case "expiryTime":
			if in.IsNull() {
				in.Skip()
				out.ExpiryTime = nil
			} else {
				if out.ExpiryTime == nil {
					out.ExpiryTime = new(TimeSinceEpoch)
				}
				(*out.ExpiryTime).UnmarshalEasyJSON(in)
			}
		case "isThirdParty":
			out.IsThirdParty = bool(in.Bool())
		case "usageRestriction":
			(out.UsageRestriction).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp2(out *jwriter.Writer, in OriginTrialToken) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"origin\":"
		out.RawString(prefix[1:])
		out.String(string(in.Origin))
	}
	{
		const prefix string = ",\"matchSubDomains\":"
		out.RawString(prefix)
		out.Bool(bool(in.MatchSubDomains))
	}
	{
		const prefix string = ",\"trialName\":"
		out.RawString(prefix)
		out.String(string(in.TrialName))
	}
	{
		const prefix string = ",\"expiryTime\":"
		out.RawString(prefix)
		if in.ExpiryTime == nil {
			out.RawString("null")
		} else {
			(*in.ExpiryTime).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"isThirdParty\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsThirdParty))
	}
	{
		const prefix string = ",\"usageRestriction\":"
		out.RawString(prefix)
		(in.UsageRestriction).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OriginTrialToken) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OriginTrialToken) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OriginTrialToken) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OriginTrialToken) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp3(in *jlexer.Lexer, out *OriginTrial) {
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
		case "trialName":
			out.TrialName = string(in.String())
		case "status":
			(out.Status).UnmarshalEasyJSON(in)
		case "tokensWithStatus":
			if in.IsNull() {
				in.Skip()
				out.TokensWithStatus = nil
			} else {
				in.Delim('[')
				if out.TokensWithStatus == nil {
					if !in.IsDelim(']') {
						out.TokensWithStatus = make([]*OriginTrialTokenWithStatus, 0, 8)
					} else {
						out.TokensWithStatus = []*OriginTrialTokenWithStatus{}
					}
				} else {
					out.TokensWithStatus = (out.TokensWithStatus)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *OriginTrialTokenWithStatus
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(OriginTrialTokenWithStatus)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.TokensWithStatus = append(out.TokensWithStatus, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp3(out *jwriter.Writer, in OriginTrial) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"trialName\":"
		out.RawString(prefix[1:])
		out.String(string(in.TrialName))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		(in.Status).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"tokensWithStatus\":"
		out.RawString(prefix)
		if in.TokensWithStatus == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.TokensWithStatus {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OriginTrial) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OriginTrial) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OriginTrial) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OriginTrial) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp3(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp4(in *jlexer.Lexer, out *Node) {
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
		case "nodeId":
			(out.NodeID).UnmarshalEasyJSON(in)
		case "parentId":
			(out.ParentID).UnmarshalEasyJSON(in)
		case "backendNodeId":
			(out.BackendNodeID).UnmarshalEasyJSON(in)
		case "nodeType":
			(out.NodeType).UnmarshalEasyJSON(in)
		case "nodeName":
			out.NodeName = string(in.String())
		case "localName":
			out.LocalName = string(in.String())
		case "nodeValue":
			out.NodeValue = string(in.String())
		case "childNodeCount":
			out.ChildNodeCount = int64(in.Int64())
		case "children":
			if in.IsNull() {
				in.Skip()
				out.Children = nil
			} else {
				in.Delim('[')
				if out.Children == nil {
					if !in.IsDelim(']') {
						out.Children = make([]*Node, 0, 8)
					} else {
						out.Children = []*Node{}
					}
				} else {
					out.Children = (out.Children)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *Node
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(Node)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Children = append(out.Children, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "attributes":
			if in.IsNull() {
				in.Skip()
				out.Attributes = nil
			} else {
				in.Delim('[')
				if out.Attributes == nil {
					if !in.IsDelim(']') {
						out.Attributes = make([]string, 0, 4)
					} else {
						out.Attributes = []string{}
					}
				} else {
					out.Attributes = (out.Attributes)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.Attributes = append(out.Attributes, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "documentURL":
			out.DocumentURL = string(in.String())
		case "baseURL":
			out.BaseURL = string(in.String())
		case "publicId":
			out.PublicID = string(in.String())
		case "systemId":
			out.SystemID = string(in.String())
		case "internalSubset":
			out.InternalSubset = string(in.String())
		case "xmlVersion":
			out.XMLVersion = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "value":
			out.Value = string(in.String())
		case "pseudoType":
			(out.PseudoType).UnmarshalEasyJSON(in)
		case "shadowRootType":
			(out.ShadowRootType).UnmarshalEasyJSON(in)
		case "frameId":
			(out.FrameID).UnmarshalEasyJSON(in)
		case "contentDocument":
			if in.IsNull() {
				in.Skip()
				out.ContentDocument = nil
			} else {
				if out.ContentDocument == nil {
					out.ContentDocument = new(Node)
				}
				(*out.ContentDocument).UnmarshalEasyJSON(in)
			}
		case "shadowRoots":
			if in.IsNull() {
				in.Skip()
				out.ShadowRoots = nil
			} else {
				in.Delim('[')
				if out.ShadowRoots == nil {
					if !in.IsDelim(']') {
						out.ShadowRoots = make([]*Node, 0, 8)
					} else {
						out.ShadowRoots = []*Node{}
					}
				} else {
					out.ShadowRoots = (out.ShadowRoots)[:0]
				}
				for !in.IsDelim(']') {
					var v6 *Node
					if in.IsNull() {
						in.Skip()
						v6 = nil
					} else {
						if v6 == nil {
							v6 = new(Node)
						}
						(*v6).UnmarshalEasyJSON(in)
					}
					out.ShadowRoots = append(out.ShadowRoots, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "templateContent":
			if in.IsNull() {
				in.Skip()
				out.TemplateContent = nil
			} else {
				if out.TemplateContent == nil {
					out.TemplateContent = new(Node)
				}
				(*out.TemplateContent).UnmarshalEasyJSON(in)
			}
		case "pseudoElements":
			if in.IsNull() {
				in.Skip()
				out.PseudoElements = nil
			} else {
				in.Delim('[')
				if out.PseudoElements == nil {
					if !in.IsDelim(']') {
						out.PseudoElements = make([]*Node, 0, 8)
					} else {
						out.PseudoElements = []*Node{}
					}
				} else {
					out.PseudoElements = (out.PseudoElements)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *Node
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(Node)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.PseudoElements = append(out.PseudoElements, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "distributedNodes":
			if in.IsNull() {
				in.Skip()
				out.DistributedNodes = nil
			} else {
				in.Delim('[')
				if out.DistributedNodes == nil {
					if !in.IsDelim(']') {
						out.DistributedNodes = make([]*BackendNode, 0, 8)
					} else {
						out.DistributedNodes = []*BackendNode{}
					}
				} else {
					out.DistributedNodes = (out.DistributedNodes)[:0]
				}
				for !in.IsDelim(']') {
					var v8 *BackendNode
					if in.IsNull() {
						in.Skip()
						v8 = nil
					} else {
						if v8 == nil {
							v8 = new(BackendNode)
						}
						(*v8).UnmarshalEasyJSON(in)
					}
					out.DistributedNodes = append(out.DistributedNodes, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "isSVG":
			out.IsSVG = bool(in.Bool())
		case "compatibilityMode":
			(out.CompatibilityMode).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp4(out *jwriter.Writer, in Node) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nodeId\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.NodeID))
	}
	if in.ParentID != 0 {
		const prefix string = ",\"parentId\":"
		out.RawString(prefix)
		out.Int64(int64(in.ParentID))
	}
	{
		const prefix string = ",\"backendNodeId\":"
		out.RawString(prefix)
		out.Int64(int64(in.BackendNodeID))
	}
	{
		const prefix string = ",\"nodeType\":"
		out.RawString(prefix)
		(in.NodeType).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"nodeName\":"
		out.RawString(prefix)
		out.String(string(in.NodeName))
	}
	{
		const prefix string = ",\"localName\":"
		out.RawString(prefix)
		out.String(string(in.LocalName))
	}
	{
		const prefix string = ",\"nodeValue\":"
		out.RawString(prefix)
		out.String(string(in.NodeValue))
	}
	if in.ChildNodeCount != 0 {
		const prefix string = ",\"childNodeCount\":"
		out.RawString(prefix)
		out.Int64(int64(in.ChildNodeCount))
	}
	if len(in.Children) != 0 {
		const prefix string = ",\"children\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v9, v10 := range in.Children {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil {
					out.RawString("null")
				} else {
					(*v10).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Attributes) != 0 {
		const prefix string = ",\"attributes\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v11, v12 := range in.Attributes {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	if in.DocumentURL != "" {
		const prefix string = ",\"documentURL\":"
		out.RawString(prefix)
		out.String(string(in.DocumentURL))
	}
	if in.BaseURL != "" {
		const prefix string = ",\"baseURL\":"
		out.RawString(prefix)
		out.String(string(in.BaseURL))
	}
	if in.PublicID != "" {
		const prefix string = ",\"publicId\":"
		out.RawString(prefix)
		out.String(string(in.PublicID))
	}
	if in.SystemID != "" {
		const prefix string = ",\"systemId\":"
		out.RawString(prefix)
		out.String(string(in.SystemID))
	}
	if in.InternalSubset != "" {
		const prefix string = ",\"internalSubset\":"
		out.RawString(prefix)
		out.String(string(in.InternalSubset))
	}
	if in.XMLVersion != "" {
		const prefix string = ",\"xmlVersion\":"
		out.RawString(prefix)
		out.String(string(in.XMLVersion))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	if in.Value != "" {
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.String(string(in.Value))
	}
	if in.PseudoType != "" {
		const prefix string = ",\"pseudoType\":"
		out.RawString(prefix)
		(in.PseudoType).MarshalEasyJSON(out)
	}
	if in.ShadowRootType != "" {
		const prefix string = ",\"shadowRootType\":"
		out.RawString(prefix)
		(in.ShadowRootType).MarshalEasyJSON(out)
	}
	if in.FrameID != "" {
		const prefix string = ",\"frameId\":"
		out.RawString(prefix)
		out.String(string(in.FrameID))
	}
	if in.ContentDocument != nil {
		const prefix string = ",\"contentDocument\":"
		out.RawString(prefix)
		(*in.ContentDocument).MarshalEasyJSON(out)
	}
	if len(in.ShadowRoots) != 0 {
		const prefix string = ",\"shadowRoots\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v13, v14 := range in.ShadowRoots {
				if v13 > 0 {
					out.RawByte(',')
				}
				if v14 == nil {
					out.RawString("null")
				} else {
					(*v14).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.TemplateContent != nil {
		const prefix string = ",\"templateContent\":"
		out.RawString(prefix)
		(*in.TemplateContent).MarshalEasyJSON(out)
	}
	if len(in.PseudoElements) != 0 {
		const prefix string = ",\"pseudoElements\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v15, v16 := range in.PseudoElements {
				if v15 > 0 {
					out.RawByte(',')
				}
				if v16 == nil {
					out.RawString("null")
				} else {
					(*v16).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.DistributedNodes) != 0 {
		const prefix string = ",\"distributedNodes\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v17, v18 := range in.DistributedNodes {
				if v17 > 0 {
					out.RawByte(',')
				}
				if v18 == nil {
					out.RawString("null")
				} else {
					(*v18).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.IsSVG {
		const prefix string = ",\"isSVG\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsSVG))
	}
	if in.CompatibilityMode != "" {
		const prefix string = ",\"compatibilityMode\":"
		out.RawString(prefix)
		(in.CompatibilityMode).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Node) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Node) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Node) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Node) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp4(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp5(in *jlexer.Lexer, out *Frame) {
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
		case "id":
			(out.ID).UnmarshalEasyJSON(in)
		case "parentId":
			(out.ParentID).UnmarshalEasyJSON(in)
		case "loaderId":
			out.LoaderID = LoaderID(in.String())
		case "name":
			out.Name = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "urlFragment":
			out.URLFragment = string(in.String())
		case "domainAndRegistry":
			out.DomainAndRegistry = string(in.String())
		case "securityOrigin":
			out.SecurityOrigin = string(in.String())
		case "mimeType":
			out.MimeType = string(in.String())
		case "unreachableUrl":
			out.UnreachableURL = string(in.String())
		case "adFrameStatus":
			if in.IsNull() {
				in.Skip()
				out.AdFrameStatus = nil
			} else {
				if out.AdFrameStatus == nil {
					out.AdFrameStatus = new(AdFrameStatus)
				}
				(*out.AdFrameStatus).UnmarshalEasyJSON(in)
			}
		case "secureContextType":
			(out.SecureContextType).UnmarshalEasyJSON(in)
		case "crossOriginIsolatedContextType":
			(out.CrossOriginIsolatedContextType).UnmarshalEasyJSON(in)
		case "gatedAPIFeatures":
			if in.IsNull() {
				in.Skip()
				out.GatedAPIFeatures = nil
			} else {
				in.Delim('[')
				if out.GatedAPIFeatures == nil {
					if !in.IsDelim(']') {
						out.GatedAPIFeatures = make([]GatedAPIFeatures, 0, 4)
					} else {
						out.GatedAPIFeatures = []GatedAPIFeatures{}
					}
				} else {
					out.GatedAPIFeatures = (out.GatedAPIFeatures)[:0]
				}
				for !in.IsDelim(']') {
					var v19 GatedAPIFeatures
					(v19).UnmarshalEasyJSON(in)
					out.GatedAPIFeatures = append(out.GatedAPIFeatures, v19)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp5(out *jwriter.Writer, in Frame) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.ParentID != "" {
		const prefix string = ",\"parentId\":"
		out.RawString(prefix)
		out.String(string(in.ParentID))
	}
	{
		const prefix string = ",\"loaderId\":"
		out.RawString(prefix)
		out.String(string(in.LoaderID))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	if in.URLFragment != "" {
		const prefix string = ",\"urlFragment\":"
		out.RawString(prefix)
		out.String(string(in.URLFragment))
	}
	{
		const prefix string = ",\"domainAndRegistry\":"
		out.RawString(prefix)
		out.String(string(in.DomainAndRegistry))
	}
	{
		const prefix string = ",\"securityOrigin\":"
		out.RawString(prefix)
		out.String(string(in.SecurityOrigin))
	}
	{
		const prefix string = ",\"mimeType\":"
		out.RawString(prefix)
		out.String(string(in.MimeType))
	}
	if in.UnreachableURL != "" {
		const prefix string = ",\"unreachableUrl\":"
		out.RawString(prefix)
		out.String(string(in.UnreachableURL))
	}
	if in.AdFrameStatus != nil {
		const prefix string = ",\"adFrameStatus\":"
		out.RawString(prefix)
		(*in.AdFrameStatus).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"secureContextType\":"
		out.RawString(prefix)
		(in.SecureContextType).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"crossOriginIsolatedContextType\":"
		out.RawString(prefix)
		(in.CrossOriginIsolatedContextType).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"gatedAPIFeatures\":"
		out.RawString(prefix)
		if in.GatedAPIFeatures == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.GatedAPIFeatures {
				if v20 > 0 {
					out.RawByte(',')
				}
				(v21).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Frame) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Frame) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Frame) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Frame) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp5(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp6(in *jlexer.Lexer, out *BackendNode) {
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
		case "nodeType":
			(out.NodeType).UnmarshalEasyJSON(in)
		case "nodeName":
			out.NodeName = string(in.String())
		case "backendNodeId":
			(out.BackendNodeID).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp6(out *jwriter.Writer, in BackendNode) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nodeType\":"
		out.RawString(prefix[1:])
		(in.NodeType).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"nodeName\":"
		out.RawString(prefix)
		out.String(string(in.NodeName))
	}
	{
		const prefix string = ",\"backendNodeId\":"
		out.RawString(prefix)
		out.Int64(int64(in.BackendNodeID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BackendNode) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BackendNode) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BackendNode) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BackendNode) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp6(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp7(in *jlexer.Lexer, out *AdFrameStatus) {
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
		case "adFrameType":
			(out.AdFrameType).UnmarshalEasyJSON(in)
		case "explanations":
			if in.IsNull() {
				in.Skip()
				out.Explanations = nil
			} else {
				in.Delim('[')
				if out.Explanations == nil {
					if !in.IsDelim(']') {
						out.Explanations = make([]AdFrameExplanation, 0, 4)
					} else {
						out.Explanations = []AdFrameExplanation{}
					}
				} else {
					out.Explanations = (out.Explanations)[:0]
				}
				for !in.IsDelim(']') {
					var v22 AdFrameExplanation
					(v22).UnmarshalEasyJSON(in)
					out.Explanations = append(out.Explanations, v22)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp7(out *jwriter.Writer, in AdFrameStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"adFrameType\":"
		out.RawString(prefix[1:])
		(in.AdFrameType).MarshalEasyJSON(out)
	}
	if len(in.Explanations) != 0 {
		const prefix string = ",\"explanations\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v23, v24 := range in.Explanations {
				if v23 > 0 {
					out.RawByte(',')
				}
				(v24).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AdFrameStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AdFrameStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCdp7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AdFrameStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AdFrameStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCdp7(l, v)
}
