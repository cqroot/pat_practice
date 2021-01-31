#include <iostream>
#include <string>

int main() {
    int num;

    std::string max_name, min_name, max_id, min_id;
    int max_score = -1, min_score = 101;

    std::cin >> num;

    for (int i = 0; i < num; ++i) {
        std::string current_name, current_id;
        int current_score;
        std::cin >> current_name >> current_id >> current_score;

        if (current_score > max_score) {
            max_score = current_score;
            max_name = current_name;
            max_id = current_id;
        }
        if (current_score < min_score) {
            min_score = current_score;
            min_name = current_name;
            min_id = current_id;
        }
    }

    std::cout << max_name << " " << max_id << std::endl;
    std::cout << min_name << " " << min_id << std::endl;
    return 0;
}
