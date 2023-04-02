pub struct Node {
    data: u32,
    next: Option<Box<Node>>
}

struct LinkedList {
    head: Option<Box<Node>>
}

impl LinkedList {
    fn Insert(&mut self, d: u32) {
        let newNode = Node {data: d, next: Some(Box::new(Node{data: 0, next: None}))};
        if assert_eq!(*self.head, None) {
            self.head = newNode
        } else {
            
        }
    }
}

fn main(){
    
}