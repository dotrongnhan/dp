package prototype

type InterfaceNode interface {
	Clone() InterfaceNode
	Print(s string)
}
