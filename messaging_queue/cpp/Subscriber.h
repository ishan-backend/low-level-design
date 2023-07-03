#pragma once
#include "ISubscriber.h"
#include <boost/uuid/random_generator.hpp>

class Subscriber: public ISubscriber {
    private:
        static boost::uuids::random_generator rg;
        std::string id;
        std::string name;
    
    public:
        Subscriber(std::string);
        std::string getId() const override;
        void consumeMessage(Message msg) override;
};

/*
This code defines a C++ class named `Subscriber`, which is a concrete implementation of the abstract base class `ISubscriber`. It represents a subscriber in a messaging system and contains member variables and member functions for handling messages.

Let's break down each part of the `Subscriber` class definition:

1. `class Subscriber: public ISubscriber`: This line defines the `Subscriber` class, which is derived from the `ISubscriber` class. It uses the `public` inheritance mode, meaning that all public members of `ISubscriber` become public members of `Subscriber`. This indicates that `Subscriber` is a specialization of `ISubscriber` and provides concrete implementations for the pure virtual functions declared in the base class.

2. `private:`: This access specifier specifies that the following members (variables and functions) are accessible only within the `Subscriber` class and not from outside.

3. `static boost::uuids::random_generator rg;`: This line declares a static member variable named `rg`. It uses the `boost::uuids::random_generator` type from the Boost C++ Libraries. Static member variables are shared among all instances of the class, and they must be defined outside the class as well.

4. `std::string id;`: This line declares a private member variable `id`, which is of type `std::string`. It represents the unique identifier of the subscriber.

5. `std::string name;`: This line declares another private member variable `name`, which is also of type `std::string`. It represents the name of the subscriber.

6. `public:`: This access specifier specifies that the following members are accessible from outside the class.

7. `Subscriber(std::string);`: This is the constructor of the `Subscriber` class, which takes a `std::string` parameter and initializes the `name` member variable with the provided string.

8. `std::string getId() const override;`: This is a member function that overrides the `getId()` function from the base class `ISubscriber`. It returns a `std::string` (the unique identifier) and is marked as `const`, meaning it does not modify the state of the `Subscriber` object.

9. `void consumeMessage(Message msg) override;`: This is another member function that overrides the `consumeMessage()` function from the base class `ISubscriber`. It takes a `Message` object (presumably defined elsewhere) as a parameter and returns `void`, meaning it does not return any value. This function is responsible for handling the received message.

In summary, the `Subscriber` class is a concrete implementation of the `ISubscriber` interface. It defines a constructor to set the subscriber's name, and it provides implementations for the `getId()` and `consumeMessage()` functions required by the base class. The `rg` static member variable is likely used for generating unique identifiers for each `Subscriber` object.
 
*/