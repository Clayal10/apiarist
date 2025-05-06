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

function bytesToFloat(byteArray) {
    const float64Array = [];

    const dataView = new DataView(new ArrayBuffer(8));

    for (let i = 0; i < byteArray.length; i += 8) {
      for (let j = 0; j < 8; j++) {
        dataView.setUint8(j, byteArray[i + j]);
      }
      float64Array.push(dataView.getFloat64(0, false)); // Big Endian?
    }
  
    return float64Array;
}

/* WebSocket actions*/
websocket.onopen = () => {
    console.log('Connected to WebSocket');
    websocket.send('Request for canvas update');
};

websocket.onmessage = (event) => {
    console.log('Message received:', event.data);

    if (event.data instanceof Blob) {
        const reader = new FileReader();

        reader.onload = () => {
            const arrayBuffer = reader.result;
            const uint8Array = new Uint8Array(arrayBuffer);
            console.log('Received data:', uint8Array);
        }
    }

    const arrayBuffer = event.data.arrayBuffer();
    const byteArray = new Uint8Array(arrayBuffer);

    var graphValues = bytesToFloat(byteArray);

    console.log(graphValues);
};

websocket.onerror = (error) => {
    console.error('WebSocket error:', error);
};

/*Should take a list of values to map to the canvas*/
const loadGraph = () => {
    drawBackGround();
    // Get values and draw them.
}