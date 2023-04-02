#include <stdio.h>
#include <string.h>
    
int main(){
  printf("Arrays in C!\n");
  char fruits[5][10] = {"Apples", "Bananas", "Oranges", "Peaches", "Berries"};
  printf("Before: %s\n", fruits[0]);
  strcpy(fruits[0], "Pineapples");
  printf("Before: %s\n", fruits[0]);
  return 0;
}

// gcc array.c
// ./a.out