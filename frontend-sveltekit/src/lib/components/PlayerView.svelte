<script lang="ts">
    import { PlayerGame } from "$lib/services/player";
    import PlayerJoinView from "./PlayerJoinView.svelte";

    let game = new PlayerGame();
    let active = false;

    function onJoin() {
        active = true
    }
</script>

{#if active}
    <p>active</p>
{:else}
    <PlayerJoinView {onJoin} {game} />
{/if}
