<script>
  import { onMount } from "svelte";
  import Audio from "$lib/Audio/index.svelte";

  export let solution = undefined;
  let canvas;
  let canvasContext;
  let audio;
  let worker;
  let data = {
    xml: undefined,
    info: {},
    objects: [],
  };

  onMount(() => {
    canvasContext = canvas.getContext("2d");
    worker = new pdfjsLib.PDFWorker("worker");
  });

  const onSolutionChange = () => {
    // CLEAR CURRENT SOLUTION PROPERTIES
    canvas.style.background = "";
    canvasContext.clearRect(0, 0, canvas.width, canvas.height);
    data = {
      xml: undefined,
      info: {},
      objects: [],
    };

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
        // XML Parse
        data.xml = new DOMParser().parseFromString(xml, "text/xml");
        data.xml = data.xml.childNodes[0];

        // XML Parse Info
        // [w, h, s, pX, pY, unknown]
        data.info._raw = data.xml.childNodes[0].childNodes[0].data.split("|");
        data.info.solutionWidth = data.info._raw[0];
        data.info.solutionHeight = data.info._raw[1];
        data.info.scale = data.info._raw[2];
        data.info.paddingLeft = data.info._raw[3];
        data.info.paddingTop = data.info._raw[4];

        // XML Parse Objects
        data.objects = [];
        for (let i = 1; i < data.xml.childNodes.length; i++) {
          let element = data.xml.childNodes[i];
          var object = {
            id: element.getAttribute("name"),
            type: element.getAttribute("act"),
            startTime: parseFloat(element.getAttribute("t1") / 1000),
            duration: parseFloat(element.getAttribute("t2") / 1000),
            color: "#" + element.getAttribute("color").slice(2),
            highlight: element.getAttribute("hl"),
            size: parseFloat(element.getAttribute("thc")),
            status: false,
          };

          if (
            object.type == "line" ||
            object.type == "eraser" ||
            object.type == "arrow" ||
            object.type == "triangle" ||
            object.type == "rectangle"
          ) {
            object.points = [];
            for (let i = 0; i < element.childNodes.length; i++) {
              let point = element.childNodes[i].childNodes[0].data.split("|");
              object.points.push(point);
            }
          } else if (object.type == "circle") {
            // [x, y, w, h]
            object.rectangle =
              element.childNodes[0].childNodes[0].data.split("|");
            object.rectangle = object.rectangle.map((e) => parseFloat(e));
          } else if (
            object.type == "delete" ||
            object.type == "add" ||
            object.type == "scale" ||
            object.type == "swf"
          ) {
            // TODO here
            // alert("NOT IMPLEMENTED YET");
            // object.objectID = element.find("id").text();
          }

          data.objects.push(object);
        }

        data.objects.sort(function (a, b) {
          return a.startTime - b.startTime;
        });

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

            let scale =
              (Math.min(
                playerRect.width / currentViewport.width,
                playerRect.height / currentViewport.height
              ) /
                100) *
              95; // 95% scale
            let viewport = page.getViewport({ scale: scale });

            // Render at a higher resolution but show at a lower resolution
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

            // canvasContext.scale(playerScale, playerScale);
          });
      });
  };

  $: typeof window === "object" && solution && onSolutionChange();
</script>

<div class="player">
  <div class="video">
    <canvas bind:this={canvas} />
  </div>
  <Audio src={solution && solution.audio} bind:this={audio} />
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
