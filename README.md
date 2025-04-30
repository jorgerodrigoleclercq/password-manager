# 🛡️ Go Password Manager

## Overview

This is a simple command-line password manager written in Go. 
It lets you securely generate and retrieve passwords for websites and usernames. 
The passwords are stored locally in a `passwords.json` file using base64-encoded cryptographically secure random strings. 
Passwords are stored as values for keys made up of the website and the username the password is for:

```json
{
  "foo.com|bar@foo.com": "2NwkBKlLYYmR3Uns_tQzxne2vaYElQ==",
  "bar.com|foorodrigoleclercq": "EGPjL4sODnh1cGnCEVjc7lFjZQrgHQ=="
}
```

---

## 🧑‍💻 Usage

### Build and Run

```bash
make
./main set | get
```

## Potential follow-ups
- Storing the passwords in a database such as MongoDB
- Creating a browser extension
