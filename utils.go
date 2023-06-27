package utils

type SymbolType int8
type Direct int8

const (
	NotFound SymbolType = iota
	Circular SymbolType = iota
	Quadric  SymbolType = iota
	Figuric  SymbolType = iota
)

const (
	Regular Direct = iota
	Close   Direct = iota
	Open    Direct = iota
)

type Symbol struct {
	symbolType SymbolType
	direct     Direct
}

func getSymbolType(r rune) SymbolType {
	if r == '(' || r == ')' {
		return Circular
	}
	if r == '[' || r == ']' {
		return Quadric
	}
	if r == '{' || r == '}' {
		return Figuric
	}
	return NotFound
}

func getDirect(r rune) Direct {
	if r == '(' || r == '[' || r == '{' {
		return Open
	}
	if r == ')' || r == ']' || r == '}' {
		return Close
	}
	return Regular
}

func toSymbol(r rune) Symbol {

	sym := Symbol{
		symbolType: getSymbolType(r),
		direct:     getDirect(r),
	}
	return sym
}

func isClosing(ns Symbol, ps Symbol) bool {
	return ns.symbolType == ps.symbolType &&
		ns.direct == Close &&
		ps.direct == Open
}

func IsValid(s string) bool {
	var stack []Symbol

	for _, r := range s {
		sym := toSymbol(r)
		if len(stack) > 0 {
			n := len(stack) - 1 // Top element
			prev := stack[n]
			if sym.direct == Close {
				if isClosing(sym, prev) {
					stack = stack[:n]
					continue
				} else {
					return false
				}
			}
		} else {
			if sym.direct == Close {
				return false
			}
		}
		if sym.direct == Open {
			stack = append(stack, sym)
		}
	}

	return len(stack) == 0
}
