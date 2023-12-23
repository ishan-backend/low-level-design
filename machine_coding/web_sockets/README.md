### Building a web chat application and dump real time data to the subscribers feed with websockets

**Inspect web browser - multiple tabs - create clients/players**:
```javascript
let socket = new WebSocket("ws://localhost:3000/ws/chat")
// undefined
socket.onmessage = (event) => {console.log("message received from server : ", event.data)}
// (event) => {console.log("message received from server : ", event.data)}
socket.send("hello from client")
// undefined
// message received from server :  thank you for the message!
```

**Handle Document Upload Percentage Relay**:
```javascript
let socket = new WebSocket("ws://localhost:3000/ws/document-upload-perc")
socket.onmessage = (event) => {console.log("message received from server : ", event.data)}
```
