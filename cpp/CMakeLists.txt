cmake_minimum_required(VERSION 3.20)
project(telcobss-cpp LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

add_library(telcobss_common STATIC
  src/placeholder.cpp
)

target_include_directories(telcobss_common PUBLIC include)

message(STATUS "C++ shared code project configured")
