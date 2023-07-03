# low-level-design
Few generic problems on LLD/Machine Coding for practice in C++/Go


## Installing C++
xcode-select --install
g++ --version
g++ -o output_file input_file.cpp

## How to install, compile and link the binary with your project in C++
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