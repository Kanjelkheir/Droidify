import { init } from '../serverless.js';

export const handler = init((() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set([".DS_Store","favicon.svg"]),
	mimeTypes: {".svg":"image/svg+xml"},
	_: {
		client: {start:"_app/immutable/entry/start.CnQvvbLW.js",app:"_app/immutable/entry/app.DqlYXxXa.js",imports:["_app/immutable/entry/start.CnQvvbLW.js","_app/immutable/chunks/DBYW9kc1.js","_app/immutable/chunks/DRVc36eT.js","_app/immutable/chunks/-F4HFozQ.js","_app/immutable/chunks/C75fvuOP.js","_app/immutable/entry/app.DqlYXxXa.js","_app/immutable/chunks/-F4HFozQ.js","_app/immutable/chunks/DRVc36eT.js","_app/immutable/chunks/C75fvuOP.js","_app/immutable/chunks/DsnmJJEf.js"],stylesheets:[],fonts:[],uses_env_dynamic_public:false},
		nodes: [
			__memo(() => import('../server/nodes/0.js')),
			__memo(() => import('../server/nodes/1.js')),
			__memo(() => import('../server/nodes/2.js'))
		],
		routes: [
			{
				id: "/",
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 2 },
				endpoint: null
			}
		],
		prerendered_routes: new Set([]),
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})());
