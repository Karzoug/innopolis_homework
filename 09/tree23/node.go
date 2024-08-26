package tree23

import "golang.org/x/exp/constraints"

type node[T constraints.Ordered] struct {
	parent *node[T]    // родитель узла
	sons   [4]*node[T] // сыновья узла
	keys   [3]T        // ключи узла
	length int         // количество ключей
}

func (n *node[T]) Insert(x T) *node[T] {
	if n == nil {
		return &node[T]{
			parent: nil,
			keys:   [3]T{x},
			length: 1,
		}
	}

	if n.isLeaf() {
		n.insertToNode(x)
	} else if x <= n.keys[0] {
		n.sons[0].Insert(x)
	} else if (n.length == 1) || ((n.length == 2) && x <= n.keys[1]) {
		n.sons[1].Insert(x)
	} else {
		n.sons[2].Insert(x)
	}

	return n.split()
}

func (n *node[T]) Search(x T) *node[T] {
	if n == nil {
		return nil
	}

	if n.findKey(x) {
		return n
	} else if x < n.keys[0] {
		return n.sons[0].Search(x)
	} else if ((n.length == 2) && (x < n.keys[1])) || (n.length == 1) {
		return n.sons[1].Search(x)
	} else if n.length == 2 {
		return n.sons[2].Search(x)
	}

	return nil
}

func (n *node[T]) Remove(x T) *node[T] {
	item := n.Search(x)

	if item == nil {
		return n
	}

	var minNode *node[T]
	if item.keys[0] == x {
		minNode = item.sons[1].searchMin()
	} else {
		minNode = item.sons[2].searchMin()
	}

	if minNode != nil {
		var z T
		if x == item.keys[0] {
			z = item.keys[0]
		} else {
			z = item.keys[1]
		}
		z, minNode.keys[0] = minNode.keys[0], z
		item = minNode
	}

	item.removeFromNode(x)
	return item.fix() // Вызываем функцию для восстановления свойств дерева.
}

func (n *node[T]) isLeaf() bool {
	return n.sons[0] == nil && n.sons[1] == nil && n.sons[2] == nil && n.sons[3] == nil
}

func (n *node[T]) insertToNode(x T) {
	n.keys[n.length] = x
	n.length++
	n.sort()
}

func (n *node[T]) split() *node[T] {
	if n.length < 3 {
		return n
	}

	x := &node[T]{
		parent: n.parent,
		sons:   [4]*node[T]{n.sons[0], n.sons[1]},
		keys:   [3]T{n.keys[0]},
		length: 1,
	}
	y := &node[T]{
		parent: n.parent,
		sons:   [4]*node[T]{n.sons[2], n.sons[3]},
		keys:   [3]T{n.keys[2]},
		length: 1,
	}
	if x.sons[0] != nil {
		x.sons[0].parent = x
	}
	if x.sons[1] != nil {
		x.sons[1].parent = x
	}
	if y.sons[0] != nil {
		y.sons[0].parent = y
	}
	if y.sons[1] != nil {
		y.sons[1].parent = y
	}

	if n.parent != nil {
		n.parent.insertToNode(n.keys[1])

		if n.parent.sons[0] == n {
			n.parent.sons[0] = nil
		} else if n.parent.sons[1] == n {
			n.parent.sons[1] = nil
		} else if n.parent.sons[2] == n {
			n.parent.sons[2] = nil
		}

		if n.parent.sons[0] == nil {
			n.parent.sons[3] = n.parent.sons[2]
			n.parent.sons[2] = n.parent.sons[1]
			n.parent.sons[1] = y
			n.parent.sons[0] = x
		} else if n.parent.sons[1] == nil {
			n.parent.sons[3] = n.parent.sons[2]
			n.parent.sons[2] = y
			n.parent.sons[1] = x
		} else {
			n.parent.sons[3] = y
			n.parent.sons[2] = x
		}

		return n.parent
	} else {
		x.parent = n
		y.parent = n

		n.sons = [4]*node[T]{x, y, nil, nil}
		n.keys[0] = n.keys[1]
		n.length = 1

		return n
	}
}

func (n *node[T]) sort() {
	switch n.length {
	case 1:
		return
	case 2:
		if n.keys[0] > n.keys[1] {
			n.keys[0], n.keys[1] = n.keys[1], n.keys[0]
		}
	case 3:
		if n.keys[0] > n.keys[1] {
			n.keys[0], n.keys[1] = n.keys[1], n.keys[0]
		}
		if n.keys[0] > n.keys[2] {
			n.keys[0], n.keys[2] = n.keys[2], n.keys[0]
		}
		if n.keys[1] > n.keys[2] {
			n.keys[1], n.keys[2] = n.keys[2], n.keys[1]
		}
	}
}

