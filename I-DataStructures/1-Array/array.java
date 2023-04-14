public class array {
    public static void main(String[] args)
    {
        System.out.println("Arrays in Java");
        String[] fruits = {"Apples", "Bananas", "Oranges", "Peaches", "Berries"};
        System.out.printf("Before: %s\n", fruits[0]);
        fruits[0] = "Pineapples";
        System.out.printf("After: %s\n", fruits[0]);
    }
}