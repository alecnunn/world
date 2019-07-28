package charge

import svg "github.com/ajstarks/svgo"

var StagLodged = Charge{
	Identifier: "stag-lodged",
	Name:       "stag lodged",
	Noun:       "stag",
	NounPlural: "stags",
	Descriptor: "lodged",
	SingleOnly: false,
	Tags: []string{
		"animal",
		"hunting",
		"stag",
	},
	Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
		pathData0 := "M242.1,249.3c3.3,3.8,5.8,7.7,10.3,10.1c3.4,1.8-0.4,3.5-1.4,4.7c-2.5,2.8-6.6,2.3-9.5,2.2c-5.8-0.2-10.8,2.2-16.3,2.9    c-8.1,1.1-16.2,2.5-24.2,4c-10,1.8-20.2,2.9-29.9,5.9c-4.1,1.3-7.3,4.5-10.2,7.8c-2.2,2.4-5,4.9-8.4,3.4c-1.2-0.5-2.2-0.3-3.2-0.5    c-1.3-0.2-3.2,0.7-3.9-0.5c-0.8-1.5,0.6-3,1.6-4.3c5.2-6.3,12.7-9.5,19.1-14.1c5.6-4,12.1-4.7,18.7-5.3c11-1,21.7-3.3,32.4-6.1    c0.5-0.1,1-0.7,1.9-1.4c-12.1,0.4-23.5,0.7-34.9,3.2c-7.3,1.6-14.8,3.1-22.5,2.8c-4.6-0.2-7.4,3.9-11.1,5.9    c-2.6,1.4-4.5,2.7-7.7,1.7c-2.5-0.8-5.4-0.2-9.1-0.2c3.3-4.7,7.4-7.1,11.1-10.1c5.7-4.7,12.9-4.9,19.6-6.4    c8.6-1.9,17.5-2.9,26.3-4.4c-0.8-1.9-2-2.4-3.9-2.5c-8.8-0.3-17.4,0.8-25.9,2.6c-5.1,1-10.1,2.3-15.3,3.2    c-2.1,0.4-3.9,1.2-5.5,2.4c-4.4,3.4-9.4,3.9-14.6,3.5c-4.4-0.3-8.6,1.5-13,1.6c4.4,0,8.9,0,13.3,0c2.6,0,5.2,0.5,6.8,2.8    c1.5,2.1,0.5,4.4-2,4.5c-3,0.1-5.8,1-8.7,1.6c-2.2,0.5-4.7,0.7-6.8-0.8c-1.4-1-2.9-0.7-4.3-0.7c-20.6,0-41.2-0.1-61.8,0.1    c-3.5,0-6.3-1.4-9.1-2.7c-3-1.3-3-3.7-0.8-6.1c1.3-1.4,2.9-2.3,4.7-3.2c-3.5-1.7-7.2-2.6-11-2.1c-4.8,0.6-7.2-2.6-7.7-7.3    c-0.5-5.1,0.7-6.2,6.6-6.8c12.4-1.4,24.5-4,36.4-7.9c1.9-0.6,3.8,0.1,5.9-0.7c-0.5-2.3-2.4-3.8-4-4.8c-3.1-1.8-4.8-5.2-8.2-6.8    c-2.5-1.1-1.9-3.7-1.9-5.8c0-5.5,0-5.5-5.3-5.5c-0.8,0-1.7,0-2.7,0c0.3-1.9,2.1-2.4,2.9-3.8c1.8-3.1,3-6.2,2.7-9.9    c-0.3-3.2-0.4-6.3,2.2-9.2c2.2-2.3-1.5-3.3-2.3-5.6c5.8-0.7,5.4-5.6,6.7-9.6c2.3-7,4.6-14.3,11.3-18.7c1.2-0.8,2.7-1.3,0-3.3    c-4.7-3.4-9.7-3.6-14.7-2.9c-3,0.5-6.2,0.2-9.2,1.3c-2.7,1-5.1-0.4-6.5-3c-0.4-0.8-0.7-1.6-1.5-1.8c-2.6-0.7-4.5-2.1-3.5-4.9    c1.1-2.9,2.3-6,6.1-6.8c5.8-1.2,11.6-1.7,17.4-2.9c6-1.3,10.2-5.7,15.6-8c2.9-1.3,6.9-0.9,8.7-4.4c-1.6-2.4-4.6-2.6-6.8-2.7    c-8.1-0.2-11.2-5.3-13.9-11.8c-2.2-5.4-1.3-10.7-1.5-16.1c-0.1-1.9,1.2-3.9,3-4.5c1.9-0.6,2,1.9,2.8,3c2.8,3.9,2.9,8.7,4.1,13.2    c0.8,3.1,2.5,5.5,5.5,7c2.9,1.5,4.4,0.4,4.4-2.2c0-5.6,1.3-11.6-5.2-15.5c-3.2-2-3.7-6.4-3.8-10.2c-0.2-4.6-0.1-9.1,0-13.7    c0-1.8-0.7-4.3,2.4-4.1c2.5,0.1,4,1.7,4.2,4.2c0.2,3,0.2,6.1,0,9.1c-0.1,1.7,0.5,3.2,0.9,4.7c0.5,2,2.3,3.2,3.9,3.1    c1.5-0.1,1.1-2.2,1.1-3.5c0-0.9,0.6-1.6,0.6-2.3c0.1-3.7,3.5-7.3-1.9-11c-4.2-3-3.8-11.9-1.2-16.2c1.8-2.8,3.2-2.8,4.9,0.2    c0.7,1.2,0.9,2.5,0.8,3.9c-0.1,1.2,0,2.5,0,3.7c-0.1,1.6,0.2,3.3,1.5,3.8c1.3,0.5,1.4-1.6,2.2-2.4c3-3.2,1.9-7.4,2-11.3    c0.1-4.5,1-5.3,5.3-5.2c3.1,0,3.4,2.1,3.8,4.3c0.6,3.6,1.7,4.2,4.7,2.2c2-1.3,3.9-2.9,5.9-4.1c1.8-1.1,4-2.8,5.9-0.4    c2.4,3,0,4.9-2.1,6.3c-3.5,2.4-7.4,4.1-11.3,5.8c-7.7,3.2-10.2,10.3-12.1,17.5c-2.4,9.3-3.9,18.8-3.6,28.5    c0.2,5.6,2.3,10.1,5.1,14.6c0.6,1,1.3,1,2,0.8c1-0.4,0.8-1.5,0.5-2.2c-0.7-2-1.5-4-2.5-5.9c-2.5-5-1.8-9.7,0.6-14.6    c1.8-3.5,3.7-7.1,4.7-11c0.3-1.2,0.9-3.4,2.9-2.7c1.7,0.6,2,2.5,1.7,4.1c-0.6,2.6-1.4,5.1-2.2,7.7c-1.3,4.1-1.5,8.1-0.3,12.3    c1.5,5.1,4.8,6.6,9.2,3.6c3.5-2.5,6.8-5.3,10.2-7.9c6.1-4.5,13.5-6.4,19.6-10.8c4.6-3.3,9.2-6.7,8.4-13.5    c-0.1-0.8,0.5-1.8,0.8-2.6c1.4-4.4,1.9-4.5,5.7-1.2c2.4-4.7,4.1-9.5,4.3-14.8c0.1-5.2,1-10.1,4.2-14.4c0.5-0.7,0.9-1.6,1.8-1.3    c0.8,0.2,1.2,0.9,1.5,1.7c0.9,2.6,1.4,5.3,0.4,8c-0.8,2.4-1,5,0.8,6.4c1.7,1.4,3.7-0.4,5.3-1.6c3.9-2.9,7.7-6,10-10.5    c1.2-2.4,3.3-3.2,5.4-2.1c2.7,1.3,0.7,3.4-0.2,4.8c-4.9,7.9-10.6,15.1-19.2,19.2c-3,1.5-3.3,4.6-5,6.8c-0.4,0.5-0.6,2.1,0.4,2.6    c1.1,0.6,2.5,0.6,3.6-0.1c2.3-1.3,4.6-2.8,6.5-4.7c1.5-1.5,2.8-3.3,4.9-4.3c1.6-0.7,3-1.2,4,0.2c0.9,1.2,1,2.9-0.1,4.2    c-3.9,4.6-7.5,9.4-12.6,12.7c-3.8,2.5-8.1,3.8-12.4,3.6c-4.2-0.2-6.1,1.7-7.5,5c-2.6,6-6.9,10.3-13,12.4c-4,1.4-8.4,1.5-12.5,2.6    c-6.4,1.6-11.3,5.3-15.9,9.7c-0.3,0.3-1,0.5-0.4,1.2c0.6,0.6,1,0.5,1.7,0.1c8.6-4.3,17.9-2.1,27-2.6c1.5-0.1,3.1,0.6,3.8,1.9    c0.9,1.6-1.1,2.3-2,3.1c-7.4,6.9-15.6,12.2-26,13.2c-1.4,0.1-3.4,0.2-3.2,2.2c0.7,7.2-3.7,13-5.5,19.4c-2.3,8-3.5,16.4-6.8,24.2    c-1.6,3.8,1.5,7,4.2,9.7c1.3,1.3,2.7,0.8,4,0c5.7-3.6,12.1-4.2,18.7-4.2c10.2,0.1,20.4-0.7,30.5,1c12.2,2,24.5,2.3,36.8,3.7    c8.7,1,17.6,1.6,26.3,3c14.3,2.3,26.5,8.8,34.5,21.4c3.3,5.2,8.5,9.5,9.3,16.2c0.3,2.2,1.2,4.4,1.5,6.6c1,5.6,2.6,10.7,7.5,14.2    c1,0.7,2.3,1.8,1.5,3c-0.9,1.4-1.9,3-4,3.1c-5.3,0.4-10-0.4-13.6-4.8c-1.7-2.1-4-3.8-5.8-5.9c-2.2-2.7-3.9-2.7-6.6-0.4    C246.7,247,245.1,249.2,242.1,249.3z"
		pathData1 := "M96.1,186.8c2.6,1.7-1.6,3.9,0.5,5.6c2.4-0.5-0.1-4.8,3.5-4.6c-0.2,3.1-2,6.3,0.7,9.2    c0.6,0.6,0.1,1.2-0.3,1.7c-1.1,1.3-3.4,2.2-1.2,4.7c0.8,1-2.6,1.8-2.1,4.1c1.8-0.1,3.3-1.3,3-2.6c-1-4.8,5-6.5,4.6-11    c-0.1-0.6,0.6-0.7,0.8-0.5c3.4,2,4.7-1.6,6.6-2.7c6.5-3.7,13.4-3.6,20.4-4c14.6-0.8,28.9,1.9,43.3,2.7c9.3,0.5,18.7,1.5,28,2.6    c12.3,1.5,24.9,1.7,36.3,7.9c7.8,4.3,12.5,11.1,17.8,17.9c5.2,6.7,7.6,14.3,9,22.3c0.9,5.1,3.6,8.6,7.7,11.3    c0.6,0.4,1.4,0.7,1.2,1.4c-0.1,0.6-0.7,1-1.5,1.1c-4.1,0.4-7.9-0.5-10.9-3.3c-2.4-2.3-1.3-5.2-0.5-7.9c-1-0.1-1.3,0.5-1.7,1    c-0.5,0.8-0.6,2.6-1.8,2c-1.2-0.6,0-2,0.6-2.6c4.1-4.3,0.5-9.1,1.1-13.6c-1.1,3.1-1.1,3.7-0.8,7.1c0.3,2.6-1.9,4-3.2,5.8    c-0.5,0.7-0.9,0.3-1.2-0.1c-0.6-0.6-1.3-1-0.7-2.3c1.3-3.3,2.2-6.8,3.1-10.3c1.2-4.4-0.1-8.2-3.5-11.4c2.9,6,0.8,11.6-1.1,17.2    c-3,8.7-15.5,14.3-22.7,10.5c5.2-2,11.3-2.5,15.8-7c-9,4.8-18.2,5.8-27.9,1.5c2.7,2.2,4.4,4.9,6.3,7.3c2.1,2.7,6.3,2.4,9.6,1.7    c3.2-0.7,4.9,1,6.2,3c1.7,2.5,3.7,4.6,6.2,6.3c2.2,1.4,1.5,2.2-0.4,3.4c-2.1,1.3-4.7,1.2-6.6,1.3c-6,0-11.2,2.9-16.9,3.7    c-8.2,1.1-16.5,2.5-24.6,4.1c-11,2.2-22.3,2.8-32.9,7c-1.9,0.7-2.2-0.5-2.9-1.2c-1.2-1.3,0.1-1.8,1-2.5c6.5-4.9,14-6.7,22-7.6    c11.8-1.3,23.6-3.2,34.9-7.1c2.2-0.8,5-1,7.2,0.3c2.8,1.5,5.7,1.1,8.6,0.4c-1.4-0.1-2.9-0.2-4.3-0.3c0-0.1,0-0.2,0-0.4    c2.1-0.4,4.2-0.7,6.3-1.2c1-0.2,2-0.9,1.6-1.9c-0.5-1.2-1.3-0.1-2,0.2c-3.7,1.3-7.6,1.4-11.5,1.6c2.4-1,4.8-1.8,7.4-1.7    c-2.6-0.5-5.1-0.5-7.8,0.1c-4.3,0.9-8.3-1.5-12.3-3.1c-8.2-3.2-15.8-7.2-20.7-15c-1.1-1.8-1.6-3.5-1.2-5.4    c0.6-2.8-0.8-4.4-2.6-6.5c-0.3,0.7-0.5,1-0.6,1.3c-0.3,1.3,0.7,3.3-0.7,3.7c-2,0.5-2.7-2.1-3-3.2c-3.2-8.8-1.3-22,10-26.2    c2.7-1,5.9-0.9,8.9-1c11.3-0.5,22.5,0,32.8,5.3c0.8,0.4,1.5,0.4,2.6,0.2c-3.8-4.7-9.5-5.9-14.6-7c-9.9-2.2-20.1-1.7-30.1-0.1    c-7.6,1.2-13.1,8.5-13.3,16.4c-0.1,4.6-0.7,9.3,0.7,13.9c0.8,2.7-0.3,4.1-3.4,3.6c-0.8-0.1-2-0.6-2.3,0.8    c-0.4,1.7,1.1,1.9,1.9,2.1c1.8,0.5,3.3,1,4,3.4c-4.9-0.2-8.4-3.9-12.7-4.8c-0.8,0.6-0.2,0.7,0.1,0.9c1.7,1.4,5.3,3.1,4.4,4.4    c-1.5,2.2-3.4-1.1-5.3-1.6c-1.8-0.5-3.8-0.4-5.2-3c0.2,5.1,1.3,6.1,7.9,8.1c-1.5,3.1-4.1,1.5-5.6,0.8c-2.5-1.2-4.4-2.6-6.6,0.3    c-0.4,0.5-1.5,0.1-2.3-0.4c-1.1-0.7-2.1-0.4-2.8,1c1.3,0.9,3.1-0.3,4.8,1.2c-7.4-0.2-14.4,0.6-20.9-3.2c-1.2-0.7-2.1-1.3-1.4-2.1    c0.9-0.9-0.2-3,2-3.1c1.7-0.1,1.8,0.7,2.1,2.2c0.5,3.2,2.9,1.8,5.4,1.9c-1.8-2.2-5.3-3.1-3.4-6.5c2.6,0.7,4.4,2.8,6.9,4    c0.6-2.9-3.8-3.1-2.8-6.4c1.6,0,3.5,0,5.3,0c0.6,0,1.5,0.3,1.7-0.7c0.2-0.9,0.1-1.9-0.8-2.1c-1.6-0.4-2.8-1.4-4.2-2.1    c-0.1-0.1-0.3,0-0.4,0c-1.6,0-0.4,3.9-3.1,2.4c-1.6-0.9-2.8-2.6-5.1-2.3c-0.1,1.6,1.7,2.7,1.3,4.4c-0.2,0.7,0.4,1.9-0.7,2.1    c-1.7,0.2-0.5-1.6-1.3-2c-1.3,1.4-1,5.8-5,2.3c-1.1-1-1.7,0.9-2.5,1.4c-0.8,0.6-1.4,1.2-0.5,2.2c2,2.1,0.8,3.1-1.4,3.5    c-0.6,0.1-1,0.3-1,1.1c1,1.3,2.9-0.2,4.2,1.2c-2.1,1.7-4.4,3.2-5.1,6.2c-0.4,1.6-2.3,1.9-3.9,1.1c-0.8-0.4-1.6-0.8-2.3-0.4    c-9.7,5.7-20.5,3.7-30.9,4.7c-11.2,1-22.2,3.4-33.5,3.5c-1.8,0-3.7,1.3-5.4,2.3c8.5-1.2,17-0.6,25.5-0.7    c5.3-0.1,10.5,0.3,15.8-0.1c8.1-0.6,16,1.2,24,1.6c2.3,0.1,1.3,2,2.2,2.8c0.5,0.4,0.5,1.4,0,1.9c-1.1,1-2.3-0.3-2.7-0.8    c-1.5-1.6-3.2-1.4-4.9-1.4c-20.6,0-41.2,0-61.8,0c-0.7,0-1.4,0-2.1,0c-3.2,0.1-6.9-1.2-7.4-3.8c-0.6-2.9,3.4-3.8,5.8-4.4    c10.1-2.6,20.1-5,30-8.4c5.8-2,11.6-3.9,17.4-5.8c1.7-0.6,3-1.4,2.8-3.4c-0.2-2.7,0.9-5,1.6-7.4c1.2-4.1,2.5-8.1,2.8-12.5    c0.2-3.8,1.9-7.7,6.6-9c-5.4,0.8-8.7,4.6-8.5,9.7c0.1,3,0,5.5-3.7,5.9c-0.1,0-0.3,0.1-0.4,0.1c-1.8,0-2.1,3.6-4.2,2    c-1.1-0.9-1.4-2.8-2-4.2c-2.7,2.7-0.3,5,0.7,7.2c0.8,1.6,1.8-0.1,2.7-0.2c0.6-0.1,1.1-1,1.7-0.1c0.3,0.4,0.3,1.3,0.1,1.5    c-1.6,1.2-4.4,2.7-5.2,1.8c-2.5-2.5-3.9-1.2-6,0.2c-1.5,0.9-3.1,0-4.4-1.4c-1.7-1.9-0.5-3.3,0.8-4.6c-1.7-0.9-2,1-3,1.1    c-1-1.6,0.5-2.5,1.5-4c-2.9,1.3-5,1.3-5.1-2.2c0-0.2,0.1-1.1-0.6-0.6c-3.4,2.7-4.6-0.9-5.7-2.5c-1.3-2-2-4.7-0.4-7.2    c0.2-0.3-0.1-0.8-0.2-1.5c-0.7,1.3-1.3,2.3-1.8,3.4c-0.4,1-0.5,2.1-0.9,3.6c-4.8-7.5-0.9-13.7,2.8-20.8c-5.2,2.2-4.4,9.2-10.7,8.9    c4.9-5.2,5.6-11.7,8-18.1c-3.1,2-3.1,2-4.4,8.3c-1.4-8.1,3.8-12.8,8-18.3c-2.4-0.2-2.9,2.5-4.8,2.8c-0.7,0.1-1,0-1.3-0.4    c-0.6-0.7,0.1-0.6,0.6-0.9c3.2-2,3.9-5.2,4-8.7c0.1-4.7,3-8.1,7-10c1.8-0.9,3.2-1.9,4.4-3.3c-0.6-0.2-1-0.4-1.6,0.1    c-0.7,0.6-1.7,0.8-2.2,0.1c-0.7-0.9,0.3-1.5,0.8-2c2.7-2.4,4.9-6.5,9.4-2.7c3.8-3.2,7.8,1.7,11.8-0.3c-2.1-0.9-4.2-1.7-5.8-2.4    c3.7-3,8.8-1.6,13.1-3.6c2.6-1.2,5.2-2,7.2-4.1c0.5-0.6,1.1-1.6,2-0.8c0.6,0.5,0.4,1.5-0.1,2.3c-1.8,3.4-3.2,7.1-6.4,9.7    c-1.3,1-2.9,2-3.3,4.3c2.5,0.2,2.9-2.6,5-2.9c0.1,2.8,0.3,5.5-1.9,7.7l0.1,0.1c0.2-1.2,0.4-2.4-0.2-3.8c-3.5,8.7-7.7,17-8.7,26.4    c-0.3,0.2-0.6,0.5-0.8,0.8C95.7,187.4,95.9,187.1,96.1,186.8z"
		pathData2 := "M68,84.4c2.5,3,3.6,6.3,4.2,9.9c0.9,5,1.9,10,7.7,12.2c4.9,1.8,8,0.1,8-4.9c0-2.6,0.2-5.3,0-7.9    c-0.2-2.7-0.8-5-3.6-6.6c-4-2.3-4.9-6.7-5.3-10.8c-0.5-5-0.1-10.1-0.1-15.3c2,0.5,1.6,1.9,1.7,3.1c0,2.8,0.3,5.6-0.1,8.3    c-0.6,4.5,2.1,6.9,5.1,9.1c2.6,1.9,5.4,0.8,5.6-2.5c0.4-9.3,4.6-17.2,8.5-25.2c1.4-2.9,0.3-5.9,0.7-8.9c0.2-1.5-0.6-3.7,1.8-3.9    c2.5-0.3,2.2,1.7,2.4,3.4c0.6,4.7,3.5,5.6,8.3,2.7c2.8-1.7,5.6-3.6,8.6-5.5c1.2,3.1-1.3,3.9-2.8,5.2c-3.6,3.1-8.3,4.1-12.3,6.5    c-4.5,2.7-6.7,6.6-8.5,11.1c-5,12.5-6.5,25.6-5.8,39c0.2,3.2,1.1,6.4,2.9,8.5c1.8,2.1,1.3,3.7,0.8,5.4c-0.7,2.4-0.3,4.1,1.9,5.3    c2.6,1.4,4.6,4.6,8.2,2.1c10.6-7.5,22.5-9.2,35.1-8.2c1.2,0.1,1.9,1,3.3,1c-10.7,9.3-22.2,16-37.9,12.7c1.9,3.1,2.7,0.4,3.6,1.5    c-0.4,2.6-3.7,0.2-4.3,2.5c1.2,1.6,2.2-0.6,3.4-0.2c1,1.1-0.1,2.6,0.8,3.9c1.6-0.9,1.5-3.3,4.1-3.9c-2.3,5.8-6.1,9.9-11.4,12    c-7.1,2.9-14.7,3.8-22.4,2.7c-1.6-0.2-3.4-1.1-4.7-2.1c-2-1.5-4.1-2-6.4-2c-1.7,0-1.7-0.1-2.2-1.9c-0.3-1-1.3-0.7-1.9-0.3    c-4,2.9-9,2.4-13.4,3.8c-1.3,0.4-3.1,0.3-3.3-1.4c-0.2-1.5,1.4-1.5,2.6-1.9c3.6-1.3,7.3-0.4,11.2-0.9c-4.6-2.2-9.2-1.6-13.3-0.3    c-3.3,1-5.3-0.3-5.9-2.4c-0.5-1.9,0.8-4.2,3.3-5.7c3.7-2.1,7.6-1.5,11.4-2.1c5.9-1,11.3-2.5,16.3-6.1c3.2-2.3,6.9-3.9,10.9-4.9    c2.2-0.6,5-0.7,5.8-3.6c0.7-2.9-1.4-4.9-5.5-5.5C73.7,109.9,68,103.5,68,92.2C68,89.6,68,87,68,84.4z"
		pathData3 := "M76,241.7c-1.2-1.2-2.3-2.2-3.7-3.5c0.1,3.7,4.5,4.1,5.1,7.5c-2.8-0.8-3.7-4.2-6.7-3.9c0.4,2.4,3.6,2,4,4.3    c-0.6,0.6-1,0.7-1.6,0.3c-1.6-0.9-1.5-0.5-3.3,0.1c-1.9,0.7-4.1,0.1-6.1-0.4c-0.5-0.1-1.1-0.9-1.3-0.8c-3.8,1.8-11,3.3-14.4,2.1    c-3.5-1.3-6.9,0-10.3-0.5c8.1,2.7,16.9,2.2,24.8,5.3c-4.2,2.2-9.2,2.3-13.9,3.6c-2.7,0.7-4.8-2-7.8-1.9c-2.3,0.1-4.8-0.5-7.2-0.5    c-3.7,0-5.7-1.5-6.3-5.3c-0.7-4.7,2.9-4.8,5.6-5.5c5.1-1.3,10.4-1.2,15.6-2.2c6.8-1.3,13.3-4,20.1-5.3    C73,234.1,75.7,236.4,76,241.7z"
		pathData4 := "M152.6,84.3c-1.9,6.4-5.7,9.8-11.1,11.9c2.5,1.2,5.8-0.1,8.1-2.9c4-4.9,6.2-10.6,6.3-17    c0-1.9,0.9-3.5,1.8-5.5c3.4,11.7-3.5,24.5-11.8,29.6c-3.7,2.3-8.6,2.9-13,3.4c-6.1,0.7-10.7,4.3-15.2,7.6    c-3.9,2.9-8.2,5.5-10.8,10.1c-0.9,1.7-7.4,0.6-7.9-1.3c-0.3-1.3,0-2.7,2.1-3.1c2.9-0.5,4.4-4,3-6.5c-4.1-7.4-5.6-14.6,0.1-22    c1.6-2,1.5-5.3,3.6-7.5c-1.2,4.2-2.1,8.5-3.6,12.6c-1.5,4.2-0.3,8.1,0.7,11.9c0.9,3.4,8.6,5,11.5,2.5c8.1-6.9,16.9-12.8,26.8-16.7    C146.9,89.9,148.9,86.7,152.6,84.3z"
		pathData5 := "M193.1,43c-4.2,8.1-8.8,15.8-17.7,19.6c-6,2.6-7.2,8.9-10.1,13.8c-0.7,1.1,1.6,1.8,2.8,1.9    c4.7,0.2,8.5-1.9,12-4.7c1.8-1.5,3.4-3.2,4.7-5.1c0.6-0.9,1.3-1.1,1.9-0.8c1.2,0.5,0.4,1.2-0.1,1.9c-5.9,8.4-13.1,14.7-23.9,15.5    c-2,0.2-2.8-0.3-2.9-2.6c-0.1-4.4,0.6-8.2,3.3-12.1c3.7-5.4,5-11.8,4.3-18.5c-0.4-3.2,0.9-5.7,2.9-8.3c1.3,3.4-0.7,6.1-0.5,9    c0.2,2.5,0.9,4.8,3.1,5.8c2.1,0.9,4.1-0.3,5.9-1.6C184.3,52.6,188.5,47.6,193.1,43z"
		pathData6 := "M204.2,255.3c-1.3,2.2-2.3,2.1-3.3,2c-11.3-0.3-22.1,3.2-33.1,4.7c-5,0.7-10.1,0-15.1,1    c-1.2,0.3-2.5-0.2-2.6-1.4c-0.1-1.5,1.1-2.5,2.5-2.7c6.3-1.2,12.6-2.3,19-3.3c5.4-0.9,11.2-1.1,16.2-2.8    C194.6,250.5,198.4,256.2,204.2,255.3z"
		pathData7 := "M81,260.7c5.8-0.3,11.6-1.1,17.4-1.7c9-1,18.3-0.7,26.9-4.5c2.6-1.1,4.2,1.6,6.3,2.4c0.3,0.1,0.4,1-0.4,1.2    c-0.4,0.1-0.8,0.1-1.2,0.1c-15.3-0.8-30.4,4.1-45.7,2.5C83.2,260.6,82.1,260.7,81,260.7z"
		pathData8 := "M88.5,46.5c1.8,2.9,0.9,5.5,1,8c0.1,2.1,0.3,4.5,2.2,5.3c2.5,1,0.5,2.1,0.2,2.7c-0.9,1.7-1.7,0.2-2.3-0.5    C86.6,58.9,86.1,53.4,88.5,46.5z"
		pathData9 := "M105,164.2c0.9,2.4-0.1,4.8-1,6.6c-2.2,4.9-2.7,9.7-1.3,14.8c0.3,1.3,0.1,2.7-0.2,4.1    c-1-4.2-4.8-7.7-2.4-12.7c2-4.1,3.3-8.4,4.9-12.7C105.1,164.3,105,164.2,105,164.2z"
		pathData10 := "M130.3,265.4c-1.7,2.7-4.6,1.3-6.6,2.6c-0.7,0.5-1.5-0.6-1.7-1.3c-0.3-1,0.2-1.8,1-2.5    C125.8,262,127.9,264.1,130.3,265.4z"
		pathData11 := "M152.7,288.7c2.2-4.6,6.1-5.7,8.3-8.6c0.5-0.7,1.1-0.2,1.7,0.3c0.9,0.8,0.1,1.1-0.3,1.5    C160,284.6,157.3,286.8,152.7,288.7z"
		pathData12 := "M148,287.9c2.8-3.6,5.9-7,10.1-9.1c0.8-0.4,1.1-0.4,1.6,0.2c0.4,0.4,0,0.6-0.3,0.8    C155.4,282.2,152.5,286.1,148,287.9z"
		pathData13 := "M147.5,261.4c-1,4.6-5.9,8.3-10.1,8.2C141,267.1,144.1,264.9,147.5,261.4z"
		pathData14 := "M145.4,269.6c2.2-2.2,3.5-5.9,8.2-4.4C150.7,266.7,148.6,269.2,145.4,269.6z"
		pathData15 := "M137.6,248.3c2.2,0.1,3.8-0.6,4.9,1.3C140.6,250.5,139.2,249.6,137.6,248.3z"
		pathData16 := "M98.4,243.3c2.1,0.3,2.2,2.6,3.5,3.5c3.8,2.6,8.1,3.9,12.8,4c1,0,1.6-0.4,2.1-1.1c1.6-2.4,3.5-2.5,5.7-1    c-1.6,0.3-2.1,1-3.4,3.4c-0.3,0.6-1.3,0.9-2.1,1c-6.2,1.4-16-3.4-18.8-9.2C98.2,243.8,98.3,243.6,98.4,243.3z"
		pathData17 := "M84.3,214.2c-0.2,1.3-0.4,2.3,0,3.6c0.9,3-1.9,6.9-4.9,7.8c-1.5,0.4-2.2-0.5-2.2-1.7c0-3.9-0.9-8.1,3.2-12.2    c-0.2,4.3-2.8,7.5-1.2,11.1c2.6-0.7,3.4-2.5,3-5.2C82,216.3,81.8,214.6,84.3,214.2z"
		pathData18 := "M75.1,207c-0.3,1-2.9-0.1-1.4,1.9c1.4,1.8,0.4,4.1,1.2,5.9c0.5,1.2-0.2,1.5-0.8,1.7c-1,0.3-1.2-0.5-1.1-1.2    c0.9-4.4-2.4-7.6-4-11c-1.9-3.8,0.3-5.2,1.4-8c1.7,1.7,0,2.5-0.2,3.5c-0.4,1.4-1.6,2.9,1.3,2.2c0.6-0.1,1.5,0.7,1.2,1.1    C69.1,207.4,73.5,206.3,75.1,207z"
		pathData19 := "M110,222.5c-1.6,2.7-1,4.6,0.1,6.5c-2.6-1.5-5.9-2.1-6.5-5.6c-0.6-3.6,0.3-6.8,4.1-9.1c-0.1,4-5,7.3-1.2,11.3    C107.3,224.5,107.8,223.2,110,222.5z"
		pathData20 := "M112.9,202.1c-2.7,3.3-4.4,7-8.5,8.4c-0.5,0.2-1,0.7-1.4,0.1c-0.5-0.8,0-1,0.7-1.4c1.5-0.9,4.4-1.1,2.4-4.1    C108.5,204.5,109.7,201.3,112.9,202.1z"
		pathData21 := "M161.6,240c-3.5,1.9-5-1.1-7.3-1.4c-0.3,0-0.4-1.1-0.6-1.7c0.5-0.1,1.2-0.4,1.6-0.2C157.2,237.6,159.1,238.7,161.6,240z"
		pathData22 := "M108.7,200.7c1.1-1.6,4.1-2.1,2.5-5c-0.1-0.1,0.8-0.8,1.3-1.2c-0.3,1.4,2.5,1.7,1.3,3.3    C112.6,199.5,111.5,201.5,108.7,200.7z"
		pathData23 := "M66.6,196.2c1.6-2.9,2.9-4.9,4.8-6.5c0.3-0.2,0.8-0.4,1.1-0.3c0.2,0.1,0.3,0.6,0.4,1C70.4,191.5,70.7,195.4,66.6,196.2z"
		pathData24 := "M86.5,221.2c1.1,3-1.2,4.8-0.6,7.2C84.1,225.2,84.1,225,86.5,221.2z"
		pathData25 := "M188.7,222.1c0.6-0.4,1.2-0.5,1.9,0c-1.3,1.3,0.1,3.3-1.1,4.5c-0.4,0.4-0.9-0.1-1.4-0.5c-1.1-1-0.4-1.6,0.3-2.3    C188.8,223.2,189.4,222.8,188.7,222.1L188.7,222.1z"
		pathData26 := "M188.7,222c-0.6-2.7,0.4-4.2,3.5-4.6C190.9,219.2,189.8,220.6,188.7,222C188.7,222.1,188.7,222,188.7,222z"
		pathData27 := "M141.5,189.1c2.3-1.9,4-0.4,5.7,0C145.5,189.1,143.8,189.1,141.5,189.1z"
		pathData28 := "M101.1,151.8c1-1.1,1.4-2.2,2.7-2.2c0.2,0,0.6,0.6,0.6,0.9C103.9,151.9,102.9,152.3,101.1,151.8z"
		pathData29 := "M158,229.8c1.5-0.7,2.6-1.1,3.6,0c0.1,0.1,0,0.6-0.1,0.7C160.4,230.9,159.3,231.2,158,229.8z"
		pathData30 := "M112.2,226.7c1,0.9,2.1,1.4,2.1,2.6c0,0.2-0.6,0.6-0.8,0.6C112.1,229.4,111.6,228.5,112.2,226.7z"
		pathData31 := "M192.3,214.5c-0.9,1.6-2.4,1.7-4.6,2.8C189.4,215.2,190.1,213.7,192.3,214.5z"
		pathData32 := "M244.4,211.3c-1.6-1.1-2.2-2.1-2-4.1C244,208.3,244.6,209.3,244.4,211.3z"
		pathData33 := "M243.9,236.7c0.8-1.8,1.8-1.8,2.7-1.6c0.2,0,0.5,0.8,0.4,1C246.2,237.2,245.1,236.3,243.9,236.7z"
		pathData34 := "M96.1,186.8c0.5,1.4-1,2-1.2,3.1c0-1.2-1.5-3.2,1.4-3C96.2,186.9,96.1,186.8,96.1,186.8z"
		pathData35 := "M73.5,184.1c0.8,1.6-0.4,2.1-1.6,3c-0.2-1.7,0-2.8,1.7-2.9L73.5,184.1z"
		pathData36 := "M252.1,212.3c-0.1,1.4,0.1,2.4-0.2,3.2c-0.2,0.4-0.9-0.3-1.1-0.6C249.9,213.6,251.2,213.3,252.1,212.3z"
		pathData37 := "M246.6,210.2c1.3,0.8,1.7,1.7,1.1,3C246.3,212.7,246,211.7,246.6,210.2z"
		pathData38 := "M84.5,207.2c1.3-0.2,1.8-1.6,3.2-1C87.1,207.6,86.2,207.9,84.5,207.2z"
		pathData39 := "M73.7,184.3c0.9-0.5,0.6-2.1,2.6-2.5c-0.9,1.3-0.2,3.3-2.7,2.4C73.5,184.1,73.7,184.3,73.7,184.3z"
		pathData40 := "M101.7,154.2c1.5-0.3,2-1.5,3.3-1.2C104.7,154.6,103.7,154.7,101.7,154.2z"
		pathData41 := "M81.8,132.3c-0.3-4.6,3.5-2.8,5.2-4c2.4-1.8,4.7-2.9,7.3-0.9c2.2,1.7,4.2,3.8,6.1,6.2c-2.8-1.5-4.6-0.1-6.6,1.3    c-1.9,1.4-4.2,0.6-5.6-0.7C86.2,132.5,84.4,131.6,81.8,132.3z"
		pathData42 := "M129.8,124.5c-1.6-0.7,1.5-2.6-0.9-3.4c-0.8-0.2-1.9-0.9-2,0c-0.3,4.6-3.6,1.1-5.2,2.1c0.7,1.3,1.8,2.3,3.1,2.6    c-2.8,3.3-5.9-1.1-9,0.5c4.3-6.9,10.8-9.5,17.6-7.6c1.5,0.4,0.2,1.2,0.6,1.7c0.6,0.6,0.9,1.7-0.5,1.6    C131.5,122.1,131.1,124,129.8,124.5z"
		pathData43 := "M86.2,121.4c-1,2-3,2.6-3.4,5.4c2.6-2.7,4.8-4.7,8.2-4.2c-3.8,0.8-6,4.3-9.5,5.6c-0.5,0.2-0.9,0.5-1.3-0.1    c-0.9-1,0.5-0.8,0.6-1.2C81.9,124.6,82.1,121.4,86.2,121.4z"
		pathData44 := "M88.2,139.3c-1.6-1.5-2.4-2.2-3.4-3.1c-0.5,2.8,3.1,3.1,3,5.4c-3.8-1.2-3.4-2.1-5.6-8C84.9,134.6,87.5,135,88.2,139.3z"
		pathData45 := "M51.3,133.5c0.1,3.4-2,3.6-3.9,4.3C46.8,134.6,49.3,134.7,51.3,133.5z"
		pathData46 := "M91.3,132.8c1.9,0.2,0.9-3.1,3-1.9c0.6,0.3,0.4,1.6-0.3,2.2c-1.5,1.3-2.9,1.1-4.3-0.3    c-0.5-0.5-1.6-1.4-0.7-1.9C91.2,129.6,89.8,133.4,91.3,132.8z"
		canvas.Scale(0.8)
		canvas.Translate(50, 20)
		canvas.Path(pathData0, "fill:"+lineColor)
		canvas.Path(pathData1, "fill:"+hexCode)
		canvas.Path(pathData2, "fill:"+hexCode)
		canvas.Path(pathData3, "fill:"+hexCode)
		canvas.Path(pathData4, "fill:"+hexCode)
		canvas.Path(pathData5, "fill:"+hexCode)
		canvas.Path(pathData6, "fill:"+hexCode)
		canvas.Path(pathData7, "fill:"+hexCode)
		canvas.Path(pathData8, "fill:"+hexCode)
		canvas.Path(pathData9, "fill:"+hexCode)
		canvas.Path(pathData10, "fill:"+hexCode)
		canvas.Path(pathData11, "fill:"+hexCode)
		canvas.Path(pathData12, "fill:"+hexCode)
		canvas.Path(pathData13, "fill:"+hexCode)
		canvas.Path(pathData14, "fill:"+hexCode)
		canvas.Path(pathData15, "fill:"+hexCode)
		canvas.Path(pathData16, "fill:"+lineColor)
		canvas.Path(pathData17, "fill:"+lineColor)
		canvas.Path(pathData18, "fill:"+lineColor)
		canvas.Path(pathData19, "fill:"+lineColor)
		canvas.Path(pathData20, "fill:"+lineColor)
		canvas.Path(pathData21, "fill:"+lineColor)
		canvas.Path(pathData22, "fill:"+lineColor)
		canvas.Path(pathData23, "fill:"+lineColor)
		canvas.Path(pathData24, "fill:"+lineColor)
		canvas.Path(pathData25, "fill:"+lineColor)
		canvas.Path(pathData26, "fill:"+lineColor)
		canvas.Path(pathData27, "fill:"+lineColor)
		canvas.Path(pathData28, "fill:"+lineColor)
		canvas.Path(pathData29, "fill:"+lineColor)
		canvas.Path(pathData30, "fill:"+lineColor)
		canvas.Path(pathData31, "fill:"+lineColor)
		canvas.Path(pathData32, "fill:"+lineColor)
		canvas.Path(pathData33, "fill:"+lineColor)
		canvas.Path(pathData34, "fill:"+lineColor)
		canvas.Path(pathData35, "fill:"+lineColor)
		canvas.Path(pathData36, "fill:"+lineColor)
		canvas.Path(pathData37, "fill:"+lineColor)
		canvas.Path(pathData38, "fill:"+lineColor)
		canvas.Path(pathData39, "fill:"+lineColor)
		canvas.Path(pathData40, "fill:"+lineColor)
		canvas.Path(pathData41, "fill:"+lineColor)
		canvas.Path(pathData42, "fill:"+lineColor)
		canvas.Path(pathData43, "fill:"+lineColor)
		canvas.Path(pathData44, "fill:"+lineColor)
		canvas.Path(pathData45, "fill:"+lineColor)
		canvas.Path(pathData46, "fill:"+hexCode)
		canvas.Gend()
		canvas.Gend()
	},
}