func (n *node[T]) findKey(x T) bool {
	for i := range n.length {
		if n.keys[i] == x {
			return true
		}
	}
	return false
}

func (n *node[T]) searchMin() *node[T] {
	if n == nil {
		return n
	}
	if n.sons[0] == nil {
		return n
	} 
	
	return n.sons[0].searchMin()
}

func (n *node[T]) removeFromNode(x T) {
	if n.length >= 1 && n.keys[0] == x {
		n.keys[0] = n.keys[1]
		n.keys[1] = n.keys[2]
		n.length--
	} else if n.length == 2 && n.keys[1] == x {
		n.keys[1] = n.keys[2]
		n.length--
	}
}

func (n *node[T]) fix() *node[T] {
	if n.length == 0 && n.parent == nil { // Случай 0, когда удаляем единственный ключ в дереве
		n = nil
		return nil
	}
	if n.length != 0 { // Случай 1, когда вершина, в которой удалили ключ, имела два ключа
		if n.parent != nil {
			return n.parent.fix()
		} else {
			return n
		}
	}

	parent := n.parent
	// Случай 2, когда достаточно перераспределить ключи в дереве
	if parent.sons[0].length == 2 || parent.sons[1].length == 2 || parent.length == 2 {
		n = n.redistribute()
	} else if parent.length == 2 && parent.sons[2].length == 2 { // Аналогично
		n = n.redistribute()
	} else {
		n = n.merge() // Случай 3, когда нужно произвести склеивание и пройтись вверх по дереву как минимум на еще одну вершину
	}

	return n.fix()
}

