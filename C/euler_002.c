#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int fibonacci(int limit, bool (*check)(int)) {
    int a = 1;
    int b = 2;
    int sum = 2;
    int nextNum = a + b;

    do {
        //printf("%d + %d = %d\n", a,b,nextNum);
        if (check(nextNum)) {
            printf("sum: %d, num: %d\n", sum, nextNum);
            sum += nextNum;
        }

        a = b;
        b = nextNum;
        nextNum = a + b;
    } while (nextNum < limit);

    return sum;
}

int main (int argc, char** argv) {
    bool check(int n) {
        return n % 2 == 0;
    }
    
    int sum = fibonacci(4000000, check);
    printf("%d\n", sum);
}
