# Project Overview

## Purpose
This is a deliberately vulnerable SNS (Social Networking Service) application similar to Twitter/X, created for security learning purposes. It contains intentional security vulnerabilities for educational purposes and should NOT be used in production environments.

## Key Features
- Post viewing (no login required)
- User login functionality
- Create, edit, delete posts (logged-in users only)
- Profile editing (nickname changes)
- Like functionality
- Hashtag search

## Vulnerability Categories
The application includes intentional vulnerabilities for learning:
- SQL Injection (login, posts)
- Cross-Site Scripting (XSS) in posts and profiles
- OS Command Injection in avatar upload
- Authentication token exposure
- Password logging

## Test Accounts
- alice/alice
- bob/bob
- charlie/charlie

## Service URLs
- Frontend: http://localhost:5173
- Backend API: http://localhost:9090
- API endpoints prefix: /api

## Database Access
- User: root
- Password: password
- Database: vuln_app
- Character set: utf8mb4

## Important Notes
- Vulnerabilities are documented in docs/vulnerabilities.md
- This is a learning tool for security education
- DO NOT fix vulnerabilities unless explicitly asked
- DO NOT add security measures unless requested