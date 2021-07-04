<script>
  import { onMount } from "svelte";

  export let solution = undefined;
  let canvas;
  let canvasContext;
  let worker;

  let audio;
  let audioDurationText = "";
  let data = {
    xml: undefined,
    info: {},
    objects: [],
  };

  onMount(() => {
    canvasContext = canvas.getContext("2d");
    worker = new pdfjsLib.PDFWorker("worker");
  });

  const calculateTime = (secs) => {
    const minutes = Math.floor(secs / 60);
    const seconds = Math.floor(secs % 60);
    const returnedSeconds = seconds < 10 ? `0${seconds}` : `${seconds}`;
    return `${minutes}:${returnedSeconds}`;
  };

  const onSolutionChange = () => {
    // CLEAR CURRENT SOLUTION PROPERTIES
    canvasContext.clearRect(0, 0, canvas.width, canvas.height);
    audio = undefined;
    audioDurationText = "";
    data = {
      xml: undefined,
      info: {},
      objects: [],
    };

    // FETCH PDF AND RENDER
    fetch(solution.pdf)
      .then((res) => res.arrayBuffer())
      .then((data) => pdfjsLib.getDocument({ data, worker }).promise)
      .then((pdf) => pdf.getPage(1))
      .then((page) => {
        let resolution = 2;

        let desiredSizes = canvas.parentNode.getBoundingClientRect(); // .player rect
        let currentViewport = page.getViewport({ scale: 1 });

        let scale =
          (Math.min(
            desiredSizes.width / currentViewport.width,
            desiredSizes.height / currentViewport.height
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
      });

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
        data.info._raw = data.xml.childNodes[0].childNodes[0].data.split("|");
        data.info.scale = data.info._raw[2];

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
      })

      .then(() => {
        audio = new Audio(solution.audio);
        audio.play();

        audio.addEventListener("playing", function (event) {
          return;
          if (_seekStatus) {
            _sound.pause();
            return;
          }
          _tween = new TweenLite(_sound, _sound.duration, {
            onUpdate: sound_playProgress,
          });
          playTween();
        });

        audio.addEventListener("pause", function (event) {
          return;
          if (_seekStatus) return;
          _tween.pause();
          _tween.onUpdate = null;
          _sound.pause();
          pauseTween();
        });

        audio.addEventListener("seeked", function (event) {
          return;
        });
      });
  };

  const onToggle = () => {
    if (!audio) return;

    audio.paused ? audio.play() : audio.pause();
    audio = audio;
  };

  $: typeof window === "object" && solution && onSolutionChange();
</script>

<div class="player">
  <div class="video">
    <canvas bind:this={canvas} />
  </div>
  <div class="audio">
    <button class="toggle" on:click={onToggle}>
      {!audio || audio.paused ? "\u25BA" : "\u2759\u00A0\u2759"}
    </button>
    <div class="progress">
      <span class="duration">{audioDurationText}&nbsp;&nbsp;</span>
    </div>
    <button class="speed" />
  </div>
</div>

<style>
  .player {
    width: 100%;
  }

  .player .video {
    display: flex;
    justify-content: center;
    align-items: center;

    height: calc(100% - 40pt);
  }

  .player .audio {
    display: flex;
    justify-content: space-between;

    width: 100%;
    height: 40pt;

    box-sizing: border-box;
    border-top: 1px solid var(--accents-3);
    background-color: white;
  }
  .player .audio button {
    width: 40pt;

    margin: 0;
    padding: 0;
    border: 0;
    font-size: 150%;
    background-color: white;
  }
  .player .audio .toggle {
  }
  .player .audio .progress {
    display: flex;
    justify-content: end;
    align-items: center;
    flex: 1 1 auto;
  }
  .player .audio .speed {
  }
</style>
