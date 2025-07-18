<script lang="ts">
    import { currentQuestion, tick, type HostGame } from "$lib/services/host";


    export let game: HostGame;

</script>

<div>
    <p>Clock: {$tick}</p>
    {#if $currentQuestion != null}
        <div class="card p-4 m-2 preset-filled flex flex-col text-center gap-4">
          <h2 class="text-lg p-2">{$currentQuestion.name}</h2>
          <div class="flex justify-around gap-4">
            {#each $currentQuestion.choices as choice}
              <!-- <div class="btn preset-filled-secondary-50-950"> -->
              <div class="">
                {choice.name}
              </div>
            {/each}
          </div>
        </div>
      {/if}
</div>