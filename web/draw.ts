const valueAPI = "/graph/";
const userAPI = "/submit/";

const startTraining = "Start Training";
const training = "Training..."

interface UserInput {
  Inertia: number;
  CogCoef: number;
  SocCoef: number;
}
interface GraphData {
  Data: number[];
}

window.onload = (event) => {
  try {
    const canvas = setupCanvas("canvas");
    const ctx = setupContext(canvas)
    drawBackGround(canvas, ctx);
  } catch (error) {
    console.error(error)
  }
};

async function sendInfo() {
  try {
    const user: UserInput = {
      Inertia: getElementValue("inertia-value"),
      CogCoef: getElementValue("cog-coef-value"),
      SocCoef: getElementValue("soc-coef-value")
    };

    const response = await fetch(userAPI, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
      body: JSON.stringify(user)
    });

    if (!response.ok) {
      const err = await response.json();
      throw new Error(err.message);
    }
    const button = document.getElementById("rest-button-send") as HTMLButtonElement;
    button.innerText = training;
  } catch (e) {
    console.error("Error in parsing html input: ", e);
    return;
  }
}

async function loadAndGraph() {
  try {
    const data = await getData();
    console.log(data);
    if (data) {
      loadGraph(data);
    }
    const button = document.getElementById("rest-button-send") as HTMLButtonElement;
    button.innerText = startTraining;
  } catch (e) {
    console.error("Error in retreiving data: ", e);
    return;
  }
};

async function getData(): Promise<number[]> {
  const response = await fetch(valueAPI);
  if (!response.ok) {
    throw new Error(`Error code: ${response.status}`);
  }
  const data = await (response.json()) as GraphData;
  return data.Data;
}


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

const clearGraph = () => {
  const canvas = setupCanvas("canvas");
  const ctx = setupContext(canvas);
  ctx.beginPath();
  ctx.fillStyle = '#F7FBFF';
  ctx.rect(0, 0, canvas.width, canvas.height)
  ctx.fill();
  ctx.closePath();
  drawBackGround(canvas, ctx)
};

// Main operating logic
const loadGraph = (values: number[]) => {
  try {
    const canvas = setupCanvas("canvas");
    const ctx = setupContext(canvas);
    drawBackGround(canvas, ctx);
    drawGraph(canvas, ctx, values)
  } catch (error) {
    console.log(error);
  }
};

const drawGraph = (canvas: HTMLCanvasElement, ctx: CanvasRenderingContext2D, values: number[]) => {
  const interval = canvas.width / values.length;
  /* Graph of the canvas
    0 100 200 300 400 500 600 700 800  
  100
  200
  300
  400
  500
  600
  */

  ctx.beginPath()
  ctx.lineWidth = 1;
  ctx.strokeStyle = '#08306B';
  let pos = 0
  for (let i = 0; i < values.length - 1; i++) {
    ctx.moveTo(pos, rangeCalc(values[i], canvas.height));
    ctx.lineTo(pos + interval, rangeCalc(values[i + 1], canvas.height));
    ctx.stroke();
    pos += interval
  }
  ctx.closePath();
};

// We'll get values between -1 and 1, and we need to map those to:
//  -1 = canvas.height (600)
//   1 = 0
// Function will take in the value and the canvas height (since it could be dynamic),
// and return the canvas height pos needed.
const rangeCalc = (val: number, height: number): number => {
  const zero = height / 2; // need to move above or below
  const scale = height / 2;
  if (val === 0) {
    return zero
  }

  var diff = Math.abs(val) * scale
  if (val < 0) {
    return zero + diff
  }
  return zero - diff
}

const setupCanvas = (id: string): HTMLCanvasElement => {
  var canvas = document.getElementById(id) as HTMLCanvasElement;
  if (!canvas) {
    throw new Error("Could not get canvas.");
  }
  return canvas;
}

const setupContext = (canvas: HTMLCanvasElement): CanvasRenderingContext2D => {
  var ctx = canvas.getContext("2d");
  if (!ctx) {
    throw new Error("Could not get context.");
  }
  return ctx;
}

const drawBackGround = (canvas: HTMLCanvasElement, ctx: CanvasRenderingContext2D) => {
  ctx.beginPath();
  ctx.strokeStyle = '#010101';
  ctx.lineWidth = 3;
  ctx.moveTo(0, canvas.height / 2);
  ctx.lineTo(canvas.width, canvas.height / 2);
  ctx.stroke();

  ctx.moveTo(canvas.width / 2, 0);
  ctx.lineTo(canvas.width / 2, canvas.height);

  ctx.stroke();
  ctx.closePath();

  const interval = canvas.width / 1000;
  ctx.beginPath();
  ctx.lineWidth = 1;
  let pos = 0;
  ctx.strokeStyle = '#888888';
  var piInterval = (Math.PI * 6) / 1000;
  for (let pi = -Math.PI * 3; pi < Math.PI * 3; pi += piInterval) {
    // Draw Real sine wave
    ctx.moveTo(pos, rangeCalc(Math.sin(pi), canvas.height));
    ctx.lineTo(pos + interval, rangeCalc(Math.sin(pi + piInterval), canvas.height));
    ctx.stroke();
    pos += interval
  }
  ctx.closePath()
}

const getElementValue = (id: string): number => {
  const elem = document.getElementById(id) as HTMLInputElement;
  if (!elem) {
    throw new Error("Could not get element");
  }
  return Number(elem.value);
}