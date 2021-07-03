<script>
  import { onMount } from "svelte";

  export let solution = undefined;

  let canvas;

  onMount(() => {
    pdfjsLib.GlobalWorkerOptions.workerSrc =
      "https://cdn.jsdelivr.net/npm/pdfjs-dist@latest/build/pdf.worker.min.js";
  });

  // XML PARSING
  let data = {
    xml: "",
    info: {},
    objects: [],
  };

  let onSolutionChange = () =>
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
            alert("NOT IMPLEMENTED YET");
            object.objectID = element.find("id").text();
          }

          data.objects.push(object);
        }

        data.objects.sort(function (a, b) {
          return a.startTime - b.startTime;
        });
      })
      .then(() => pdfjsLib.getDocument(solution.pdf).promise)
      .then((pdf) => pdf.getPage(1))
      .then((page) => {
        //"/api/proxy?publisher=cap&path=/uploads/assets/questions/37511/3fc2fe867ebab060983a9ea4b6e5ed81-copy.pdf"
        let scale = 1;
        let resolution = 4;

        let viewport = page.getViewport({ scale });

        // Render at a higher resolution but show at a lower resolution
        canvas.width = resolution * viewport.width;
        canvas.height = resolution * viewport.height;

        canvas.style.width = viewport.width + "px";
        canvas.style.height = viewport.height + "px";

        canvas.getContext("2d").scale(resolution, resolution);
        return page.render({
          canvasContext: canvas.getContext("2d"),
          viewport,
        });
      });

  $: typeof window === "object" && solution && onSolutionChange();
</script>

<div id="player">
  <canvas bind:this={canvas} />
</div>

<style>
</style>
