// Typed API client for BobMediaManager
// All fetch calls proxy through the Go backend at :48040

const BASE = '/api';

async function request<T>(method: string, path: string, body?: unknown): Promise<T> {
	const res = await fetch(BASE + path, {
		method,
		headers: body ? { 'Content-Type': 'application/json' } : {},
		body: body ? JSON.stringify(body) : undefined
	});
	if (!res.ok) throw new Error(`${method} ${path} → ${res.status}`);
	return res.json() as Promise<T>;
}

// Config
export const getConfig = () => request<Config>('GET', '/config');
export const putConfig = (cfg: Config) => request<Config>('PUT', '/config', cfg);

// Scan
export const scan = (req: ScanRequest) => request<MediaFile[]>('POST', '/scan', req);

// Match & search
export const match = (req: unknown) => request<unknown>('POST', '/match', req);
export const search = (req: SearchRequest) => request<SearchResult[]>('POST', '/search', req);

// Rename
export const renamePreview = (req: unknown) => request<unknown>('POST', '/rename/preview', req);
export const renameExecute = (req: unknown) => request<unknown>('POST', '/rename/execute', req);

// --- Type stubs (full types in Phase 9) ---

export interface Config {
	port: number;
	theme: string;
	sources: Record<string, boolean>;
	apiKeys: Record<string, string>;
	companions: Record<string, unknown>;
	rename: { defaultOperation: string; renameFolders: boolean };
}

export interface MediaFile {
	id: string;
	path: string;
	name: string;
	size: number;
	isFolder: boolean;
	mediaType: string;
	matchStatus: 'pending' | 'matched' | 'unmatched';
	confidence: number;
	resolvedIds: Record<string, string>;
	renamedPath?: string;
}

export interface SearchResult {
	id: string;
	sourceSlug: string;
	title: string;
	year?: number;
	overview?: string;
	posterUrl?: string;
	ids?: Record<string, string>;
	confidence: number;
}

export interface ScanRequest {
	path: string;
	recursive: boolean;
	mediaType: string;
}

export interface SearchRequest {
	query: string;
	sources: string[];
	mediaType: string;
}
