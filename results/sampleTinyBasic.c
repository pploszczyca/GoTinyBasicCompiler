#include <stdio.h>
int main() {
    label_10:
    printf("%s\n", "HELLO, WORLD!");
    label_20:
    int A = 10;
    label_30:
    int B = 20;
    label_40:
    int C = A+B;
    label_50:
    printf("%s%d\n", "A + B = ", C);
    label_60:
    printf("%s\n", "ENTER A NUMBER: ");
    label_70:
    int D;
    scanf("%d", &D);
    label_80:
    if (D>10) goto label_110;
    label_90:
    printf("%s\n", "THE NUMBER IS LESS THAN OR EQUAL TO 10");
    label_100:
    goto label_120;
    label_110:
    printf("%s\n", "THE NUMBER IS GREATER THAN 10");
    label_120:
    int E = 1+2+3+4+5;
    label_130:
    printf("%s%d\n", "SUM OF 1 TO 5 = ", E);
    label_140:
    printf("%s\n", "PROGRAM END");
    label_150:
    return 0;
}
