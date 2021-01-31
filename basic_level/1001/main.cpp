#include <iostream>

int main() {
    int n, count = 0;
    std::cin >> n;
    while (n != 1) {
        if (n % 2 == 1) {
            n = (3 * n + 1) / 2;
        } else {
            n = n / 2;
        }
        count += 1;
    }
    std::cout << count << std::endl;
    return 0;
}
