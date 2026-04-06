# Rapdev Portfolio Server — API Documentation

## Overview

GraphQL API server for portfolio data built with **GoFiber** + **graphql-go** + **PostgreSQL (raw query)**.

**Base URL:** `/graphql`
**Method:** `POST`
**Content-Type:** `application/json`
**Playground:** Available at `/` (GraphiQL)
**Health Check:** `GET /health`

---

## GraphQL Schema

### Scalar Types

| Type | Description |
|------|-------------|
| `String` | UTF-8 string |
| `Int` | 32-bit integer |
| `Boolean` | true/false |
| `DateTime` | ISO 8601 date-time string (RFC3339) |

### Types

```graphql
type User {
  id: String!
  name: String!
  email: String!
  role: String!
  image: String
  imageFileId: String
  createdAt: DateTime!
  updatedAt: DateTime!
}

type Article {
  id: String!
  title: String!
  content: String!
  path: String!
  viewCount: Int!
  likes: Int!
  authorId: String!
  published: Boolean!
  author: User
  createdAt: DateTime!
  updatedAt: DateTime!
}

type Education {
  id: String!
  institution: String!
  degree: String!
  logo: String
  startDate: DateTime!
  endDate: DateTime
  gpa: String
  achievements: [String]
  createdAt: DateTime!
  updatedAt: DateTime!
}

type Experience {
  id: String!
  company: String!
  position: String!
  type: String!
  logo: String
  startDate: DateTime!
  endDate: DateTime
  tags: [String]
  description: [String]
  createdAt: DateTime!
  updatedAt: DateTime!
}

type Project {
  id: String!
  title: String!
  subtitle: String
  description: String!
  image: String
  tags: [String]
  demoUrl: String
  githubUrl: String
  createdAt: DateTime!
  updatedAt: DateTime!
}

type Gallery {
  id: String!
  image: String!
  imageFileId: String!
  caption: String
  createdAt: DateTime!
  updatedAt: DateTime!
}

type SocialLink {
  id: String!
  title: String!
  url: String!
  order: Int!
  isActive: Boolean!
  createdAt: DateTime!
  updatedAt: DateTime!
}
```

### Query Definitions

```graphql
type Query {
  articles: [Article]
  educations: [Education]
  experiences: [Experience]
  projects: [Project]
  galleries: [Gallery]
  socialLinks: [SocialLink]
}
```

---

## Query Examples

### Get All Articles

**Request:**
```json
{
  "query": "{ articles { id title path viewCount likes published author { id name email image } createdAt updatedAt } }"
}
```

**Response:**
```json
{
  "data": {
    "articles": [
      {
        "id": "clx1abc123",
        "title": "Getting Started with Go",
        "path": "getting-started-with-go",
        "viewCount": 150,
        "likes": 42,
        "published": true,
        "author": {
          "id": "clx1user01",
          "name": "John Doe",
          "email": "john@example.com",
          "image": "https://example.com/avatar.jpg"
        },
        "createdAt": "2025-01-15T10:30:00Z",
        "updatedAt": "2025-01-20T08:00:00Z"
      }
    ]
  }
}
```

### Get All Projects

**Request:**
```json
{
  "query": "{ projects { id title subtitle description image tags demoUrl githubUrl createdAt } }"
}
```

**Response:**
```json
{
  "data": {
    "projects": [
      {
        "id": "clx2proj01",
        "title": "Portfolio Website",
        "subtitle": "Personal portfolio built with Next.js",
        "description": "A modern portfolio website with blog support.",
        "image": "https://example.com/project.jpg",
        "tags": ["Next.js", "TypeScript", "Tailwind"],
        "demoUrl": "https://portfolio.example.com",
        "githubUrl": "https://github.com/user/portfolio",
        "createdAt": "2025-02-01T12:00:00Z"
      }
    ]
  }
}
```

### Get All Experiences

**Request:**
```json
{
  "query": "{ experiences { id company position type logo startDate endDate tags description createdAt } }"
}
```

**Response:**
```json
{
  "data": {
    "experiences": [
      {
        "id": "clx3exp01",
        "company": "Tech Corp",
        "position": "Senior Software Engineer",
        "type": "FULLTIME",
        "logo": "https://example.com/logo.png",
        "startDate": "2023-01-01T00:00:00Z",
        "endDate": null,
        "tags": ["Go", "PostgreSQL", "GraphQL"],
        "description": ["Led backend architecture", "Improved API performance by 40%"],
        "createdAt": "2025-01-01T00:00:00Z"
      }
    ]
  }
}
```

### Get All Educations

**Request:**
```json
{
  "query": "{ educations { id institution degree logo startDate endDate gpa achievements createdAt } }"
}
```

### Get All Social Links

**Request:**
```json
{
  "query": "{ socialLinks { id title url order isActive } }"
}
```

**Response:**
```json
{
  "data": {
    "socialLinks": [
      {
        "id": "clx4sl01",
        "title": "GitHub",
        "url": "https://github.com/username",
        "order": 1,
        "isActive": true
      },
      {
        "id": "clx4sl02",
        "title": "LinkedIn",
        "url": "https://linkedin.com/in/username",
        "order": 2,
        "isActive": true
      }
    ]
  }
}
```

### Get All Galleries

**Request:**
```json
{
  "query": "{ galleries { id image imageFileId caption createdAt } }"
}
```

---

## Error Handling

Errors follow the standard GraphQL error format:

```json
{
  "data": null,
  "errors": [
    {
      "message": "error description",
      "locations": [{ "line": 1, "column": 2 }],
      "path": ["fieldName"]
    }
  ]
}
```

---

## Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `DATABASE_URL` | PostgreSQL connection string | Yes |
| `PORT` | Server port (default: 3000) | No |

---

## Running Locally

```bash
cp .env.example .env
# Edit .env and set your DATABASE_URL

# Using air (live-reload)
air

# Or using go run
go run cmd/main.go
```

Open the playground at `http://localhost:3000/`

---

## Deployment (Vercel)

The project is configured for Vercel deployment via `vercel.json`. Set `DATABASE_URL` in Vercel environment variables.
All requests are routed to `api/index.go` which exposes a standard `net/http` handler via GoFiber adaptor.
