#include <stdio.h>
int main() {
    label_10:
    printf("%s\n", "WELCOME TO TINYBASIC");
    label_20:
    printf("%s\n", "THIS PROGRAM PERFORMS A SERIES OF OPERATIONS");
    label_30:
    int X = 5 ;
    label_40:
    int Y = 15 ;
    label_50:
    int Z = X*Y ;
    label_60:
    printf("%s%d\n", "X * Y = ", Z);
    label_70:
    printf("%s\n", "ENTER YOUR NAME: ");
    label_80:
    int N;
    scanf("%d", &N);
    label_90:
    printf("%s%d\n", "HELLO, ", N);
    label_100:
    printf("%s\n", "LET'S DO SOME MATH!");
    label_110:
    int S = 0 ;
    label_120:
    int I = 1 ;
    label_130:
    if (I>10) goto label_150;
    label_140:
    int S = S+I ;
    label_145:
    int I = I+1 ;
    label_146:
    goto label_130;
    label_150:
    printf("%s%d\n", "SUM OF 1 TO 10 = ", S);
    label_160:
    int F = 1 ;
    label_170:
    int J = 1 ;
    label_180:
    if (J>5) goto label_210;
    label_190:
    int F = F*J ;
    label_200:
    int J = J+1 ;
    label_205:
    goto label_180;
    label_210:
    printf("%s%d\n", "FACTORIAL OF 5 IS ", F);
    label_220:
    int K = 1 ;
    label_230:
    printf("%s\n", "SQUARE NUMBERS FROM 1 TO 10");
    label_240:
    if (K>10) goto label_270;
    label_250:
    int Q = K*K ;
    label_260:
    printf("%d%s%d\n", K, " SQUARED IS ", Q);
    label_265:
    int K = K+1 ;
    label_266:
    goto label_240;
    label_330:
    printf("%s\n", "ENTER ANOTHER NUMBER TO FIND ITS SQUARE ROOT: ");
    label_340:
    int P;
    scanf("%d", &P);
    label_350:
    int R = P*P ;
    label_360:
    printf("%s%d%s%d\n", "SQUARE OF ", P, " IS ", R);
    label_370:
    printf("%s\n", "END OF PROGRAM");
    label_380:
    return 0;
}
