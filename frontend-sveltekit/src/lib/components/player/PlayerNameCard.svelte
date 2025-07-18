<script lang="ts">
    import type { Player } from "$lib/model/quiz";

    export let player: Player;
</script>

<div class="card p-2 preset-filled flex items-center justify-center">
    <p class="text-xl text-center">{player.name}</p>
</div>

