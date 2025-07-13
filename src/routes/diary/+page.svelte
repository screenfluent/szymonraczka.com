<script lang="ts">
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<div class="container mx-auto px-4 py-8">
	<header class="mb-8">
		<h1 class="text-4xl font-bold mb-4">My Diary</h1>
		<p class="text-lg text-gray-600">
			Uncensored snapshots of my life and rebuild, inspired by Rick Rubin.
		</p>
	</header>

	{#await data.entries}
		<div class="text-center py-8">
			<p class="text-gray-500">Loading diary entries...</p>
		</div>
	{:then entries}
		{#if entries.length > 0}
			<div class="space-y-6">
				{#each entries as entry}
					<article class="border-b border-gray-200 pb-6">
						<h2 class="text-2xl font-semibold mb-2">
							<a href="/diary/{entry.slug}" class="hover:text-blue-600 transition-colors">
								{entry.title}
							</a>
						</h2>
						<div class="flex items-center text-sm text-gray-500 mb-3">
							<time datetime={entry.date.toISOString()}>
								{entry.date.toLocaleDateString('en-US', {
									year: 'numeric',
									month: 'long',
									day: 'numeric'
								})}
							</time>
							{#if entry.tags.length > 0}
								<span class="mx-2">•</span>
								<div class="flex gap-2">
									{#each entry.tags as tag}
										<span class="bg-gray-100 px-2 py-1 rounded text-xs">{tag}</span>
									{/each}
								</div>
							{/if}
						</div>
						{#if entry.excerpt}
							<p class="text-gray-700">{entry.excerpt}</p>
						{/if}
					</article>
				{/each}
			</div>
		{:else}
			<div class="text-center py-12">
				<p class="text-gray-500 text-lg">No diary entries yet.</p>
				<p class="text-gray-400 text-sm mt-2">
					Start sharing your honest thoughts and experiences.
				</p>
			</div>
		{/if}
	{:catch error}
		<div class="text-center py-8">
			<p class="text-red-600">Failed to load diary entries: {error.message}</p>
		</div>
	{/await}
</div>