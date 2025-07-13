<script lang="ts">
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	const { entry } = data;
</script>

<div class="container mx-auto px-4 py-8">
	<article class="prose prose-lg max-w-4xl mx-auto">
		<header class="mb-8">
			<h1 class="text-4xl font-bold mb-4">{entry.title}</h1>
			<div class="flex items-center text-gray-600 mb-4">
				<time datetime={entry.date.toISOString()}>
					{entry.date.toLocaleDateString('en-US', {
						year: 'numeric',
						month: 'long',
						day: 'numeric',
						hour: '2-digit',
						minute: '2-digit'
					})}
				</time>
				{#if entry.tags.length > 0}
					<span class="mx-2">•</span>
					<div class="flex gap-2">
						{#each entry.tags as tag}
							<span class="bg-gray-100 px-2 py-1 rounded text-sm">{tag}</span>
						{/each}
					</div>
				{/if}
			</div>
		</header>

		{#if entry.Component}
			{@const TheComponent = entry.Component}
			<div class="content">
				<TheComponent />
			</div>
		{/if}

		<footer class="mt-12 pt-8 border-t border-gray-200">
			<a href="/diary" class="text-blue-600 hover:text-blue-800 transition-colors">
				← Back to all diary entries
			</a>
		</footer>
	</article>
</div>