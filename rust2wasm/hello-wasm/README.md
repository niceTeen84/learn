# Rust 编译 wasm 文件

## 1 安装依赖
```shell
cargo install wasm-pack --no-default-features # 忽略 OpenSSL
```

## 2 编译项目
```shell
wasm-pack build --target web
```

