cmake_minimum_required(VERSION 3.20)
project(rating-engine LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

find_package(PkgConfig QUIET)

if(PKG_CONFIG_FOUND)
  pkg_check_modules(RDKAFKA REQUIRED librdkafka)
  pkg_check_modules(COUCHBASE QUIET libcouchbase)
endif()

add_executable(rating-engine main.cpp)

target_compile_definitions(rating-engine PRIVATE RATING_ENGINE_VERSION="\"0.1.0\"")

target_include_directories(rating-engine PRIVATE ${CMAKE_CURRENT_SOURCE_DIR})

if(RDKAFKA_FOUND)
  target_include_directories(rating-engine PRIVATE ${RDKAFKA_INCLUDE_DIRS})
  target_link_libraries(rating-engine PRIVATE ${RDKAFKA_LIBRARIES})
endif()

if(COUCHBASE_FOUND)
  target_include_directories(rating-engine PRIVATE ${COUCHBASE_INCLUDE_DIRS})
  target_link_libraries(rating-engine PRIVATE ${COUCHBASE_LIBRARIES})
endif()

if(NOT RDKAFKA_FOUND)
  message(WARNING "librdkafka not found; compile-time Kafka integration is disabled. Install librdkafka-dev for real integration.")
endif()

if(NOT COUCHBASE_FOUND)
  message(WARNING "libcouchbase not found; Couchbase C++ integration is disabled. Install libcouchbase-dev for real integration.")
endif()
