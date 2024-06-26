import debug from 'debug';

const DEBUG_NAMESPACES = [
    'form-input',
    'page-controller',
]

for (const ns of DEBUG_NAMESPACES) {

    debug.enable(`${ns}`);
    debug.enable(`${ns}:*`);
}