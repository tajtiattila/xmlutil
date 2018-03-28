package xmlutil

import "encoding/xml"

func (n *Node) Translate() {
	ctx := new(context)
	n.translate(ctx)
}

func (n *Node) translate(ctx *context) {
	top := len(ctx.ns)

	for _, a := range n.Attr {
		if a.Name.Space == "xmlns" {
			ctx.addURIPrefix(a.Value, a.Name.Local)
		}
	}

	n.XMLName = ctx.translate(n.XMLName)
	for i := range n.Attr {
		a := &n.Attr[i]
		a.Name = ctx.translate(a.Name)
	}
	for _, child := range n.Child {
		child.translate(ctx)
	}

	ctx.pop(top)
}

type context struct {
	ns []namespace

	uriPrefix map[string]string
}

type namespace struct {
	prefix string
	uri    string
}

func (ctx *context) translate(n xml.Name) xml.Name {
	if n.Space == "" {
		return n
	}

	if n.Space == "xmlns" {
		return xml.Name{
			Local: n.Space + ":" + n.Local,
		}
	}

	if ctx.uriPrefix == nil {
		ctx.uriPrefix = make(map[string]string)
		for _, ns := range ctx.ns {
			ctx.uriPrefix[ns.uri] = ns.prefix
		}
	}

	if p, ok := ctx.uriPrefix[n.Space]; ok {
		return xml.Name{
			Local: p + ":" + n.Local,
		}
	}

	return n
}

func (ctx *context) addURIPrefix(uri, prefix string) {
	ctx.ns = append(ctx.ns, namespace{
		prefix: prefix,
		uri:    uri,
	})
	if ctx.uriPrefix != nil {
		ctx.uriPrefix[uri] = prefix
	}
}

func (ctx *context) pop(n int) {
	if n == len(ctx.ns) {
		return
	}
	ctx.ns = ctx.ns[:n]
	ctx.uriPrefix = nil
}

/*
type stackRef struct {
	enc *xml.Encoder
	ns []namespace
	top int // len(ns) on getStack call
}

func getStack(e *xml.Encoder, n *Node) (stackRef, xml.StartElement) {
	stackMap.mu.Lock()
	defer stackMap.mu.Unlock()

	ns := stackMap.m[e]

	stk := stackRef{
		enc: e,
		top: len(ns),
	}
	for _, a := range n.Attr {
		if a.Name.Space == "xmlns" {
			ns = append(ns, namespace{
				prefix: a.Name.Local,
				uri: a.Value,
			})

		}
	}

	stk.ns = ns

	if len(ns) != stk.top {
		if stackMap.m == nil {
			stackMap.m = make(map[*xml.Encoder][]namespace)
		}
		stackMap.m[e] = ns
	}
}

func (stk stackRef) release() {
	stackMap.mu.Lock()
	defer stackMap.mu.Unlock()

	if stk.top == 0 {
		delete(stackMap.m, stk.enc)
		return
	}

	stackMap.m[stk.enc] = stk.ns[:stk.top]
}

var stackMap struct {
	mu sync.Mutex
	m map[*xml.Encoder][]namespace
}
*/
