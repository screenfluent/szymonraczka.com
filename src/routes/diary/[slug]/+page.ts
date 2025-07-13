import { loadDiaryEntry } from '$lib/content';
import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {
	const entry = await loadDiaryEntry(params.slug, 'en');

	// If the entry doesn't exist, SvelteKit will render the
	// nearest +error.svelte page with a 404 status.
	if (!entry) {
		error(404, 'Diary entry not found');
	}

	return {
		entry
	};
};