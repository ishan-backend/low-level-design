#pragma once
#include <atomic>
#include <memory>
#include "ISubscriber.h"
#include <functional>

class TopicSubscriber {
    private:
        std::atomic<int> offset;
        std::shared_ptr<ISubscriber> sub_ptr;
    
    public:
        TopicSubscriber(std::shared_ptr<ISubscriber> sub_ptr, int offset = 0);
        int getOffset() const;
        std::shared_ptr<ISubscriber> getSubscriber() const;
        void setOffset(int);
        void incrementOffset();
};

struct TopicSubscriberCompare {
    bool operator()(const std::shared_ptr<TopicSubscriber> & tsptr1, const std::shared_ptr<TopicSubscriber> & tsptr2) const {
        return tsptr1->getSubscriber() < tsptr2->getSubscriber();
    }
};

/*
The code provided defines a C++ class named `TopicSubscriber`, which seems to represent a subscriber associated with a specific topic. It is designed to keep track of the subscriber and an offset value that may have significance in the context of message consumption.

Let's break down each part of the `TopicSubscriber` class definition:

1. `#pragma once`: This is a preprocessor directive that ensures the header file is included only once during compilation, preventing potential issues with duplicate declarations.

2. `#include <memory>`, `#include <atomic>`, `#include <functional>`, `#include "Subscriber.h"`: These are include directives, allowing the class to use functionality from the `<memory>`, `<atomic>`, and `<functional>` standard C++ headers. It also includes the header file for the `Subscriber` class, which seems to be a dependency.

3. `std::atomic<int> offset;`: This member variable is an `std::atomic` object of type `int`. It is likely used to store an offset value associated with this subscriber. The use of `std::atomic` ensures that access to this variable is thread-safe.

4. `std::shared_ptr<ISubscriber> sub_ptr;`: This member variable is a shared pointer to an `ISubscriber` object. It represents the subscriber associated with this `TopicSubscriber`. The use of `std::shared_ptr` implies that this object is part of shared ownership.

5. `TopicSubscriber(std::shared_ptr<ISubscriber>, int offset = 0);`: This is the constructor of the `TopicSubscriber` class. It takes two parameters: a shared pointer to an `ISubscriber` object, representing the subscriber, and an optional `int` parameter, representing the initial offset value (defaulting to 0). It initializes the `sub_ptr` and `offset` member variables with the provided values.

6. `int getOffset() const;`: This member function is a getter method that returns the current offset value associated with this `TopicSubscriber`.

7. `std::shared_ptr<ISubscriber> getSubscriber() const;`: This member function is a getter method that returns the shared pointer to the `ISubscriber` object representing the subscriber associated with this `TopicSubscriber`.

8. `void setOffset(int);`: This member function is used to set the offset value associated with this subscriber. It takes an `int` parameter and updates the `offset` member variable accordingly.

9. `void incrementOffset();`: This member function is used to increment the offset value by 1. It is likely used when consuming messages from the topic, and the offset represents the message index or position.

10. `struct TopicSubscriberCompare { ... };`: This is a custom comparison functor used for comparing shared pointers to `TopicSubscriber` objects. It is likely used to order or sort subscribers based on their underlying subscriber objects.

In summary, the `TopicSubscriber` class represents a subscriber associated with a specific topic. It maintains an offset value (possibly used for message consumption tracking) and a shared pointer to the `ISubscriber` interface representing the subscriber itself. The use of `std::atomic` ensures that concurrent access to the offset variable is thread-safe. The custom comparison functor TopicSubscriberCompare can be used to order or sort subscribers based on their underlying subscriber objects.`

*/