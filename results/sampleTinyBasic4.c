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

int main() {
    LabelMap labels[] = {
        {10, &&label_10},
        {20, &&label_20},
        {30, &&label_30},
        {40, &&label_40},
        {50, &&label_50},
        {60, &&label_60},
    };
    int numLabels = sizeof(labels) / sizeof(labels[0]);
    label_10:
    printf("%s\n", "FOR LOOP EXAMPLE");
    label_20:
    for (int A = 1; A <= 6; A++) {
    label_30:
    printf("%s%d\n", "A = ", A);
    label_40:
    }
    label_50:
    printf("%s\n", "PROGRAM END");
    label_60:
    return 0;
default_label_addr:
}
