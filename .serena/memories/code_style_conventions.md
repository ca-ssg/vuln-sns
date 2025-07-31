# Code Style and Conventions

## Frontend (Vue/TypeScript)
- **Framework**: Vue 3 Composition API with TypeScript
- **Component Style**: Single File Components (.vue)
- **Naming**: PascalCase for components, camelCase for functions/variables
- **State Management**: Pinia stores with TypeScript
- **Styling**: Combination of Quasar components, Tailwind utility classes, and scoped SCSS
- **API Calls**: Axios with base URL from VITE_API_URL environment variable
- **Type Safety**: TypeScript with strict type checking enabled
- **Code Formatting**: Prettier configuration in .prettierrc.json
- **Linting**: ESLint with Vue plugin

## Backend (Go)
- **Package Structure**: Standard Go module layout
- **Error Handling**: Direct error returns, logged but not always properly handled (intentionally vulnerable)
- **Database**: Raw SQL queries with string concatenation (intentionally vulnerable)
- **API Response**: JSON format with gin.H{} maps
- **Middleware**: Custom authentication middleware
- **Naming**: Go conventions (exported/unexported based on capitalization)
- **No Go-specific linting/formatting commands defined in project**

## General Conventions
- **Comments**: Minimal commenting (avoid adding unless requested)
- **Security**: Intentionally vulnerable - DO NOT fix unless asked
- **File Organization**: Feature-based organization in both frontend and backend
- **Environment Variables**: Used for configuration (API URLs, DB credentials)
- **Docker**: All services run in containers, no local development setup