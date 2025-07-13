import { getAllTags } from '$lib/content';
import type { ParamMatcher } from '@sveltejs/kit';

export const match: ParamMatcher = async (param) => {
	// Fetch all available tags from your content
	const allTags = await getAllTags('en');
	
	// Check if the parameter from the URL is in the list of valid tags
	return allTags.includes(param);
};