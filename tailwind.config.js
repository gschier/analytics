const colorVariants = ['primary', 'secondary', 'danger', 'gray'];

const backgroundColors = {};
for (const c of colorVariants) {
    backgroundColors[c] = {};
    for (const n of [0, 50, 100, 200, 300, 400, 500, 600, 700, 800, 900]) {
        backgroundColors[c][n] = function({ opacityVariable, opacityValue }) {
            if (opacityValue !== undefined) {
                return `hsla(var(--color-${c}-${n}), ${opacityValue})`;
            }
            if (opacityVariable !== undefined) {
                return `hsla(var(--color-${c}-${n}), var(${opacityVariable}, 1))`;
            }
            return `hsl(var(--color-${c}-${n}))`;
        };
    }
}


const white = ({ opacityVariable, opacityValue }) => {
    if (opacityValue !== undefined) {
        return `hsla(0, 100%, 100%, ${opacityValue})`;
    }
    if (opacityVariable !== undefined) {
        return `hsla(0, 100%, 100%, var(${opacityVariable}, 1))`;
    }
    return `hsl(1, 100%, 100%)`;
};

module.exports = {
    mode: 'jit',
    purge: ['./src/**/*.{js,ts,jsx,tsx}'],
    darkMode: 'class',
    theme: {
        colors: {
            white,
            ...backgroundColors,
        },
    },
    variants: {},
    plugins: [],
};
