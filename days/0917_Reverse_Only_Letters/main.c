#include <stdlib.h>
#include <stdio.h>

#include <stddef.h>
#include <stdint.h>
#include <stdbool.h>

char* reverseOnlyLetters(char* s) {
    size_t n = strlen(s); // c moment.

    int i = 0;
    int j = n-1;

    while (i < j) {
        while (i < j && !isalpha(s[i])) i++;
        while (i < j && !isalpha(s[j])) j--;

        if (i >= j) break;

        char tmp = s[i];
        s[i] = s[j];
        s[j] = tmp;
        i++;
        j--;
    }
    return s;
}


int main(void) {
    printf("hello world!\n");

    return 0;
}
