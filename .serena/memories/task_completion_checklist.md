# Task Completion Checklist

When completing a task in this project, follow these steps:

## Frontend Tasks
1. **Type Checking**: Run `npm run type-check` in the frontend directory
2. **Linting**: Run `npm run lint` to check and fix linting issues
3. **Formatting**: Run `npm run format` to ensure consistent code style
4. **Build Test**: Run `npm run build` to ensure the production build works
5. **Manual Testing**: Test the feature in the browser at http://localhost:5173

## Backend Tasks
1. **Compilation**: The Go code should compile without errors when Docker rebuilds
2. **Manual Testing**: Test API endpoints using curl or the frontend
3. **Docker Logs**: Check `docker-compose logs backend` for any runtime errors

## General Tasks
1. **Git Status**: Check `git status` to see all changes
2. **Docker Services**: Ensure all services are running with `docker-compose ps`
3. **Database Changes**: If schema changed, may need to run `make reset` to reinitialize
4. **Documentation**: Update relevant documentation if the change affects usage

## Important Notes
- The project contains intentional vulnerabilities - DO NOT fix them unless asked
- No automated tests are configured in this project
- Backend has no linting/formatting commands defined
- Always test changes through the running Docker containers
- If type checking or linting commands are missing, ask the user for the correct commands