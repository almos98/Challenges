#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int sumOfMultiples (int* multiples, int multiples_len, int limit) {
    int sum = 0;

    for (int i = 1; i < limit; i++) {
        bool multiple = false;
        for (int j = 0; j < multiples_len; j++) {
            int n = multiples[j];
            multiple = multiple || (i % n == 0);
        }
        sum = multiple ? sum + i : sum;
    }

    return sum;
}

int main (int argc, char** argv) {
   int multiples [2] = {3,5};
   printf("%d\n", sumOfMultiples(multiples, 2, 1000));
}
