import { loadDiaryEntries, getAllTags } from '$lib/content';
import type { PageLoad } from './$types';

export const load: PageLoad = () => {
	// This page now only shows all entries.
	// The selectedTag is null because no tag is selected here.
	return {
		entries: loadDiaryEntries('en'),
		allTags: getAllTags('en'),
		selectedTag: null
	};
};