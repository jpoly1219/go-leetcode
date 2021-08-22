<script>
    import Nav from "../components/nav.svelte"
</script>

<div class="p-8 h-screen">
    <Nav/>
    <slot></slot>
</div>

<style>
    @tailwind base;
    @tailwind components;
    @tailwind utilities;
</style>