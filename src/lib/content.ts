import type { ComponentType } from 'svelte';

export interface DiaryEntry {
	slug: string;
	title: string;
	date: Date;
	tags: string[];
	Component: ComponentType; // Svelte component from mdsvex; render with <svelte:component this={entry.Component} />
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
	module: { metadata: DiaryEntryFrontmatter; default: ComponentType },
	lang: 'en' | 'pl'
): DiaryEntry {
	try {
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
	} catch (error) {
		throw new Error(`Failed to create diary entry from ${filename}: ${error}`);
	}
}

// Load diary entries from content directory using Vite's glob import
// This works with Vite's static analysis and will be bundled at build time
export async function loadDiaryEntries(lang: 'en' | 'pl' = 'en'): Promise<DiaryEntry[]> {
	const entries: DiaryEntry[] = [];

	try {
		// Use Vite's glob import to load all markdown files (mdsvex processes them)
		const modules = import.meta.glob('/content/*/diary/*.md', { eager: false });

		for (const [path, moduleLoader] of Object.entries(modules)) {
			// Extract language and filename from path
			const pathMatch = path.match(/\/content\/(\w+)\/diary\/(.+)$/);
			if (!pathMatch) continue;

			const [, pathLang, filename] = pathMatch;

			// Skip if language doesn't match
			if (pathLang !== lang) continue;

			try {
				// Load the module to get metadata and Component
				const module = (await moduleLoader()) as {
					metadata: DiaryEntryFrontmatter;
					default: ComponentType;
				};

				const entry = createDiaryEntry(filename, module, lang);
				entries.push(entry);
			} catch (error) {
				console.warn(`Failed to load diary entry ${filename}:`, error);
				// Continue processing other files
			}
		}

		// Sort by date (newest first)
		entries.sort((a, b) => b.date.getTime() - a.date.getTime());

		return entries;
	} catch (error) {
		console.error('Failed to load diary entries:', error);
		return [];
	}
}

// Load a single diary entry by slug and language
export async function loadDiaryEntry(
	slug: string,
	lang: 'en' | 'pl' = 'en'
): Promise<DiaryEntry | null> {
	try {
		const entries = await loadDiaryEntries(lang);
		return entries.find((entry) => entry.slug === slug) || null;
	} catch (error) {
		console.error(`Failed to load diary entry ${slug}:`, error);
		return null;
	}
}

// Get diary entries filtered by tags
export async function getDiaryEntriesByTag(
	tag: string,
	lang: 'en' | 'pl' = 'en'
): Promise<DiaryEntry[]> {
	try {
		const entries = await loadDiaryEntries(lang);
		return entries.filter((entry) => entry.tags.includes(tag));
	} catch (error) {
		console.error(`Failed to get diary entries by tag ${tag}:`, error);
		return [];
	}
}

// Get all unique tags from diary entries
export async function getAllTags(lang: 'en' | 'pl' = 'en'): Promise<string[]> {
	try {
		const entries = await loadDiaryEntries(lang);
		const tagSet = new Set<string>();

		entries.forEach((entry) => {
			entry.tags.forEach((tag) => tagSet.add(tag));
		});

		return Array.from(tagSet).sort();
	} catch (error) {
		console.error('Failed to get all tags:', error);
		return [];
	}
}
