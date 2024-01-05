/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["views/**/*.templ"],
	theme: {
		extend: {
			spacing: {
				"25": "100px",
				"50": "200px",
				"content": "calc(100% - 48px)"
			},
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
			gridTemplateColumns: {
				'auto-280': 'repeat(auto-fill, 280px)'
			},
			aspectRatio: {
				'banner': "4 / 1",
				"thumbnail": "4 / 3"
			}
		},
	},
	darkMode: "class",
	plugins: [],
};
