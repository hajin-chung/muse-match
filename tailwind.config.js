/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["views/**/*.templ"],
	theme: {
		extend: {
			fontSize: {
				xxs: "10px",
			},
			fontFamily: {
				sans: "Pretendard Variable",
				writing: "Satisfy",
			},
			colors: {
				kakao: "#FAE100",
				naver: "#2DB400",
			},
			width: {
				layout: "1344px",
			},
		},
	},
	darkMode: "class",
	plugins: [],
};
