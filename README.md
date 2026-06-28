Steps to setup the starter template

Clone the project
git clone https://github.com/singhsanket143/Express-Typescript-Starter-Project.git <ProjectName>
Move in to the folder structure
cd <ProjectName>
Install npm dependencies
npm i
Create a new .env file in the root directory and add the PORT env variable
echo PORT=3000 >> .env
Start the express server
npm run dev
Room Availability Extension Scheduler

The HotelService includes an automated room availability extension scheduler that runs every minute to ensure continuous room availability.

How it works

Automatic Extension: The scheduler runs every minute using a cron job (* * * * *)
Latest Date Detection: For each room category, it finds the latest date of availability
One Day Extension: It automatically adds one more day to the latest availability date
Queue Processing: New room instances are added to the Redis queue for processing
Duplicate Prevention: The system checks for existing rooms before creating new ones
API Endpoints

Start Scheduler

POST /api/v1/scheduler/start
Starts the room availability extension scheduler.

Stop Scheduler

POST /api/v1/scheduler/stop
Stops the room availability extension scheduler.

Get Scheduler Status

GET /api/v1/scheduler/status
Returns the current status of the scheduler.

Manual Extension

POST /api/v1/scheduler/extend
Manually triggers room availability extension for all room categories.

Example Usage

# Start the scheduler
curl -X POST http://localhost:3000/api/v1/scheduler/start

# Check scheduler status
curl -X GET http://localhost:3000/api/v1/scheduler/status

# Manually trigger extension
curl -X POST http://localhost:3000/api/v1/scheduler/extend

# Stop the scheduler
curl -X POST http://localhost:3000/api/v1/scheduler/stop
Configuration

The scheduler automatically starts when the server initializes. The cron expression is set to run every minute (* * * * *) and uses UTC timezone.

Logging

The scheduler provides detailed logging for:

Scheduler start/stop events
Room category processing
Job queue additions
Error handling and recovery
