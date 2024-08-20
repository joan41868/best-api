const WebSocket = require('ws');

const ws = new WebSocket('ws://localhost:8080/ws');

ws.on('open', () => {
  console.log('Connected to WebSocket server');
});

ws.on('message', (data) => {
  console.log('Received message:', new String(data).toString());
});
