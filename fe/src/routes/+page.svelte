<script>
	import { onMount } from 'svelte';

	let message = "ola"
    
	onMount(async () => {
		console.log("onMount")

		try {
			const response = await fetch("http://localhost:9080/api")
			if (response.ok) {
					const data = await response.json()
					console.log(data)
					message = data["name"]
			} else {
					message = "error"
			}
		} catch (error) {
			console.error("Error fetching data", error)
		}
	});


</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Demo sveltkit app!" />
</svelte:head>

<section>
	<h1>
		Hey {message}!
	</h1>
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}

	h1 {
		width: 100%;
	}

</style>
