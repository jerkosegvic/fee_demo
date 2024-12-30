#include <stdio.h>
#include <unistd.h>

int main() {
    for (int i = 100; i >= 0; i--) {
        printf("Broj: %d\n", i);
        sleep(1);
    }
    return 0;
}