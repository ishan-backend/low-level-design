# low-level-design
Few generic problems on LLD/Machine Coding for practice in C++/Go


## Installing C++
```
xcode-select --install
g++ --version
g++ -o output_file input_file.cpp
```

## Installing Boost Library in C++
```
Functionalities:

1. *Smart Pointers* : Boost provides smart pointers like shared_ptr, unique_ptr, and weak_ptr, which help manage memory automatically and efficiently, preventing memory leaks and simplifying memory management.

2. Containers and Data Structures: Boost offers various containers and data structures, including unordered_map, unordered_set, multi_index_container, and variant, which extend the capabilities of standard C++ containers.

3. Algorithms: Boost includes additional algorithms for tasks such as sorting, searching, and working with ranges, complementing the C++ Standard Library algorithms.

4. Regular Expressions: Boost provides a powerful and flexible regular expression library, boost::regex, for pattern matching and text processing tasks.

5. Multi-threading: Boost.Thread offers a portable way to work with multithreading and provides synchronization mechanisms like mutexes and condition variables.

6. Filesystem Operations: Boost.Filesystem facilitates file and directory manipulation, making it easier to work with the file system.

7. Serialization: The Boost.Serialization library enables object serialization, allowing objects to be converted into a portable binary or text format.

8. Numeric and Math Functions: Boost.Math provides additional mathematical functions, including special functions, random number generators, and various numerical tools.

9. Date and Time: Boost.Date_Time offers classes and functions to work with dates, times, and time zones, beyond what's available in the standard library.

10. Interprocess Communication: Boost.Interprocess provides mechanisms for communication and synchronization between different processes.

11. Networking: Boost.Asio is a versatile networking library that supports various network protocols and asynchronous programming.

12. Testing: Boost.Test provides a powerful framework for unit testing C++ code, making it easier to write and execute test cases.

13. Parsing: Boost.Spirit offers a powerful parsing library that allows you to define grammars directly in C++ code.

```


## How to install, compile and link the binary with your project in C++

```
Step 1: Install a C++ Compiler and Build Tools.
-  The C++ compiler is responsible for translating your C++ source code into object files, and the linker combines these object files together to create the final executable or library.
- On macOS: Install Xcode Command Line Tools, which includes Clang

Step 2: Organize Your Project Files. 
- Create a directory for your project and organize your C++ source files within it.

Step 3: Write Your C++ Code.
- Write your C++ code in the source files with the .cpp extension and any necessary header files with the .h extension

Step 4: Step 4: Compile Your Code
- Open a terminal or command prompt, navigate to your project folder, and use the C++ compiler to compile your source files. For example, using GCC:
g++ -c main.cpp other_file.cpp

- This will generate object files main.o and other_file.o in your project folder.

Step 5: Link Your Object Files
- After compiling, you need to link the object files to create the final executable:
g++ -o my_program main.o other_file.o
- This will generate an executable named my_program (or whatever name you provide after -o)

Step 6: Run Your Program
- On Linux/macOS, run your program from the terminal:
./my_program

```