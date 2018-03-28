package xmlutil

import "encoding/xml"

// translateMap maps namespace URIs to their prefixes.
type translateMap map[string]string

func (tr translateMap) translate(n xml.Name) xml.Name {
	if n.Space == "" {
		return n
	}

	if n.Space == "xmlns" {
		return xml.Name{
			Local: n.Space + ":" + n.Local,
		}
	}

	if prefix, ok := tr[n.Space]; ok {
		return xml.Name{
			Local: prefix + ":" + n.Local,
		}
	}

	return n
}

// startElement returns the translated start element for n.
func (n *Node) startElement() xml.StartElement {
	if n.tr == nil {
		return xml.StartElement{
			Name: n.Name,
			Attr: n.Attr,
		}
	}

	var start xml.StartElement
	start.Name = n.tr.translate(n.Name)
	start.Attr = make([]xml.Attr, len(n.Attr))
	for i, a := range n.Attr {
		start.Attr[i] = xml.Attr{
			Name:  n.tr.translate(a.Name),
			Value: a.Value,
		}
	}
	return start
}

func (n *Node) translate(parent translateMap) {
	var m translateMap
	for _, a := range n.Attr {
		if a.Name.Space == "xmlns" {
			if m == nil {
				m = make(translateMap)
				for k, v := range parent {
					m[k] = v
				}
			}
			m[a.Value] = a.Name.Local
		}
	}

	// no new namespace, use parent
	if m == nil {
		m = parent
	}

	n.tr = m

	for _, child := range n.Child {
		child.translate(m)
	}
}
