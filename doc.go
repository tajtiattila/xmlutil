// Package xmlutil provides the Document type that uses
// namespace prefixes when marshaling XML.
package xmlutil

import "encoding/xml"

// Document is an XML document that keeps namespaces prefixes when marshaling.
type Document struct {
	Root *Node
}

// MarshalXML marshals d in XML format.
func (d Document) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	d.Root.translate(nil)
	return d.Root.MarshalXML(enc, start)
}

// UnmarshalXML unmarshals d from XML format.
func (d *Document) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	d.Root = new(Node)
	return d.Root.UnmarshalXML(dec, start)
}
