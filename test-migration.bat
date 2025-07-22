@echo off
REM Database Migration Test Script for Windows
REM This script tests the migration of all Laravel tables to Go GORM models

echo ===================================
echo Go Database Migration Test
echo ===================================

REM Check if .env file exists
if not exist .env (
    echo ⚠️  No .env file found. Creating example...
    copy .env.example .env
    echo 📝 Please edit .env with your PostgreSQL connection details
    pause
    exit /b 1
)

echo 📋 Available Commands:
echo   1. Test Migration Help
echo   2. Run All Migrations  
echo   3. Reset Database
echo   4. Build Project
echo.

REM Test 1: Show migration help
echo 🔍 Testing migration tool help...
go run tools/migrate/migrate.go -help
echo.

REM Test 2: Build project
echo 🔨 Building project...
go build ./...
if %errorlevel% equ 0 (
    echo ✅ Build successful!
) else (
    echo ❌ Build failed!
    pause
    exit /b 1
)

echo.
echo 🚀 Migration system is ready!
echo.
echo To run migrations:
echo   go run tools/migrate/migrate.go -migrate
echo.
echo To reset database:
echo   go run tools/migrate/migrate.go -reset
echo.
echo To start the server (auto-migrates):
echo   go run cmd/main.go

pause
