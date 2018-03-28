package xmlutil

import (
	"encoding/xml"
	"strings"
)

type Node struct {
	XMLName xml.Name   // node name and namespace
	Attr    []xml.Attr // captures all unbound attributes and XMP qualifiers
	Value   string     // node char data
	Child   []*Node    // child nodes

	tr translateMap
}

func (n *Node) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if n.XMLName.Local == "" {
		return nil
	}

	start = n.startElement()
	return e.EncodeElement(struct {
		Data  string `xml:",chardata"`
		Child []*Node
	}{
		Data:  n.Value,
		Child: n.Child,
	}, start)
}

func (n *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	n.XMLName = start.Name
	n.Attr = start.Attr

	for {
		t, err := d.Token()
		if err != nil {
			return err
		}
		switch t := t.(type) {
		case xml.CharData:
			n.Value = strings.TrimSpace(string(t))
		case xml.StartElement:
			x := new(Node)
			x.UnmarshalXML(d, t)
			n.Child = append(n.Child, x)
		case xml.EndElement:
			return nil
		}
	}
}
