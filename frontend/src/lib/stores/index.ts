import { writable, derived, get } from 'svelte/store';
import type { MediaFile, SearchResult } from '$lib/api';

// ── UI state ──────────────────────────────────────────────────────────────

export const activeTab = writable<string>('rename');
export const rightTab = writable<string>('metadata');
export const rightCollapsed = writable<boolean>(false);
export const theme = writable<string>('dark');

// ── File lists ────────────────────────────────────────────────────────────

export const sourceFiles = writable<MediaFile[]>([]);
export const sourceFilter = writable<string>('');
export const renamedFiles = writable<MediaFile[]>([]);
export const renamedFilter = writable<string>('');

// Multi-select: set of selected file IDs
export const selection = writable<Set<string>>(new Set());
let lastClickedIndex = -1;

export function selectFile(file: MediaFile, index: number, ctrl: boolean, shift: boolean, files: MediaFile[]) {
  selection.update(sel => {
    const newSel = new Set(sel);
    if (shift && lastClickedIndex >= 0) {
      const lo = Math.min(lastClickedIndex, index);
      const hi = Math.max(lastClickedIndex, index);
      for (let i = lo; i <= hi; i++) newSel.add(files[i].id);
    } else if (ctrl) {
      if (newSel.has(file.id)) newSel.delete(file.id);
      else newSel.add(file.id);
    } else {
      newSel.clear();
      newSel.add(file.id);
    }
    lastClickedIndex = index;
    return newSel;
  });
}

export function clearSelection() {
  selection.set(new Set());
  lastClickedIndex = -1;
}

export function selectAllMatched() {
  const files = get(sourceFiles);
  selection.set(new Set(files.filter(f => f.matchStatus === 'matched').map(f => f.id)));
}

export const filteredSource = derived([sourceFiles, sourceFilter], ([files, q]) => {
  if (!q) return files;
  const lq = q.toLowerCase();
  return files.filter(f => f.name.toLowerCase().includes(lq));
});

export const filteredRenamed = derived([renamedFiles, renamedFilter], ([files, q]) => {
  if (!q) return files;
  const lq = q.toLowerCase();
  return files.filter(f => (f.renamedPath ?? f.name).toLowerCase().includes(lq));
});

// Stats
export const fileStats = derived(sourceFiles, files => ({
  total: files.length,
  matched: files.filter(f => f.matchStatus === 'matched').length,
  pending: files.filter(f => f.matchStatus === 'pending').length,
  unmatched: files.filter(f => f.matchStatus === 'unmatched').length,
}));

// ── Rename settings ───────────────────────────────────────────────────────

export const renameOp = writable<'move' | 'copy' | 'rename'>('move');
export const destination = writable<string>('');
export const renameFolders = writable<boolean>(false);
export const formatTemplate = writable<string>('{n} ({y})/Season {s00}/{n} - S{s00}E{e00} - {t}');
export const mediaType = writable<string>('anime');

// ── Inspector (right panel) ───────────────────────────────────────────────

export const selectedFile = writable<MediaFile | null>(null);

// ── Modals ────────────────────────────────────────────────────────────────

export const idPickerOpen = writable<boolean>(false);
export const sourcesModalOpen = writable<boolean>(false);
export const idPickerTarget = writable<MediaFile | null>(null);

// ── Context menu ─────────────────────────────────────────────────────────

export interface CtxMenuState {
  open: boolean;
  x: number;
  y: number;
  file: MediaFile | null;
}
export const ctxMenu = writable<CtxMenuState>({ open: false, x: 0, y: 0, file: null });

export function openCtxMenu(e: MouseEvent, file: MediaFile) {
  e.preventDefault();
  ctxMenu.set({ open: true, x: e.clientX, y: e.clientY, file });
}
export function closeCtxMenu() {
  ctxMenu.update(s => ({ ...s, open: false, file: null }));
}

// ── Progress / WebSocket ──────────────────────────────────────────────────

export interface ProgressEvent {
  operation: string;
  fileId?: string;
  total: number;
  done: number;
  error?: string;
  status: 'running' | 'completed' | 'error';
}
export const activeProgress = writable<ProgressEvent | null>(null);

let ws: WebSocket | null = null;

export function connectWS() {
  if (ws) return;
  const url = `ws://${location.host}/ws/progress`;
  ws = new WebSocket(url);
  ws.onmessage = (e) => {
    try {
      const evt = JSON.parse(e.data) as ProgressEvent;
      activeProgress.set(evt.status === 'completed' ? null : evt);
    } catch {}
  };
  ws.onclose = () => {
    ws = null;
    setTimeout(connectWS, 3000);
  };
}

// ── Enabled sources (badges) ──────────────────────────────────────────────

export const enabledSources = writable<Record<string, boolean>>({
  anidb: true, tvdb: true, tmdb: true, imdb: true, musicbrainz: false,
});

// ── Search results (ID picker) ────────────────────────────────────────────

export const searchResults = writable<SearchResult[]>([]);
export const searchQuery = writable<string>('');
export const pickedResult = writable<SearchResult | null>(null);
