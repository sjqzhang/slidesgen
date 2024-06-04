#!/bin/bash

# 设置目标平台和架构
targetPlatforms=("windows" "linux" "darwin")
targetArchs=("amd64" "386")

# 遍历目标平台和架构
for platform in "${targetPlatforms[@]}"; do
  for arch in "${targetArchs[@]}"; do
    # 设置构建目标
    target="slidesgen_${platform}_${arch}"

    # 执行构建命令
    GOOS=${platform} GOARCH=${arch} go build -o ${target} slidesgen.go

    # 检查构建结果
    if [ $? -eq 0 ]; then
      echo "构建成功，生成的可执行文件: ${target}"
    else
      echo "构建失败"
    fi
  done
done