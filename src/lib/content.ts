import type { Component } from 'svelte';

export interface DiaryEntry {
	slug: string;
	title: string;
	date: Date;
	tags: string[];
	Component: Component; // Svelte component from mdsvex; render as <TheComponent /> where TheComponent = entry.Component
	excerpt?: string;
	lang: 'en' | 'pl';
	filename: string;
}

export interface DiaryEntryFrontmatter {
	title: string;
	date?: string;
	tags?: string[];
	excerpt?: string;
}

// Define the structure of an eagerly loaded module
type EagerModule = {
	metadata: DiaryEntryFrontmatter;
	default: Component;
};

// Parse filename to extract date and slug from YYYYMMDDHHMM-slug.md format
export function parseFilename(filename: string): { date: Date; slug: string } {
	const match = filename.match(/^(\d{12})-(.+)\.md$/);

	if (!match) {
		throw new Error(`Invalid filename format: ${filename}. Expected YYYYMMDDHHMM-slug.md`);
	}

	const [, dateStr, slug] = match;

	// Parse YYYYMMDDHHMM into Date
	const year = parseInt(dateStr.slice(0, 4));
	const month = parseInt(dateStr.slice(4, 6)) - 1; // Month is 0-indexed
	const day = parseInt(dateStr.slice(6, 8));
	const hour = parseInt(dateStr.slice(8, 10));
	const minute = parseInt(dateStr.slice(10, 12));

	const date = new Date(year, month, day, hour, minute);

	if (isNaN(date.getTime())) {
		throw new Error(`Invalid date in filename: ${filename}`);
	}

	return { date, slug };
}

// Create a DiaryEntry from mdsvex module
export function createDiaryEntry(
	filename: string,
	module: { metadata: DiaryEntryFrontmatter; default: Component },
	lang: 'en' | 'pl'
): DiaryEntry {
	const { date, slug } = parseFilename(filename);
	const { metadata, default: Component } = module;

	// Use metadata title or fallback to slug
	const title = metadata.title || slug.replace(/-/g, ' ');

	// Use metadata date or fallback to filename date
	const entryDate = metadata.date ? new Date(metadata.date) : date;

	// Validate date
	if (isNaN(entryDate.getTime())) {
		throw new Error(`Invalid date in metadata for ${filename}`);
	}

	return {
		slug,
		title,
		date: entryDate,
		tags: metadata.tags || [],
		Component,
		excerpt: metadata.excerpt,
		lang,
		filename
	};
}

// Load all modules eagerly (synchronously) at module initialization time
const modules = import.meta.glob<EagerModule>('/content/*/diary/*.md', { eager: true });

// Process all entries immediately when this module is imported
const allEntries: DiaryEntry[] = Object.entries(modules)
	.map(([path, module]) => {
		// Extract language and filename from path
		const pathMatch = path.match(/\/content\/(\w+)\/diary\/(.+)$/);
		if (!pathMatch) {
			console.warn(`Skipping invalid path: ${path}`);
			return null;
		}

		// Cast lang to supported languages
		const [, lang, filename] = pathMatch as [string, 'en' | 'pl', string];

		try {
			return createDiaryEntry(filename, module, lang);
		} catch (error) {
			console.warn(`Failed to process diary entry ${filename}:`, error);
			return null;
		}
	})
	// Filter out any null entries (failed processing)
	.filter((entry): entry is DiaryEntry => entry !== null)
	// Sort by date (newest first)
	.sort((a, b) => b.date.getTime() - a.date.getTime());

/**
 * Loads diary entries synchronously from the pre-processed list.
 */
export function loadDiaryEntries(lang: 'en' | 'pl' = 'en'): DiaryEntry[] {
	return allEntries.filter(entry => entry.lang === lang);
}

/**
 * Loads a single diary entry synchronously.
 */
export function loadDiaryEntry(
	slug: string,
	lang: 'en' | 'pl' = 'en'
): DiaryEntry | null {
	const entries = loadDiaryEntries(lang);
	return entries.find((entry) => entry.slug === slug) || null;
}

/**
 * Gets filtered diary entries synchronously.
 */
export function getDiaryEntriesByTag(
	tag: string,
	lang: 'en' | 'pl' = 'en'
): DiaryEntry[] {
	const entries = loadDiaryEntries(lang);
	return entries.filter((entry) => entry.tags.includes(tag));
}

/**
 * Gets all unique tags synchronously.
 */
export function getAllTags(lang: 'en' | 'pl' = 'en'): string[] {
	const entries = loadDiaryEntries(lang);
	const tagSet = new Set<string>();

	entries.forEach((entry) => {
		entry.tags.forEach((tag) => tagSet.add(tag));
	});

	return Array.from(tagSet).sort();
}
