#include <stdio.h>
int main() {
    label_10:
    printf("%s\n", "WHILE LOOP EXAMPLE");
    label_20:
    int A = 1;
    label_30:
    while (A<=5) {
    label_40:
    printf("%s%d\n", "A = ", A);
    label_50:
    A = A+1;
    label_60:
    }
    label_70:
    printf("%s\n", "PROGRAM END");
    label_80:
    return 0;
}
