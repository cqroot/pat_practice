#include <iostream>
#include <string>


int main() {
    std::string num;
    std::cin >> num;

    if (num == "0") {
        std::cout << "ling" << std::endl;
    }

    std::string numstr[] = {
        "ling", "yi", "er", "san", "si",
        "wu", "liu", "qi", "ba", "jiu"
    };

    int sum = 0;
    for (int i = 0; i < num.size(); ++i) {
        sum += static_cast<int>(num[i]) - '0';
    }

    if (sum < 10) {
        std::cout << numstr[sum] << std::endl;
    } else if (sum < 100) {
        std::cout << numstr[sum/10] << " " << numstr[sum%10] << std::endl;
    } else {
        std::cout << numstr[sum/100] << " " << numstr[sum/10%10]
            << " " << numstr[sum%10] << std::endl;
    }

    return 0;
}
