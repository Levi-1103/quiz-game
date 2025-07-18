<script lang="ts">
  import { currentQuestion, tick, type HostGame } from "$lib/services/host";
  import { ProgressRing } from '@skeletonlabs/skeleton-svelte'

  import { COLORS } from "$lib/model/quiz";

  export let game: HostGame;
</script>

<div>
  <div class="grid grid-cols-1 gap-4 justify-center">
      <ProgressRing value={$tick} max={60}>
        <p class="text-2xl">{$tick}</p>
      </ProgressRing>
    </div>
  
  {#if $currentQuestion != null}
    <div class="card p-4 m-2 preset-filled flex flex-col text-center gap-4">
      <h2 class="text-4xl p-2">{$currentQuestion.name}</h2>
      <div class="flex justify-around gap-4">
        {#each COLORS as color, i}
            <div class="card p-20 {color}">
              <p class="text-primary-50 text-4xl">{$currentQuestion.choices[i].name}</p>
            </div>
        {/each}
      </div>
    </div>
  {/if}
</div>
