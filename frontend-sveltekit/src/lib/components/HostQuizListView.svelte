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

{#each quizzes as quiz}
    <div class="card p-4 m-2 preset-filled flex justify-between items-center">
        {quiz.name}
        <button onclick={()=> {host(quiz.id)}} class="btn preset-filled-secondary-50-950">Host</button>
    </div>
{/each}
