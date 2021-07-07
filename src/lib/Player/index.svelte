<script>
  import { onMount } from "svelte";
  import Audio from "$lib/Audio/index.svelte";
  import utils from "./utils.js";

  export let solution = undefined;
  let canvas;
  let canvasContext;
  let canvasTween;
  let worker;
  let tween;
  let data = { xml: undefined, info: {}, objects: [] };

  onMount(() => {
    canvasContext = canvas.getContext("2d");
    worker = new pdfjsLib.PDFWorker("worker");
  });

  const onSolutionChange = () => {
    // CLEAR CURRENT SOLUTION PROPERTIES
    canvas.style.background = "";
    canvasContext.clearRect(0, 0, canvas.width, canvas.height);
    data = { xml: undefined, info: {}, objects: [] };

    // FETCH XML AND AUDIO
    fetch(solution.xml)
      .then((res) => res.text())
      .then((res) => {
        // XML Decrypt
        res = res.split("");
        res.splice(10, 1);
        res.splice(20, 1);
        let xml = atob(res.join("")).replaceAll(",", ".");

        return xml;
      })
      .then((xml) => {
        data = utils.parse_fernus_xml(xml);

        console.log("data.objects", data.objects);
      })
      .then(() => {
        // FETCH PDF AND RENDER
        fetch(solution.pdf)
          .then((res) => res.arrayBuffer())
          .then((data) => pdfjsLib.getDocument({ data, worker }).promise)
          .then((pdf) => pdf.getPage(1))
          .then((page) => {
            data.info.questionWidth = page.view[2];
            data.info.questionHeight = page.view[3];

            let resolution = window.devicePixelRatio; // * 2;

            let playerRect = canvas.parentNode.getBoundingClientRect(); // .player rect
            let currentViewport = page.getViewport({ scale: 1 });

            // Fit canvas as much as possible
            let scale = Math.min(
              playerRect.width / currentViewport.width,
              playerRect.height / currentViewport.height
            );
            let viewport = page.getViewport({ scale: scale });

            // Render at a higher resolution but show at a lower resolution -> sharper
            canvas.width = resolution * viewport.width;
            canvas.height = resolution * viewport.height;

            canvas.style.width = viewport.width + "px";
            canvas.style.height = viewport.height + "px";

            return page.render({
              canvasContext,
              viewport,
              transform: [resolution, 0, 0, resolution, 0, 0], // force it bigger size
              background: "transparent",
            }).promise;
          })
          .then(() => {
            let solutionPNG = canvas.toDataURL("image/png");
            // TODO: CLEAN
            // IDK After here, i just translated from fernus's source

            let playerRect = canvas.parentNode.getBoundingClientRect(); // .player rect
            let playerScale = Math.min(
              playerRect.width / data.info.solutionWidth,
              playerRect.height / data.info.solutionHeight
            );

            var _infoPaddingX = data.info.paddingLeft * playerScale;
            var _infoPaddingY = data.info.paddingTop * playerScale;
            var _infoSolutionW = data.info.solutionWidth * playerScale;
            var _infoSolutionH = data.info.solutionHeight * playerScale;
            var _x = data.info.questionWidth * data.info.scale * playerScale;
            var _y = data.info.questionHeight * data.info.scale * playerScale;

            canvas.width = _infoSolutionW;
            canvas.height = _infoSolutionH;

            canvas.style.backgroundImage = "url('" + solutionPNG + "')";
            canvas.style.backgroundPosition = `${_infoPaddingX}px ${_infoPaddingY}px`;
            canvas.style.backgroundSize = `${_x}px ${_y}px`;
            canvas.style.backgroundRepeat = "no-repeat";

            canvasContext.scale(playerScale, playerScale);
          });
      });
  };

  function canvasObjectAdd(object, _status, _addObject) {
    if (canvasTween) {
      canvasTween.clear();
      canvasTween.kill();
      canvasTween = null;
    }
    switch (object.type) {
      case "line":
        let position = { x: object.points[0].x, y: object.points[0].y };

        canvasContext.beginPath();
        canvasContext.name = object.id;
        canvasContext.lineWidth = object.size;
        canvasContext.strokeStyle = object.color;
        canvasContext.globalAlpha = object.highlight == "false" ? 1 : 0.5;
        canvasContext.globalCompositeOperation = "source-over";
        canvasContext.lineJoin = "round";
        canvasContext.lineCap = "round";
        // canvasContext.fillStyle = "none";
        canvasContext.moveTo(position.x, position.y);

        if (!_status) {
          canvasTween = new TimelineLite();
          canvasTween.add(
            TweenLite.to(position, object.duration, {
              bezier: {
                curviness: 0,
                values: object.points,
                autoRotate: false,
              },
              ease: Linear.easeNone,
              onUpdate: () => {
                canvasContext.lineTo(position.x, position.y);
                canvasContext.stroke();
              },
              onComplete: () => {
                canvasContext.closePath();
                canvasContext.stroke();
                // removeAfter(_objectArray);
              },
            })
          );
          // _objectArray.push(object);
        } else {
          fastShape(object, false);
        }
        //console.log('line');
        break;
      case "eraser":
        addEraser(object, _status);
        //console.log('eraser');
        break;
      case "arrow":
        addArrow(object, _status);
        //console.log(_object);
        break;
      case "triangle":
        addTriangle(object, _status);
        //console.log('triangle');
        break;
      case "rectangle":
        addRectangle(object, _status);
        //console.log('rectangle');
        break;
      case "circle":
        addCircle(object, _status);
        //console.log('circle');
        break;
      case "delete":
        removeShape(object);
        //console.log('delete');
        break;
      case "add":
        fastShape(_addObject, _status);
        //console.log(_addObject,1);
        break;
      case "scale":
        changeScale(object);
        break;
      case "swf":
        changeSWF(object);
        break;
    }
  }

  let onTimeUpdate = (event) => {};
  let onPlaying = (event) => {
    let audio = event.detail.target;
    tween = new TweenLite(audio, audio.duration, {
      onUpdate: () => {
        for (let i = 0; i < data.objects.length; i++) {
          let object = data.objects[i];
          if (!object.status && object.startTime <= audio.currentTime) {
            object.status = true;
            if (object.type == "add") {
              alert("UNTESTED");
              continue;
              let addObject = data.objects.find((e) => {
                e.id === object.objectID;
              });
              console.log(addObject);
              // _object.status = true;
              // canvas_addNewProcess(object, true, _object);
            } else {
              console.log("ADD", object.duration);
              canvasObjectAdd(object);
            }
          }
        }
      },
    });

    if (canvasTween) canvasTween.play();
  };

  let onPause = () => {
    tween.pause();
    tween.onUpdate = undefined;

    if (canvasTween) canvasTween.pause();
  };

  $: typeof window === "object" && solution && onSolutionChange();
</script>

<div class="player">
  <div class="video">
    <canvas bind:this={canvas} />
  </div>
  <Audio
    src={solution && solution.audio}
    on:timeupdate={onTimeUpdate}
    on:playing={onPlaying}
    on:pause={onPause}
  />
</div>

<style>
  .player {
    width: 100%;
    height: 100%;
  }

  .player .video {
    display: flex;
    justify-content: center;
    align-items: center;

    height: calc(100% - var(--header-height));
  }
</style>
