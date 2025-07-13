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

	{#if data.allTags.length > 0}
		<div class="mb-8">
			<h2 class="text-sm font-medium text-gray-700 mb-3">Filter by tag:</h2>
			<div class="flex flex-wrap gap-2">
				<a 
					href="/diary" 
					class="px-3 py-1 rounded-full text-sm transition-colors bg-gray-100 text-gray-700 hover:bg-gray-200"
				>
					All
				</a>
				{#each data.allTags as tag}
					<a 
						href="/{encodeURIComponent(tag)}" 
						class="px-3 py-1 rounded-full text-sm transition-colors {data.selectedTag === tag ? 'bg-blue-100 text-blue-800' : 'bg-gray-100 text-gray-700 hover:bg-gray-200'}"
					>
						{tag}
					</a>
				{/each}
			</div>
			{#if data.selectedTag}
				<p class="text-sm text-gray-600 mt-2">
					Showing entries tagged with "{data.selectedTag}"
				</p>
			{/if}
		</div>
	{/if}

	{#if data.entries.length > 0}
		<div class="space-y-6">
			{#each data.entries as entry}
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
									<a 
										href="/{encodeURIComponent(tag)}" 
										class="bg-gray-100 hover:bg-gray-200 px-2 py-1 rounded text-xs transition-colors"
									>
										{tag}
									</a>
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
			<p class="text-gray-500 text-lg">No diary entries found for "{data.selectedTag}".</p>
			<p class="text-gray-400 text-sm mt-2">
				<a href="/diary" class="text-blue-600 hover:underline">View all entries</a>
			</p>
		</div>
	{/if}
</div>