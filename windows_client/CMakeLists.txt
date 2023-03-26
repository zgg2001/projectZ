cmake_minimum_required(VERSION 3.5)

project(windows_client LANGUAGES CXX)

set(CMAKE_INCLUDE_CURRENT_DIR ON)

set(CMAKE_AUTOUIC ON)
set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# QtCreator supports the following variables for Android, which are identical to qmake Android variables.
# Check http://doc.qt.io/qt-5/deployment-android.html for more information.
# They need to be set before the find_package(Qt5 ...) call.

#if(ANDROID)
#    set(ANDROID_PACKAGE_SOURCE_DIR "${CMAKE_CURRENT_SOURCE_DIR}/android")
#    if (ANDROID_ABI STREQUAL "armeabi-v7a")
#        set(ANDROID_EXTRA_LIBS
#            ${CMAKE_CURRENT_SOURCE_DIR}/path/to/libcrypto.so
#            ${CMAKE_CURRENT_SOURCE_DIR}/path/to/libssl.so)
#    endif()
#endif()

find_package(Qt5 COMPONENTS Widgets REQUIRED)
find_package(gRPC CONFIG REQUIRED)
find_package(protobuf CONFIG REQUIRED)

if(ANDROID)
  add_library(windows_client SHARED
    main.cpp
    widget.cpp
    widget.h
  )
else()
  add_executable(windows_client
    main.cpp
    widget.cpp
    widget.h
    pbfile/service.pb.h
    pbfile/service.grpc.pb.h
    pbfile/service.pb.cc
    pbfile/service.grpc.pb.cc
  )
endif()

target_link_libraries(windows_client PRIVATE Qt5::Widgets)
target_link_libraries(windows_client PRIVATE gRPC::gpr gRPC::grpc gRPC::grpc++ gRPC::grpc++_alts)
target_link_libraries(windows_client PRIVATE protobuf::libprotoc protobuf::libprotobuf protobuf::libprotobuf-lite)

