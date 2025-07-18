<script lang="ts">
    import { players, type HostGame } from "$lib/services/host";
    import PlayerNameCard from "../player/PlayerNameCard.svelte";


    export let game: HostGame;

    function start() {
        game.startGame();
    }
</script>

<div>
   
</div>


<div class="flex flex-col gap-4 mt-10">
      <h1 class="text-center text-5xl">Host Lobby</h1>
       <button class="btn preset-filled" onclick={start}>StartGame</button>
        <h2 class="text-center text-5xl">Players</h2>
        {#each $players as player (player.id)}
            <PlayerNameCard {player} />
        {/each}
       
</div>