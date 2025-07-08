
<script lang="ts">
    import type { HostGame } from "$lib/services/host";


    export let game: HostGame;

    function start() {
        game.startGame();
    }
</script>

<div>
    <button class="btn preset-filled" onclick={start}>StartGame</button>
      <p>HOST</p>
      <p>lobby State</p>
</div>