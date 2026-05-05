cmake_minimum_required(VERSION 3.20)
project(compute_engine VERSION 0.1.0 LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

add_library(engine_lib
    src/pricing.cpp
    src/stats.cpp
)
target_include_directories(engine_lib PUBLIC include)

add_executable(engine src/main.cpp)
target_link_libraries(engine PRIVATE engine_lib)

enable_testing()
add_executable(engine_tests tests/test_main.cpp)
target_link_libraries(engine_tests PRIVATE engine_lib)
add_test(NAME engine_tests COMMAND engine_tests)
