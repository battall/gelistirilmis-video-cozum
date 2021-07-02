<script>
  import { onMount } from "svelte";

  export let solution = {};

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
    };
    let ruffle = window.RufflePlayer.newest();
    player = ruffle.createPlayer();
    let container = document.getElementById("player");
    container.appendChild(player);
  });

  $: solution && solution.swf && player && player.load(solution.swf);
</script>

<div id="player">
  <canvas />
</div>

<style>
</style>
