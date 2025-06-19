const websocket = new WebSocket("ws://localhost:8201/ws")



/* WebSocket actions*/
websocket.onopen = () => {
    console.log('Connected to WebSocket');
    websocket.send('Request for canvas update');
};

websocket.onmessage = async(event) => {
    console.log('Message received:', event.data);

    const arrayBuffer = await event.data.arrayBuffer();
    const byteArray = new Uint8Array(arrayBuffer);

    console.log(byteArray);

    var graphValues = bytesToFloat(byteArray);

    console.log(graphValues);
    loadGraph(graphValues);
};

websocket.onerror = (error) => {
    console.error('WebSocket error:', error);
};

/*Should take a list of values to map to the canvas*/
const bytesToFloat = (byteArray: Uint8Array): number[] => {
    var arr: number[] = [];
    const dataView = new DataView(new ArrayBuffer(8));

    for (let i = 0; i < byteArray.length; i += 8) {
      for (let j = 0; j < 8; j++) {
        dataView.setUint8(j, byteArray[i + j]);
      }
      arr.push(dataView.getFloat64(0, true));
    }
  
    return arr;
}

const loadGraph = (values: number[]) => {
    try{
        var canvas = setupCanvas("canvas");
        var ctx = setupContext(canvas);
        drawBackGround(canvas, ctx);
        console.log(values);
    } catch(error){
        console.log(error);
    }
}

const setupCanvas = (id: string): HTMLCanvasElement => {
    var canvas = document.getElementById(id) as HTMLCanvasElement;
    if (!canvas){
        throw new Error("Could not get canvas.");
    }
    return canvas;
}

const setupContext = (canvas: HTMLCanvasElement): CanvasRenderingContext2D => {
    var ctx = canvas.getContext("2d");
    if (!ctx){
        throw new Error("Could not get context.");
    }
    return ctx;
}

const drawBackGround = (canvas: HTMLCanvasElement, ctx: CanvasRenderingContext2D) => {
    ctx.beginPath();
    ctx.fillStyle = "#a0a0a0";
    ctx.lineWidth = 5;
    ctx.moveTo(0, canvas.height/2);
    ctx.lineTo(canvas.width, canvas.height/2);
    ctx.stroke();

    ctx.moveTo(canvas.width/2, 0);
    ctx.lineTo(canvas.width/2, canvas.height);

    ctx.stroke();
}

