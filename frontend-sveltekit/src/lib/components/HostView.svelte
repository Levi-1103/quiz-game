<script lang="ts">
    import { GameState } from "$lib/model/net";
    import { HostGame, gameState } from "$lib/services/host";
    import HostLobbyView from "./HostLobbyView.svelte";
    import HostQuizListView from "./HostQuizListView.svelte";
    import HostPlayView from "./HostPlayView.svelte";

    let game = new HostGame();
    let active = false;

    function onHost(quizID: string) {
        game.hostQuiz(quizID);
        active = true;
    }

    let views: Record<GameState, any> = {
        [GameState.LobbyState]: HostLobbyView,
        [GameState.PlayState]: HostPlayView,
        [GameState.RevealState]: undefined,
        [GameState.EndState]: undefined
    };
</script>

<button
    onclick={() => {
        console.log($gameState as GameState);
    }}>Game State Test</button
>

{#if active}
    {@const SvelteComponent = views[$gameState]}
    {#key SvelteComponent}
        <SvelteComponent {game} />
    {/key}
{:else}
    <HostQuizListView host={onHost} />
{/if}
