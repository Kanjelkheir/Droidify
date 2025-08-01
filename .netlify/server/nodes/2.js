

export const index = 2;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/2.DQfONYMh.js","_app/immutable/chunks/DsnmJJEf.js","_app/immutable/chunks/DmYPTLUW.js","_app/immutable/chunks/-F4HFozQ.js"];
export const stylesheets = [];
export const fonts = [];
