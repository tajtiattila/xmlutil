package xmlutil_test

import (
	"encoding/xml"

	"github.com/tajtiattila/xmlutil"
)

func buildNode(parent *xmlutil.Node, space, local string, opt ...nodeOption) *xmlutil.Node {
	n := &xmlutil.Node{
		Name: xml.Name{
			Space: space,
			Local: local,
		},
	}
	for _, o := range opt {
		o.set(n)
	}
	if parent != nil {
		parent.Child = append(parent.Child, n)
	}
	return n
}

type nodeOption interface {
	set(n *xmlutil.Node)
}

type attrOption xml.Attr

func attr(space, local, value string) nodeOption {
	return attrOption(xml.Attr{
		Name: xml.Name{
			Space: space,
			Local: local,
		},
		Value: value,
	})
}

func (o attrOption) set(n *xmlutil.Node) {
	n.Attr = append(n.Attr, xml.Attr(o))
}

func ns(prefix, uri string) nodeOption {
	return attr("xmlns", prefix, uri)
}

type valueOption string

func value(v string) nodeOption {
	return valueOption(v)
}

func (o valueOption) set(n *xmlutil.Node) {
	n.Value = string(o)
}
