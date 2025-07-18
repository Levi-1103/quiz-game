<script lang="ts">
    import type { PlayerGame } from "$lib/services/player";

  let code = "";
  let name = "";
  export let game:PlayerGame;
  export let onJoin;

  function joinGame(){
    onJoin()
    game.join(code, name);
  }
</script>

<div class="flex flex-col gap-4 mt-10">
  <h1 class="text-center text-5xl">Join Game</h1>
  <input class="input" type="text" placeholder="Game Code" bind:value={code} />
  <input
    class="input"
    type="text"
    placeholder="Player Name"
    bind:value={name}
  />
  <button class="btn preset-filled" onclick={joinGame}>Join Game </button>
</div>
