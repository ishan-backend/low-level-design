#pragma once
#include <string>
#include <chrono>

class Message {
        std::string msg;
        std::chrono::time_point<std::chrono::system_clock> ts;

    public:
        Message(std::string);
        std::string getMessage() const;
};

/*
    std::string getMessage() const -> getMessage is const member function, which promises to not modify any of the member variables of the class, and its return type is string.
*/