#pragma once
#include <string>
#include "Message.h"

class ISubscriber {
    public:
        virtual std::string getId() const = 0;
        virtual void consumeMessage(Message msg) = 0;
        virtual ~ISubscriber(){}
};

/*
The code defines an abstract C++ class called ISubscriber, which serves as an interface for implementing subscribers in a messaging system. This class contains three virtual member functions and a virtual destructor.

class ISubscriber: This line declares the start of the class definition, and ISubscriber is the name of the class.

public:: This access specifier specifies that the following members (functions and the destructor) are accessible from outside the class.

virtual std::string getId() const = 0;: This is a pure virtual function, indicated by the = 0 at the end. Pure virtual functions have no implementation in the base class and must be implemented by derived classes. Here, getId() is a function that returns a std::string (a string), and it is marked as const, which means it does not modify the state of the object. Any class inheriting from ISubscriber must provide a concrete implementation of getId().

virtual void consumeMessage(Message msg) = 0;: This is another pure virtual function. It takes a Message object (presumably defined elsewhere) as a parameter and returns void, meaning it does not return any value. Like getId(), this function is also marked as virtual and must be implemented by derived classes.

virtual ~ISubscriber(){}: This is a virtual destructor. It is marked as virtual to ensure that the proper destructor of the derived class is called when deleting a pointer to ISubscriber. The = {} at the end means the destructor has an empty body. This class is designed as an interface, and thus, it should have a virtual destructor to ensure proper memory cleanup in derived classes.

In summary, the ISubscriber class serves as an interface in C++. It provides two pure virtual functions (getId() and consumeMessage()) that must be implemented by any class that inherits from ISubscriber. This design allows for polymorphism, allowing different implementations of subscribers to be used interchangeably in the messaging system.
*/