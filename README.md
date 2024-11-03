# Quicknet

**Quicknet** is a web-based network speed testing application built using Go (backend) and React (frontend). The app measures internet speed by testing **latency (ping)**, **download speed**, and **upload speed**. It provides an intuitive interface for users to get accurate network metrics, similar to services like Speedtest.net.

## Features

- **Ping Test**: Measures the latency between the user and the server.
- **Download Speed Test**: Simulates a file download to calculate download speed in real-time.
- **Upload Speed Test**: Measures upload speed by simulating an upload from the user's machine to the server.
- **Responsive UI**: Built with React for a fast, responsive experience on all device sizes.

## Tech Stack

- **Frontend**: React, Axios
- **Backend**: Go (using `net/http` package)

## Architecture

Quicknet uses a client-server architecture:
- The **React frontend** makes API calls to the **Go backend** to initiate tests and display results.
- The **Go backend** hosts the endpoints to serve large data for download speed testing, receive upload data, and respond to latency requests.

## Getting Started

Follow these instructions to set up and run the project locally.

### Prerequisites

- [Go](https://golang.org/doc/install) 1.18 or higher
- [Node.js](https://nodejs.org/) 12 or higher
- [npm](https://www.npmjs.com/) (included with Node.js)

### Installation

#### 1. Clone the Repository

```bash
git clone https://github.com/BarakaCaleb/Quicknet.git
cd Quicknet

2. Set Up the Backend (Go)

Navigate to the backend folder and install Go dependencies.

bash

cd Quicknet-backend
go mod tidy

Run the Go server:

bash

go run main.go

The backend server should start on http://localhost:8080.
3. Set Up the Frontend (React)

Open a new terminal window, navigate to the frontend folder, and install npm dependencies.

bash

cd ../Quicknet-frontend
npm install

Start the React development server:

bash

npm start

The frontend server will start on http://localhost:3000.
Usage

    Open your browser and navigate to http://localhost:3000.
    Use the Ping Test, Download Test, and Upload Test buttons to measure different network metrics.
    View results in real-time on the interface.

API Endpoints

The Go backend provides the following endpoints:

    /ping: Responds with a simple "pong" to measure latency.
    /download: Streams a 50MB file to the client to simulate download speed.
    /upload: Receives a large data payload from the client to simulate upload speed and measures the time taken.

Project Structure

php

Speedlyzer
│
├── Quicknet-backend           # Go backend
│   ├── main.go                  # Main Go server file with endpoints
│   └── go.mod                   # Go module dependencies
│
└── Quicknet-frontend           # React frontend
    ├── public                   # Public assets
    ├── src
    │   ├── components           # React components
    │   ├── App.js               # Main App component
    │   ├── SpeedTest.js         # Speed test component
    │   └── index.js             # Entry point
    └── package.json             # NPM dependencies

Contributing

Contributions are welcome! To contribute:

    Fork the project.
    Create a new branch (git checkout -b feature/YourFeature).
    Commit your changes (git commit -m 'Add some feature').
    Push to the branch (git push origin feature/YourFeature).
    Open a pull request.

Future Enhancements

    Enhanced Error Handling: Add better error messages for connectivity issues.
    Deployment: Deploy the app to cloud services for global access.
    Real-Time Charting: Add visualizations for download/upload speed over time.
    Multi-Server Support: Allow users to choose test servers based on their location.

License

This project is licensed under the MIT License - see the LICENSE file for details.


