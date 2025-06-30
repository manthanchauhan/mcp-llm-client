
# About This Project

This project is a command-line interface (CLI) application written in Go that interacts with a Large Language Model (LLM). It's designed to be a conversational AI that can understand user requests, manage session information, and potentially interact with external tools.

## Key Components

*   **`main.go`**: The entry point of the application. It initializes the connection to the LLM and starts the interactive chat session.
*   **`cli` package**: Handles the user-facing command-line interface, including reading user input and displaying AI responses.
*   **`llm` package**: Manages all communication with the LLM. It's responsible for sending user messages, handling system-level instructions, and processing the LLM's responses. It connects to a local LLM server, likely running a model like Phi-3.
*   **`sessionmanager` package**: Creates and manages user sessions. It stores session-specific data, such as conversation history and extracted user information.
*   **`thinktank` package**: This is the core logic unit of the application. It orchestrates the conversation flow, working to understand the user's intent (e.g., booking a new loan, closing an existing one). It uses the `llm` package to achieve this.
*   **`tool` package**: Defines a set of tools that the AI can potentially use. Currently, it includes a tool for retrieving user information based on a mobile number.
*   **`go.mod`**: Defines the project's dependencies, which include `godotenv` for managing environment variables and `lumberjack` for logging.

## How it Works

1.  The application starts by loading environment variables and establishing a connection with the LLM.
2.  A new chat session is created, and the `thinktank` initiates the conversation.
3.  The user interacts with the application through the CLI.
4.  The `thinktank` processes the user's input, attempting to identify the user's needs and extract relevant information (like their mobile number).
5.  The `llm` package sends and receives messages from the LLM.
6.  The `sessionmanager` keeps track of the conversation and any data collected during the session.
7.  The AI's response is then displayed to the user in the CLI.

The ultimate goal of the application appears to be to serve as an intelligent assistant that can understand and fulfill user requests, potentially by integrating with other systems through the defined tools.
