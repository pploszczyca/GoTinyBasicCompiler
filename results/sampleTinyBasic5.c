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
		{100, &&label_100},
		{110, &&label_110},
		{120, &&label_120},
		{1000, &&label_1000},
		{1010, &&label_1010},
		{1020, &&label_1020},
		{2000, &&label_2000},
		{2010, &&label_2010},
	};
	int numLabels = sizeof(labels) / sizeof(labels[0]);
	label_10:
	push(&gosubStack, &&label_gosub_1);
	goto *find_label(100, labels, numLabels);
	label_gosub_1: ;
	label_20:
	return 0;
	label_100:
	push(&gosubStack, &&label_gosub_2);
	goto *find_label(1000, labels, numLabels);
	label_gosub_2: ;
	label_110:
	printf("%s\n", "END GOSUB 1000");
	label_120:
	goto *pop(&gosubStack);;
	label_1000:
	push(&gosubStack, &&label_gosub_3);
	goto *find_label(2000, labels, numLabels);
	label_gosub_3: ;
	label_1010:
	printf("%s\n", "END GOSUB 2000");
	label_1020:
	goto *pop(&gosubStack);;
	label_2000:
	printf("%s\n", "HELLO, WORLD!");
	label_2010:
	goto *pop(&gosubStack);;
}
