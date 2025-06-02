# Cachito

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/achetronic/cachito)
![GitHub](https://img.shields.io/github/license/achetronic/cachito)

![YouTube Channel Subscribers](https://img.shields.io/youtube/channel/subscribers/UCeSb3yfsPNNVr13YsYNvCAw?label=achetronic&link=http%3A%2F%2Fyoutube.com%2Fachetronic)
![GitHub followers](https://img.shields.io/github/followers/achetronic?label=achetronic&link=http%3A%2F%2Fgithub.com%2Fachetronic)
![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/achetronic?style=flat&logo=twitter&link=https%3A%2F%2Ftwitter.com%2Fachetronic)

**Cachito** is a collection of ready-to-use, thread-safe cache algorithms for Go with customizable handlers and metadata support. 
Built with simplicity and extensibility in mind.

## ✨ Features

- **Thread-safe** implementations
- **Customizable handlers** for insert, delete, access, and eviction events
- **User-defined metadata** accessible in all handlers
- **Zero external dependencies**
- **Generic types** support (Go 1.18+)
- **Comprehensive examples** for each algorithm

## 🎯 Supported Algorithms

| Algorithm                            | Status     | Description                            |
|--------------------------------------|------------|----------------------------------------|
| **LRU** (Least Recently Used)        | ✅ Complete | Evicts least recently accessed items   |
| **LFU** (Least Frequently Used)      | 🚧 Planned | Evicts least frequently accessed items |
| **FIFO** (First In, First Out)       | 📋 Planned | Evicts oldest items first              |
| **LIFO** (Last In, First Out)        | 📋 Planned | Evicts newest items first              |
| **TTL** (Time To Live)               | 📋 Planned | Time-based expiration                  |
| **ARC** (Adaptive Replacement Cache) | 📋 Planned | Self-tuning between LRU and LFU        |

## 🚀 Quick Start

### Installation

```bash
go get github.com/achetronic/cachito
```

## 📚 Examples

Detailed examples for each algorithm can be found in the `/examples` directory:

- `/examples/lru/` - LRU cache examples
    - `based-on-count.go` - Size-based eviction
    - `based-on-disk.go`  - Disk-based eviction
- `/examples/lfu/` - LFU cache examples (planned)
- `/examples/ttl/` - TTL cache examples (planned)

## 🎛️ Handler System

Cachito allows you to hook into different cache operations:

### Available Handlers

Handlers are user-defined functions that are triggered in different moments. Some data are passed to those functions.
Do you need some examples?

| Handler       | Trigger                      | Use Cases                    |
|---------------|------------------------------|------------------------------|
| `OnInsert`    | When a new entry is created  | Logging, metrics, validation |
| `OnDelete`    | When an entry is removed     | Cleanup, notifications       |
| `OnAccess`    | When an entry is accessed    | Analytics, usage tracking    |
| `ShouldEvict` | Before insertion (if needed) | Custom eviction logic        |

## 🗺️ Roadmap

### Core Algorithms
- [x] LRU (Least Recently Used)
- [ ] LFU (Least Frequently Used)
- [ ] FIFO (First In, First Out)
- [ ] LIFO (Last In, First Out)
- [ ] TTL (Time To Live)
- [ ] ARC (Adaptive Replacement Cache)

### Features
- [x] Thread-safe operations
- [x] Generic type support
- [x] Customizable handlers
- [x] User-defined metadata
- [ ] Disk persistence
- [ ] Cache warming strategies
- [ ] Metrics and monitoring

## 🤝 Contributing

All contributions are welcome! Whether you're reporting bugs, suggesting features, or submitting code — thank you! Here’s how to get involved:

▸ [Open an issue](https://github.com/achetronic/cachito/issues/new) to report bugs or request features

▸ [Submit a pull request](https://github.com/achetronic/cachito/pulls) to contribute improvements

<!---
▸ [Ask a question or start a discussion](https://github.com/achetronic/cachito/discussions)
-->

▸ [Check open milestones](https://github.com/achetronic/cachito/milestones) to see what’s coming

▸ [Read the contributing guide](./docs/CONTRIBUTING.md) to get started smoothly


## 📄 License

Cachito is licensed under the [Apache 2.0 License](./LICENSE).

## 🙏 Acknowledgments

- Inspired by the need for flexible, production-ready cache implementations
- Built with ❤️ for the Go community
