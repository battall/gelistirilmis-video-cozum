<script context="module">
  const calculateTime = (secs) => {
    const minutes = Math.floor(secs / 60);
    const seconds = Math.floor(secs % 60);
    const returnedSeconds = seconds < 10 ? `0${seconds}` : `${seconds}`;
    return `${minutes}:${returnedSeconds}`;
  };
</script>

<script>
  import { onMount } from "svelte";

  export let src = "";

  let audio;
  let audioPaused = true;
  let audioCurrentTimeText = "";
  let audioPlaybackRate = 1;

  let track;
  let trackWidth = "0";

  onMount(() => {
    audio = new Audio();

    // This emits when src changed
    audio.addEventListener("timeupdate", () => {
      audioCurrentTimeText = calculateTime(audio.currentTime);
      trackWidth = (audio.currentTime / audio.duration) * 100 || 0;
    });

    // audio.play();
    // audio = audio;

    audio.addEventListener("playing", function (event) {
      audioPaused = false;

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
      audioPaused = true;

      return;
      if (_seekStatus) return;
      _tween.pause();
      _tween.onUpdate = null;
      _sound.pause();
      pauseTween();
    });

    audio.addEventListener("seeked", function (event) {
      return;
      _seekStatus = true;
      stopAll();
      clearAllShapes();

      var _time = _sound.currentTime;

      aa.each(_mainArray, function (_key, _object) {
        if (
          _object.startTime <= _time &&
          _time <= _object.startTime + _object.duration
        ) {
          _time = _object.startTime - 0.001;
          return false;
        }
      });

      aa.each(_mainArray, function (_key, _object) {
        if (_object.startTime < _time) {
          _object.status = true;

          if (_object.type == "add") {
            var _tempObject = getObjectForAdd(_object.objectID);
            _tempObject.status = true;
            canvas_addNewProcess(_object, true, _tempObject, false);
          } else {
            canvas_addNewProcess(_object, true, false);
          }
        } else {
          _object.status = false;
        }
      });

      _seekStatus = false;
    });
  });

  const onSrcChange = () => {
    audioPaused = true;
    audio.src = src;
    audio.play();
  };

  const onToggle = () => {
    if (!audio) return;

    audio.paused ? audio.play() : audio.pause();
  };

  const onChangeSpeed = () => {
    if (audioPlaybackRate === 2) audioPlaybackRate = 0.5;
    else audioPlaybackRate += 0.5;
    audio.playbackRate = audioPlaybackRate;
  };

  $: typeof window === "object" && src && onSrcChange();
</script>

<div class="audio">
  <button class="toggle" on:click={onToggle}>
    {audioPaused ? "\u25BA" : "\u2759\u00A0\u2759"}
  </button>
  <div class="progress">
    <div class="track" style={`width: ${trackWidth}%`} />
    <span class="current-time">{audioCurrentTimeText}</span>
  </div>
  <button class="speed" on:click={onChangeSpeed}>{audioPlaybackRate}×</button>
</div>

<style>
  .audio {
    display: flex;
    justify-content: space-between;

    width: 100%;
    height: 40pt;

    box-sizing: border-box;
    border-top: 1px solid var(--accents-3);
    background-color: white;
  }
  .audio button {
    width: 40pt;

    margin: 0;
    padding: 0;
    border: 0;
    font-size: 150%;
    background-color: white;
  }
  .audio .progress {
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex: 1 1 auto;
  }
  .audio .progress .track {
    height: 100%;

    background-color: lightblue;

    transition: width 500ms ease;
  }
  .audio .progress .current-time {
    position: absolute;
    right: 40pt;

    width: 32pt;
  }
</style>
