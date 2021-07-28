package linkedlist

func (nodelist1 *NodeList) Mergetwolinkedlists(nodelist2 *NodeList) {
	if nodelist1.Head == nil {
		nodelist1.Head = nodelist2.Head
		Logger.Info("nodelist1.Head is null in MergeTwoLinkedLists method")
	} else {
		head1 := nodelist1.Head
		for head1.nextnode != nil {
			head1 = head1.nextnode
		}
		head1.nextnode = nodelist2.Head
	}

}
