cmake_minimum_required(VERSION 3.22)
project("Message Queue")
set(SRC_DIR "cpp")

# Find Boost libraries
find_package(Boost REQUIRED COMPONENTS system thread)

# Locate GTest (if needed)
# find_package(GTest REQUIRED)

# Add include directories for Boost and GTest (if needed)
include_directories(${SRC_DIR} ${Boost_INCLUDE_DIRS})
# include_directories(${GTEST_INCLUDE_DIRS})

file(GLOB src_files ${SRC_DIR}/*.cpp)
list(FILTER src_files EXCLUDE REGEX ".*\\.test\\.cpp$")

# file(GLOB tst_files ${SRC_DIR}/*.cpp)
# list(FILTER tst_files EXCLUDE REGEX ".*main\\.cpp$")

add_executable(messageQueue ${src_files})
# add_executable(programTests ${tst_files})

# Link Boost libraries
target_link_libraries(messageQueue PRIVATE ${Boost_LIBRARIES})
# target_link_libraries(programTests ${GTEST_LIBRARIES} pthread)
