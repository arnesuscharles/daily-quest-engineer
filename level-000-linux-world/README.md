# Level 000 - Linux World

## Quest
Entering the Linux world using WSL.

## Objectives
- Install WSL2
- Install Ubuntu
- Open Linux terminal
- Learn basic Linux commands
- Connect VS Code to WSL

## Commands Learned

### Check current directory
```bash
pwd
```
---

# GitHub SSH Authentication

## Problem Encountered

When pushing to GitHub using HTTPS:

```bash
git push -u origin main
```

Error:
```
remote: Invalid username or token.
Password authentication is not supported for Git operations.
```
## Solution

Using SSH authentication instead of password authentication.

Generate SSH Key
```
ssh-keygen -t ed25519 -C "your_email@example.com"
```
Start SSH Agent
```
eval "$(ssh-agent -s)"
```
Add SSH Key
```
ssh-add ~/.ssh/id_ed25519
```
Show Public Key
```
cat ~/.ssh/id_ed25519.pub
```
Add SSH Key to GitHub
```
GitHub Settings → SSH and GPG Keys.
```
Test Connection
```
ssh -T git@github.com
```
Change Git Remote to SSH
```
git remote set-url origin git@github.com:USERNAME/daily-quest-engineer.git
```
## What I Learned
- Difference between HTTPS and SSH authentication
- SSH public/private key concepts
- GitHub secure authentication
- SSH agent usage
- Remote repository configuration