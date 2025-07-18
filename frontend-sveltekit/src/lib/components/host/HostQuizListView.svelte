<script lang="ts">
    import { type Quiz } from "$lib/model/quiz";

    let quizzes: Quiz[] = [];
    export let host;

    async function getQuizzes() {
        const response = await fetch("http://localhost:3000/api/quizzes");
        if (!response.ok) {
            console.error(response);
            return;
        }
        const json = await response.json();

        return json;
    }

    (async () => {
        quizzes = await getQuizzes();
    })();
</script>

<div class="flex flex-col gap-4 mt-10">
      <h1 class="text-center text-5xl">Quiz List</h1>
    {#each quizzes as quiz}
        <div
            class="card p-2 preset-filled flex justify-between items-center"
        >
            {quiz.name}
            <button
                onclick={() => {
                    host(quiz.id);
                }}
                class="btn preset-filled-secondary-50-950">Host</button
            >
        </div>
    {/each}
</div>

<!-- <div class="flex flex-col gap-4 mt-10">
  <h1 class="text-center text-5xl">Join Game</h1>
  <input class="input" type="text" placeholder="Game Code" bind:value={code} />
  <input
    class="input"
    type="text"
    placeholder="Player Name"
    bind:value={name}
  />
  <button class="btn preset-filled" onclick={joinGame}>Join Game </button>
</div> -->
