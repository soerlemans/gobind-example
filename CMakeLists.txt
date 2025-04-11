# Set minimum CMake version.
cmake_minimum_required(VERSION 3.22)

# Define the project.
project(
  gobind-example
  VERSION 0.1.0
  DESCRIPTION "Example usage of the gobind library."
  LANGUAGES CXX C
)

# Create target.
set(TARGET_GOBIND_EXAMPLE ${PROJECT_NAME})

# Create a library.
add_library(
  ${TARGET_GOBIND_EXAMPLE} SHARED
	hello_world.cpp
)

# Specify the C++ standard.
set_target_properties(
	${TARGET_GOBIND}
	PROPERTIES
	CXX_STANDARD 20
)

# Define include directories for headers
target_include_directories(
  ${TARGET_GOBIND_EXAMPLE}
	PUBLIC /usr/local/include
)

# Linking
link_directories(/usr/local/lib)

target_link_libraries(
	${TARGET_GOBIND_EXAMPLE}
  PRIVATE gobind gobind
)
