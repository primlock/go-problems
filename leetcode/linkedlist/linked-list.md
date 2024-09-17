# Linked Lists in Go

Linked lists are nothing more than a bunch of nodes containing data and a pointer to access the next node. Unlike arrays, the nodes do not need to be stored sequentially in memory. They can exist in multiple blocks of memory and be accessed with the pointer address stored in a node. There are 3 types of linked lists:

* Singly Linked Lists
* Doubly Linked Lists
* Circular Linked Lists

### Singly Linked List (SLL)

This type of list contains the data and only one pointer to its next node. You can iterate forward, not in reverse. You always need to keep a reference to the head node in a SLL, losing it's address means losing access to your linked list.

The last node in an SLL is `nil` representing no more nodes after the previous node.

To create a node in our SLL that's data is an int, we can use a struct.
```go
type Node struct {
    val  int
    next *Node
}
```

To hold a collection of nodes, we can define a linked list as a struct to hold the head pointer.
```go
type LinkedList struct {
    head *Node
}
```

### Adding Nodes

Nodes can be added to a linked list quite easily. To add a new node to the tail of a linked list, iterate over the list until you arrive at the last node. You can be sure that you're at the last node when the next pointer is `nil`.

If there is currently no head of the linked list, assign the list head to the new node.

```go
newNode := &Node{
    val:  val,
    next: nil,
}

if list.head == nil {
    list.head = newNode
}
```

If the list head exists, iterate to the end of the list where the following node will be `nil` and set its next node equal to the new node.

```go
newNode := &Node{
    val:  val,
    next: nil,
}

current := list.head
for current.next != nil {
    current = current.next
}

current.next = newNode
```

### Removing Nodes

Removing nodes is as simple as reorganizing the next pointers. To remove a node from a list, you need the help of an additional pointer that trails behind the current pointer. When you're at the node to be removed, the previous pointer's next pointer is assigned to the current pointer's next pointer.

```go
var prev *Node
current := list.head

for current != nil {
    if current == targetNode {
        prev.next = current.next
        return nil
    }

    prev = current
    current = current.next
}
```

If the target node is the head node then we can reassign the head of the linked list to be the next node.

```go
if targetNode == list.head {
    list.head = list.head.next
}
```

