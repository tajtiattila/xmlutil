// Package xmlutil provides the Document type that uses
// namespace prefixes when marshaling XML.
package xmlutil

import "encoding/xml"

// Document is an XML document that keeps namespaces prefixes when marshaling.
type Document struct {
	*Node // document root
}

// MarshalXML marshals d in XML format.
func (d Document) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	d.Node.translate(nil)
	return d.Node.MarshalXML(enc, start)
}

// UnmarshalXML unmarshals d from XML format.
func (d *Document) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	d.Node = new(Node)
	return d.Node.UnmarshalXML(dec, start)
}
