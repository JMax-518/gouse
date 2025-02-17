<div align="center">
    <img src="https://github.com/thuongtruong1009/gouse/actions/workflows/analysis.yml/badge.svg?branch=main" alt="ci_status">
    <img src="https://github.com/thuongtruong1009/gouse/actions/workflows/ci.yml/badge.svg?branch=main" alt="ci_status">
    <a href="https://sonarcloud.io/summary/new_code?id=thuongtruong1009_gouse"><img src="https://sonarcloud.io/api/project_badges/measure?project=thuongtruong1009_gouse&metric=alert_status" alt="sonar"></a>
    <a href="https://app.codacy.com/gh/thuongtruong1009/gouse/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade"><img src="https://app.codacy.com/project/badge/Grade/21f940894abd4e0384ef8b84adc294da" alt="codacy"></a>
    <a href="https://app.deepsource.com/gh/thuongtruong1009/gouse/" target="_blank"><img alt="DeepSource" title="DeepSource" src="https://app.deepsource.com/gh/thuongtruong1009/gouse.svg/?label=resolved+issues&show_trend=false&token=VqBk6AsowWePd3khy3AUkvXJ"/></a>
    <a href="https://goreportcard.com/report/thuongtruong1009/gouse"><img src="https://goreportcard.com/badge/github.com/thuongtruong1009/gouse" alt="go_report_card"></a>
    <a href="https://codecov.io/gh/thuongtruong1009/gouse"><img src="https://codecov.io/gh/thuongtruong1009/gouse/branch/main/graph/badge.svg" alt="codecov"></a>
    <a href="https://dl.circleci.com/status-badge/redirect/gh/thuongtruong1009/gouse/tree/main"><img src="https://dl.circleci.com/status-badge/img/gh/thuongtruong1009/gouse/tree/main.svg?style=svg" alt="circleci"></a>
    <a href="https://pkg.go.dev/github.com/thuongtruong1009/gouse"><img src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square" alt="go.dev"></a>
    <img src="./public/count.svg" alt="gouse_functions_count">
    <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/thuongtruong1009/gouse">
    <a href=""><img src="https://img.shields.io/github/license/thuongtruong1009/gouse" alt="license"></a>
</div>

# ![](/public/banner.png)

## 🧠 Why Gouse?

