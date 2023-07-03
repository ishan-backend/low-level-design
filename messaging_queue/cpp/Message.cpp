#include "Message.h"

Message::Message(std::string message): msg(message) {
    ts = std::chrono::system_clock::now();
}

/*
Message::Message(std::string message): This line defines the constructor for the class Message. Constructors have the same name as the class, and they are called when an object of the class is created. This constructor takes a std::string parameter named message.

: message(message): This part is the member initializer list. It initializes the member variable message with the value of the message parameter. Member initializer lists are used to set the initial values of class members before the constructor's body is executed.

timestamp = std::chrono::system_clock::now();: This line sets the value of the timestamp member variable. It uses the <chrono> library to obtain the current system time as a std::chrono::time_point object using std::chrono::system_clock::now(). The timestamp member variable is likely declared with a type compatible with std::chrono::time_point, allowing it to store the current system time when the Message object is created.
*/

std::string Message::getMessage() const {
    return msg;
}