# 🛡️ Go Password Manager

## Overview

This is a simple command-line password manager written in Go. It lets you securely generate and retrieve passwords for websites and usernames. The passwords are stored locally in a `passwords.json` file using base64-encoded cryptographically secure random strings.

---

## 🧑‍💻 Usage

### Build and Run

```bash
make
./main set | get
