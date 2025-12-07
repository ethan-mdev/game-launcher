<script lang="ts">
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    export let username = "";
    export let password = "";

    let isRegistering = false;
    let confirmPassword = "";
</script>

<div class="h-full flex items-center justify-center p-8">
    <div class="w-full max-w-xs space-y-6">
        <!-- Logo/Title -->
        <div class="text-center mb-8">
            <!-- Logo would go here -->
            <h1 class="text-white text-lg font-medium">
                {isRegistering ? "Create an account" : "Sign in to continue"}
            </h1>
        </div>

        <!-- Form -->
        <div class="space-y-3">
            <input
                type="text"
                bind:value={username}
                placeholder={isRegistering ? "Username" : "Username or Email"}
                class="w-full bg-white/5 text-white px-4 py-3.5 rounded-xl border border-white/10 focus:border-yellow-500/50 focus:bg-white/10 focus:outline-none transition placeholder:text-neutral-500"
            />

            {#if isRegistering}
                <input
                    type="email"
                    placeholder="Email"
                    class="w-full bg-white/5 text-white px-4 py-3.5 rounded-xl border border-white/10 focus:border-yellow-500/50 focus:bg-white/10 focus:outline-none transition placeholder:text-neutral-500"
                />
            {/if}

            <input
                type="password"
                bind:value={password}
                placeholder="Password"
                class="w-full bg-white/5 text-white px-4 py-3.5 rounded-xl border border-white/10 focus:border-yellow-500/50 focus:bg-white/10 focus:outline-none transition placeholder:text-neutral-500"
            />

            {#if isRegistering}
                <input
                    type="password"
                    bind:value={confirmPassword}
                    placeholder="Confirm Password"
                    class="w-full bg-white/5 text-white px-4 py-3.5 rounded-xl border border-white/10 focus:border-yellow-500/50 focus:bg-white/10 focus:outline-none transition placeholder:text-neutral-500"
                />
            {/if}
        </div>

        <button
            on:click={() => dispatch(isRegistering ? "register" : "login")}
            class="w-full py-3.5 bg-gradient-to-r from-yellow-500 to-yellow-600 text-black font-semibold rounded-xl hover:opacity-90 transition shadow-lg shadow-yellow-500/20"
        >
            {isRegistering ? "Create Account" : "Sign In"}
        </button>

        <p class="text-center text-neutral-600 text-sm">
            {#if isRegistering}
                Already have an account?
                <button
                    on:click={() => (isRegistering = false)}
                    class="text-yellow-500 hover:text-yellow-400 transition"
                >
                    Sign in
                </button>
            {:else}
                Need an account?
                <button
                    on:click={() => (isRegistering = true)}
                    class="text-yellow-500 hover:text-yellow-400 transition"
                >
                    Register
                </button>
            {/if}
        </p>
    </div>
</div>
