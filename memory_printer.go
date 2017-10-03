package todolist

type MemoryPrinter struct {
	Groups *GroupedTodos
}

func (m *MemoryPrinter) Print(groupedTodos *GroupedTodos, printNotes bool) {
	f.Groups = groupedTodos
}