- Gouse is a modern essential [`Golang`](https://golang.org/) utility package delivering consistency, modularity, performance, & extras presets inspired by [`Lodash`](https://lodash.com/)
- `Javascript` user-friendly syntax.
- No config - import as utility functions.
- Lightweight package: easy to use, chainable, and extendable, and available in various builds & module formats.
- Gouse provides a wide variety of available methods, taking the hassle out of working with arrays, numbers, objects, strings, etc. Each method has different features, so you can pick the ones that fit your project best.
- Comprehensive documentation and examples.
- Powerful package that is suitable for small to large projects and compatible with all OS platforms
- Thanks to Gouse, you can:
  - Set up and scale projects rapidly.
  - Handle complex logic use-cases such as database connection, build APIs, error handling, log management...
  - Optimize performance and increase productivity.
  - Build easily consistent systems with available functions.
  - Avoid writing repetitive code and a unified code style.
  - Reduce the number of lines of code and make it easier to read, understand, and maintain.
  - Avoid compatibility conflicts and unexpected errors.
  - And more...

## ✨ Motivation

- Go has emerged as a server language, but it still doesn't have complete and consistent packages available to support coding development.
  - Developers must write by hand or search manually. That wastes time and even causes many compatibility conflicts
  - Must update each dependent package every time you update
  - Unexpected errors can easily arise during execution
  - Performance is not optimized
  - The number of lines of code is very long that not easy to read and understand
  - Code logic may not be consistent, making it difficult to maintain and scale

👉 To address that need, Gouse was created as a powerful toolkit for Go developers, a collection of built-in functions and best practices that provide comprehensive, powerful, and reliable solutions. Trusted to build services, APIs, and web applications.

## 🚀 Features

> Below is a list of modules that Gouse supports. This project is still in development stage, so not all features are available.

✅ Array <br/>
✅ Cache <br/>
✅ Chart <br/>
✅ Config <br/>
✅ Connection <br/>
✅ Console <br/>
✅ Date <br/>
✅ Function <br/>
✅ Helper <br/>
✅ I/O <br/>
✅ Log <br/>
✅ Math <br/>
✅ Net <br/>
✅ Number <br/>
✅ Regex <br/>
✅ Struct <br/>
✅ String <br/>
✅ Tool <br/>
✅ Type <br/>

## 📋 Requirements

> Compatibility with Go >= **`1.18`**

## 📦 Installation

```go
go get github.com/thuongtruong1009/gouse
```

## 🕯️ Quick Start

```go
package main

import "github.com/thuongtruong1009/gouse"

func main() {
    gouse.Stater()
}
```

## 🦄 Usage

- Using package directly in your module as ultra-lightweight utility functions.

- View more examples at [`sample`](sample) folder.

## 📖 Documentation

- To read the completely package documentation guide, reference at [![Dev package](https://pkg.go.dev/badge/github.com/thuongtruong1009/gouse)](https://pkg.go.dev/github.com/thuongtruong1009/gouse)

## 📊 Benchmark

`GOOS`: windows

`GOARCH`: amd64

`CPU`: AMD Ryzen 5 5600U with Radeon Graphics (12) @ 2.300GHz

`ITERATIONS`: 5

`INPUT`: 1000000

<!-- ## 📚 Examples -->

## 📁 Project Structure

![Project Structure](https://raw.githubusercontent.com/thuongtruong1009/gouse/diagram/project-structure.svg)

## 🛠️ Development

```bash
$ git clone https://github.com/thuongtruong1009/gouse.git
```

- Refer to the [`Contributing Commands`](.github/CONTRIBUTING.md#commands) section for more information on how to develop the project.

## 📝 Contributing

- We welcome your contributions! If you're looking for issues to work on, try looking at the good first issue list. We do our best to tag issues suitable for new external contributors with that label, so it's a great way to find something you can help with!

- Please read our [`Code of Conduct`](.github/CODE_OF_CONDUCT.md) before contributing.

- Refer to the [`Contributing Guide`](.github/CONTRIBUTING.md) for more information on how to get started.

## 📄 License

- Gouse is released under the [`MIT License`](LICENSE). See the LICENSE file for more information.

- For more information, see the [`Licensing FAQs`](https://opensource.org/faq#mit-vs-bsd).

## 📌 Support

- The tool has been tested on a variety of inputs, but it's not perfect.
- For support in using Gouse, please reach out in the following venues:
  - [`Raise Issues`](https://github.com/thuongtruong1009/gouse/issues/new) - For generally applicable issues and feedback.
  - [`Join Discussions`](https://github.com/thuongtruong1009/gouse/discussions) - For ideas, questions, or issues regarding Gouse's design, development, and future.

## 📜 Changelog

- Gouse is under active development. This means that new features, bug fixes, and breaking changes will be released frequently. We encourage you to keep the [`CHANGELOG`](CHANGELOG.md) open while upgrading to see what's new!

- For more information on how to use the changelog, please refer to [`Keeping a Changelog`](https://keepachangelog.com/en/1.0.0/).

## 🧬 Dependencies

- Gouse is built on top of the following below and others open-source projects

- Special thanks to the following dependencies that helped make this project possible:
  - [Google UUID](github.com/google/uuid) - A fast and simple UUID library for Go 🔑
  - [Survey v2](github.com/AlecAivazis/survey/v2) - A golang library for building interactive prompts with full support for windows and posix terminals 🙋
  - [Bubbletea](github.com/charmbracelet/bubbletea) - A powerful little TUI framework 🏗
  - [Go Cache](github.com/patrickmn/go-cache) - An in-memory key:value store/cache (similar to Memcached) 🗄
  - [Crypto](golang.org/x/crypto) - A collection of cryptographic algorithms and protocols for Go 📦
  - [Go eCharts](github.com/go-echarts/go-echarts) - 🎨 The adorable charts library for Golang 📊
  - [Env config](github.com/kelseyhightower/envconfig) x [Go Toml](github.com/pelletier/go-toml) - A Go library for managing configuration data from files 📄
  - [Minio Client Go](github.com/minio/minio-go/v7) - MinIO Go Library for Amazon S3 compatible cloud storage 📦

## 📚 References

- [Golang](https://golang.org/)
- [Golang Docs](https://pkg.go.dev/)
- [Lodash.js](https://lodash.com/)
- [Lodash collection](https://www.geeksforgeeks.org/lodash/?ref=header_search)
- [Javascript methods reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference)
- [File handling in Golang](https://www.cloudhadoop.com/2018/11/learn-golang-tutorials-file-directory.html)
- [Golang x Github Actions](https://github.com/brpaz/github-actions-demo-go)
- [Dynamic HTML](https://css-tricks.com/dynamic-page-replacing-content)
- [Profiling Go Programs](https://blog.golang.org/pprof) - [Example](https://dev.to/immortalt/use-pprof-for-golang-program-memory-analysis-2cj6)
