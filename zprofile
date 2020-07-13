
PROGRAM="$HOME/Programs"
# PATH="$HOME/.local/bin:$PATH"

# Go support
#export GOROOT="$PROGRAM/go1.12"
#PATH="$GOROOT/bin:$PATH"
GOPATH="$HOME/.gopath"
export PATH="$GOPATH/bin:$PATH"

# Java JDK support
#export JAVA_HOME="$PROGRAM/jdk1.8.0_162"
#PATH="$JAVA_HOME/bin:$PATH"

# maven support
#PATH="$PROGRAM/apache-maven-3.5.3/bin:$PATH"

# gradle support
#PATH="$PROGRAM/gradle-4.6/bin:$PATH"

# protobuf support
#PATH="$PROGRAM/protoc-3.5.1/bin:$PATH"

# docker compose
#PATH="$PROGRAM/bin:$PATH"

# cuda
#CUDAPATH="$PROGRAM/cuda-10.0"
#PATH="$CUDAPATH/bin:$PATH"

#export ENV="/home/jam/env"
#

# Android
export ANDROID_HOME="$HOME/Library/Android/sdk"
export ANDROID_NDK_HOME="$ANDROID_HOME/ndk/21.0.6113669"
export PATH="$ANDROID_HOME/tools:$ANDROID_HOME/platform-tools:$PATH"

# flutter
export PATH="$PROGRAM/flutter/.pub-cache/bin:$PROGRAM/flutter/bin:$PATH"
export PUB_HOSTED_URL="https://pub.flutter-io.cn"
export FLUTTER_STORAGE_BASE_URL="https://storage.flutter-io.cn"

# homebrew support
export HOMEBREW_BOTTLE_DOMAIN="https://mirrors.ustc.edu.cn/homebrew-bottles"


export PATH="$HOME/.cargo/bin:$PATH"
