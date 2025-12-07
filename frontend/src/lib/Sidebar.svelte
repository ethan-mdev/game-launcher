<script lang="ts">
    import { currentPage, type Page } from "./stores/navigation";
    import icon from "../assets/images/icon.png";

    function navigate(page: Page) {
        currentPage.set(page);
    }
</script>

<!-- Sidebar: fixed width -->
<aside
    class="w-56 flex flex-col bg-neutral-900 p-4 border-r border-neutral-800"
>
    <!-- Logo area -->
    <div class="flex items-center gap-3 mb-6">
        <div class="w-10 h-10 bg-gradient-to-br from-yellow-400 to-yellow-600 rounded-lg">
            <img src={icon} alt="Logo" class="w-10 h-10 rounded-lg" />
        </div>
        <div class="text-white font-semibold">Launcher</div>
    </div>

    <!-- Navigation -->
    <nav class="flex flex-col gap-1">
        <button
            on:click={() => navigate("play")}
            class="px-3 py-2 rounded-lg text-left {$currentPage === 'play'
                ? 'bg-neutral-800 text-white border-l-2 border-yellow-500'
                : 'text-neutral-400 hover:bg-neutral-800'}">Play</button
        >
        <button
            on:click={() => navigate("news")}
            class="px-3 py-2 rounded-lg text-left {$currentPage === 'news'
                ? 'bg-neutral-800 text-white border-l-2 border-yellow-500'
                : 'text-neutral-400 hover:bg-neutral-800'}"
            >News & Events</button
        >
        <button
            on:click={() => navigate("settings")}
            class="px-3 py-2 rounded-lg text-left {$currentPage === 'settings'
                ? 'bg-neutral-800 text-white border-l-2 border-yellow-500'
                : 'text-neutral-400 hover:bg-neutral-800'}">Settings</button
        >
    </nav>

    <!-- Bottom status -->
    <div class="mt-auto text-sm text-neutral-500">
        <div class="flex justify-between">
            <span>Ping</span><span class="text-green-400">28 ms</span>
        </div>
    </div>
</aside>