func (n *node[T]) redistribute() *node[T] {
	parent := n.parent
	first := parent.sons[0]
	second := parent.sons[1]
	third := parent.sons[2]

	if (parent.length == 2) && (first.length < 2) && (second.length < 2) && (third.length < 2) {
		if first == n {
			parent.sons[0] = parent.sons[1]
			parent.sons[1] = parent.sons[2]
			parent.sons[2] = nil
			parent.sons[0].insertToNode(parent.keys[0])
			parent.sons[0].sons[2] = parent.sons[0].sons[1]
			parent.sons[0].sons[1] = parent.sons[0].sons[0]

			if n.sons[0] != nil {
				parent.sons[0].sons[0] = n.sons[0]
			} else if n.sons[1] != nil {
				parent.sons[0].sons[0] = n.sons[1]
			}

			if parent.sons[0].sons[0] != nil {
				parent.sons[0].sons[0].parent = parent.sons[0]
			}

			parent.removeFromNode(parent.keys[0])
		} else if second == n {
			first.insertToNode(parent.keys[0])
			parent.removeFromNode(parent.keys[0])
			if n.sons[0] != nil {
				first.sons[2] = n.sons[0]
			} else if n.sons[1] != nil {
				first.sons[2] = n.sons[1]
			}

			if first.sons[2] != nil {
				first.sons[2].parent = first
			}

			parent.sons[1] = parent.sons[2]
			parent.sons[2] = nil

		} else if third == n {
			second.insertToNode(parent.keys[1])
			parent.sons[2] = nil
			parent.removeFromNode(parent.keys[1])
			if n.sons[0] != nil {
				second.sons[2] = n.sons[0]
			} else if n.sons[1] != nil {
				second.sons[2] = n.sons[1]
			}

			if second.sons[2] != nil {
				second.sons[2].parent = second
			}
		}
	} else if (parent.length == 2) && ((first.length == 2) || (second.length == 2) || (third.length == 2)) {
		if third == n {
			if n.sons[0] != nil {
				n.sons[1] = n.sons[0]
				n.sons[0] = nil
			}

			n.insertToNode(parent.keys[1])
			if second.length == 2 {
				parent.keys[1] = second.keys[1]
				second.removeFromNode(second.keys[1])
				n.sons[0] = second.sons[2]
				second.sons[2] = nil
				if n.sons[0] != nil {
					n.sons[0].parent = n
				}
			} else if first.length == 2 {
				parent.keys[1] = second.keys[0]
				n.sons[0] = second.sons[1]
				second.sons[1] = second.sons[0]
				if n.sons[0] != nil {
					n.sons[0].parent = n
				}

				second.keys[0] = parent.keys[0]
				parent.keys[0] = first.keys[1]
				first.removeFromNode(first.keys[1])
				second.sons[0] = first.sons[2]
				if second.sons[0] != nil {
					second.sons[0].parent = second
				}
				first.sons[2] = nil
			}
		} else if second == n {
			if third.length == 2 {
				if n.sons[0] == nil {
					n.sons[0] = n.sons[1]
					n.sons[1] = nil
				}
				second.insertToNode(parent.keys[1])
				parent.keys[1] = third.keys[0]
				third.removeFromNode(third.keys[0])
				second.sons[1] = third.sons[0]
				if second.sons[1] != nil {
					second.sons[1].parent = second
				}
				third.sons[0] = third.sons[1]
				third.sons[1] = third.sons[2]
				third.sons[2] = nil
			} else if first.length == 2 {
				if n.sons[1] == nil {
					n.sons[1] = n.sons[0]
					n.sons[0] = nil
				}
				second.insertToNode(parent.keys[0])
				parent.keys[0] = first.keys[1]
				first.removeFromNode(first.keys[1])
				second.sons[0] = first.sons[2]
				if second.sons[0] != nil {
					second.sons[0].parent = second
				}
				first.sons[2] = nil
			}
		} else if first == n {
			if n.sons[0] == nil {
				n.sons[0] = n.sons[1]
				n.sons[1] = nil
			}
			first.insertToNode(parent.keys[0])
			if second.length == 2 {
				parent.keys[0] = second.keys[0]
				second.removeFromNode(second.keys[0])
				first.sons[1] = second.sons[0]
				if first.sons[1] != nil {
					first.sons[1].parent = first
				}
				second.sons[0] = second.sons[1]
				second.sons[1] = second.sons[2]
				second.sons[2] = nil
			} else if third.length == 2 {
				parent.keys[0] = second.keys[0]
				second.keys[0] = parent.keys[1]
				parent.keys[1] = third.keys[0]
				third.removeFromNode(third.keys[0])
				first.sons[1] = second.sons[0]
				if first.sons[1] != nil {
					first.sons[1].parent = first
				}
				second.sons[0] = second.sons[1]
				second.sons[1] = third.sons[0]
				if second.sons[1] != nil {
					second.sons[1].parent = second
				}
				third.sons[0] = third.sons[1]
				third.sons[1] = third.sons[2]
				third.sons[2] = nil
			}
		}
	} else if parent.length == 1 {
		n.insertToNode(parent.keys[0])

		if first == n && second.length == 2 {
			parent.keys[0] = second.keys[0]
			second.removeFromNode(second.keys[0])

			if n.sons[0] == nil {
				n.sons[0] = n.sons[1]
			}

			n.sons[1] = second.sons[0]
			second.sons[0] = second.sons[1]
			second.sons[1] = second.sons[2]
			second.sons[2] = nil
			if n.sons[1] != nil {
				n.sons[1].parent = n
			}
		} else if second == n && first.length == 2 {
			parent.keys[0] = first.keys[1]
			first.removeFromNode(first.keys[1])

			if n.sons[1] == nil {
				n.sons[1] = n.sons[0]
			}

			n.sons[0] = first.sons[2]
			first.sons[2] = nil
			if n.sons[0] != nil {
				n.sons[0].parent = n
			}
		}
	}
	return parent
}

func (n *node[T]) merge() *node[T] {
	parent := n.parent

	if parent.sons[0] == n {
		parent.sons[1].insertToNode(parent.keys[0])
		parent.sons[1].sons[2] = parent.sons[1].sons[1]
		parent.sons[1].sons[1] = parent.sons[1].sons[0]

		if n.sons[0] != nil {
			parent.sons[1].sons[0] = n.sons[0]
		} else if n.sons[1] != nil {
			parent.sons[1].sons[0] = n.sons[1]
		}

		if parent.sons[1].sons[0] != nil {
			parent.sons[1].sons[0].parent = parent.sons[1]
		}

		parent.removeFromNode(parent.keys[0])
		parent.sons[0] = nil
	} else if parent.sons[1] == n {
		parent.sons[0].insertToNode(parent.keys[0])

		if n.sons[0] != nil {
			parent.sons[0].sons[2] = n.sons[0]
		} else if n.sons[1] != nil {
			parent.sons[0].sons[2] = n.sons[1]
		}

		if parent.sons[0].sons[2] != nil {
			parent.sons[0].sons[2].parent = parent.sons[0]
		}

		parent.removeFromNode(parent.keys[0])
		parent.sons[1] = nil
	}

	if parent.parent == nil {
		var tmp *node[T]
		if parent.sons[0] != nil {
			tmp = parent.sons[0]
		} else {
			tmp = parent.sons[1]
		}
		tmp.parent = nil
		return tmp
	}
	return parent
}
