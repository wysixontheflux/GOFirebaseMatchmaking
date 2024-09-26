!!!  THIS PROJECT IS UNDER DEVELOPMENT !!!

DEMO : https://youtu.be/Hdq7-OThoYg

# GoMatchmaking

GoMatchmaking is a matchmaking system built with Go and Firebase. It allows players to be added to a queue and automatically creates rooms when enough players are available.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/wysixontheflux/goMatchmaking.git
    cd goMatchmaking
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Set up Firebase:**

    - Place your Firebase credentials file (`codseries.json`) in the `firebase` directory.

## Usage

1. **Start the matchmaking system:**

    ```sh
    go run main.go
    ```

2. **Start the HTTP server:**

   The HTTP server will start automatically on port `8282`.

3. **(Optional) Start the WebSocket server:**

   Uncomment the relevant code in `server/ws.go` and start the WebSocket server.

## Project Structure

- `main.go`: Entry point of the application.
- `firebase/`: Contains Firebase initialization and helper functions.
- `matchmaking/`: Contains the matchmaking logic.
- `server/`: Contains the HTTP and WebSocket server implementations.
- `models/`: Contains the data models used in the project.

## Configuration

- **Firebase:**
    - Ensure your Firebase credentials file (`codseries.json`) is correctly placed in the `firebase` directory.
    - Update the Firebase database URL in `firebase.go` if necessary.

## Contributing

1. **Fork the repository.**
2. **Create a new branch:**

    ```sh
    git checkout -b feature-branch
    ```

3. **Make your changes and commit them:**

    ```sh
    git commit -m "Description of changes"
    ```

4. **Push to the branch:**

    ```sh
    git push origin feature-branch
    ```

5. **Create a pull request.**

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.