class Node:
    def __init__(self, data=None):
        self.data = data
        self.next = None
     
     
class LinkedList:
    def __init__(self):
        self.head = None
        
    def insert(self, data=None):
        new_node = Node(data)
        if self.head is None:
            self.head = new_node
        else:
            pointer = self.head
            while pointer.next is not None:
                pointer = pointer.next
            pointer.next = new_node
            
    def show(self):
        pointer = self.head
        while pointer is not None:
            print(pointer.data, "-> ", end='')
            pointer = pointer.next
        print("NULL")


def main():
    print("Linked Lists in Python")
    linkedlist = LinkedList()
    linkedlist.insert(data=22)
    linkedlist.insert(data=30)
    linkedlist.insert(data=57)
    linkedlist.show()
        
        
if __name__ == "__main__":
    main()
