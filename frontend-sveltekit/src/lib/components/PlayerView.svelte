<script lang="ts">
    import { GameState } from "$lib/model/net";
    import { gameState, PlayerGame } from "$lib/services/player";
    import PlayerJoinView from "./PlayerJoinView.svelte";
    import PlayerLobbyView from "./PlayerLobbyView.svelte";
    import PlayerPlayView from "./PlayerPlayView.svelte";

    let game = new PlayerGame();
    let active = false;

    function onJoin() {
        active = true
    }

     let views: Record<GameState, any> = {
        [GameState.LobbyState]: PlayerLobbyView,
        [GameState.PlayState]: PlayerPlayView,
        [GameState.RevealState]: undefined,
        [GameState.EndState]: undefined
    };
</script>

{#if active}
{@const SvelteComponent = views[$gameState]}
    {#key SvelteComponent}
        <SvelteComponent {game} />
    {/key}
{:else}
    <PlayerJoinView {onJoin} {game} />
{/if}
