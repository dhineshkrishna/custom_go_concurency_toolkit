# 🚀 Go Concurrency Primitives (Custom Implementation)

## 📌 Overview

This project demonstrates **custom implementations of Go concurrency primitives** such as Channels, Mutex, and WaitGroup.

The goal is to deeply understand how Go concurrency works internally by rebuilding core synchronization mechanisms from scratch.

Instead of only using `sync` package utilities, this project focuses on **how they can be designed and implemented internally**.

---

## 🎯 Objectives

- Understand Go concurrency model deeply
- Implement synchronization primitives manually
- Explore race conditions and thread safety
- Learn goroutine coordination patterns
- Build strong system design fundamentals

---

## 🧠 Concepts Implemented

### 🔹 Custom Channel
- Blocking send and receive behavior
- Buffered and unbuffered simulation
- Goroutine synchronization
- Channel close handling

### 🔹 Custom Mutex
- Lock / Unlock mechanism
- Race condition prevention
- Basic spinlock / blocking behavior concepts

### 🔹 Custom WaitGroup
- Counter-based goroutine tracking
- Wait until all goroutines complete
- Safe concurrent increment/decrement

---

## 🛠️ Tech Stack

- Go (Golang)
- Goroutines
- Low-level concurrency concepts
- Standard library (`sync`, `runtime` concepts for reference)

---
