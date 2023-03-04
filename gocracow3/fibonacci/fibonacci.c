#include <stdio.h>

unsigned long int fibbonacci(unsigned long int n) {
   if(n == 0){
      return 0;
   } 
   
   if(n == 1) {
      return 1;
   }
     
   return (fibbonacci(n-1) + fibbonacci(n-2));
}

int main() {
   printf("Fibonacci of 50 is %li: " , fibbonacci(50));
}
