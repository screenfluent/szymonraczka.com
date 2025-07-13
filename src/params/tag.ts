import { getAllTags } from '$lib/content';
import type { ParamMatcher } from '@sveltejs/kit';

export const match: ParamMatcher = (param) => {
	// Fetch all available tags from your content (now synchronous)
	const allTags = getAllTags('en');
	
	// Check if the parameter from the URL is in the list of valid tags
	return allTags.includes(param);
};