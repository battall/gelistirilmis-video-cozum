<script>
  import { onMount } from "svelte";

  export let solution = {};

  // XAML PARSING
  const data = {
    xml: "",
    info: {},
    objects: [],
  };

  // XAML Read
  let bytes = fs.readFileSync("cozum_502146_05_02_2021_12_04_10.xaml");

  // XAML Decrypt
  var _strings = bytes.toString().split("");
  _strings.splice(10, 1);
  _strings.splice(20, 1);
  let xaml = atob(_strings.join("")).replaceAll(",", ".");

  // XAML Parse
  data.xml = new DOMParser().parseFromString(xaml, "text/xml").childNodes[0];

  // XAML Parse Info
  data.info._raw = data.xml.childNodes[0].childNodes[0].data.split("|");
  data.info.scale = data.info._raw[2];

  // XAML Parse Objects
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
        // console.log(element.childNodes[i].childNodes[0].data)
        let point = element.childNodes[i].childNodes[0].data.split("|");
        object.points.push({
          x: parseFloat(point[0]),
          y: parseFloat(point[1]),
        });
      }
    } else if (object.type == "circle") {
      // [x, y, w, h]
      object.rectangle = element.childNodes[0].childNodes[0].data.split("|");
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

  console.log(data.objects);

  let player;

  onMount(() => {
    window.RufflePlayer = window.RufflePlayer || {};
    window.RufflePlayer.config = {
      autoplay: "auto",
      unmuteOverlay: "visible",
      backgroundColor: null,
      letterbox: "fullscreen",
      warnOnUnsupportedContent: false,
      contextMenu: true,
      upgradeToHttps: window.location.protocol === "https:",
      maxExecutionDuration: { secs: 15, nanos: 0 },
      logLevel: "debug",
    };
    let ruffle = window.RufflePlayer.newest();
    player = ruffle.createPlayer();
    console.log(player);
    let container = document.getElementById("player");
    container.appendChild(player);

    player.load(
      "/api/proxy?publisher=cap&path=uploads/assets/questions/13545/e1ef4952ff89a69f0175942d577d1d21.swf"
    );
  });

  $: solution && solution.swf && player && player.load(solution.swf);
</script>

<div id="player">
  <canvas />
</div>

<style>
</style>
