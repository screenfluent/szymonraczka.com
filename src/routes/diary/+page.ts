import { loadDiaryEntries } from '$lib/content';
import type { PageLoad } from './$types';

export const load: PageLoad = () => {
	// Return the promise directly. SvelteKit will stream the data.
	return {
		entries: loadDiaryEntries('en')
	};
};