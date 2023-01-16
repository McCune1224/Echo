<script lang="ts">
    import logo from "./assets/svelte.png"
    import Counter from "./lib/Counter.svelte"
    import "./App.css"
</script>

<main>
    <img src={logo} alt="Svelte Logo" />
    <!-- Make a div to have the entire page have a dark background using tailwindcss-->
    <div class="font-extrabold">
        <h1>Hello Typescript!</h1>

        <Counter />

        <p>
            Visit <a href="https://svelte.dev">svelte.dev</a> to learn how to build
            Svelte apps.
        </p>

        <p>
            Check out <a href="https://github.com/sveltejs/kit#readme"
                >SvelteKit</a
            >
            for the officially supported framework, also powered by Vite!
        </p>
    </div>
  import logo from './assets/svelte.png'
  import Auth from './lib/Auth.svelte';
  import Counter from './lib/Counter.svelte'
</script>

<main> 
  <Auth />
 
</main>
