
### To Start

1. Clone this repository, and remove the remote origin
2. Follow the instruction in REAME.md under both /app and /server to set up the project


### Project Structure
* /app       - a desktop client application (reactJS, electron.js) 
* /server    - a local server running at 127.0.0.1:8080 (golang)

### Work flow
1. When client loads, a websocket connection is initiated. The server then assigns a UUID to that client, and sends it.
2. Client then requests a task to perform. Server retrieves the first unfinished task from database, and sends it back to the client.  (Task: listen to a peice of audio and input correct transcript)
3. Client inspect the origina data and prediction from ML model, and inputs the correction.
4. Client sends the correction back to the server, and server updates the database.  

### Todo
1. (App) Displays the client ID on the top left corner of the window
2. (App) Creates a button to send a websocket(ws) message to server which initiates a request

   (Server) Receives the request, and retrieves the first unfinished task (id, url to the original audio, prediction) from database.

   (Server) Sends the task info back to client
3. (App) Displays the task info (audio, prediction), and creates a text input field to allow wizard to type in correct transcript. 
4. (App) Create a button to send the correct transcript back to the server (through ws)

   (Server) Receives the correct transcript, and update the task table in database

### Database
[ElephantSQL](https://api.elephantsql.com/console/962e8cd4-3633-4335-996d-72ef1727d80c/details)

Task

| id  | url | prediction  | correct_transcript | completion  | 
| ------------- | ------------- | ------------- | ------------- | ------------- |
| int  | string | string  | string | boolean  | 

