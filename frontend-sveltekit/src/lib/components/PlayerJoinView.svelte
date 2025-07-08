<script lang="ts">
    import type { PlayerGame } from "$lib/services/player";

  let code = "";
  let name = "";
  export let game:PlayerGame;
  export let join;

  function joinGame(){
    game.join(code, name);
  }
</script>

<div class="flex gap-4 items-center p-2">
  <input class="input" type="text" placeholder="Game Code" bind:value={code} />
  <input
    class="input"
    type="text"
    placeholder="Player Name"
    bind:value={name}
  />
  <button class="btn preset-filled" onclick={joinGame}>Join Game </button>
</div>
