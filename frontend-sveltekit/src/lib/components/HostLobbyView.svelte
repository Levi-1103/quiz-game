<script lang="ts">
    import { players, type HostGame } from "$lib/services/host";
    import PlayerNameCard from "./PlayerNameCard.svelte";

    export let game: HostGame;

    function start() {
        game.startGame();
    }
</script>

<div>
    <button class="btn preset-filled" onclick={start}>StartGame</button>
    <p>HOST</p>
    <div class="flex flex-col">
        <h1>Players</h1>
        {#each $players as player (player.id)}
            <PlayerNameCard {player} />
        {/each}
    </div>
    <p>lobby State</p>
</div>
