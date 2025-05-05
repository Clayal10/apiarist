const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");
const websocket = new WebSocket("ws://localhost:8201/ws")

const drawBackGround = () => {
    ctx.beginPath();
    ctx.fillStyle = "#a0a0a0";
    ctx.lineWidth = 5;
    ctx.moveTo(0, canvas.height/2);
    ctx.lineTo(canvas.width, canvas.height/2);
    ctx.stroke();

    ctx.moveTo(canvas.width/2, 0);
    ctx.lineTo(canvas.width/2, canvas.height);

    ctx.stroke()
}

drawBackGround();

websocket.onopen = () => {
    console.log('Connected to WebSocket');
    websocket.send('Request for canvas update');
};

websocket.onmessage = (event) => {
    console.log('Message received:', event.data);
    
};

websocket.onerror = (error) => {
    console.error('WebSocket error:', error);
};

const loadGraph = () => {
    drawBackGround();
    // Get values and draw them.
}
