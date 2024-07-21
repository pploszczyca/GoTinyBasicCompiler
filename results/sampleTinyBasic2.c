#include <stdio.h>

typedef struct {
	int lineNumber;
	void *labelAddr;
} LabelMap;

void* find_label(int lineNumber, LabelMap labels[], int numLabels) {
	for (int i = 0; i < numLabels; ++i) {
		if (labels[i].lineNumber == lineNumber) {
			return labels[i].labelAddr;
		}
	}
}

#define MAX 100 // Define the maximum size of the stack

typedef struct {
    int top;
    void* items[MAX]; // Array to store label addresses
} Stack;

// Initialize the stack
void initStack(Stack* s) {
    s->top = -1;
}

// Check if the stack is empty
int isEmpty(Stack* s) {
    return s->top == -1;
}

// Check if the stack is full
int isFull(Stack* s) {
    return s->top == MAX - 1;
}

// Push a label onto the stack
void push(Stack* s, void* label) {
    if (isFull(s)) {
        return;
    }
    s->items[++(s->top)] = label;
}

// Pop a label from the stack
void* pop(Stack* s) {
    if (isEmpty(s)) {
        return NULL; // Return NULL to indicate stack is empty
    }
    return s->items[(s->top)--];
}

// Peek at the top label of the stack without removing it
void* peek(Stack* s) {
    if (isEmpty(s)) {
        return NULL; // Return NULL to indicate stack is empty
    }
    return s->items[s->top];
}


int main() {
	Stack gosubStack;
    initStack(&gosubStack);
	LabelMap labels[] = {
		{10, &&label_10},
		{20, &&label_20},
		{30, &&label_30},
		{40, &&label_40},
		{50, &&label_50},
		{60, &&label_60},
		{70, &&label_70},
		{80, &&label_80},
		{90, &&label_90},
		{100, &&label_100},
		{110, &&label_110},
		{120, &&label_120},
		{130, &&label_130},
		{140, &&label_140},
		{145, &&label_145},
		{146, &&label_146},
		{150, &&label_150},
		{160, &&label_160},
		{170, &&label_170},
		{180, &&label_180},
		{190, &&label_190},
		{200, &&label_200},
		{205, &&label_205},
		{210, &&label_210},
		{220, &&label_220},
		{230, &&label_230},
		{240, &&label_240},
		{250, &&label_250},
		{260, &&label_260},
		{265, &&label_265},
		{266, &&label_266},
		{330, &&label_330},
		{340, &&label_340},
		{350, &&label_350},
		{360, &&label_360},
		{370, &&label_370},
		{380, &&label_380},
	};
	int numLabels = sizeof(labels) / sizeof(labels[0]);
	label_10:
	printf("%s\n", "WELCOME TO TINYBASIC");
	label_20:
	printf("%s\n", "THIS PROGRAM PERFORMS A SERIES OF OPERATIONS");
	label_30:
	int X = 5;
	label_40:
	int Y = 15;
	label_50:
	int Z = X*Y;
	label_60:
	printf("%s%d\n", "X * Y = ", Z);
	label_70:
	printf("%s\n", "ENTER YOUR FAVOURITE NUMBER: ");
	label_80:
	int N;
	scanf("%d", &N);
	label_90:
	printf("%s%d\n", "YOUR NUMBER IS: ", N);
	label_100:
	printf("%s\n", "LET'S DO SOME MATH!");
	label_110:
	int S = 0;
	label_120:
	int I = 1;
	label_130:
	if (I>10) goto *find_label(150, labels, numLabels);
	label_140:
	S = S+I;
	label_145:
	I = I+1;
	label_146:
	goto *find_label(130, labels, numLabels);
	label_150:
	printf("%s%d\n", "SUM OF 1 TO 10 = ", S);
	label_160:
	int F = 1;
	label_170:
	int J = 1;
	label_180:
	if (J>5) goto *find_label(210, labels, numLabels);
	label_190:
	F = F*J;
	label_200:
	J = J+1;
	label_205:
	goto *find_label(180, labels, numLabels);
	label_210:
	printf("%s%d\n", "FACTORIAL OF 5 IS ", F);
	label_220:
	int K = 1;
	label_230:
	printf("%s\n", "SQUARE NUMBERS FROM 1 TO 10");
	label_240:
	if (K>10) goto *find_label(330, labels, numLabels);
	label_250:
	int Q = K*K;
	label_260:
	printf("%d%s%d\n", K, " SQUARED IS ", Q);
	label_265:
	K = K+1;
	label_266:
	goto *find_label(240, labels, numLabels);
	label_330:
	printf("%s\n", "ENTER ANOTHER NUMBER TO FIND ITS SQUARE ROOT: ");
	label_340:
	int P;
	scanf("%d", &P);
	label_350:
	int R = P*P;
	label_360:
	printf("%s%d%s%d\n", "SQUARE OF ", P, " IS ", R);
	label_370:
	printf("%s\n", "END OF PROGRAM");
	label_380:
	return 0;
}
