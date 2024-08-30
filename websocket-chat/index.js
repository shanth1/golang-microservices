const socket = new WebSocket("ws://localhost:3000/ws");

socket.onmessage = (event) => console.log("response:", event.data);
socket.send("hi");
