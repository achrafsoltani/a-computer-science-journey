fn main(){
    println!("Arrays in Rust\n");
    let mut fruits: [&str; 5] = ["Apples", "Bananas", "Oranges", "Peaches", "Berries"];
    println!("Before: {}", fruits[0]);
    fruits[0] = "Pineapples";
    println!("After: {}", fruits[0]);
}