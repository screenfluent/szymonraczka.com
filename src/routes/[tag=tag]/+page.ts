import { getDiaryEntriesByTag, getAllTags } from '$lib/content';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {
	const tag = params.tag;
	
	return {
		entries: getDiaryEntriesByTag(tag, 'en'),
		allTags: getAllTags('en'),
		selectedTag: tag
	};
};