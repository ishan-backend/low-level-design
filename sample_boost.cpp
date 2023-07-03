#include <iostream>
#include <boost/algorithm/string.hpp>

int main() {
    std::string inputString = "Hello, boost C++ Libraries!";
    
    // Convert the string to uppercase using Boost
    std::string upperString = boost::algorithm::to_upper_copy(inputString);
    std::cout << "Uppercase string: " << upperString << std::endl;
    
    // Check if the string contains a specific substring using Boost
    std::string subString = "Boost";
    if (boost::algorithm::contains(inputString, subString)) {
        std::cout << "The input string contains the substring \"" << subString << "\"" << std::endl;
    } else {
        std::cout << "The input string does not contain the substring \"" << subString << "\"" << std::endl;
    }
    
    return 0;
}
