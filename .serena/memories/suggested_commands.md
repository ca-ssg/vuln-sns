# Suggested Commands

## Docker & Infrastructure Commands
```bash
# Start all services
docker-compose up -d
# OR
make start

# Stop all services
docker-compose down
# OR
make stop

# Rebuild and restart services
make restart

# Connect to MySQL database
docker-compose exec db mysql -uroot -ppassword -Dvuln_app --default-character-set=utf8mb4
# OR
make db

# Reset everything (remove volumes and restart)
make reset

# Clear all containers and volumes
make clear

# Build containers
docker-compose build
# OR
make build
```

## Frontend Development (in frontend/ directory)
```bash
# Install dependencies
npm install

# Start development server with hot reload
npm run dev

# Build for production
npm run build

# Type checking
npm run type-check

# Linting
npm run lint

# Format code
npm run format

# Preview production build
npm run preview
```

## System Commands (Darwin/macOS)
```bash
# Git commands
git status
git diff
git log
git add .
git commit -m "message"

# Directory navigation
cd <directory>
ls -la  # List with colors (aliased)

# Search commands
grep <pattern> <files>  # Has color and exclude dirs
find <path> -name <pattern>

# Docker logs
docker-compose logs -f <service>
docker-compose ps
```

## API Testing
```bash
# Health check
curl http://localhost:9090/api/health

# Login (example)
curl -X POST http://localhost:9090/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"alice"}'
```